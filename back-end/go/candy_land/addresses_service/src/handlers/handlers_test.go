package handlers

import (
	"candy_land/addresses_service/src/usecases"
	candyland_grpc "candy_land/grpc"
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestCreateAddress(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := usecases.NewMockCreateAddress(mockCtrl)

	hand := &Handler{
		CreateAddressUsecase: mockUsecase,
	}

	mockUsecase.EXPECT().Create("123", "123", "1234", "12345", "12325", "23", "HOUSE")

	ctx := context.Background()

	res, err := hand.CreateAddress(ctx, &candyland_grpc.CreateAddressRequest{
		UserId:     "123",
		CountryId:  "123",
		StateId:    "1234",
		CityId:     "12345",
		StreetId:   "12325",
		Number:     "23",
		Complement: candyland_grpc.Complements_HOUSE,
	})

	if err != nil || !res.WasCreated {
		t.Errorf("Fail to create address")
	}
}
