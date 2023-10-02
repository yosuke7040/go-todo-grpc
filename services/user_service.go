package services

import (
	"context"
	"fmt"

	"github.com/yosuke7040/go-todo-grpc/models"
	"github.com/yosuke7040/go-todo-grpc/repositories"
)

type UserServiceInterface interface {
	FindUser(context.Context, int32) (*models.User, error)
}

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(r repositories.UserRepositoryInterface) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) FindUser(ctx context.Context, id int32) (*models.User, error) {
	isValid, err := s.repo.IsValidUserId(id)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, fmt.Errorf("invalid id: %d", id)
	}

	user, err := s.repo.SelectUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
