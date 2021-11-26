package usecases

import (
	"log"

	"users-service/entities"
	"users-service/repositories"
	"users-service/utils/appstrings"
)

type GetUser interface {
	GetUserByID(userID string) (*entities.User, error)
}

type GetUserImpl struct {
	CacheRepo repositories.UsersCacheRepository
	Repo      repositories.UsersRepository
}

func (uc *GetUserImpl) GetUserByID(userID string) (*entities.User, error) {
	//Look In Cache
	user, err := uc.CacheRepo.GetUser(userID)

	//CHECK FOR NOT IN CACHE
	if err == nil {
		return user, nil
	}

	if err != appstrings.ErrorNotFoundOnCache {
		//There was an error different than it not being on cache.
		log.Println("Error on Cache.", err.Error())
	}

	//Not in cache
	user, err = uc.Repo.GetUserByID(userID)

	if err != nil {
		if err == appstrings.ErrorNotFoundOnDB {
			return nil, appstrings.ErrorNotFound
		}
		return nil, err
	}

	//Update cache for future requets
	err = uc.CacheRepo.SetUser(userID, user)

	if err != nil {
		//There was an error while working with the cache
		log.Println("Error on Cache.", err.Error())
	}

	return user, nil
}
