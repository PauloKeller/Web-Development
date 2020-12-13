package usecases

import (
	"candy_land/carts_service/src/repositories"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestAddProductToCart(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := repositories.NewMockCartsRepository(mockCtrl)

	uc := &AddProductToCartProvider{
		Repository: mockRepository,
	}

	mockRepository.EXPECT().InsertProductIntoCart("123", "1231", 4).Return(nil).Times(1)

	err := uc.Add("123", "1231", 4)

	if err != nil {
		t.Errorf("Fail to add product to cart")
	}
}
