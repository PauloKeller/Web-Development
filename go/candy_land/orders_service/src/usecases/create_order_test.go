package usecases

import (
	"candy_land/orders_service/src/repositories"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := repositories.NewMockOrdersRepository(mockCtrl)

	uc := &CreateOrderProvider{
		Repository: mockRepository,
	}

	mockRepository.EXPECT().InsertOrder("1231", "123123").Return(nil).Times(1)

	e := uc.Create("1231", "123123")

	if e != nil {
		t.Errorf("Fail to create order")
	}
}
