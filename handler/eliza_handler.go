package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"connectrpc.com/connect"
	elizav1 "github.com/yosuke7040/go-todo-grpc/gen/eliza/v1"
	"github.com/yosuke7040/go-todo-grpc/services"
)

type ElizaHandler struct {
	ser         services.ElizaServiceInterface
	streamDelay time.Duration
}

func NewElizaHandler(s services.ElizaServiceInterface, streamDelay time.Duration) *ElizaHandler {
	return &ElizaHandler{
		ser:         s,
		streamDelay: streamDelay,
	}
}

func (h *ElizaHandler) Say(
	_ context.Context,
	req *connect.Request[elizav1.SayRequest],
) (*connect.Response[elizav1.SayResponse], error) {
	reply, _ := h.ser.Reply(req.Msg.Sentence)
	return connect.NewResponse(&elizav1.SayResponse{
		Sentence: reply,
	}), nil
}

func (h *ElizaHandler) Converse(
	ctx context.Context,
	stream *connect.BidiStream[elizav1.ConverseRequest, elizav1.ConverseResponse],
) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		request, err := stream.Receive()
		if err != nil && errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("receive request: %w", err)
		}
		reply, endSession := h.ser.Reply(request.Sentence)
		if err := stream.Send(&elizav1.ConverseResponse{Sentence: reply}); err != nil {
			return fmt.Errorf("send response: %w", err)
		}
		if endSession {
			return nil
		}
	}
}

func (h *ElizaHandler) Introduce(
	ctx context.Context,
	req *connect.Request[elizav1.IntroduceRequest],
	stream *connect.ServerStream[elizav1.IntroduceResponse],
) error {
	name := req.Msg.Name
	if name == "" {
		name = "Anonymous User"
	}
	intros := h.ser.GetIntroResponses(name)
	var ticker *time.Ticker
	if h.streamDelay > 0 {
		ticker = time.NewTicker(h.streamDelay)
		defer ticker.Stop()
	}
	for _, resp := range intros {
		if ticker != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
			}
		}
		if err := stream.Send(&elizav1.IntroduceResponse{Sentence: resp}); err != nil {
			return err
		}
	}
	return nil
}
