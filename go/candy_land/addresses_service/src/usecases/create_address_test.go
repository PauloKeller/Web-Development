package usecases

import (
	"candy_land/addresses_service/src/repositories"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestCreateAddress(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := repositories.NewMockAddressesRepository(mockCtrl)

	uc := &CreateAddressProvider{
		Repository: mockRepository,
	}

	mockRepository.EXPECT().InsertAddress("123", "123", "1234", "12345", "12325", "23", "HOUSE")

	err := uc.Create("123", "123", "1234", "12345", "12325", "23", "HOUSE")

	if err != nil {
		t.Errorf("Fail to create address")
	}
}
