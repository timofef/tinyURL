// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: server.proto

package server_proto

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

// TinyUrlServiceClient is the client API for TinyUrlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TinyUrlServiceClient interface {
	Add(ctx context.Context, in *FullUrl, opts ...grpc.CallOption) (*TinyUrl, error)
	Get(ctx context.Context, in *TinyUrl, opts ...grpc.CallOption) (*FullUrl, error)
}

type tinyUrlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTinyUrlServiceClient(cc grpc.ClientConnInterface) TinyUrlServiceClient {
	return &tinyUrlServiceClient{cc}
}

func (c *tinyUrlServiceClient) Add(ctx context.Context, in *FullUrl, opts ...grpc.CallOption) (*TinyUrl, error) {
	out := new(TinyUrl)
	err := c.cc.Invoke(ctx, "/TinyUrlService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyUrlServiceClient) Get(ctx context.Context, in *TinyUrl, opts ...grpc.CallOption) (*FullUrl, error) {
	out := new(FullUrl)
	err := c.cc.Invoke(ctx, "/TinyUrlService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TinyUrlServiceServer is the server API for TinyUrlService service.
// All implementations must embed UnimplementedTinyUrlServiceServer
// for forward compatibility
type TinyUrlServiceServer interface {
	Add(context.Context, *FullUrl) (*TinyUrl, error)
	Get(context.Context, *TinyUrl) (*FullUrl, error)
	mustEmbedUnimplementedTinyUrlServiceServer()
}

// UnimplementedTinyUrlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTinyUrlServiceServer struct {
}

func (UnimplementedTinyUrlServiceServer) Add(context.Context, *FullUrl) (*TinyUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTinyUrlServiceServer) Get(context.Context, *TinyUrl) (*FullUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTinyUrlServiceServer) mustEmbedUnimplementedTinyUrlServiceServer() {}

// UnsafeTinyUrlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TinyUrlServiceServer will
// result in compilation errors.
type UnsafeTinyUrlServiceServer interface {
	mustEmbedUnimplementedTinyUrlServiceServer()
}

func RegisterTinyUrlServiceServer(s grpc.ServiceRegistrar, srv TinyUrlServiceServer) {
	s.RegisterService(&TinyUrlService_ServiceDesc, srv)
}

func _TinyUrlService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyUrlServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TinyUrlService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyUrlServiceServer).Add(ctx, req.(*FullUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyUrlService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TinyUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyUrlServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TinyUrlService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyUrlServiceServer).Get(ctx, req.(*TinyUrl))
	}
	return interceptor(ctx, in, info, handler)
}

// TinyUrlService_ServiceDesc is the grpc.ServiceDesc for TinyUrlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TinyUrlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TinyUrlService",
	HandlerType: (*TinyUrlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TinyUrlService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TinyUrlService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}