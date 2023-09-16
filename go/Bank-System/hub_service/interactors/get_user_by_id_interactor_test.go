package interactors

import (
	"hub_service/repositories"
	"testing"
)

func TestFailToGetUserByID(t *testing.T) {
	interactor := NewGetUserByIDInteractor(&repositories.MockUserRepository{})

	data, err := interactor.GetByID("")

	if data != nil {
		t.Errorf("Test failed. Shouldn't return data.")
	}

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}

func TestGetUserByID(t *testing.T) {
	interactor := &GetUserByIDInteractor{
		repository: &repositories.MockUserRepository{
			ShouldFailGetAll: true,
		},
	}

	data, err := interactor.GetByID("1111")

	if err != nil {
		t.Errorf("Test failed. Shouldn't return error.")
	}

	if data == nil {
		t.Errorf("Test failed. Shouldn't return data.")
	}
}
