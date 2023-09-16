package interactors

import (
	"fmt"
	"hub_service/dtos"
	"hub_service/utils"
)

type MockCreateUserInteractor struct{}

func (interactor MockCreateUserInteractor) Create(model *dtos.CreateUserDto) *utils.Error {
	if model.Username == "fail" {
		return &utils.Error{
			Code:    utils.InvalidDataErrorCode,
			Message: utils.InvalidFirstNameMessage.Value(),
			Err:     fmt.Errorf("some error"),
		}
	}

	if model != nil {
		return nil
	}

	return &utils.Error{
		Code:    utils.UnknownErrorCode,
		Message: utils.UnknownMessage.Value(),
		Err:     fmt.Errorf("some error"),
	}
}
