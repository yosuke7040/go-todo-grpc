package services

import (
	"context"
	"fmt"
	"log"

	"github.com/yosuke7040/go-todo-grpc/models"
	"github.com/yosuke7040/go-todo-grpc/repositories"
)

type TodoServiceInterface interface {
	GetTodoList(context.Context) ([]*models.Todo, error)
	GetTodo(context.Context, int32) (*models.Todo, error)
	CreateTodo(context.Context, string) error
	UpdateTodo(context.Context, int32, string, int32) error
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

func (s *TodoService) GetTodo(ctx context.Context, id int32) (*models.Todo, error) {
	isValid, err := s.repo.IsValidTodoId(id)
	if err != nil {
		log.Printf("GetTodo error: %v", err)
		return nil, err
	}
	if !isValid {
		log.Printf("invalid id: %d", err)
		return nil, fmt.Errorf("invalid id: %d", id)
	}

	todo, err := s.repo.SelectTodo(id)
	if err != nil {
		log.Printf("GetTodo error: %v", err)
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, title string) error {
	err := s.repo.InsertTodo(title)
	if err != nil {
		log.Printf("GetTodoList error: %v", err)
		return err
	}
	return nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, id int32, title string, status int32) error {
	// TODO:　errの使い方がアンチパターンなので修正する
	// どういう構成が良いのか。。
	isValid, err := s.repo.IsValidTodoId(id)
	if err != nil {
		log.Printf("UpdateTodo error: %v", err)
		return err
	}
	if !isValid {
		log.Printf("invalid id: %d", id)
		return fmt.Errorf("invalid id: %d", id)
	}

	err = s.repo.UpdateTodo(id, title, status)
	if err != nil {
		log.Printf("UpdateTodo error: %v", err)
		return err
	}
	return nil
}
