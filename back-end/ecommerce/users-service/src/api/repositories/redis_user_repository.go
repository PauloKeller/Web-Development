package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"

	"users-service/entities"
	"users-service/utils/appstrings"
)

type UsersCacheRepository interface {
	GetUser(userID string) (*entities.User, error)
	SetUser(userID string, User *entities.User) error
}

type RedisUsersRepository struct {
	pool *pool.Pool
}

const keySyntax = "user-%s"

func NewRedisUsersRepository() *RedisUsersRepository {
	pool, err := pool.New("tcp", "users-cache-redis:6379", 100)
	if err != nil {
		log.Fatal(err)
	}
	return &RedisUsersRepository{
		pool: pool,
	}
}

func (repo *RedisUsersRepository) GetUser(userID string) (*entities.User, error) {
	conn, err := repo.pool.Get()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf(keySyntax, userID)

	res, err := conn.Cmd("GET", key).Str()
	if err != nil {
		if err == redis.ErrRespNil {
			return nil, appstrings.ErrorNotFoundOnCache
		}
		return nil, err
	}
	var User *entities.User
	err = json.Unmarshal([]byte(res), &User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *RedisUsersRepository) SetUser(userID string, User *entities.User) error {
	conn, err := repo.pool.Get()
	if err != nil {
		return err
	}
	key := fmt.Sprintf(keySyntax, userID)

	jsonBytes, err := json.Marshal(User)
	if err != nil {
		return err
	}
	_, err = conn.Cmd("SET", key, string(jsonBytes)).Str()
	if err != nil {
		return err
	}
	return nil
}
