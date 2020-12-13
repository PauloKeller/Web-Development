package repositories

// UsersRepository interface
type UsersRepository interface {
	InsertUser(firstName string, lastName string, username string, email string) error
}

// PostgresUsersRepository driver
type PostgresUsersRepository struct{}

// NewPostgresUsersRepository constructor
// Create the database handle, confirm driver is present
func NewPostgresUsersRepository() *PostgresUsersRepository {
	repo := &PostgresUsersRepository{}

	return repo
}

// InsertUser into Postgres database
func (repo *PostgresUsersRepository) InsertUser(firstName string, lastName string, username string, email string) error {
	return nil
}
