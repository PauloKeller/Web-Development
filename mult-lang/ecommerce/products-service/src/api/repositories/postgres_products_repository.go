package repositories

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"products-service/entities"
)

// ProductsRepository interface
type ProductsRepository interface {
	GetAllProducts() ([]entities.Product, error)
}

// PostgresProductsRepository driver
type PostgresProductsRepository struct {
	db *sql.DB
}

// NewPostgresProductsRepository constructor
// Create the database handle, confirm driver is present
func NewPostgresProductsRepository() *PostgresProductsRepository {
	db, err := sql.Open("postgres", "postgres://ecommerce:mysecret@products-postgres/ecommerce-products?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	repo := &PostgresProductsRepository{
		db,
	}

	return repo
}

func (repo *PostgresProductsRepository) GetAllProducts() ([]entities.Product, error) {
	products := make([]entities.Product, 0)
	rows, err := repo.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		product := entities.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
