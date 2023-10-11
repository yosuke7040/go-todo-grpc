// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chat/v1/chat.proto

package chatv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/yosuke7040/go-todo-grpc/gen/chat/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "chat.v1.ChatService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChatServiceGetMessageStreamProcedure is the fully-qualified name of the ChatService's
	// GetMessageStream RPC.
	ChatServiceGetMessageStreamProcedure = "/chat.v1.ChatService/GetMessageStream"
	// ChatServiceCreateMessageProcedure is the fully-qualified name of the ChatService's CreateMessage
	// RPC.
	ChatServiceCreateMessageProcedure = "/chat.v1.ChatService/CreateMessage"
	// ChatServiceChatMessageStreamProcedure is the fully-qualified name of the ChatService's
	// ChatMessageStream RPC.
	ChatServiceChatMessageStreamProcedure = "/chat.v1.ChatService/ChatMessageStream"
)

// ChatServiceClient is a client for the chat.v1.ChatService service.
type ChatServiceClient interface {
	GetMessageStream(context.Context, *connect.Request[v1.GetMessageStreamRequest]) (*connect.ServerStreamForClient[v1.GetMessageStreamResponse], error)
	// rpc GetMessageStream (google.protobuf.Empty) returns (stream GetMessageStreamResponse) {};
	CreateMessage(context.Context, *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error)
	ChatMessageStream(context.Context) *connect.BidiStreamForClient[v1.ChatMessageStreamRequest, v1.ChatMessageStreamResponse]
}

// NewChatServiceClient constructs a client for the chat.v1.ChatService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chatServiceClient{
		getMessageStream: connect.NewClient[v1.GetMessageStreamRequest, v1.GetMessageStreamResponse](
			httpClient,
			baseURL+ChatServiceGetMessageStreamProcedure,
			opts...,
		),
		createMessage: connect.NewClient[v1.CreateMessageRequest, v1.CreateMessageResponse](
			httpClient,
			baseURL+ChatServiceCreateMessageProcedure,
			opts...,
		),
		chatMessageStream: connect.NewClient[v1.ChatMessageStreamRequest, v1.ChatMessageStreamResponse](
			httpClient,
			baseURL+ChatServiceChatMessageStreamProcedure,
			opts...,
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	getMessageStream  *connect.Client[v1.GetMessageStreamRequest, v1.GetMessageStreamResponse]
	createMessage     *connect.Client[v1.CreateMessageRequest, v1.CreateMessageResponse]
	chatMessageStream *connect.Client[v1.ChatMessageStreamRequest, v1.ChatMessageStreamResponse]
}

// GetMessageStream calls chat.v1.ChatService.GetMessageStream.
func (c *chatServiceClient) GetMessageStream(ctx context.Context, req *connect.Request[v1.GetMessageStreamRequest]) (*connect.ServerStreamForClient[v1.GetMessageStreamResponse], error) {
	return c.getMessageStream.CallServerStream(ctx, req)
}

// CreateMessage calls chat.v1.ChatService.CreateMessage.
func (c *chatServiceClient) CreateMessage(ctx context.Context, req *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error) {
	return c.createMessage.CallUnary(ctx, req)
}

// ChatMessageStream calls chat.v1.ChatService.ChatMessageStream.
func (c *chatServiceClient) ChatMessageStream(ctx context.Context) *connect.BidiStreamForClient[v1.ChatMessageStreamRequest, v1.ChatMessageStreamResponse] {
	return c.chatMessageStream.CallBidiStream(ctx)
}

// ChatServiceHandler is an implementation of the chat.v1.ChatService service.
type ChatServiceHandler interface {
	GetMessageStream(context.Context, *connect.Request[v1.GetMessageStreamRequest], *connect.ServerStream[v1.GetMessageStreamResponse]) error
	// rpc GetMessageStream (google.protobuf.Empty) returns (stream GetMessageStreamResponse) {};
	CreateMessage(context.Context, *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error)
	ChatMessageStream(context.Context, *connect.BidiStream[v1.ChatMessageStreamRequest, v1.ChatMessageStreamResponse]) error
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	chatServiceGetMessageStreamHandler := connect.NewServerStreamHandler(
		ChatServiceGetMessageStreamProcedure,
		svc.GetMessageStream,
		opts...,
	)
	chatServiceCreateMessageHandler := connect.NewUnaryHandler(
		ChatServiceCreateMessageProcedure,
		svc.CreateMessage,
		opts...,
	)
	chatServiceChatMessageStreamHandler := connect.NewBidiStreamHandler(
		ChatServiceChatMessageStreamProcedure,
		svc.ChatMessageStream,
		opts...,
	)
	return "/chat.v1.ChatService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChatServiceGetMessageStreamProcedure:
			chatServiceGetMessageStreamHandler.ServeHTTP(w, r)
		case ChatServiceCreateMessageProcedure:
			chatServiceCreateMessageHandler.ServeHTTP(w, r)
		case ChatServiceChatMessageStreamProcedure:
			chatServiceChatMessageStreamHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) GetMessageStream(context.Context, *connect.Request[v1.GetMessageStreamRequest], *connect.ServerStream[v1.GetMessageStreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.GetMessageStream is not implemented"))
}

func (UnimplementedChatServiceHandler) CreateMessage(context.Context, *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.CreateMessage is not implemented"))
}

func (UnimplementedChatServiceHandler) ChatMessageStream(context.Context, *connect.BidiStream[v1.ChatMessageStreamRequest, v1.ChatMessageStreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.ChatMessageStream is not implemented"))
}
