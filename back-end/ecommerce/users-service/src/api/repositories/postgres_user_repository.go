package repositories

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"users-service/entities"
	"users-service/utils/appstrings"
)

// UsersRepository interface
type UsersRepository interface {
	GetUserByID(userID string) (*entities.User, error)
	InsertUser(firstName string, lastName string, username string, email string, birthdate time.Time) error
}

// PostgresUsersRepository driver
type PostgresUsersRepository struct {
	db *sql.DB
}

// NewPostgresUsersRepository constructor
// Create the database handle, confirm driver is present
func NewPostgresUsersRepository() *PostgresUsersRepository {
	db, err := sql.Open("postgres", "postgres://ecommerce:mysecret@users-postgres/ecommerce-users?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	repo := &PostgresUsersRepository{
		db,
	}

	return repo
}

func (repo *PostgresUsersRepository) GetUserByID(userID string) (*entities.User, error) {
	user := &entities.User{}
	row := repo.db.QueryRow("Select id, username, first_name, last_name, email, birthdate, created_at, updated_at, balance from users where id=$1", userID)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Birthdate, &user.CreatedAt, &user.UpdatedAt, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appstrings.ErrorNotFoundOnDB
		}
		return nil, err
	}
	return user, nil
}

func (repo *PostgresUsersRepository) InsertUser(firstName string, lastName string, username string, email string, birthdate time.Time) error {
	log.Println(firstName, lastName, username, email, birthdate)
	_, err := repo.db.
		Exec("INSERT INTO users (first_name, last_name, username, email, birthdate) VALUES ($1,$2,$3,$4,$5)", firstName, lastName, username, email, birthdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return appstrings.ErrorInsertUser
		}
		return err
	}

	return nil
}
