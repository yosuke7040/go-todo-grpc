package handler

// import (
// 	"context"
// 	"log"

// 	"connectrpc.com/connect"
// 	chatv1 "github.com/yosuke7040/go-todo-grpc/gen/chat/v1"
// )

// type ChatHandler struct{}

// func NewChatHandler() *ChatHandler {
// 	return &ChatHandler{}
// }

// func (h *ChatHandler) GetMessages(
// 	ctx context.Context,
// 	req *connect.Request[chatv1.GetMessageStreamRequest],
// ) (*connect.ServerStream[chatv1.GetMessageStreamResponse], error) {
// 	log.Println("Request headers: ", req.Header())

// }
