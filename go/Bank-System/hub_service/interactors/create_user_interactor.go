package interactors

import (
	"hub_service/dtos"
	"hub_service/entities"
	"hub_service/repositories"
	"hub_service/utils"
)

type CreateUserInteractorInterface interface {
	Create(model *dtos.CreateUserDto) *utils.Error
}

type CreateUserInteractor struct {
	repository repositories.UserRepositoryInterface
}

func NewCreateUserInteractor(repo repositories.UserRepositoryInterface) *CreateUserInteractor {
	return &CreateUserInteractor{
		repo,
	}
}

func (interactor *CreateUserInteractor) Create(model *dtos.CreateUserDto) *utils.Error {
	var serviceError *utils.Error

	isValid, msg := model.IsValid()

	if !isValid {
		serviceError = &utils.Error{
			Code:    utils.InvalidDataErrorCode,
			Err:     nil,
			Message: msg,
		}
		return serviceError
	}

	entity := &entities.UserEntity{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Username:  model.Username,
		Email:     model.Email,
		Password:  model.Password,
	}

	err := interactor.repository.Insert(entity)

	if err != nil {
		serviceError = &utils.Error{
			Code:    utils.UnknownErrorCode,
			Err:     err,
			Message: "Cannot insert into database.",
		}
		return serviceError
	}

	return serviceError
}
