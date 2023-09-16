package repositories

import (
	"hub_service/entities"

	"gorm.io/gorm"
)

type AccountRepositoryInterface interface {
	Insert(*entities.AccountEntity) error
	// GetByID(string) (*entities.AccountEntity, error)
	// GetAll() ([]entities.AccountEntity, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (repository *AccountRepository) Insert(entity *entities.AccountEntity) error {
	err := repository.db.Debug().Create(&entity).Error

	if err != nil {
		return err
	}

	return nil
}
