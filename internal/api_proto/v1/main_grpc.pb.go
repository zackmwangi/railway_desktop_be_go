// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: main.proto

package v1

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
	MybackendSvc_GetUserInfoById_FullMethodName       = "/svc.shell.v1.MybackendSvc/GetUserInfoById"
	MybackendSvc_GetUserInfoByUsername_FullMethodName = "/svc.shell.v1.MybackendSvc/GetUserInfoByUsername"
)

// MybackendSvcClient is the client API for MybackendSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MybackendSvcClient interface {
	GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error)
	GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoByUsernameResponse, error)
}

type mybackendSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewMybackendSvcClient(cc grpc.ClientConnInterface) MybackendSvcClient {
	return &mybackendSvcClient{cc}
}

func (c *mybackendSvcClient) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error) {
	out := new(GetUserInfoByIdResponse)
	err := c.cc.Invoke(ctx, MybackendSvc_GetUserInfoById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mybackendSvcClient) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoByUsernameResponse, error) {
	out := new(GetUserInfoByUsernameResponse)
	err := c.cc.Invoke(ctx, MybackendSvc_GetUserInfoByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MybackendSvcServer is the server API for MybackendSvc service.
// All implementations should embed UnimplementedMybackendSvcServer
// for forward compatibility
type MybackendSvcServer interface {
	GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdResponse, error)
	GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest) (*GetUserInfoByUsernameResponse, error)
}

// UnimplementedMybackendSvcServer should be embedded to have forward compatible implementations.
type UnimplementedMybackendSvcServer struct {
}

func (UnimplementedMybackendSvcServer) GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoById not implemented")
}
func (UnimplementedMybackendSvcServer) GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest) (*GetUserInfoByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUsername not implemented")
}

// UnsafeMybackendSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MybackendSvcServer will
// result in compilation errors.
type UnsafeMybackendSvcServer interface {
	mustEmbedUnimplementedMybackendSvcServer()
}

func RegisterMybackendSvcServer(s grpc.ServiceRegistrar, srv MybackendSvcServer) {
	s.RegisterService(&MybackendSvc_ServiceDesc, srv)
}

func _MybackendSvc_GetUserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MybackendSvcServer).GetUserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MybackendSvc_GetUserInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MybackendSvcServer).GetUserInfoById(ctx, req.(*GetUserInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MybackendSvc_GetUserInfoByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MybackendSvcServer).GetUserInfoByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MybackendSvc_GetUserInfoByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MybackendSvcServer).GetUserInfoByUsername(ctx, req.(*GetUserInfoByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MybackendSvc_ServiceDesc is the grpc.ServiceDesc for MybackendSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MybackendSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.shell.v1.MybackendSvc",
	HandlerType: (*MybackendSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfoById",
			Handler:    _MybackendSvc_GetUserInfoById_Handler,
		},
		{
			MethodName: "GetUserInfoByUsername",
			Handler:    _MybackendSvc_GetUserInfoByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}