// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/proto/examples/helloworld.proto

package helloworld_go_proto

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// A unary RPC.
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// A client-side streaming RPC.
	SayHelloManyRequests(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloManyRequestsClient, error)
	// A server-side streaming RPC.
	SayHelloManyReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloManyRepliesClient, error)
	// A bidirectional streaming RPC.
	SayHelloConversation(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloConversationClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/s2a.examples.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloManyRequests(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloManyRequestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/s2a.examples.Greeter/SayHelloManyRequests", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloManyRequestsClient{stream}
	return x, nil
}

type Greeter_SayHelloManyRequestsClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloManyRequestsClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloManyRequestsClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloManyRequestsClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloManyReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloManyRepliesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/s2a.examples.Greeter/SayHelloManyReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloManyRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloManyRepliesClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloManyRepliesClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloManyRepliesClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloConversation(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloConversationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/s2a.examples.Greeter/SayHelloConversation", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloConversationClient{stream}
	return x, nil
}

type Greeter_SayHelloConversationClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloConversationClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloConversationClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloConversationClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// A unary RPC.
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// A client-side streaming RPC.
	SayHelloManyRequests(Greeter_SayHelloManyRequestsServer) error
	// A server-side streaming RPC.
	SayHelloManyReplies(*HelloRequest, Greeter_SayHelloManyRepliesServer) error
	// A bidirectional streaming RPC.
	SayHelloConversation(Greeter_SayHelloConversationServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloManyRequests(Greeter_SayHelloManyRequestsServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloManyRequests not implemented")
}
func (UnimplementedGreeterServer) SayHelloManyReplies(*HelloRequest, Greeter_SayHelloManyRepliesServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloManyReplies not implemented")
}
func (UnimplementedGreeterServer) SayHelloConversation(Greeter_SayHelloConversationServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloConversation not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s2a.examples.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloManyRequests_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloManyRequests(&greeterSayHelloManyRequestsServer{stream})
}

type Greeter_SayHelloManyRequestsServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloManyRequestsServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloManyRequestsServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloManyRequestsServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloManyReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloManyReplies(m, &greeterSayHelloManyRepliesServer{stream})
}

type Greeter_SayHelloManyRepliesServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloManyRepliesServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloManyRepliesServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayHelloConversation_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloConversation(&greeterSayHelloConversationServer{stream})
}

type Greeter_SayHelloConversationServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloConversationServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloConversationServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloConversationServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "s2a.examples.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloManyRequests",
			Handler:       _Greeter_SayHelloManyRequests_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloManyReplies",
			Handler:       _Greeter_SayHelloManyReplies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloConversation",
			Handler:       _Greeter_SayHelloConversation_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/examples/helloworld.proto",
}
