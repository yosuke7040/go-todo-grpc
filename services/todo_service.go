package services

import (
	"context"
	"log"

	"github.com/yosuke7040/go-todo-grpc/models"
	"github.com/yosuke7040/go-todo-grpc/repositories"
)

type TodoServiceInterface interface {
	GetTodoList(context.Context) ([]*models.Todo, error)
}

type TodoService struct {
	repo repositories.TodoRepoInterface
}

func NewTodoService(r repositories.TodoRepoInterface) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) GetTodoList(ctx context.Context) ([]*models.Todo, error) {
	todos, err := s.repo.SelectTodoList()
	if err != nil {
		log.Printf("GetTodoList error: %v", err)
		return nil, err
	}
	return todos, nil
}
