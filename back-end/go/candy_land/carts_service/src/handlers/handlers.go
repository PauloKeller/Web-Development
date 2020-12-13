package handlers

import (
	candyland_grpc "candy_land/grpc"

	"candy_land/carts_service/src/usecases"

	"context"
)

// Handler interface layer to handle the communication between repositories and use cases
type Handler struct {
	AddProductToCartUsecase usecases.AddProductToCart
}

// Handlers interface
type Handlers interface {
	AddProductToCart(ctx context.Context, in *candyland_grpc.AddProductToCartRequest) (*candyland_grpc.AddProductToCartReply, error)
}

// AddProductToCart from a valid grpc request
func (handler *Handler) AddProductToCart(ctx context.Context, in *candyland_grpc.AddProductToCartRequest) (*candyland_grpc.AddProductToCartReply, error) {
	res := &candyland_grpc.AddProductToCartReply{
		WasCreated: true,
	}

	return res, nil
}
