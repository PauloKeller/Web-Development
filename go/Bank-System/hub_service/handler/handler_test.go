package handler

import (
	"context"
	"hub_service/interactors"
	hub "hub_service/proto/out"
	"testing"
)

func TestCreateUser(t *testing.T) {
	handler := &Handler{
		CreateUserInteractor: &interactors.MockCreateUserInteractor{},
	}

	reply, err := handler.Create(context.TODO(), &hub.CreateUserRequest{})

	if err != nil {
		t.Errorf("Test failed. %v.", err)
	}

	if reply == nil {
		t.Errorf("Test failed. Should reply a user.")
	}
}

func TestFailToCreateUser(t *testing.T) {
	handler := &Handler{
		CreateUserInteractor: &interactors.MockCreateUserInteractor{},
	}

	reply, err := handler.Create(context.TODO(), &hub.CreateUserRequest{
		Username: "fail",
	})

	if reply != nil {
		t.Errorf("Test failed. Shouldn't reply a user.")
	}

	if err == nil {
		t.Errorf("Test failed. Should return a error.")
	}
}
