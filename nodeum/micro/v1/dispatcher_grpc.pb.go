// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: nodeum/micro/v1/dispatcher.proto

package microv1

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
	DispatcherService_Start_FullMethodName   = "/nodeum.micro.v1.DispatcherService/Start"
	DispatcherService_Pause_FullMethodName   = "/nodeum.micro.v1.DispatcherService/Pause"
	DispatcherService_Resume_FullMethodName  = "/nodeum.micro.v1.DispatcherService/Resume"
	DispatcherService_Stop_FullMethodName    = "/nodeum.micro.v1.DispatcherService/Stop"
	DispatcherService_ReadDir_FullMethodName = "/nodeum.micro.v1.DispatcherService/ReadDir"
)

// DispatcherServiceClient is the client API for DispatcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherServiceClient interface {
	Start(ctx context.Context, in *DispatcherServiceStartRequest, opts ...grpc.CallOption) (*DispatcherServiceStartResponse, error)
	Pause(ctx context.Context, in *DispatcherServicePauseRequest, opts ...grpc.CallOption) (*DispatcherServicePauseResponse, error)
	Resume(ctx context.Context, in *DispatcherServiceResumeRequest, opts ...grpc.CallOption) (*DispatcherServiceResumeResponse, error)
	Stop(ctx context.Context, in *DispatcherServiceStopRequest, opts ...grpc.CallOption) (*DispatcherServiceStopResponse, error)
	ReadDir(ctx context.Context, in *DispatcherServiceReadDirRequest, opts ...grpc.CallOption) (DispatcherService_ReadDirClient, error)
}

type dispatcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherServiceClient(cc grpc.ClientConnInterface) DispatcherServiceClient {
	return &dispatcherServiceClient{cc}
}

func (c *dispatcherServiceClient) Start(ctx context.Context, in *DispatcherServiceStartRequest, opts ...grpc.CallOption) (*DispatcherServiceStartResponse, error) {
	out := new(DispatcherServiceStartResponse)
	err := c.cc.Invoke(ctx, DispatcherService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) Pause(ctx context.Context, in *DispatcherServicePauseRequest, opts ...grpc.CallOption) (*DispatcherServicePauseResponse, error) {
	out := new(DispatcherServicePauseResponse)
	err := c.cc.Invoke(ctx, DispatcherService_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) Resume(ctx context.Context, in *DispatcherServiceResumeRequest, opts ...grpc.CallOption) (*DispatcherServiceResumeResponse, error) {
	out := new(DispatcherServiceResumeResponse)
	err := c.cc.Invoke(ctx, DispatcherService_Resume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) Stop(ctx context.Context, in *DispatcherServiceStopRequest, opts ...grpc.CallOption) (*DispatcherServiceStopResponse, error) {
	out := new(DispatcherServiceStopResponse)
	err := c.cc.Invoke(ctx, DispatcherService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) ReadDir(ctx context.Context, in *DispatcherServiceReadDirRequest, opts ...grpc.CallOption) (DispatcherService_ReadDirClient, error) {
	stream, err := c.cc.NewStream(ctx, &DispatcherService_ServiceDesc.Streams[0], DispatcherService_ReadDir_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dispatcherServiceReadDirClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DispatcherService_ReadDirClient interface {
	Recv() (*DispatcherServiceReadDirResponse, error)
	grpc.ClientStream
}

type dispatcherServiceReadDirClient struct {
	grpc.ClientStream
}

func (x *dispatcherServiceReadDirClient) Recv() (*DispatcherServiceReadDirResponse, error) {
	m := new(DispatcherServiceReadDirResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DispatcherServiceServer is the server API for DispatcherService service.
// All implementations must embed UnimplementedDispatcherServiceServer
// for forward compatibility
type DispatcherServiceServer interface {
	Start(context.Context, *DispatcherServiceStartRequest) (*DispatcherServiceStartResponse, error)
	Pause(context.Context, *DispatcherServicePauseRequest) (*DispatcherServicePauseResponse, error)
	Resume(context.Context, *DispatcherServiceResumeRequest) (*DispatcherServiceResumeResponse, error)
	Stop(context.Context, *DispatcherServiceStopRequest) (*DispatcherServiceStopResponse, error)
	ReadDir(*DispatcherServiceReadDirRequest, DispatcherService_ReadDirServer) error
	mustEmbedUnimplementedDispatcherServiceServer()
}

// UnimplementedDispatcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherServiceServer struct {
}

func (UnimplementedDispatcherServiceServer) Start(context.Context, *DispatcherServiceStartRequest) (*DispatcherServiceStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedDispatcherServiceServer) Pause(context.Context, *DispatcherServicePauseRequest) (*DispatcherServicePauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedDispatcherServiceServer) Resume(context.Context, *DispatcherServiceResumeRequest) (*DispatcherServiceResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedDispatcherServiceServer) Stop(context.Context, *DispatcherServiceStopRequest) (*DispatcherServiceStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedDispatcherServiceServer) ReadDir(*DispatcherServiceReadDirRequest, DispatcherService_ReadDirServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadDir not implemented")
}
func (UnimplementedDispatcherServiceServer) mustEmbedUnimplementedDispatcherServiceServer() {}

// UnsafeDispatcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherServiceServer will
// result in compilation errors.
type UnsafeDispatcherServiceServer interface {
	mustEmbedUnimplementedDispatcherServiceServer()
}

func RegisterDispatcherServiceServer(s grpc.ServiceRegistrar, srv DispatcherServiceServer) {
	s.RegisterService(&DispatcherService_ServiceDesc, srv)
}

func _DispatcherService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherServiceStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Start(ctx, req.(*DispatcherServiceStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherServicePauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherService_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Pause(ctx, req.(*DispatcherServicePauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherServiceResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherService_Resume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Resume(ctx, req.(*DispatcherServiceResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherServiceStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Stop(ctx, req.(*DispatcherServiceStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_ReadDir_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DispatcherServiceReadDirRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DispatcherServiceServer).ReadDir(m, &dispatcherServiceReadDirServer{stream})
}

type DispatcherService_ReadDirServer interface {
	Send(*DispatcherServiceReadDirResponse) error
	grpc.ServerStream
}

type dispatcherServiceReadDirServer struct {
	grpc.ServerStream
}

func (x *dispatcherServiceReadDirServer) Send(m *DispatcherServiceReadDirResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DispatcherService_ServiceDesc is the grpc.ServiceDesc for DispatcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodeum.micro.v1.DispatcherService",
	HandlerType: (*DispatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _DispatcherService_Start_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _DispatcherService_Pause_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _DispatcherService_Resume_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _DispatcherService_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadDir",
			Handler:       _DispatcherService_ReadDir_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nodeum/micro/v1/dispatcher.proto",
}
