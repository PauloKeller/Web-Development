package usecases

import "candy_land/carts_service/src/repositories"

// AddProductToCartProvider repositories
type AddProductToCartProvider struct {
	Repository repositories.CartsRepository
}

// AddProductToCart interface
type AddProductToCart interface {
	Add(cartID string, productID string, quantity int) error
}

// Add implementation
func (uc *AddProductToCartProvider) Add(cartID string, productID string, quantity int) error {
	err := uc.Repository.InsertProductIntoCart(cartID, productID, quantity)

	if err != nil {
		return err
	}

	return nil
}
