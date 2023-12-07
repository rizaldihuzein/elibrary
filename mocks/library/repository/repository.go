// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks_repository is a generated GoMock package.
package mocks_repository

import (
	context "context"
	domain "cosmart-library/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAPIInterface is a mock of APIInterface interface.
type MockAPIInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAPIInterfaceMockRecorder
}

// MockAPIInterfaceMockRecorder is the mock recorder for MockAPIInterface.
type MockAPIInterfaceMockRecorder struct {
	mock *MockAPIInterface
}

// NewMockAPIInterface creates a new mock instance.
func NewMockAPIInterface(ctrl *gomock.Controller) *MockAPIInterface {
	mock := &MockAPIInterface{ctrl: ctrl}
	mock.recorder = &MockAPIInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIInterface) EXPECT() *MockAPIInterfaceMockRecorder {
	return m.recorder
}

// FetchBookListBySubject mocks base method.
func (m *MockAPIInterface) FetchBookListBySubject(ctx context.Context, subject string) ([]domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookListBySubject", ctx, subject)
	ret0, _ := ret[0].([]domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookListBySubject indicates an expected call of FetchBookListBySubject.
func (mr *MockAPIInterfaceMockRecorder) FetchBookListBySubject(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookListBySubject", reflect.TypeOf((*MockAPIInterface)(nil).FetchBookListBySubject), ctx, subject)
}