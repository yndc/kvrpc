// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// KVRPCClient is the client API for KVRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVRPCClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Empty, error)
}

type kVRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewKVRPCClient(cc grpc.ClientConnInterface) KVRPCClient {
	return &kVRPCClient{cc}
}

func (c *kVRPCClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/pb.KVRPC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRPCClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/pb.KVRPC/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRPCClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.KVRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRPCClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.KVRPC/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVRPCServer is the server API for KVRPC service.
// All implementations must embed UnimplementedKVRPCServer
// for forward compatibility
type KVRPCServer interface {
	Ping(context.Context, *Empty) (*PingResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Del(context.Context, *DelRequest) (*Empty, error)
	mustEmbedUnimplementedKVRPCServer()
}

// UnimplementedKVRPCServer must be embedded to have forward compatible implementations.
type UnimplementedKVRPCServer struct {
}

func (UnimplementedKVRPCServer) Ping(context.Context, *Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKVRPCServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKVRPCServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKVRPCServer) Del(context.Context, *DelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedKVRPCServer) mustEmbedUnimplementedKVRPCServer() {}

// UnsafeKVRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVRPCServer will
// result in compilation errors.
type UnsafeKVRPCServer interface {
	mustEmbedUnimplementedKVRPCServer()
}

func RegisterKVRPCServer(s grpc.ServiceRegistrar, srv KVRPCServer) {
	s.RegisterService(&_KVRPC_serviceDesc, srv)
}

func _KVRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRPCServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRPC_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRPCServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVRPC/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRPCServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRPCServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRPC_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRPCServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KVRPC/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRPCServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.KVRPC",
	HandlerType: (*KVRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _KVRPC_Ping_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _KVRPC_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KVRPC_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _KVRPC_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}
