// Code generated by MockGen. DO NOT EDIT.
// Source: ./url_shortener_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "shorted/model"
	shorted_error "shorted/shorted_error"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockURLShortenerService is a mock of URLShortenerService interface.
type MockURLShortenerService struct {
	ctrl     *gomock.Controller
	recorder *MockURLShortenerServiceMockRecorder
}

// MockURLShortenerServiceMockRecorder is the mock recorder for MockURLShortenerService.
type MockURLShortenerServiceMockRecorder struct {
	mock *MockURLShortenerService
}

// NewMockURLShortenerService creates a new mock instance.
func NewMockURLShortenerService(ctrl *gomock.Controller) *MockURLShortenerService {
	mock := &MockURLShortenerService{ctrl: ctrl}
	mock.recorder = &MockURLShortenerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLShortenerService) EXPECT() *MockURLShortenerServiceMockRecorder {
	return m.recorder
}

// GetFullURL mocks base method.
func (m *MockURLShortenerService) GetFullURL(ctx *gin.Context, url string) (string, *shorted_error.ShortedError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullURL", ctx, url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*shorted_error.ShortedError)
	return ret0, ret1
}

// GetFullURL indicates an expected call of GetFullURL.
func (mr *MockURLShortenerServiceMockRecorder) GetFullURL(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullURL", reflect.TypeOf((*MockURLShortenerService)(nil).GetFullURL), ctx, url)
}

// GetShortenedURL mocks base method.
func (m *MockURLShortenerService) GetShortenedURL(ctx *gin.Context, url string) model.ShortUrlResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortenedURL", ctx, url)
	ret0, _ := ret[0].(model.ShortUrlResponse)
	return ret0
}

// GetShortenedURL indicates an expected call of GetShortenedURL.
func (mr *MockURLShortenerServiceMockRecorder) GetShortenedURL(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortenedURL", reflect.TypeOf((*MockURLShortenerService)(nil).GetShortenedURL), ctx, url)
}
