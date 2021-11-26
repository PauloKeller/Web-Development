package repositories

// PostgresAddressesRepository driver
type PostgresAddressesRepository struct {
}

// AddressesRepository interface
type AddressesRepository interface {
	InsertAddress(userID string, countryID string, stateID string, cityID string, streetID string, number string, complement string) error
}

// NewPostgresAddressesRepository constructor
// Create the database handle, confirm driver is present
func NewPostgresAddressesRepository() *PostgresAddressesRepository {
	repo := &PostgresAddressesRepository{}

	return repo
}

// InsertAddress into Postgres database
func (repo *PostgresAddressesRepository) InsertAddress(userID string, countryID string, stateID string, cityID string, streetID string, number string, complement string) error {
	return nil
}
