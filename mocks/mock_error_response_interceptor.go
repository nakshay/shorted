// Code generated by MockGen. DO NOT EDIT.
// Source: ./error_response_interceptor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	shorted_error "shorted/shorted_error"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockErrorResponseInterceptor is a mock of ErrorResponseInterceptor interface.
type MockErrorResponseInterceptor struct {
	ctrl     *gomock.Controller
	recorder *MockErrorResponseInterceptorMockRecorder
}

// MockErrorResponseInterceptorMockRecorder is the mock recorder for MockErrorResponseInterceptor.
type MockErrorResponseInterceptorMockRecorder struct {
	mock *MockErrorResponseInterceptor
}

// NewMockErrorResponseInterceptor creates a new mock instance.
func NewMockErrorResponseInterceptor(ctrl *gomock.Controller) *MockErrorResponseInterceptor {
	mock := &MockErrorResponseInterceptor{ctrl: ctrl}
	mock.recorder = &MockErrorResponseInterceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockErrorResponseInterceptor) EXPECT() *MockErrorResponseInterceptorMockRecorder {
	return m.recorder
}

// HandleBadRequest mocks base method.
func (m *MockErrorResponseInterceptor) HandleBadRequest(ctx *gin.Context, bindErr error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleBadRequest", ctx, bindErr)
}

// HandleBadRequest indicates an expected call of HandleBadRequest.
func (mr *MockErrorResponseInterceptorMockRecorder) HandleBadRequest(ctx, bindErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBadRequest", reflect.TypeOf((*MockErrorResponseInterceptor)(nil).HandleBadRequest), ctx, bindErr)
}

// HandleServiceErr mocks base method.
func (m *MockErrorResponseInterceptor) HandleServiceErr(ctx *gin.Context, bindErr *shorted_error.ShortedError) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleServiceErr", ctx, bindErr)
}

// HandleServiceErr indicates an expected call of HandleServiceErr.
func (mr *MockErrorResponseInterceptorMockRecorder) HandleServiceErr(ctx, bindErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleServiceErr", reflect.TypeOf((*MockErrorResponseInterceptor)(nil).HandleServiceErr), ctx, bindErr)
}
