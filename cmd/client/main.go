package main

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	chatv1 "github.com/yosuke7040/go-todo-grpc/gen/chat/v1"
// )

// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func run() error {
// 	client := chatv1.NewChatServiceClient(
// 		http.DefaultClient,
// 		"http://localhost:8080",
// 	)
// 	res, err := client.CreateMessage(
// 		context.Background(),
// 		&chatv1.CreateMessageRequest{
// 			Message: &chatv1.Message{
// 				Body: "Hello",
// 			},
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println(res)
// 	return nil
// }
