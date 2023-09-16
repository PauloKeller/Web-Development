package repositories

import (
	"fmt"
	"hub_service/entities"
	"time"

	"github.com/google/uuid"
)

type MockUserRepository struct {
	ShouldFailGetAll bool
}

func (m *MockUserRepository) Insert(entity *entities.UserEntity) error {
	if entity.FirstName == "fail" {
		return fmt.Errorf("some error")
	}

	return nil
}

func (m *MockUserRepository) GetByID(id string) (*entities.UserEntity, error) {
	if len(id) > 0 {
		return &entities.UserEntity{}, nil
	}

	return nil, fmt.Errorf("some error")
}

func (m *MockUserRepository) GetAll() ([]entities.UserEntity, error) {
	if m.ShouldFailGetAll {
		return nil, fmt.Errorf("some error")
	}

	id, _ := uuid.NewUUID()

	var data = []entities.UserEntity{
		{
			FirstName: "Paulo",
			LastName:  "Keller",
			ID:        id,
			Username:  "paulokeller",
			Email:     "paulo@keller.com",
			Password:  "123456",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return data, nil
}
