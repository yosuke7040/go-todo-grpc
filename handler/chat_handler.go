package handler

import (
	"context"
	"errors"
	"io"
	"log"

	"connectrpc.com/connect"
	chatv1 "github.com/yosuke7040/go-todo-grpc/gen/chat/v1"
)

type ChatHandler struct {
	// ser      services.ChatServiceInterface
	requests []*string
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		// ser: s,
	}
}

func (h *ChatHandler) GetMessageStream(
	ctx context.Context,
	req *connect.Request[chatv1.GetMessageStreamRequest],
	stream *connect.ServerStream[chatv1.GetMessageStreamResponse],
) error {
	log.Println("Request headers: ", req.Header())

	for _, r := range h.requests {
		if err := stream.Send(&chatv1.GetMessageStreamResponse{Message: *r}); err != nil {
			return err
		}
	}

	previousCount := len(h.requests)

	for {
		currentCount := len(h.requests)
		if currentCount > previousCount {
			r := h.requests[currentCount-1]
			log.Printf("Sent: %v", *r)
			if err := stream.Send(&chatv1.GetMessageStreamResponse{Message: *r}); err != nil {
				return err
			}
		}
		previousCount = currentCount
	}
	// return nil
}

func (h *ChatHandler) CreateMessage(
	ctx context.Context,
	req *connect.Request[chatv1.CreateMessageRequest],
) (*connect.Response[chatv1.CreateMessageResponse], error) {
	// log.Println("Request headers: ", req.Header())
	// log.Println("Request body: ", req.Body())
	h.requests = append(h.requests, &req.Msg.Message)
	return connect.NewResponse(&chatv1.CreateMessageResponse{
		Result: req.Msg.Message + "を送信しました",
	}), nil
}

func (h *ChatHandler) ChatMessageStream(
	ctx context.Context,
	stream *connect.BidiStream[chatv1.ChatMessageStreamRequest, chatv1.ChatMessageStreamResponse],
) error {
	log.Printf("ChatStart\n")
	for _, r := range h.requests {
		if err := stream.Send(&chatv1.ChatMessageStreamResponse{Message: *r}); err != nil {
			return err
		}
	}

	previousCount := len(h.requests)

	go func() {
		for {
			if err := ctx.Err(); err != nil {
				return
			}
			// Receiveはリクエストを受け取るまでブロックする
			reqest, err := stream.Receive()
			h.requests = append(h.requests, &reqest.Message)
			log.Printf("request: %v", reqest)
			if err != nil && errors.Is(err, io.EOF) {
				return
			} else if err != nil {
				log.Printf("receive request: %s", err)
				return
			}
		}
	}()

	for {
		currentCount := len(h.requests)
		if currentCount > previousCount {
			r := h.requests[currentCount-1]
			log.Printf("Sent: %v", *r)
			if err := stream.Send(&chatv1.ChatMessageStreamResponse{Message: *r}); err != nil {
				return err
			}
		}
		previousCount = currentCount
	}
}
