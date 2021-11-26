package entities

import "testing"

func TestUser(t *testing.T) {
	u := new(User)
	u.Balance = 500
	u.FirstName = "Paulo"

	if u.Balance != 500 {
		t.Errorf("Test failed. Expected Normal Account.")
	}

	if u.FirstName != "Paulo" {
		t.Errorf("Test failed. Expected Silver Account.")
	}
}
