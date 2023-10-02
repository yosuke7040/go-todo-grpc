package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/rs/cors"
	"github.com/yosuke7040/go-todo-grpc/gen/todo/v1/todov1connect"
	"github.com/yosuke7040/go-todo-grpc/gen/user/v1/userv1connect"
	"github.com/yosuke7040/go-todo-grpc/handler"
	"github.com/yosuke7040/go-todo-grpc/interceptor"
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
	var ok bool
	var url, audience string
	if audience, ok = os.LookupEnv("AUDIENCE"); !ok {
		return fmt.Errorf("audience not set: %s", audience)
	}
	if url, ok = os.LookupEnv("AUTH0_DOMAIN"); !ok {
		return fmt.Errorf("auth0 domain url not set: %s", url)
	}

	db, err := db.NewDBClients().ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}

	// JWTの有効期限
	// timeout := 1 * time.Hour

	// TODO: dbの接続をまとめる
	tRepo := repositories.NewTodo(db)
	uRepo := repositories.NewUser(db)
	tSer := services.NewTodoService(tRepo)
	uSer := services.NewUserService(uRepo)
	todoHandler := handler.NewTodoHandler(tSer)
	userHandler := handler.NewUserHandler(uSer)

	authInterceptor := connect.WithInterceptors(interceptor.NewAuthInterceptor(url, audience))
	fmt.Println(authInterceptor)

	mux := http.NewServeMux()
	// mux.Handle(auth_v1connect.NewAuthServiceHandler(authServer))
	mux.Handle(userv1connect.NewUserServiceHandler(userHandler))
	// mux.Handle(todov1connect.NewTodoServiceHandler(todoHandler))
	mux.Handle(todov1connect.NewTodoServiceHandler(todoHandler, authInterceptor))

	return http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		cors.AllowAll().Handler(
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
