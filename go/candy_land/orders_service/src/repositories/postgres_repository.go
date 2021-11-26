package repositories

// PostgresOrdersRepository driver
type PostgresOrdersRepository struct{}

// OrdersRepository interface
type OrdersRepository interface {
	InsertOrder(userID string, productID string) error
}

// NewPostgresOrdersRepository constructor
// Create the database handle, confirm driver is present
func NewPostgresOrdersRepository() *PostgresOrdersRepository {
	repo := &PostgresOrdersRepository{}

	return repo
}

// InsertOrder Postgres database
func (repo *PostgresOrdersRepository) InsertOrder(userID string, productID string) error {
	return nil
}
