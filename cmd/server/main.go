package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/yosuke7040/go-todo-grpc/gen/todo/v1/todov1connect"
	"github.com/yosuke7040/go-todo-grpc/handler"
	"github.com/yosuke7040/go-todo-grpc/internal/db"
	"github.com/yosuke7040/go-todo-grpc/repositories"
	"github.com/yosuke7040/go-todo-grpc/services"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	db, err := db.NewDBClients().ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}

	repo := repositories.NewTodo(db)
	ser := services.NewTodoService(repo)
	todoHandler := handler.NewTodoHandler(ser)

	mux := http.NewServeMux()
	mux.Handle(todov1connect.NewTodoServiceHandler(todoHandler))

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		cors.AllowAll().Handler(
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
