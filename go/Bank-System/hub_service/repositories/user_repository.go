package repositories

import (
	"hub_service/entities"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Insert(*entities.UserEntity) error
	GetByID(string) (*entities.UserEntity, error)
	GetAll() ([]entities.UserEntity, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) GetAll() ([]entities.UserEntity, error) {
	var data []entities.UserEntity

	result := repository.db.Debug().Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (repository *UserRepository) GetByID(ID string) (*entities.UserEntity, error) {
	data := &entities.UserEntity{}

	result := repository.db.Debug().First(&data, "id = ?", ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (repository *UserRepository) Insert(entity *entities.UserEntity) error {
	err := repository.db.Debug().Create(&entity).Error

	if err != nil {
		return err
	}

	return nil
}
