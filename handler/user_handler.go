package handler

import (
	"context"
	"log"

	"connectrpc.com/connect"

	userv1 "github.com/yosuke7040/go-todo-grpc/gen/user/v1"
	"github.com/yosuke7040/go-todo-grpc/services"
)

type UserHandler struct {
	ser services.UserServiceInterface
}

func NewUserHandler(ser services.UserServiceInterface) *UserHandler {
	return &UserHandler{ser: ser}
}

func (h *UserHandler) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	log.Println("Request headers: ", req.Header())
	user, err := h.ser.FindUser(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&userv1.GetUserResponse{
		User: &userv1.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

// func (h *UserHandler) Login(
// 	ctx context.Context,
// 	req *connect.Request[userv1.LoginRequest],
// ) (*connect.Response[userv1.LoginResponse], error) {
// 	log.Println("Request headers: ", req.Header())
// 	user, err := h.ser.Login(ctx, req.Msg.Email, req.Msg.Password)
// 	if err != nil {
// 		return nil, connect.NewError(connect.CodeInvalidArgument, err)
// 	}

// 	res := connect.NewResponse(&userv1.LoginResponse{
// 			Id:    user.Id,
// 			Token: user.Token,
// 			Email: user.Email,
// 		s
// 	})

// 	res.Header().Set("User-Version", "v1")
// 	return res, nil
// }
