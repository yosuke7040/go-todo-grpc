package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/yosuke7040/go-todo-grpc/gen/chat/v1/chatv1connect"
	"github.com/yosuke7040/go-todo-grpc/gen/eliza/v1/elizav1connect"
	"github.com/yosuke7040/go-todo-grpc/gen/todo/v1/todov1connect"
	"github.com/yosuke7040/go-todo-grpc/gen/user/v1/userv1connect"
	"github.com/yosuke7040/go-todo-grpc/handler"
	"github.com/yosuke7040/go-todo-grpc/internal/db"
	"github.com/yosuke7040/go-todo-grpc/repositories"
	"github.com/yosuke7040/go-todo-grpc/services"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// 環境変数の取得
	// var ok bool
	// var url, audience string
	// if audience, ok = os.LookupEnv("AUDIENCE"); !ok {
	// 	return fmt.Errorf("audience not set: %s", audience)
	// }
	// if url, ok = os.LookupEnv("AUTH0_DOMAIN"); !ok {
	// 	return fmt.Errorf("auth0 domain url not set: %s", url)
	// }
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("AUTH0_DOMAIN")
	audience := os.Getenv("AUDIENCE")
	fmt.Println("url, audience", url, audience)

	db, err := db.NewDBClients().ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}

	// JWTの有効期限
	// timeout := 1 * time.Hour

	compress1KB := connect.WithCompressMinBytes(1024)

	// TODO: dbの接続をまとめる
	tRepo := repositories.NewTodo(db)
	uRepo := repositories.NewUser(db)

	tSer := services.NewTodoService(tRepo)
	uSer := services.NewUserService(uRepo)
	// cSer := services.NewChatService()
	eSer := services.NewElizaService()

	todoHandler := handler.NewTodoHandler(tSer)
	userHandler := handler.NewUserHandler(uSer)
	chatHandler := handler.NewChatHandler()
	elizaHandler := handler.NewElizaHandler(eSer, 5*time.Second)

	// authInterceptor := connect.WithInterceptors(interceptor.NewAuthInterceptor(url, audience))
	// fmt.Println(authInterceptor)

	mux := http.NewServeMux()
	// mux.Handle(auth_v1connect.NewAuthServiceHandler(authServer))
	mux.Handle(userv1connect.NewUserServiceHandler(userHandler))
	mux.Handle(todov1connect.NewTodoServiceHandler(todoHandler))
	mux.Handle(chatv1connect.NewChatServiceHandler(chatHandler))
	// mux.Handle(todov1connect.NewTodoServiceHandler(todoHandler, authInterceptor))
	mux.Handle(elizav1connect.NewElizaServiceHandler(elizaHandler, compress1KB))

	// mux.Handle(grpchealth.NewHandler(
	// 	grpchealth.NewStaticChecker(elizav1connect.ElizaServiceName),
	// 	compress1KB,
	// ))
	// mux.Handle(grpcreflect.NewHandlerV1(
	// 	grpcreflect.NewStaticReflector(elizav1connect.ElizaServiceName),
	// 	compress1KB,
	// ))
	// mux.Handle(grpcreflect.NewHandlerV1Alpha(
	// 	grpcreflect.NewStaticReflector(elizav1connect.ElizaServiceName),
	// 	compress1KB,
	// ))

	return http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		cors.AllowAll().Handler(
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
