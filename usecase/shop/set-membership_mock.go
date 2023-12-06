// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_shop is a generated GoMock package.
package shop

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKYC is a mock of KYC interface.
type MockKYC struct {
	ctrl     *gomock.Controller
	recorder *MockKYCMockRecorder
}

// MockKYCMockRecorder is the mock recorder for MockKYC.
type MockKYCMockRecorder struct {
	mock *MockKYC
}

// NewMockKYC creates a new mock instance.
func NewMockKYC(ctrl *gomock.Controller) *MockKYC {
	mock := &MockKYC{ctrl: ctrl}
	mock.recorder = &MockKYCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKYC) EXPECT() *MockKYCMockRecorder {
	return m.recorder
}

// CheckKYC mocks base method.
func (m *MockKYC) CheckKYC(ctx context.Context, shopID int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckKYC", ctx, shopID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckKYC indicates an expected call of CheckKYC.
func (mr *MockKYCMockRecorder) CheckKYC(ctx, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckKYC", reflect.TypeOf((*MockKYC)(nil).CheckKYC), ctx, shopID)
}

// MockShopScore is a mock of ShopScore interface.
type MockShopScore struct {
	ctrl     *gomock.Controller
	recorder *MockShopScoreMockRecorder
}

// MockShopScoreMockRecorder is the mock recorder for MockShopScore.
type MockShopScoreMockRecorder struct {
	mock *MockShopScore
}

// NewMockShopScore creates a new mock instance.
func NewMockShopScore(ctrl *gomock.Controller) *MockShopScore {
	mock := &MockShopScore{ctrl: ctrl}
	mock.recorder = &MockShopScoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShopScore) EXPECT() *MockShopScoreMockRecorder {
	return m.recorder
}

// GetShopScore mocks base method.
func (m *MockShopScore) GetShopScore(ctx context.Context, shopID int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShopScore", ctx, shopID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShopScore indicates an expected call of GetShopScore.
func (mr *MockShopScoreMockRecorder) GetShopScore(ctx, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShopScore", reflect.TypeOf((*MockShopScore)(nil).GetShopScore), ctx, shopID)
}

// MockOrder is a mock of Order interface.
type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

// MockOrderMockRecorder is the mock recorder for MockOrder.
type MockOrderMockRecorder struct {
	mock *MockOrder
}

// NewMockOrder creates a new mock instance.
func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

// GetNetIncome mocks base method.
func (m *MockOrder) GetNetIncome(ctx context.Context, shopID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetIncome", ctx, shopID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetIncome indicates an expected call of GetNetIncome.
func (mr *MockOrderMockRecorder) GetNetIncome(ctx, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetIncome", reflect.TypeOf((*MockOrder)(nil).GetNetIncome), ctx, shopID)
}

// GetTotalFinishOrder mocks base method.
func (m *MockOrder) GetTotalFinishOrder(ctx context.Context, shopID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalFinishOrder", ctx, shopID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalFinishOrder indicates an expected call of GetTotalFinishOrder.
func (mr *MockOrderMockRecorder) GetTotalFinishOrder(ctx, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalFinishOrder", reflect.TypeOf((*MockOrder)(nil).GetTotalFinishOrder), ctx, shopID)
}

// MockShopClass is a mock of ShopClass interface.
type MockShopClass struct {
	ctrl     *gomock.Controller
	recorder *MockShopClassMockRecorder
}

// MockShopClassMockRecorder is the mock recorder for MockShopClass.
type MockShopClassMockRecorder struct {
	mock *MockShopClass
}

// NewMockShopClass creates a new mock instance.
func NewMockShopClass(ctrl *gomock.Controller) *MockShopClass {
	mock := &MockShopClass{ctrl: ctrl}
	mock.recorder = &MockShopClassMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShopClass) EXPECT() *MockShopClassMockRecorder {
	return m.recorder
}

// SetMembership mocks base method.
func (m *MockShopClass) SetMembership(ctx context.Context, shopID int64, newStatus int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMembership", ctx, shopID, newStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMembership indicates an expected call of SetMembership.
func (mr *MockShopClassMockRecorder) SetMembership(ctx, shopID, newStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMembership", reflect.TypeOf((*MockShopClass)(nil).SetMembership), ctx, shopID, newStatus)
}
