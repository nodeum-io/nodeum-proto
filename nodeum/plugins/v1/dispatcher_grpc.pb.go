// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nodeum/plugins/v1/dispatcher.proto

package pluginsv1

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

// DispatcherPluginServiceClient is the client API for DispatcherPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherPluginServiceClient interface {
	// Before the task is started, you can change the task and/or the trigger metadata
	BeforeTaskExecuted(ctx context.Context, in *BeforeTaskExecutedRequest, opts ...grpc.CallOption) (*BeforeTaskExecutedResponse, error)
	// Before a request is dispatched, you can change it or discard it
	BeforeRequestDispatched(ctx context.Context, in *BeforeRequestDispatchedRequest, opts ...grpc.CallOption) (*BeforeRequestDispatchedResponse, error)
}

type dispatcherPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherPluginServiceClient(cc grpc.ClientConnInterface) DispatcherPluginServiceClient {
	return &dispatcherPluginServiceClient{cc}
}

func (c *dispatcherPluginServiceClient) BeforeTaskExecuted(ctx context.Context, in *BeforeTaskExecutedRequest, opts ...grpc.CallOption) (*BeforeTaskExecutedResponse, error) {
	out := new(BeforeTaskExecutedResponse)
	err := c.cc.Invoke(ctx, "/nodeum.plugins.v1.DispatcherPluginService/BeforeTaskExecuted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherPluginServiceClient) BeforeRequestDispatched(ctx context.Context, in *BeforeRequestDispatchedRequest, opts ...grpc.CallOption) (*BeforeRequestDispatchedResponse, error) {
	out := new(BeforeRequestDispatchedResponse)
	err := c.cc.Invoke(ctx, "/nodeum.plugins.v1.DispatcherPluginService/BeforeRequestDispatched", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherPluginServiceServer is the server API for DispatcherPluginService service.
// All implementations must embed UnimplementedDispatcherPluginServiceServer
// for forward compatibility
type DispatcherPluginServiceServer interface {
	// Before the task is started, you can change the task and/or the trigger metadata
	BeforeTaskExecuted(context.Context, *BeforeTaskExecutedRequest) (*BeforeTaskExecutedResponse, error)
	// Before a request is dispatched, you can change it or discard it
	BeforeRequestDispatched(context.Context, *BeforeRequestDispatchedRequest) (*BeforeRequestDispatchedResponse, error)
	mustEmbedUnimplementedDispatcherPluginServiceServer()
}

// UnimplementedDispatcherPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherPluginServiceServer struct {
}

func (UnimplementedDispatcherPluginServiceServer) BeforeTaskExecuted(context.Context, *BeforeTaskExecutedRequest) (*BeforeTaskExecutedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeforeTaskExecuted not implemented")
}
func (UnimplementedDispatcherPluginServiceServer) BeforeRequestDispatched(context.Context, *BeforeRequestDispatchedRequest) (*BeforeRequestDispatchedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeforeRequestDispatched not implemented")
}
func (UnimplementedDispatcherPluginServiceServer) mustEmbedUnimplementedDispatcherPluginServiceServer() {
}

// UnsafeDispatcherPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherPluginServiceServer will
// result in compilation errors.
type UnsafeDispatcherPluginServiceServer interface {
	mustEmbedUnimplementedDispatcherPluginServiceServer()
}

func RegisterDispatcherPluginServiceServer(s grpc.ServiceRegistrar, srv DispatcherPluginServiceServer) {
	s.RegisterService(&DispatcherPluginService_ServiceDesc, srv)
}

func _DispatcherPluginService_BeforeTaskExecuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeforeTaskExecutedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherPluginServiceServer).BeforeTaskExecuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeum.plugins.v1.DispatcherPluginService/BeforeTaskExecuted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherPluginServiceServer).BeforeTaskExecuted(ctx, req.(*BeforeTaskExecutedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherPluginService_BeforeRequestDispatched_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeforeRequestDispatchedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherPluginServiceServer).BeforeRequestDispatched(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeum.plugins.v1.DispatcherPluginService/BeforeRequestDispatched",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherPluginServiceServer).BeforeRequestDispatched(ctx, req.(*BeforeRequestDispatchedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherPluginService_ServiceDesc is the grpc.ServiceDesc for DispatcherPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodeum.plugins.v1.DispatcherPluginService",
	HandlerType: (*DispatcherPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeforeTaskExecuted",
			Handler:    _DispatcherPluginService_BeforeTaskExecuted_Handler,
		},
		{
			MethodName: "BeforeRequestDispatched",
			Handler:    _DispatcherPluginService_BeforeRequestDispatched_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeum/plugins/v1/dispatcher.proto",
}
