package handlers

import (
	"candy_land/addresses_service/src/usecases"
	candyland_grpc "candy_land/grpc"
	"context"
	"log"
)

// Handler interface layer to handle the communication between repositories and use cases
type Handler struct {
	CreateAddressUsecase usecases.CreateAddress
}

// Handlers interface
type Handlers interface {
	CreateAddress(ctx context.Context, in *candyland_grpc.CreateAddressRequest) (*candyland_grpc.CreateAddressReply, error)
}

// CreateAddress from a valid grpc request
func (handler *Handler) CreateAddress(ctx context.Context, in *candyland_grpc.CreateAddressRequest) (*candyland_grpc.CreateAddressReply, error) {
	err := handler.CreateAddressUsecase.Create(in.UserId, in.CountryId, in.StateId, in.CityId, in.StreetId, in.Number, in.Complement.String())

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := &candyland_grpc.CreateAddressReply{
		WasCreated: true,
	}

	return res, nil
}
