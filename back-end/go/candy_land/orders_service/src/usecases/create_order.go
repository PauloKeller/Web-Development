package usecases

import "candy_land/orders_service/src/repositories"

// CreateOrderProvider repositories
type CreateOrderProvider struct {
	Repository repositories.OrdersRepository
}

// CreateOrder interface
type CreateOrder interface {
	Create(userID string, productID string) error
}

// Create implementation
func (uc *CreateOrderProvider) Create(userID string, productID string) error {
	err := uc.Repository.InsertOrder(userID, productID)

	if err != nil {
		return err
	}

	return nil
}
