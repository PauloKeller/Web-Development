package handlers

import (
	candyland_grpc "candy_land/grpc"

	"candy_land/users_service/src/usecases"
	"context"
	"log"
)

// Handler interface layer to handle the communication between repositories and use cases
type Handler struct {
	CreateUserUsecase usecases.CreateUser
}

// Handlers interface
type Handlers interface {
	CreateUser(ctx context.Context, in *candyland_grpc.CreateUserRequest) (*candyland_grpc.CreateUserReply, error)
}

// CreateUser from a valid grpc request
func (handler *Handler) CreateUser(ctx context.Context, in *candyland_grpc.CreateUserRequest) (*candyland_grpc.CreateUserReply, error) {
	err := handler.CreateUserUsecase.Create(in.FirstName, in.LastName, in.Username, in.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := &candyland_grpc.CreateUserReply{
		WasCreated: true,
	}
	return res, nil
}
