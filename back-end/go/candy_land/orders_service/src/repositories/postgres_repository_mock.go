// Code generated by MockGen. DO NOT EDIT.
// Source: postgres_repository.go

// Package repositories is a generated GoMock package.
package repositories

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrdersRepository is a mock of OrdersRepository interface
type MockOrdersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersRepositoryMockRecorder
}

// MockOrdersRepositoryMockRecorder is the mock recorder for MockOrdersRepository
type MockOrdersRepositoryMockRecorder struct {
	mock *MockOrdersRepository
}

// NewMockOrdersRepository creates a new mock instance
func NewMockOrdersRepository(ctrl *gomock.Controller) *MockOrdersRepository {
	mock := &MockOrdersRepository{ctrl: ctrl}
	mock.recorder = &MockOrdersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrdersRepository) EXPECT() *MockOrdersRepositoryMockRecorder {
	return m.recorder
}

// InsertOrder mocks base method
func (m *MockOrdersRepository) InsertOrder(userID, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrder", userID, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrder indicates an expected call of InsertOrder
func (mr *MockOrdersRepositoryMockRecorder) InsertOrder(userID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrder", reflect.TypeOf((*MockOrdersRepository)(nil).InsertOrder), userID, productID)
}
