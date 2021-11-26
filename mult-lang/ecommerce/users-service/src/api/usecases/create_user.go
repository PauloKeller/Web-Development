package usecases

import (
	"time"
	"users-service/repositories"
)

type CreateUser interface {
	CreateUser(firstName string, lastName string, username string, email string, birthdate time.Time) error
}

type CreateUserImpl struct {
	CacheRepo repositories.UsersCacheRepository
	Repo      repositories.UsersRepository
}

func (uc *CreateUserImpl) CreateUser(firstName string, lastName string, username string, email string, birthdate time.Time) error {
	err := uc.Repo.InsertUser(firstName, lastName, username, email, birthdate)

	if err != nil {
		return err
	}

	return nil
}
