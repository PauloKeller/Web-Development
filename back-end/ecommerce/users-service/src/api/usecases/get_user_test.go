package usecases

import (
	"testing"
	"time"

	"users-service/entities"
	"users-service/utils/appstrings"
)

const (
	uIDCache      = "cache_user_id"
	uNameCache    = "test_cache"
	uBalanceCache = 1000
	uEmailCache   = "test@cache.com"
	uIDRepo       = "repo_user_id"
	uNameRepo     = "test_repo"
	uBalanceRepo  = 500
	uEmailRepo    = "test@repo.com"
)

func TestGetUserInCache(t *testing.T) {
	uc := &GetUserImpl{
		CacheRepo: &MockUsersCacheRepository{},
		Repo:      &MockUsersRepository{},
	}

	userDTO, err := uc.GetUserByID(uIDCache)
	if err != nil {
		t.Errorf("Test failed. Unexpected error.")
	}

	if userDTO.Username != uNameCache {
		t.Errorf("Test failed. Wrong Username.")
	}
	if userDTO.Email != uEmailCache {
		t.Errorf("Test failed. Wrong Email.")
	}
}

func TestGetUserNotInCache(t *testing.T) {
	cacheRepo := &MockUsersCacheRepository{}

	uc := &GetUserImpl{
		CacheRepo: cacheRepo,
		Repo:      &MockUsersRepository{},
	}

	userDTO, err := uc.GetUserByID(uIDRepo)
	if err != nil {
		t.Errorf("Test failed. Unexpected error.")
	}

	if userDTO.Username != uNameRepo {
		t.Errorf("Test failed. Wrong Username.")
	}
	if userDTO.Email != uEmailRepo {
		t.Errorf("Test failed. Wrong Email.")
	}
	if !cacheRepo.WasSetCalled {
		t.Errorf("Test failed. Expected Cache Set.")
	}
}

func TestGetUserNotInCacheOrDB(t *testing.T) {
	cacheRepo := &MockUsersCacheRepository{}

	uc := &GetUserImpl{
		CacheRepo: cacheRepo,
		Repo:      &MockUsersRepository{},
	}

	_, err := uc.GetUserByID("Test3")
	if err != appstrings.ErrorNotFound {
		t.Errorf("Test failed. Expected not Found Error.")
	}

}

type MockUsersRepository struct{}

func (m *MockUsersRepository) GetUserByID(userID string) (*entities.User, error) {
	if userID == uIDRepo {
		return &entities.User{
			Username: uNameRepo,
			Balance:  uBalanceRepo,
			Email:    uEmailRepo,
		}, nil
	}
	return nil, appstrings.ErrorNotFound
}

func (m *MockUsersRepository) InsertUser(firstName string, lastName string, username string, email string, birthdate time.Time) error {
	if firstName != "" {
		return nil
	}
	return appstrings.ErrorNotFound
}

type MockUsersCacheRepository struct {
	WasSetCalled bool
}

func (m *MockUsersCacheRepository) GetUser(userID string) (*entities.User, error) {
	if userID == uIDCache {
		return &entities.User{
			Username: uNameCache,
			Balance:  uBalanceCache,
			Email:    uEmailCache,
		}, nil
	}
	return nil, appstrings.ErrorNotFoundOnDB
}

func (m *MockUsersCacheRepository) SetUser(userID string, User *entities.User) error {
	m.WasSetCalled = true
	return nil
}
