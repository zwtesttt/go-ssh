// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/sshService.proto

package proto

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

// ConServiceClient is the client API for ConService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConServiceClient interface {
	GetSshInfo(ctx context.Context, in *SendId, opts ...grpc.CallOption) (*SshInfo, error)
}

type conServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConServiceClient(cc grpc.ClientConnInterface) ConServiceClient {
	return &conServiceClient{cc}
}

func (c *conServiceClient) GetSshInfo(ctx context.Context, in *SendId, opts ...grpc.CallOption) (*SshInfo, error) {
	out := new(SshInfo)
	err := c.cc.Invoke(ctx, "/proto.ConService/GetSshInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConServiceServer is the server API for ConService service.
// All implementations must embed UnimplementedConServiceServer
// for forward compatibility
type ConServiceServer interface {
	GetSshInfo(context.Context, *SendId) (*SshInfo, error)
	mustEmbedUnimplementedConServiceServer()
}

// UnimplementedConServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConServiceServer struct {
}

func (UnimplementedConServiceServer) GetSshInfo(context.Context, *SendId) (*SshInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSshInfo not implemented")
}
func (UnimplementedConServiceServer) mustEmbedUnimplementedConServiceServer() {}

// UnsafeConServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConServiceServer will
// result in compilation errors.
type UnsafeConServiceServer interface {
	mustEmbedUnimplementedConServiceServer()
}

func RegisterConServiceServer(s grpc.ServiceRegistrar, srv ConServiceServer) {
	s.RegisterService(&ConService_ServiceDesc, srv)
}

func _ConService_GetSshInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConServiceServer).GetSshInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConService/GetSshInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConServiceServer).GetSshInfo(ctx, req.(*SendId))
	}
	return interceptor(ctx, in, info, handler)
}

// ConService_ServiceDesc is the grpc.ServiceDesc for ConService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConService",
	HandlerType: (*ConServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSshInfo",
			Handler:    _ConService_GetSshInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sshService.proto",
}

// VerifyAuthClient is the client API for VerifyAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerifyAuthClient interface {
	AuthToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type verifyAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyAuthClient(cc grpc.ClientConnInterface) VerifyAuthClient {
	return &verifyAuthClient{cc}
}

func (c *verifyAuthClient) AuthToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.VerifyAuth/authToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyAuthServer is the server API for VerifyAuth service.
// All implementations must embed UnimplementedVerifyAuthServer
// for forward compatibility
type VerifyAuthServer interface {
	AuthToken(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedVerifyAuthServer()
}

// UnimplementedVerifyAuthServer must be embedded to have forward compatible implementations.
type UnimplementedVerifyAuthServer struct {
}

func (UnimplementedVerifyAuthServer) AuthToken(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthToken not implemented")
}
func (UnimplementedVerifyAuthServer) mustEmbedUnimplementedVerifyAuthServer() {}

// UnsafeVerifyAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerifyAuthServer will
// result in compilation errors.
type UnsafeVerifyAuthServer interface {
	mustEmbedUnimplementedVerifyAuthServer()
}

func RegisterVerifyAuthServer(s grpc.ServiceRegistrar, srv VerifyAuthServer) {
	s.RegisterService(&VerifyAuth_ServiceDesc, srv)
}

func _VerifyAuth_AuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyAuthServer).AuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VerifyAuth/authToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyAuthServer).AuthToken(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// VerifyAuth_ServiceDesc is the grpc.ServiceDesc for VerifyAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerifyAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VerifyAuth",
	HandlerType: (*VerifyAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "authToken",
			Handler:    _VerifyAuth_AuthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sshService.proto",
}
