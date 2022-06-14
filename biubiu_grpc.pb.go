// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/helloworld/v1/biubiu.proto

package biubiupb

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

// BiuBiuClient is the client API for BiuBiu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiuBiuClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type biuBiuClient struct {
	cc grpc.ClientConnInterface
}

func NewBiuBiuClient(cc grpc.ClientConnInterface) BiuBiuClient {
	return &biuBiuClient{cc}
}

func (c *biuBiuClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.v1.BiuBiu/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiuBiuServer is the server API for BiuBiu service.
// All implementations must embed UnimplementedBiuBiuServer
// for forward compatibility
type BiuBiuServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedBiuBiuServer()
}

// UnimplementedBiuBiuServer must be embedded to have forward compatible implementations.
type UnimplementedBiuBiuServer struct {
}

func (UnimplementedBiuBiuServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBiuBiuServer) mustEmbedUnimplementedBiuBiuServer() {}

// UnsafeBiuBiuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiuBiuServer will
// result in compilation errors.
type UnsafeBiuBiuServer interface {
	mustEmbedUnimplementedBiuBiuServer()
}

func RegisterBiuBiuServer(s grpc.ServiceRegistrar, srv BiuBiuServer) {
	s.RegisterService(&BiuBiu_ServiceDesc, srv)
}

func _BiuBiu_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiuBiuServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.BiuBiu/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiuBiuServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiuBiu_ServiceDesc is the grpc.ServiceDesc for BiuBiu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiuBiu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.BiuBiu",
	HandlerType: (*BiuBiuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _BiuBiu_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/helloworld/v1/biubiu.proto",
}
