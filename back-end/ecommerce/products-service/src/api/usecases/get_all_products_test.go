package usecases

import (
	"testing"
	"time"

	"products-service/entities"
)

var (
	name        = "product"
	description = "desc"
	createdAt   = time.Date(2018, 03, 11, 0, 0, 0, 0, time.UTC)
	updatedAt   = time.Date(2019, 03, 11, 0, 0, 0, 0, time.UTC)
)

func TestGetAllProducts(t *testing.T) {
	uc := &GetAllProductsImpl{
		Repo: &MockProductsRepository{},
	}

	productDTO, err := uc.GetAllProducts()
	if err != nil {
		t.Errorf("Test failed. Unexpected error.")
	}

	if len(productDTO.Results) != 10 {
		t.Errorf("Test failed. Wrong Size.")
	}

	if productDTO.Results[0].Description != description {
		t.Errorf("Test failed. Wrong Description.")
	}

	if productDTO.Results[3].Quantity != 3 {
		t.Errorf("Test failed. Wrong Quantity.")
	}
}

type MockProductsRepository struct{}

func (m *MockProductsRepository) GetAllProducts() ([]entities.Product, error) {
	products := make([]entities.Product, 0)

	for i := 0; i < 10; i++ {
		price := 12 * i

		p := entities.Product{
			Name:        "Test" + string(i),
			Description: description,
			Quantity:    i,
			Price:       float64(price),
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}

		products = append(products, p)
	}

	return products, nil
}
