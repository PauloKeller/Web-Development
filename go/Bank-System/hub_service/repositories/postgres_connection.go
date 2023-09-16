package repositories

import (
	"fmt"

	"hub_service/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnectionData struct {
	DBName   string
	Driver   string
	Username string
	Port     string
	Host     string
	Password string
}

type Repositories struct {
	db      *gorm.DB
	User    UserRepositoryInterface
	Account AccountRepositoryInterface
}

func NewRepositories(connectionData *PostgresConnectionData) (*Repositories, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", connectionData.Host, connectionData.Port, connectionData.Username, connectionData.DBName, connectionData.Password)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Repositories{
		User:    NewUserRepository(db),
		Account: NewAccountRepository(db),
		db:      db,
	}, nil
}

func (repository *Repositories) Automigrate() error {
	return repository.db.AutoMigrate(&entities.UserEntity{}, &entities.AccountEntity{})
}
