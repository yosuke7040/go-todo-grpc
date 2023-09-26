package handler

import (
	"context"
	"log"

	"connectrpc.com/connect"

	todov1 "github.com/yosuke7040/go-todo-grpc/gen/protos/todo/v1"
	"github.com/yosuke7040/go-todo-grpc/services"
)

type TodoHander struct {
	ser services.TodoServiceInterface
}

func NewTodoHandler(ser services.TodoServiceInterface) *TodoHander {
	return &TodoHander{ser: ser}
}

func (h *TodoHander) Create(
	ctx context.Context,
	req *connect.Request[todov1.CreateRequest],
) (*connect.Response[todov1.CreateResponse], error) {
	log.Println("Request headers: ", req.Header())
	err := h.ser.CreateTodo(ctx, req.Msg.Title)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&todov1.CreateResponse{})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}

func (h *TodoHander) Read(
	ctx context.Context,
	req *connect.Request[todov1.ReadRequest],
) (*connect.Response[todov1.ReadResponse], error) {
	log.Println("Request headers: ", req.Header())
	todo, err := h.ser.GetTodo(ctx, req.Msg.Id)
	if err != nil {
		log.Printf("Read error: %v", err)
		return nil, err
	}

	res := connect.NewResponse(&todov1.ReadResponse{
		Todo: &todov1.Todo{
			Id:     todo.Id,
			Title:  todo.Title,
			Status: todov1.Status(todo.Status),
		},
	})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}

func (h *TodoHander) Update(
	ctx context.Context,
	req *connect.Request[todov1.UpdateRequest],
) (*connect.Response[todov1.UpdateResponse], error) {
	log.Println("Request headers: ", req.Header())
	err := h.ser.UpdateTodo(ctx, req.Msg.Id, req.Msg.Title, int32(req.Msg.Status))
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&todov1.UpdateResponse{})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}

func (h *TodoHander) Delete(
	ctx context.Context,
	req *connect.Request[todov1.DeleteRequest],
) (*connect.Response[todov1.DeleteResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&todov1.DeleteResponse{})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}

func (h *TodoHander) List(
	ctx context.Context,
	req *connect.Request[todov1.ListRequest],
) (*connect.Response[todov1.ListResponse], error) {
	log.Println("Request headers: ", req.Header())
	todos, err := h.ser.GetTodoList(ctx)
	if err != nil {
		log.Printf("List error: %v", err)
		return nil, err
	}

	grpcTodos := make([]*todov1.Todo, 0)
	for _, todo := range todos {
		grpcTodos = append(grpcTodos, &todov1.Todo{
			Id:     todo.Id,
			Title:  todo.Title,
			Status: todov1.Status(todo.Status),
		})
	}

	res := connect.NewResponse(&todov1.ListResponse{Todos: grpcTodos})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}
