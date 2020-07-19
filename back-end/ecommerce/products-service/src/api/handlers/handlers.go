package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"products-service/usecases"
)

type Handlers struct {
	GetAllProductsUsecase usecases.GetAllProducts
}

func (handler *Handlers) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.GetAllProductsUsecase.GetAllProducts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	prodcutsJSON, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(prodcutsJSON))
}
