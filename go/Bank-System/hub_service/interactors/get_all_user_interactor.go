package interactors

import (
	"hub_service/entities"
	"hub_service/repositories"
)

type GetAllUserInteractorInterface interface {
	GetAll() ([]entities.UserEntity, error)
}

type GetAllUserInteractor struct {
	repository repositories.UserRepositoryInterface
}

func NewGetAllUserInteractor(repo repositories.UserRepositoryInterface) *GetAllUserInteractor {
	return &GetAllUserInteractor{
		repo,
	}
}

func (interactor *GetAllUserInteractor) GetAll() ([]entities.UserEntity, error) {
	data, err := interactor.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}
