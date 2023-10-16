// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: redpanda/api/dataplane/v1alpha1/topic.proto

package dataplanev1alpha1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TopicService_CreateTopic_FullMethodName              = "/redpanda.api.dataplane.v1alpha1.TopicService/CreateTopic"
	TopicService_ListTopics_FullMethodName               = "/redpanda.api.dataplane.v1alpha1.TopicService/ListTopics"
	TopicService_DeleteTopic_FullMethodName              = "/redpanda.api.dataplane.v1alpha1.TopicService/DeleteTopic"
	TopicService_GetTopicConfiguration_FullMethodName    = "/redpanda.api.dataplane.v1alpha1.TopicService/GetTopicConfiguration"
	TopicService_UpdateTopicConfiguration_FullMethodName = "/redpanda.api.dataplane.v1alpha1.TopicService/UpdateTopicConfiguration"
	TopicService_SetTopicConfiguration_FullMethodName    = "/redpanda.api.dataplane.v1alpha1.TopicService/SetTopicConfiguration"
)

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error)
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error)
	GetTopicConfiguration(ctx context.Context, in *GetTopicConfigurationRequest, opts ...grpc.CallOption) (*GetTopicConfigurationResponse, error)
	UpdateTopicConfiguration(ctx context.Context, in *UpdateTopicConfigurationRequest, opts ...grpc.CallOption) (*UpdateTopicConfigurationResponse, error)
	SetTopicConfiguration(ctx context.Context, in *SetTopicConfigurationRequest, opts ...grpc.CallOption) (*SetTopicConfigurationResponse, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_CreateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, TopicService_ListTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error) {
	out := new(DeleteTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_DeleteTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopicConfiguration(ctx context.Context, in *GetTopicConfigurationRequest, opts ...grpc.CallOption) (*GetTopicConfigurationResponse, error) {
	out := new(GetTopicConfigurationResponse)
	err := c.cc.Invoke(ctx, TopicService_GetTopicConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopicConfiguration(ctx context.Context, in *UpdateTopicConfigurationRequest, opts ...grpc.CallOption) (*UpdateTopicConfigurationResponse, error) {
	out := new(UpdateTopicConfigurationResponse)
	err := c.cc.Invoke(ctx, TopicService_UpdateTopicConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) SetTopicConfiguration(ctx context.Context, in *SetTopicConfigurationRequest, opts ...grpc.CallOption) (*SetTopicConfigurationResponse, error) {
	out := new(SetTopicConfigurationResponse)
	err := c.cc.Invoke(ctx, TopicService_SetTopicConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations must embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error)
	GetTopicConfiguration(context.Context, *GetTopicConfigurationRequest) (*GetTopicConfigurationResponse, error)
	UpdateTopicConfiguration(context.Context, *UpdateTopicConfigurationRequest) (*UpdateTopicConfigurationResponse, error)
	SetTopicConfiguration(context.Context, *SetTopicConfigurationRequest) (*SetTopicConfigurationResponse, error)
	mustEmbedUnimplementedTopicServiceServer()
}

// UnimplementedTopicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedTopicServiceServer) ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedTopicServiceServer) DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedTopicServiceServer) GetTopicConfiguration(context.Context, *GetTopicConfigurationRequest) (*GetTopicConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicConfiguration not implemented")
}
func (UnimplementedTopicServiceServer) UpdateTopicConfiguration(context.Context, *UpdateTopicConfigurationRequest) (*UpdateTopicConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopicConfiguration not implemented")
}
func (UnimplementedTopicServiceServer) SetTopicConfiguration(context.Context, *SetTopicConfigurationRequest) (*SetTopicConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTopicConfiguration not implemented")
}
func (UnimplementedTopicServiceServer) mustEmbedUnimplementedTopicServiceServer() {}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_ListTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopicConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopicConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetTopicConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopicConfiguration(ctx, req.(*GetTopicConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopicConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopicConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_UpdateTopicConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopicConfiguration(ctx, req.(*UpdateTopicConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_SetTopicConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTopicConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).SetTopicConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_SetTopicConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).SetTopicConfiguration(ctx, req.(*SetTopicConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redpanda.api.dataplane.v1alpha1.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _TopicService_ListTopics_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
		{
			MethodName: "GetTopicConfiguration",
			Handler:    _TopicService_GetTopicConfiguration_Handler,
		},
		{
			MethodName: "UpdateTopicConfiguration",
			Handler:    _TopicService_UpdateTopicConfiguration_Handler,
		},
		{
			MethodName: "SetTopicConfiguration",
			Handler:    _TopicService_SetTopicConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redpanda/api/dataplane/v1alpha1/topic.proto",
}