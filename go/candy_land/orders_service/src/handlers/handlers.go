package handlers

import (
	candyland_grpc "candy_land/grpc"
	"candy_land/orders_service/src/usecases"
	"context"
)

// Handler interface layer to handle the communication between repositories and use cases
type Handler struct {
	CreateOrderUsecase usecases.CreateOrder
}

// Handlers interface
type Handlers interface {
	CreateOrder(ctx context.Context, in *candyland_grpc.CreateOrderRequest) (*candyland_grpc.CreateOrderReply, error)
}

// CreateOrder from a valid grpc request
func (handler *Handler) CreateOrder(ctx context.Context, in *candyland_grpc.CreateOrderRequest) (*candyland_grpc.CreateOrderReply, error) {
	res := &candyland_grpc.CreateOrderReply{
		WasCreated: true,
	}

	return res, nil
}
