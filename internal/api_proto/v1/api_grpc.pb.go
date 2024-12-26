// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MybackendGrpcSvc_GetUserInfoById_FullMethodName        = "/mybackend.v1.MybackendGrpcSvc/GetUserInfoById"
	MybackendGrpcSvc_GetUserInfoByUsername_FullMethodName  = "/mybackend.v1.MybackendGrpcSvc/GetUserInfoByUsername"
	MybackendGrpcSvc_ServiceCreateFromImage_FullMethodName = "/mybackend.v1.MybackendGrpcSvc/ServiceCreateFromImage"
	MybackendGrpcSvc_DeleteService_FullMethodName          = "/mybackend.v1.MybackendGrpcSvc/DeleteService"
)

// MybackendGrpcSvcClient is the client API for MybackendGrpcSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MybackendGrpcSvcClient interface {
	GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error)
	GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoByUsernameResponse, error)
	ServiceCreateFromImage(ctx context.Context, in *ServiceCreateFromImageRequest, opts ...grpc.CallOption) (*ServiceCreateFromImageResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
}

type mybackendGrpcSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewMybackendGrpcSvcClient(cc grpc.ClientConnInterface) MybackendGrpcSvcClient {
	return &mybackendGrpcSvcClient{cc}
}

func (c *mybackendGrpcSvcClient) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoByIdResponse)
	err := c.cc.Invoke(ctx, MybackendGrpcSvc_GetUserInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mybackendGrpcSvcClient) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoByUsernameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoByUsernameResponse)
	err := c.cc.Invoke(ctx, MybackendGrpcSvc_GetUserInfoByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mybackendGrpcSvcClient) ServiceCreateFromImage(ctx context.Context, in *ServiceCreateFromImageRequest, opts ...grpc.CallOption) (*ServiceCreateFromImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceCreateFromImageResponse)
	err := c.cc.Invoke(ctx, MybackendGrpcSvc_ServiceCreateFromImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mybackendGrpcSvcClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, MybackendGrpcSvc_DeleteService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MybackendGrpcSvcServer is the server API for MybackendGrpcSvc service.
// All implementations should embed UnimplementedMybackendGrpcSvcServer
// for forward compatibility.
type MybackendGrpcSvcServer interface {
	GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdResponse, error)
	GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest) (*GetUserInfoByUsernameResponse, error)
	ServiceCreateFromImage(context.Context, *ServiceCreateFromImageRequest) (*ServiceCreateFromImageResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
}

// UnimplementedMybackendGrpcSvcServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMybackendGrpcSvcServer struct{}

func (UnimplementedMybackendGrpcSvcServer) GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoById not implemented")
}
func (UnimplementedMybackendGrpcSvcServer) GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest) (*GetUserInfoByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUsername not implemented")
}
func (UnimplementedMybackendGrpcSvcServer) ServiceCreateFromImage(context.Context, *ServiceCreateFromImageRequest) (*ServiceCreateFromImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceCreateFromImage not implemented")
}
func (UnimplementedMybackendGrpcSvcServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedMybackendGrpcSvcServer) testEmbeddedByValue() {}

// UnsafeMybackendGrpcSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MybackendGrpcSvcServer will
// result in compilation errors.
type UnsafeMybackendGrpcSvcServer interface {
	mustEmbedUnimplementedMybackendGrpcSvcServer()
}

func RegisterMybackendGrpcSvcServer(s grpc.ServiceRegistrar, srv MybackendGrpcSvcServer) {
	// If the following call pancis, it indicates UnimplementedMybackendGrpcSvcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MybackendGrpcSvc_ServiceDesc, srv)
}

func _MybackendGrpcSvc_GetUserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MybackendGrpcSvcServer).GetUserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MybackendGrpcSvc_GetUserInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MybackendGrpcSvcServer).GetUserInfoById(ctx, req.(*GetUserInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MybackendGrpcSvc_GetUserInfoByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MybackendGrpcSvcServer).GetUserInfoByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MybackendGrpcSvc_GetUserInfoByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MybackendGrpcSvcServer).GetUserInfoByUsername(ctx, req.(*GetUserInfoByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MybackendGrpcSvc_ServiceCreateFromImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceCreateFromImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MybackendGrpcSvcServer).ServiceCreateFromImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MybackendGrpcSvc_ServiceCreateFromImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MybackendGrpcSvcServer).ServiceCreateFromImage(ctx, req.(*ServiceCreateFromImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MybackendGrpcSvc_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MybackendGrpcSvcServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MybackendGrpcSvc_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MybackendGrpcSvcServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MybackendGrpcSvc_ServiceDesc is the grpc.ServiceDesc for MybackendGrpcSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MybackendGrpcSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mybackend.v1.MybackendGrpcSvc",
	HandlerType: (*MybackendGrpcSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfoById",
			Handler:    _MybackendGrpcSvc_GetUserInfoById_Handler,
		},
		{
			MethodName: "GetUserInfoByUsername",
			Handler:    _MybackendGrpcSvc_GetUserInfoByUsername_Handler,
		},
		{
			MethodName: "ServiceCreateFromImage",
			Handler:    _MybackendGrpcSvc_ServiceCreateFromImage_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _MybackendGrpcSvc_DeleteService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
