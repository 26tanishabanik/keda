// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: LiiklusService.proto

package liiklus

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LiiklusServiceClient is the client API for LiiklusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiiklusServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LiiklusService_SubscribeClient, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (LiiklusService_ReceiveClient, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error)
	GetEndOffsets(ctx context.Context, in *GetEndOffsetsRequest, opts ...grpc.CallOption) (*GetEndOffsetsReply, error)
}

type liiklusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiiklusServiceClient(cc grpc.ClientConnInterface) LiiklusServiceClient {
	return &liiklusServiceClient{cc}
}

func (c *liiklusServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error) {
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LiiklusService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &LiiklusService_ServiceDesc.Streams[0], "/com.github.bsideup.liiklus.LiiklusService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &liiklusServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LiiklusService_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type liiklusServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *liiklusServiceSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *liiklusServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (LiiklusService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &LiiklusService_ServiceDesc.Streams[1], "/com.github.bsideup.liiklus.LiiklusService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &liiklusServiceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LiiklusService_ReceiveClient interface {
	Recv() (*ReceiveReply, error)
	grpc.ClientStream
}

type liiklusServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *liiklusServiceReceiveClient) Recv() (*ReceiveReply, error) {
	m := new(ReceiveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *liiklusServiceClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error) {
	out := new(GetOffsetsReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/GetOffsets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) GetEndOffsets(ctx context.Context, in *GetEndOffsetsRequest, opts ...grpc.CallOption) (*GetEndOffsetsReply, error) {
	out := new(GetEndOffsetsReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/GetEndOffsets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiiklusServiceServer is the server API for LiiklusService service.
// All implementations must embed UnimplementedLiiklusServiceServer
// for forward compatibility
type LiiklusServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishReply, error)
	Subscribe(*SubscribeRequest, LiiklusService_SubscribeServer) error
	Receive(*ReceiveRequest, LiiklusService_ReceiveServer) error
	Ack(context.Context, *AckRequest) (*empty.Empty, error)
	GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error)
	GetEndOffsets(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error)
	mustEmbedUnimplementedLiiklusServiceServer()
}

// UnimplementedLiiklusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiiklusServiceServer struct {
}

func (UnimplementedLiiklusServiceServer) Publish(context.Context, *PublishRequest) (*PublishReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedLiiklusServiceServer) Subscribe(*SubscribeRequest, LiiklusService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedLiiklusServiceServer) Receive(*ReceiveRequest, LiiklusService_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedLiiklusServiceServer) Ack(context.Context, *AckRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedLiiklusServiceServer) GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffsets not implemented")
}
func (UnimplementedLiiklusServiceServer) GetEndOffsets(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndOffsets not implemented")
}
func (UnimplementedLiiklusServiceServer) mustEmbedUnimplementedLiiklusServiceServer() {}

// UnsafeLiiklusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiiklusServiceServer will
// result in compilation errors.
type UnsafeLiiklusServiceServer interface {
	mustEmbedUnimplementedLiiklusServiceServer()
}

func RegisterLiiklusServiceServer(s grpc.ServiceRegistrar, srv LiiklusServiceServer) {
	s.RegisterService(&LiiklusService_ServiceDesc, srv)
}

func _LiiklusService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiiklusServiceServer).Subscribe(m, &liiklusServiceSubscribeServer{stream})
}

type LiiklusService_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type liiklusServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *liiklusServiceSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

func _LiiklusService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiiklusServiceServer).Receive(m, &liiklusServiceReceiveServer{stream})
}

type LiiklusService_ReceiveServer interface {
	Send(*ReceiveReply) error
	grpc.ServerStream
}

type liiklusServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *liiklusServiceReceiveServer) Send(m *ReceiveReply) error {
	return x.ServerStream.SendMsg(m)
}

func _LiiklusService_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_GetOffsets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).GetOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/GetOffsets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).GetOffsets(ctx, req.(*GetOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_GetEndOffsets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).GetEndOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/GetEndOffsets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).GetEndOffsets(ctx, req.(*GetEndOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LiiklusService_ServiceDesc is the grpc.ServiceDesc for LiiklusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiiklusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.bsideup.liiklus.LiiklusService",
	HandlerType: (*LiiklusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _LiiklusService_Publish_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _LiiklusService_Ack_Handler,
		},
		{
			MethodName: "GetOffsets",
			Handler:    _LiiklusService_GetOffsets_Handler,
		},
		{
			MethodName: "GetEndOffsets",
			Handler:    _LiiklusService_GetEndOffsets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _LiiklusService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Receive",
			Handler:       _LiiklusService_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "LiiklusService.proto",
}
