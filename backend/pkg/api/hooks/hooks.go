// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package hooks contains the interface definitions of different hooks
package hooks

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/cloudhut/common/rest"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/redpanda-data/console/backend/pkg/api/httptypes"
	pkgconnect "github.com/redpanda-data/console/backend/pkg/connect"
	"github.com/redpanda-data/console/backend/pkg/console"
	"github.com/redpanda-data/console/backend/pkg/redpanda"
)

// ConfigConnectRPCRequest is the config object that is passed into the
// hook to configure the Connect API. The hook implementation can use
// this to control the behaviour of the connect API (e.g. change order,
// add additional interceptors, mount more routes etc).
type ConfigConnectRPCRequest struct {
	BaseInterceptors []connect.Interceptor
	GRPCGatewayMux   *runtime.ServeMux
}

// ConfigConnectRPCResponse configures connect services.
type ConfigConnectRPCResponse struct {
	// Instructs OSS to use these intercptors for all connect services
	Interceptors []connect.Interceptor

	// Instructs OSS to register these services in addition to the OSS ones
	AdditionalServices []ConnectService
}

// ConnectService is a Connect handler along with its metadata
// that is required to mount the service in the mux as well
// as advertise it in the gRPC reflector.
type ConnectService struct {
	ServiceName string
	MountPath   string
	Handler     http.Handler
}

// RouteHooks allow you to modify the Router
type RouteHooks interface {
	// ConfigAPIRouter allows you to modify the router responsible for all /api routes
	ConfigAPIRouter(router chi.Router)

	// ConfigAPIRouterPostRegistration allows you to modify the router responsible for
	// all /api routes after all routes have been registered.
	ConfigAPIRouterPostRegistration(router chi.Router)

	// ConfigInternalRouter allows you to modify the router responsible for all internal /admin/* routes
	ConfigInternalRouter(router chi.Router)

	// ConfigRouter allows you to modify the router responsible for all non /api and non /admin routes.
	// By default we serve the frontend on these routes.
	ConfigRouter(router chi.Router)

	// ConfigConnectRPC receives the basic interceptors used by OSS.
	// The hook can modify the interceptors slice, i.e. adding new interceptors, removing some, re-ordering, and return it in ConnectConfig.
	// The hook can return additional connect services that shall be mounted by OSS.
	ConfigConnectRPC(ConfigConnectRPCRequest) ConfigConnectRPCResponse

	// InitConnectRPCRouter is used to initialize the ConnectRPC router with any top level middleware.
	InitConnectRPCRouter(router chi.Router)
}

// AuthorizationHooks include all functions which allow you to intercept the requests at various
// endpoints where RBAC rules may be applied.
type AuthorizationHooks interface {
	// Topic Hooks
	CanSeeTopic(ctx context.Context, topicName string) (bool, *rest.Error)
	CanCreateTopic(ctx context.Context, topicName string) (bool, *rest.Error)
	CanEditTopicConfig(ctx context.Context, topicName string) (bool, *rest.Error)
	CanDeleteTopic(ctx context.Context, topicName string) (bool, *rest.Error)
	CanPublishTopicRecords(ctx context.Context, topicName string) (bool, *rest.Error)
	CanDeleteTopicRecords(ctx context.Context, topicName string) (bool, *rest.Error)
	CanViewTopicPartitions(ctx context.Context, topicName string) (bool, *rest.Error)
	CanViewTopicConfig(ctx context.Context, topicName string) (bool, *rest.Error)
	CanViewTopicMessages(ctx context.Context, req *httptypes.ListMessagesRequest) (bool, *rest.Error)
	CanUseMessageSearchFilters(ctx context.Context, req *httptypes.ListMessagesRequest) (bool, *rest.Error)
	CanViewTopicConsumers(ctx context.Context, topicName string) (bool, *rest.Error)
	AllowedTopicActions(ctx context.Context, topicName string) ([]string, *rest.Error)
	PrintListMessagesAuditLog(ctx context.Context, r any, req *console.ListMessageRequest)

	// ACL Hooks
	CanListACLs(ctx context.Context) (bool, *rest.Error)
	CanCreateACL(ctx context.Context) (bool, *rest.Error)
	CanDeleteACL(ctx context.Context) (bool, *rest.Error)

	// Quotas Hookas
	CanListQuotas(ctx context.Context) (bool, *rest.Error)

	// ConsumerGroup Hooks
	CanSeeConsumerGroup(ctx context.Context, groupName string) (bool, *rest.Error)
	CanEditConsumerGroup(ctx context.Context, groupName string) (bool, *rest.Error)
	CanDeleteConsumerGroup(ctx context.Context, groupName string) (bool, *rest.Error)
	AllowedConsumerGroupActions(ctx context.Context, groupName string) ([]string, *rest.Error)

	// Operations Hooks
	CanPatchPartitionReassignments(ctx context.Context) (bool, *rest.Error)
	CanPatchConfigs(ctx context.Context) (bool, *rest.Error)

	// Kafka Connect Hooks
	CanViewConnectCluster(ctx context.Context, clusterName string) (bool, *rest.Error)
	CanEditConnectCluster(ctx context.Context, clusterName string) (bool, *rest.Error)
	CanDeleteConnectCluster(ctx context.Context, clusterName string) (bool, *rest.Error)
	AllowedConnectClusterActions(ctx context.Context, clusterName string) ([]string, *rest.Error)

	// Kafka User Hooks
	CanListKafkaUsers(ctx context.Context) (bool, *rest.Error)
	CanCreateKafkaUsers(ctx context.Context) (bool, *rest.Error)
	CanDeleteKafkaUsers(ctx context.Context) (bool, *rest.Error)
	IsProtectedKafkaUser(userName string) bool

	// Schema Registry Hooks
	CanViewSchemas(ctx context.Context) (bool, *rest.Error)
	CanCreateSchemas(ctx context.Context) (bool, *rest.Error)
	CanDeleteSchemas(ctx context.Context) (bool, *rest.Error)
	CanManageSchemaRegistry(ctx context.Context) (bool, *rest.Error)

	// Redpanda Role Hooks
	CanListRedpandaRoles(ctx context.Context) (bool, *rest.Error)
	CanCreateRedpandaRoles(ctx context.Context) (bool, *rest.Error)
	CanDeleteRedpandaRoles(ctx context.Context) (bool, *rest.Error)
}

// ConsoleHooks are hooks for providing additional context to the Frontend where needed.
// This could be information about what license is used, what enterprise features are
// enabled etc.
type ConsoleHooks interface {
	// ConsoleLicenseInformation returns the license information for Console.
	// Based on the returned license the frontend will display the
	// appropriate UI and also warnings if the license is (about to be) expired.
	ConsoleLicenseInformation(ctx context.Context) redpanda.License

	// EnabledFeatures returns a list of string enums that indicate what features are enabled.
	// Only toggleable features that require conditional rendering in the Frontend will be returned.
	// The information will be baked into the index.html so that the Frontend knows about it
	// at startup, which might be important to not block rendering (e.g. SSO enabled -> render login).
	EnabledFeatures() []string

	// EnabledConnectClusterFeatures returns a list of features that are supported on this
	// particular Kafka connect cluster.
	EnabledConnectClusterFeatures(ctx context.Context, clusterName string) []pkgconnect.ClusterFeature

	// EndpointCompatibility returns information what endpoints are available to the frontend.
	// This considers the active configuration (e.g. is secret store enabled), target cluster
	// version and what features are supported by our upstream systems.
	// The response of this hook will be merged into the response that was originally
	// composed by Console.
	EndpointCompatibility() []console.EndpointCompatibilityEndpoint

	// CheckWebsocketConnection extracts metadata from the websocket request.
	// Because some metadata is part of the HTTP request and other metadata is part
	// of the first websocket message sent, a middleware can not be used here.
	// The returned context must be used for subsequent requests. The Websocket
	// connection must be closed if an error is returned.
	CheckWebsocketConnection(r *http.Request, req httptypes.ListMessagesRequest) (context.Context, error)
}
