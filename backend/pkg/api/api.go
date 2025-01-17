// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package api is the first layer that processes an incoming HTTP request. It is in charge
// of validating the user input, verifying whether the user is authorized (by calling hooks),
// as well as setting up all the routes and dependencies for subsequently called
// route handlers & services.
package api

import (
	"context"
	"fmt"
	"io/fs"
	"time"

	"github.com/cloudhut/common/logging"
	"github.com/cloudhut/common/rest"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/redpanda-data/console/backend/pkg/config"
	"github.com/redpanda-data/console/backend/pkg/connect"
	"github.com/redpanda-data/console/backend/pkg/console"
	"github.com/redpanda-data/console/backend/pkg/embed"
	"github.com/redpanda-data/console/backend/pkg/git"
	"github.com/redpanda-data/console/backend/pkg/redpanda"
	"github.com/redpanda-data/console/backend/pkg/version"
)

//go:generate mockgen -destination=./mocks/authz_hooks.go -package=mocks . AuthorizationHooks

// API represents the server and all it's dependencies to serve incoming user requests
type API struct {
	Cfg *config.Config

	Logger      *zap.Logger
	ConsoleSvc  console.Servicer
	ConnectSvc  *connect.Service
	GitSvc      *git.Service
	RedpandaSvc *redpanda.Service

	// FrontendResources is an in-memory Filesystem with all go:embedded frontend resources.
	// The index.html is expected to be at the root of the filesystem. This prop will only be accessed
	// if the config property serveFrontend is set to true.
	FrontendResources fs.FS

	// License is the license information for Console that will be used for logging
	// and visibility purposes inside the open source version. License protected features
	// are not checked with this license.
	License redpanda.License

	// Hooks to add additional functionality from the outside at different places
	Hooks *Hooks

	// internal server intance
	server *rest.Server
}

// New creates a new API instance
func New(cfg *config.Config, opts ...Option) *API {
	logger := logging.NewLogger(&cfg.Logger, cfg.MetricsNamespace)

	logger.Info("started Redpanda Console",
		zap.String("version", version.Version),
		zap.String("built_at", version.BuiltAt))

	redpandaSvc, err := redpanda.NewService(cfg.Redpanda, logger)
	if err != nil {
		logger.Fatal("failed to create Redpanda service", zap.Error(err))
	}

	connectSvc, err := connect.NewService(cfg.Connect, logger)
	if err != nil {
		logger.Fatal("failed to create Kafka connect service", zap.Error(err))
	}

	var consoleSvc console.Servicer
	if cfg.Console.Enabled {
		consoleSvc, err = console.NewService(cfg, logger, redpandaSvc, connectSvc)
		if err != nil {
			logger.Fatal("failed to create console service", zap.Error(err))
		}
	}

	// Use default frontend resources from embeds. They may be overridden via functional options.
	// We don't use hooks here because we may want to use the API struct without providing all hooks.
	fsys, err := fs.Sub(embed.FrontendFiles, "frontend")
	if err != nil {
		logger.Fatal("failed to build subtree from embedded frontend files", zap.Error(err))
	}

	year := 24 * time.Hour * 365
	a := &API{
		Cfg:               cfg,
		Logger:            logger,
		ConsoleSvc:        consoleSvc,
		ConnectSvc:        connectSvc,
		RedpandaSvc:       redpandaSvc,
		Hooks:             newDefaultHooks(),
		FrontendResources: fsys,
		License: redpanda.License{
			Source:    redpanda.LicenseSourceConsole,
			Type:      redpanda.LicenseTypeOpenSource,
			ExpiresAt: time.Now().Add(year * 10).Unix(),
		},
	}
	for _, opt := range opts {
		opt(a)
	}

	return a
}

// Start the API server and block
func (api *API) Start() {
	err := api.ConsoleSvc.Start()
	if err != nil {
		api.Logger.Fatal("failed to start console service", zap.Error(err))
	}

	mux := api.routes()

	// Server
	api.server, err = rest.NewServer(&api.Cfg.REST.Config, api.Logger, mux)
	if err != nil {
		api.Logger.Fatal("failed to create HTTP server", zap.Error(err))
	}

	// need this to make gRPC protocol work
	api.server.Server.Handler = h2c.NewHandler(mux, &http2.Server{})

	err = api.server.Start()
	if err != nil {
		api.Logger.Fatal("REST Server returned an error", zap.Error(err))
	}
}

// Stop gracefully stops the API.
func (api *API) Stop(ctx context.Context) error {
	err := api.server.Server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to shutdown HTTP server: %w", err)
	}
	return nil
}
