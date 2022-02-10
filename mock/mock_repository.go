// Code generated by MockGen. DO NOT EDIT.
// Source: repository/wallet.go

// Package mock is a generated GoMock package.
package mock

import (
	model "bootcamp/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWalletRepository is a mock of IWalletRepository interface.
type MockIWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletRepositoryMockRecorder
}

// MockIWalletRepositoryMockRecorder is the mock recorder for MockIWalletRepository.
type MockIWalletRepositoryMockRecorder struct {
	mock *MockIWalletRepository
}

// NewMockIWalletRepository creates a new mock instance.
func NewMockIWalletRepository(ctrl *gomock.Controller) *MockIWalletRepository {
	mock := &MockIWalletRepository{ctrl: ctrl}
	mock.recorder = &MockIWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWalletRepository) EXPECT() *MockIWalletRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIWalletRepository) Create(username string, amount float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", username, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIWalletRepositoryMockRecorder) Create(username, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIWalletRepository)(nil).Create), username, amount)
}

// GetAll mocks base method.
func (m *MockIWalletRepository) GetAll() ([]*model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIWalletRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIWalletRepository)(nil).GetAll))
}

// GetByUsername mocks base method.
func (m *MockIWalletRepository) GetByUsername(username string) (*model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockIWalletRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockIWalletRepository)(nil).GetByUsername), username)
}

// Update mocks base method.
func (m *MockIWalletRepository) Update(username string, amount float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", username, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIWalletRepositoryMockRecorder) Update(username, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIWalletRepository)(nil).Update), username, amount)
}