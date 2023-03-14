// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: nodeum/micro/v1/scheduler.proto

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
	SchedulerService_ReadSchedule_FullMethodName   = "/nodeum.micro.v1.SchedulerService/ReadSchedule"
	SchedulerService_ListSchedules_FullMethodName  = "/nodeum.micro.v1.SchedulerService/ListSchedules"
	SchedulerService_WriteSchedule_FullMethodName  = "/nodeum.micro.v1.SchedulerService/WriteSchedule"
	SchedulerService_DeleteSchedule_FullMethodName = "/nodeum.micro.v1.SchedulerService/DeleteSchedule"
)

// SchedulerServiceClient is the client API for SchedulerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerServiceClient interface {
	// ReadSchedule returns information about a specific schedule, including date
	// of next occurrence.
	ReadSchedule(ctx context.Context, in *ReadScheduleRequest, opts ...grpc.CallOption) (*ReadScheduleResponse, error)
	// ListSchedules returns all actives schedules.
	ListSchedules(ctx context.Context, in *ListSchedulesRequest, opts ...grpc.CallOption) (*ListSchedulesResponse, error)
	// WriteSchedule creates or update a schedule for a task.
	// If one schedule already exists for the task, it will be overwritten.
	WriteSchedule(ctx context.Context, in *WriteScheduleRequest, opts ...grpc.CallOption) (*WriteScheduleResponse, error)
	DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*DeleteScheduleResponse, error)
}

type schedulerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerServiceClient(cc grpc.ClientConnInterface) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) ReadSchedule(ctx context.Context, in *ReadScheduleRequest, opts ...grpc.CallOption) (*ReadScheduleResponse, error) {
	out := new(ReadScheduleResponse)
	err := c.cc.Invoke(ctx, SchedulerService_ReadSchedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ListSchedules(ctx context.Context, in *ListSchedulesRequest, opts ...grpc.CallOption) (*ListSchedulesResponse, error) {
	out := new(ListSchedulesResponse)
	err := c.cc.Invoke(ctx, SchedulerService_ListSchedules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) WriteSchedule(ctx context.Context, in *WriteScheduleRequest, opts ...grpc.CallOption) (*WriteScheduleResponse, error) {
	out := new(WriteScheduleResponse)
	err := c.cc.Invoke(ctx, SchedulerService_WriteSchedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*DeleteScheduleResponse, error) {
	out := new(DeleteScheduleResponse)
	err := c.cc.Invoke(ctx, SchedulerService_DeleteSchedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServiceServer is the server API for SchedulerService service.
// All implementations must embed UnimplementedSchedulerServiceServer
// for forward compatibility
type SchedulerServiceServer interface {
	// ReadSchedule returns information about a specific schedule, including date
	// of next occurrence.
	ReadSchedule(context.Context, *ReadScheduleRequest) (*ReadScheduleResponse, error)
	// ListSchedules returns all actives schedules.
	ListSchedules(context.Context, *ListSchedulesRequest) (*ListSchedulesResponse, error)
	// WriteSchedule creates or update a schedule for a task.
	// If one schedule already exists for the task, it will be overwritten.
	WriteSchedule(context.Context, *WriteScheduleRequest) (*WriteScheduleResponse, error)
	DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleResponse, error)
	mustEmbedUnimplementedSchedulerServiceServer()
}

// UnimplementedSchedulerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServiceServer struct {
}

func (UnimplementedSchedulerServiceServer) ReadSchedule(context.Context, *ReadScheduleRequest) (*ReadScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSchedule not implemented")
}
func (UnimplementedSchedulerServiceServer) ListSchedules(context.Context, *ListSchedulesRequest) (*ListSchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchedules not implemented")
}
func (UnimplementedSchedulerServiceServer) WriteSchedule(context.Context, *WriteScheduleRequest) (*WriteScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteSchedule not implemented")
}
func (UnimplementedSchedulerServiceServer) DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchedule not implemented")
}
func (UnimplementedSchedulerServiceServer) mustEmbedUnimplementedSchedulerServiceServer() {}

// UnsafeSchedulerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServiceServer will
// result in compilation errors.
type UnsafeSchedulerServiceServer interface {
	mustEmbedUnimplementedSchedulerServiceServer()
}

func RegisterSchedulerServiceServer(s grpc.ServiceRegistrar, srv SchedulerServiceServer) {
	s.RegisterService(&SchedulerService_ServiceDesc, srv)
}

func _SchedulerService_ReadSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ReadSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchedulerService_ReadSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ReadSchedule(ctx, req.(*ReadScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ListSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ListSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchedulerService_ListSchedules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ListSchedules(ctx, req.(*ListSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_WriteSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).WriteSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchedulerService_WriteSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).WriteSchedule(ctx, req.(*WriteScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_DeleteSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).DeleteSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchedulerService_DeleteSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).DeleteSchedule(ctx, req.(*DeleteScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchedulerService_ServiceDesc is the grpc.ServiceDesc for SchedulerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchedulerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodeum.micro.v1.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadSchedule",
			Handler:    _SchedulerService_ReadSchedule_Handler,
		},
		{
			MethodName: "ListSchedules",
			Handler:    _SchedulerService_ListSchedules_Handler,
		},
		{
			MethodName: "WriteSchedule",
			Handler:    _SchedulerService_WriteSchedule_Handler,
		},
		{
			MethodName: "DeleteSchedule",
			Handler:    _SchedulerService_DeleteSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeum/micro/v1/scheduler.proto",
}
