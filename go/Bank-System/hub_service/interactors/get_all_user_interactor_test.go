package interactors

import (
	"hub_service/repositories"
	"testing"
)

func TestFailToGetAllUser(t *testing.T) {
	interactor := NewGetAllUserInteractor(&repositories.MockUserRepository{ShouldFailGetAll: true})

	data, err := interactor.GetAll()

	if len(data) > 0 {
		t.Errorf("Test failed. Shouldn't return data.")
	}

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}

func TestGetAllUser(t *testing.T) {
	interactor := &GetAllUserInteractor{
		repository: &repositories.MockUserRepository{},
	}

	data, err := interactor.GetAll()

	if err != nil {
		t.Errorf("Test failed. Shouldn't return error.")
	}

	if len(data) <= 0 {
		t.Errorf("Test failed. Should return data.")
	}
}
