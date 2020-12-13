package handlers

import (
	"candy_land/carts_service/src/usecases"
	candyland_grpc "candy_land/grpc"
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func AddProductToCart(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := usecases.NewMockAddProductToCart(mockCtrl)

	hand := &Handler{
		AddProductToCartUsecase: mockUsecase,
	}

	mockUsecase.EXPECT().Add("123", "1231", 4).Return(nil).Times(1)

	ctx := context.Background()
	res, err := hand.AddProductToCart(ctx, &candyland_grpc.AddProductToCartRequest{
		CartId:    "123",
		ProductId: "1231",
		Quantity:  4,
	})

	if err != nil || !res.WasCreated {
		t.Errorf("Fail to add product to cart")
	}
}
