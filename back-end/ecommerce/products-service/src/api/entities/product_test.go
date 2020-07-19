package entities

import "testing"

func TestProduct(t *testing.T) {
	p := new(Product)
	p.Price = 500
	p.Name = "product"

	if p.Price != 500 {
		t.Errorf("Test failed. 500 of price.")
	}

	if p.Name != "product" {
		t.Errorf("Test failed. Product name.")
	}
}
