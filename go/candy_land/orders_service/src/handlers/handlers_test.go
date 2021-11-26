package handlers

import (
	candyland_grpc "candy_land/grpc"
	"candy_land/orders_service/src/usecases"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func CreateOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := usecases.NewMockCreateOrder(mockCtrl)

	h := &Handler{
		CreateOrderUsecase: mockUsecase,
	}

	mockUsecase.EXPECT().Create("1231", "123131").Return(nil).Times(1)

	ctx := context.Background()

	res, err := h.CreateOrder(ctx, &candyland_grpc.CreateOrderRequest{
		UserId:    "1231",
		ProductId: "123131",
	})

	if err != nil || !res.WasCreated {
		t.Errorf("Fail to create order")
	}
}
