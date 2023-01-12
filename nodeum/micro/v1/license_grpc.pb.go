// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nodeum/micro/v1/license.proto

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

// LicenseServiceClient is the client API for LicenseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LicenseServiceClient interface {
	// ShowLicense shows details about the current license
	ShowLicense(ctx context.Context, in *ShowLicenseRequest, opts ...grpc.CallOption) (*ShowLicenseResponse, error)
	// UpdateLicense updates the current license
	UpdateLicense(ctx context.Context, in *UpdateLicenseRequest, opts ...grpc.CallOption) (*UpdateLicenseResponse, error)
}

type licenseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLicenseServiceClient(cc grpc.ClientConnInterface) LicenseServiceClient {
	return &licenseServiceClient{cc}
}

func (c *licenseServiceClient) ShowLicense(ctx context.Context, in *ShowLicenseRequest, opts ...grpc.CallOption) (*ShowLicenseResponse, error) {
	out := new(ShowLicenseResponse)
	err := c.cc.Invoke(ctx, "/nodeum.micro.v1.LicenseService/ShowLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseServiceClient) UpdateLicense(ctx context.Context, in *UpdateLicenseRequest, opts ...grpc.CallOption) (*UpdateLicenseResponse, error) {
	out := new(UpdateLicenseResponse)
	err := c.cc.Invoke(ctx, "/nodeum.micro.v1.LicenseService/UpdateLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicenseServiceServer is the server API for LicenseService service.
// All implementations must embed UnimplementedLicenseServiceServer
// for forward compatibility
type LicenseServiceServer interface {
	// ShowLicense shows details about the current license
	ShowLicense(context.Context, *ShowLicenseRequest) (*ShowLicenseResponse, error)
	// UpdateLicense updates the current license
	UpdateLicense(context.Context, *UpdateLicenseRequest) (*UpdateLicenseResponse, error)
	mustEmbedUnimplementedLicenseServiceServer()
}

// UnimplementedLicenseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLicenseServiceServer struct {
}

func (UnimplementedLicenseServiceServer) ShowLicense(context.Context, *ShowLicenseRequest) (*ShowLicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowLicense not implemented")
}
func (UnimplementedLicenseServiceServer) UpdateLicense(context.Context, *UpdateLicenseRequest) (*UpdateLicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLicense not implemented")
}
func (UnimplementedLicenseServiceServer) mustEmbedUnimplementedLicenseServiceServer() {}

// UnsafeLicenseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LicenseServiceServer will
// result in compilation errors.
type UnsafeLicenseServiceServer interface {
	mustEmbedUnimplementedLicenseServiceServer()
}

func RegisterLicenseServiceServer(s grpc.ServiceRegistrar, srv LicenseServiceServer) {
	s.RegisterService(&LicenseService_ServiceDesc, srv)
}

func _LicenseService_ShowLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).ShowLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeum.micro.v1.LicenseService/ShowLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).ShowLicense(ctx, req.(*ShowLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LicenseService_UpdateLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).UpdateLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeum.micro.v1.LicenseService/UpdateLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).UpdateLicense(ctx, req.(*UpdateLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LicenseService_ServiceDesc is the grpc.ServiceDesc for LicenseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LicenseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodeum.micro.v1.LicenseService",
	HandlerType: (*LicenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowLicense",
			Handler:    _LicenseService_ShowLicense_Handler,
		},
		{
			MethodName: "UpdateLicense",
			Handler:    _LicenseService_UpdateLicense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeum/micro/v1/license.proto",
}
