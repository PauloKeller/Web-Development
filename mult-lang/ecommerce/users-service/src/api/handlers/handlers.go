package handlers

import (
	"context"
	"log"
	"time"

	"users-service/entities"
	users_grpc "users-service/grpc"
	"users-service/usecases"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

type Handlers struct {
	GetUserUsecase    usecases.GetUser
	CreateUserUsecase usecases.CreateUser
}

func (handler *Handlers) transformUserRPC(u *entities.User) *users_grpc.User {
	if u == nil {
		return nil
	}

	updatedAt := &google_protobuf.Timestamp{Seconds: u.UpdatedAt.Unix()}
	createdAt := &google_protobuf.Timestamp{Seconds: u.CreatedAt.Unix()}
	birthdate := &google_protobuf.Timestamp{Seconds: u.Birthdate.Unix()}
	res := &users_grpc.User{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Balance:   u.Balance,
		Birthdate: birthdate,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
	return res
}

func (handler *Handlers) transformUserData(u *users_grpc.User) *entities.User {
	updatedAt := time.Unix(u.GetUpdatedAt().GetSeconds(), 0)
	createdAt := time.Unix(u.GetCreatedAt().GetSeconds(), 0)
	birthdate := time.Unix(u.GetBirthdate().GetSeconds(), 0)
	res := &entities.User{
		ID:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Balance:   u.Balance,
		Birthdate: birthdate,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
	return res
}

func (handler *Handlers) GetUserByID(ctx context.Context, in *users_grpc.GetUserByIdRequest) (*users_grpc.User, error) {
	id := in.UserId
	u, err := handler.GetUserUsecase.GetUserByID(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if u == nil {
		return nil, err
	}

	res := handler.transformUserRPC(u)
	return res, nil
}

func (handler *Handlers) CreateUser(ctx context.Context, in *users_grpc.CreateUserRequest) (*users_grpc.CreateUserReply, error) {
	birthdate := time.Unix(in.GetBirthdate().GetSeconds(), 0)
	err := handler.CreateUserUsecase.CreateUser(in.FirstName, in.LastName, in.Username, in.Email, birthdate)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	res := &users_grpc.CreateUserReply{}
	return res, nil
}
