package interactors

import (
	"hub_service/dtos"
	"hub_service/repositories"
	"testing"
)

func TestFailToCreate(t *testing.T) {
	interactor := NewCreateUserInteractor(&repositories.MockUserRepository{})
	err := interactor.Create(&dtos.CreateUserDto{})

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}

func TestCreate(t *testing.T) {
	interactor := &CreateUserInteractor{
		repository: &repositories.MockUserRepository{},
	}

	err := interactor.Create(&dtos.CreateUserDto{
		FirstName: "Paulo",
		LastName:  "Keller",
		Username:  "paulo",
		Email:     "paulo@keller.com",
		Password:  "123456",
	})

	if err != nil {
		t.Errorf("Test failed. Should create.")
	}
}

func TestFailToCreateWhenInsert(t *testing.T) {
	interactor := &CreateUserInteractor{
		repository: &repositories.MockUserRepository{},
	}

	err := interactor.Create(&dtos.CreateUserDto{
		FirstName: "fail",
		LastName:  "Keller",
		Username:  "paulo",
		Email:     "paulo@keller.com",
		Password:  "123456",
	})

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}
