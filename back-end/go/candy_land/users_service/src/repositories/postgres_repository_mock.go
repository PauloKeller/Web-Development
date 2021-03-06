// Code generated by MockGen. DO NOT EDIT.
// Source: postgres_repository.go

// Package repositories is a generated GoMock package.
package repositories

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsersRepository is a mock of UsersRepository interface
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// InsertUser mocks base method
func (m *MockUsersRepository) InsertUser(firstName, lastName, username, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", firstName, lastName, username, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockUsersRepositoryMockRecorder) InsertUser(firstName, lastName, username, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUsersRepository)(nil).InsertUser), firstName, lastName, username, email)
}
