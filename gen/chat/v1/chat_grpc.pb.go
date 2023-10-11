// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chat/v1/chat.proto

package chatv1

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
	ChatService_GetMessageStream_FullMethodName  = "/chat.v1.ChatService/GetMessageStream"
	ChatService_CreateMessage_FullMethodName     = "/chat.v1.ChatService/CreateMessage"
	ChatService_ChatMessageStream_FullMethodName = "/chat.v1.ChatService/ChatMessageStream"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	GetMessageStream(ctx context.Context, in *GetMessageStreamRequest, opts ...grpc.CallOption) (ChatService_GetMessageStreamClient, error)
	// rpc GetMessageStream (google.protobuf.Empty) returns (stream GetMessageStreamResponse) {};
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	ChatMessageStream(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatMessageStreamClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) GetMessageStream(ctx context.Context, in *GetMessageStreamRequest, opts ...grpc.CallOption) (ChatService_GetMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_GetMessageStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetMessageStreamClient interface {
	Recv() (*GetMessageStreamResponse, error)
	grpc.ClientStream
}

type chatServiceGetMessageStreamClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetMessageStreamClient) Recv() (*GetMessageStreamResponse, error) {
	m := new(GetMessageStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ChatMessageStream(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_ChatMessageStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatMessageStreamClient{stream}
	return x, nil
}

type ChatService_ChatMessageStreamClient interface {
	Send(*ChatMessageStreamRequest) error
	Recv() (*ChatMessageStreamResponse, error)
	grpc.ClientStream
}

type chatServiceChatMessageStreamClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatMessageStreamClient) Send(m *ChatMessageStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatMessageStreamClient) Recv() (*ChatMessageStreamResponse, error) {
	m := new(ChatMessageStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	GetMessageStream(*GetMessageStreamRequest, ChatService_GetMessageStreamServer) error
	// rpc GetMessageStream (google.protobuf.Empty) returns (stream GetMessageStreamResponse) {};
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	ChatMessageStream(ChatService_ChatMessageStreamServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) GetMessageStream(*GetMessageStreamRequest, ChatService_GetMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessageStream not implemented")
}
func (UnimplementedChatServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedChatServiceServer) ChatMessageStream(ChatService_ChatMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatMessageStream not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_GetMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessageStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetMessageStream(m, &chatServiceGetMessageStreamServer{stream})
}

type ChatService_GetMessageStreamServer interface {
	Send(*GetMessageStreamResponse) error
	grpc.ServerStream
}

type chatServiceGetMessageStreamServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetMessageStreamServer) Send(m *GetMessageStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ChatMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ChatMessageStream(&chatServiceChatMessageStreamServer{stream})
}

type ChatService_ChatMessageStreamServer interface {
	Send(*ChatMessageStreamResponse) error
	Recv() (*ChatMessageStreamRequest, error)
	grpc.ServerStream
}

type chatServiceChatMessageStreamServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatMessageStreamServer) Send(m *ChatMessageStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatMessageStreamServer) Recv() (*ChatMessageStreamRequest, error) {
	m := new(ChatMessageStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _ChatService_CreateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessageStream",
			Handler:       _ChatService_GetMessageStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChatMessageStream",
			Handler:       _ChatService_ChatMessageStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat/v1/chat.proto",
}
