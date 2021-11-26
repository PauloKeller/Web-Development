package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"products-service/entities"
	"products-service/usecases"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

var (
	id          = "aa"
	name        = "product"
	description = "desc"
	quantity    = 5
	createdAt   = time.Date(2018, 03, 11, 0, 0, 0, 0, time.UTC)
	updatedAt   = time.Date(2019, 03, 11, 0, 0, 0, 0, time.UTC)
	price       = 200

	p = entities.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Price:       float64(price),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
)

func TestAllProductsHandlerEverythingRight(t *testing.T) {
	h := &Handlers{
		GetAllProductsUsecase: &MockGetAllProductsUsecase{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/products", http.HandlerFunc(h.GetAllProducts))

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/products"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("Test failed. Expected 200 Status. Got: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	products := make([]entities.Product, 0)
	products = append(products, p)
	pdto := &usecases.ProductDTO{}
	pdto.Results = products
	stub, _ := json.Marshal(pdto)

	log.Info(string(body))
	log.Info(string(stub))

	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}
	if string(body) != string(stub) {
		t.Errorf("Test failed. Wrong body.")
	}
}

type MockGetAllProductsUsecase struct{}

func (m *MockGetAllProductsUsecase) GetAllProducts() (*usecases.ProductDTO, error) {
	products := make([]entities.Product, 0)
	products = append(products, p)
	pdto := &usecases.ProductDTO{}
	pdto.Results = products
	return pdto, nil
}
