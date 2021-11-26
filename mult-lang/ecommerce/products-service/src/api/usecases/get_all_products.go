package usecases

import (
	"products-service/entities"
	"products-service/repositories"
	"products-service/utils/appstrings"
)

type ProductDTO struct {
	Results []entities.Product `json:"results"`
}

type GetAllProducts interface {
	GetAllProducts() (*ProductDTO, error)
}

type GetAllProductsImpl struct {
	Repo repositories.ProductsRepository
}

func (uc *GetAllProductsImpl) GetAllProducts() (*ProductDTO, error) {
	products, err := uc.Repo.GetAllProducts()

	if err != nil {
		if err == appstrings.ErrorNotFoundOnDB {
			return nil, appstrings.ErrorNotFound
		}
		return nil, err
	}

	productDTO := &ProductDTO{
		Results: products,
	}

	return productDTO, nil
}
