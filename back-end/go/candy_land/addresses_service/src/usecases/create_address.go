package usecases

import "candy_land/addresses_service/src/repositories"

// CreateAddressProvider repositories
type CreateAddressProvider struct {
	Repository repositories.AddressesRepository
}

// CreateAddress interface
type CreateAddress interface {
	Create(userID string, countryID string, stateID string, cityID string, streetID string, number string, complement string) error
}

// Create implementation
func (uc *CreateAddressProvider) Create(userID string, countryID string, stateID string, cityID string, streetID string, number string, complement string) error {
	err := uc.Repository.InsertAddress(userID, countryID, stateID, cityID, streetID, number, complement)

	if err != nil {
		return err
	}

	return nil
}
