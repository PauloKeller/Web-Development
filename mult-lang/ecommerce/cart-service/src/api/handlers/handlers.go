package handlers

import (
	"context"

	cart_grpc "cart-service/grpc"
)

type Handlers struct {
}

func (handlers *Handlers) CreateCart(ctx context.Context, in *cart_grpc.CreateCartRequest) (*cart_grpc.Cart, error) {
	out := &cart_grpc.Cart{}

	return out, nil
}
