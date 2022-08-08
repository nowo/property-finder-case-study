// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/order/repository_order/irepository_order.go

// Package mocks is a generated GoMock package.
package mocks

import (
	order "property-finder-go-bootcamp-homework/internal/domain/order"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIOrderRepository is a mock of IOrderRepository interface.
type MockIOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIOrderRepositoryMockRecorder
}

// MockIOrderRepositoryMockRecorder is the mock recorder for MockIOrderRepository.
type MockIOrderRepositoryMockRecorder struct {
	mock *MockIOrderRepository
}

// NewMockIOrderRepository creates a new mock instance.
func NewMockIOrderRepository(ctrl *gomock.Controller) *MockIOrderRepository {
	mock := &MockIOrderRepository{ctrl: ctrl}
	mock.recorder = &MockIOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOrderRepository) EXPECT() *MockIOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockIOrderRepository) CreateOrder(newOrder order.Order) (order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", newOrder)
	ret0, _ := ret[0].(order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockIOrderRepositoryMockRecorder) CreateOrder(newOrder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockIOrderRepository)(nil).CreateOrder), newOrder)
}

// GetOrderByUserID mocks base method.
func (m *MockIOrderRepository) GetOrderByUserID(userID uint) ([]order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByUserID", userID)
	ret0, _ := ret[0].([]order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByUserID indicates an expected call of GetOrderByUserID.
func (mr *MockIOrderRepositoryMockRecorder) GetOrderByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByUserID", reflect.TypeOf((*MockIOrderRepository)(nil).GetOrderByUserID), userID)
}

// GetOrderFromLastMonth mocks base method.
func (m *MockIOrderRepository) GetOrderFromLastMonth(userID uint) ([]order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderFromLastMonth", userID)
	ret0, _ := ret[0].([]order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderFromLastMonth indicates an expected call of GetOrderFromLastMonth.
func (mr *MockIOrderRepositoryMockRecorder) GetOrderFromLastMonth(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderFromLastMonth", reflect.TypeOf((*MockIOrderRepository)(nil).GetOrderFromLastMonth), userID)
}
