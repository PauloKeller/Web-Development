package usecases

import (
	"candy_land/users_service/src/repositories"
)

// CreateUser interface
type CreateUser interface {
	Create(firstName string, lastName string, username string, email string) error
}

// CreateUserProvider repositories
type CreateUserProvider struct {
	Repository repositories.UsersRepository
}

// Create implementation
func (uc *CreateUserProvider) Create(firstName string, lastName string, username string, email string) error {
	err := uc.Repository.InsertUser(firstName, lastName, username, email)

	if err != nil {
		return err
	}

	return nil
}
