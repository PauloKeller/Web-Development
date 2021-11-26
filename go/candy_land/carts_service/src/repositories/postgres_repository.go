package repositories

// PostgresCartsRepository driver
type PostgresCartsRepository struct{}

// CartsRepository interface
type CartsRepository interface {
	InsertProductIntoCart(cartID string, productID string, quantity int) error
}

// NewPostgresCartsRepository constructor
// Create the database handle, confirm driver is present
func NewPostgresCartsRepository() *PostgresCartsRepository {
	repo := &PostgresCartsRepository{}

	return repo
}

// InsertProductIntoCart Postgres database
func (repo *PostgresCartsRepository) InsertProductIntoCart(cartID string, productID string, quantity int) error {
	return nil
}
