package usecases

import (
	"testing"

	"candy_land/users_service/src/repositories"

	gomock "github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := repositories.NewMockUsersRepository(mockCtrl)

	uc := &CreateUserProvider{
		Repository: mockRepository,
	}

	mockRepository.EXPECT().InsertUser("paulo", "keller", "paulokeller", "paulo@test.com").Return(nil).Times(1)

	err := uc.Create("paulo", "keller", "paulokeller", "paulo@test.com")

	if err != nil {
		t.Errorf("Fail to create user")
	}
}
