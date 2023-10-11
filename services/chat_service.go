package services

// import (
// 	"connectrpc.com/connect"
// 	chatv1 "github.com/yosuke7040/go-todo-grpc/gen/chat/v1"
// )

// type ChatServiceInterface interface {
// }

// type ChatService struct {
// 	requests []*string
// }

// func NewChatService() *ChatService {
// 	return &ChatService{}
// }

// // func (s *ChatService) CreateMessage(ctx context.Context, r *pb.GetMessageStreamRequest) (*pb.GetMessageStreamResponse, error) {
// // 	log.Printf("Received: %v", r.GetMessageStream(ctx))
// // }

// func (s *ChatService) Converse(
// 	// ctx context.Context,
// 	stream *connect.ServerStream[chatv1.GetMessageStreamResponse],
// ) error {
// 	for _, r := range s.requests {
// 		if err := stream.Send(&chatv1.GetMessageStreamResponse{Message: *r}); err != nil {
// 			return err
// 		}
// 	}

// 	previousCount := len(s.requests)
// }
