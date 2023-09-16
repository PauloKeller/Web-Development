package interactors

import (
	"hub_service/entities"
	"hub_service/repositories"
)

type GetUserByIDInteractorInterface interface {
	GetByID(ID string) (*entities.UserEntity, error)
}

type GetUserByIDInteractor struct {
	repository repositories.UserRepositoryInterface
}

func NewGetUserByIDInteractor(repo repositories.UserRepositoryInterface) *GetUserByIDInteractor {
	return &GetUserByIDInteractor{
		repo,
	}
}

func (interactor *GetUserByIDInteractor) GetByID(ID string) (*entities.UserEntity, error) {
	data, err := interactor.repository.GetByID(ID)

	if err != nil {
		return nil, err
	}

	return data, nil
}
