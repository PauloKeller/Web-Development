package handler

import (
	"context"
	"log"

	"hub_service/dtos"
	"hub_service/interactors"
	hub "hub_service/proto/out"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	CreateUserInteractor  interactors.CreateUserInteractorInterface
	GetAllUserInteractor  interactors.GetAllUserInteractorInterface
	GetUserByIDInteractor interactors.GetUserByIDInteractorInterface
}

type HandlerInterface interface {
	hub.HubServiceServer
	CreateUser(context.Context, *hub.CreateUserRequest) (*hub.CreateUserReply, error)
	GetAllUsers(*empty.Empty, hub.HubService_GetAllUsersServer) error
	GetByIdUser(context.Context, *hub.RequestGetUserById) (*hub.User, error)
	CreateAccount(context.Context, *hub.CreateAccountRequest) (*hub.CreateAccountReply, error)
	GetAllAccounts(*empty.Empty, hub.HubService_GetAllAccountsServer) error
	GetAccountById(context.Context, *hub.RequestGetAccountById) (*hub.Account, error)
}

func (handler *Handler) CreateUser(context context.Context, in *hub.CreateUserRequest) (*hub.CreateUserReply, error) {
	model := &dtos.CreateUserDto{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
	}

	err := handler.CreateUserInteractor.Create(model)
	if err != nil {
		log.Println(err.Error())
		grpcError := status.Error(codes.Code(err.Code.ConvertToGrpc()), err.Error())
		return nil, grpcError
	}

	response := &hub.CreateUserReply{
		WasCreated: true,
	}

	return response, nil
}

func (handler *Handler) GetAllUsers(in *empty.Empty, stream hub.HubService_GetAllUsersServer) error {
	data, err := handler.GetAllUserInteractor.GetAll()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, value := range data {
		if err := stream.Send(&hub.User{
			Id:        value.ID.String(),
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Username:  value.Username,
			Email:     value.Email,
			Password:  value.Password,
			CreatedAt: timestamppb.New(value.CreatedAt),
			UpdatedAt: timestamppb.New(value.UpdatedAt),
			DeletedAt: timestamppb.New(value.DeletedAt),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (handler *Handler) GetByIdUser(ctx context.Context, in *hub.RequestGetUserById) (*hub.User, error) {
	data, err := handler.GetUserByIDInteractor.GetByID(in.GetId())

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	response := &hub.User{
		Id:        data.ID.String(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		DeletedAt: timestamppb.New(data.DeletedAt),
	}

	return response, nil
}

func (handler *Handler) CreateAccount(context.Context, *hub.CreateAccountRequest) (*hub.CreateAccountReply, error) {
	return nil, nil
}

func (handler *Handler) GetAllAccounts(*empty.Empty, hub.HubService_GetAllAccountsServer) error {
	return nil
}

func (handler *Handler) GetAccountById(context.Context, *hub.RequestGetAccountById) (*hub.Account, error) {
	return nil, nil
}

func mustEmbedUnimplementedHubServiceServer() {

}
