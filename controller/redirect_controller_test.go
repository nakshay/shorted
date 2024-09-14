package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"shorted/mocks"
	shortedErr "shorted/shorted_error"
	"testing"
)

type RedirectControllerTestSuite struct {
	suite.Suite
	context                      *gin.Context
	mockCtrl                     *gomock.Controller
	recorder                     *httptest.ResponseRecorder
	mockErrorResponseInterceptor *mocks.MockErrorResponseInterceptor
	mockRedirectService          *mocks.MockRedirectService
	controller                   RedirectController
}

func TestRedirectControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RedirectControllerTestSuite))
}

func (suite *RedirectControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.context.Request = httptest.NewRequest("POST", "/api/v1/short-it", nil)
	suite.mockErrorResponseInterceptor = mocks.NewMockErrorResponseInterceptor(suite.mockCtrl)
	suite.mockRedirectService = mocks.NewMockRedirectService(suite.mockCtrl)
	suite.controller = NewRedirectController(suite.mockRedirectService, suite.mockErrorResponseInterceptor)
}

func (suite *RedirectControllerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *RedirectControllerTestSuite) TestRedirectUrl_ShouldReturnErrorWhenParamsIsMissingInRequest() {
	suite.context.Request = httptest.NewRequest("GET", "/", nil)

	err := shortedErr.BadRequestErrorWithErrorMessage("short url is missing")
	suite.mockErrorResponseInterceptor.EXPECT().HandleBadRequest(suite.context, err)
	suite.controller.RedirectToFullUrl(suite.context)

}

func (suite *RedirectControllerTestSuite) TestRedirectUrl_ShouldReturnErrorWhenServiceReturnsError() {
	suite.context.Request = httptest.NewRequest("GET", "/", nil)
	suite.context.Params = gin.Params{{Key: "shortURL", Value: "some-short-url"}}
	suite.mockRedirectService.EXPECT().GetFullURL(suite.context, "some-short-url").Return("", shortedErr.InternalServerError())
	suite.mockErrorResponseInterceptor.EXPECT().HandleServiceErr(suite.context, shortedErr.InternalServerError())
	suite.controller.RedirectToFullUrl(suite.context)

}

func (suite *RedirectControllerTestSuite) TestRedirectUrl_ShouldReturnRedirectUrlOnSuccess() {
	shortURL := "some-short-url"
	suite.context.Request = httptest.NewRequest("GET", "/"+shortURL, nil)
	suite.context.Params = gin.Params{{Key: "shortURL", Value: "some-short-url"}}
	suite.mockRedirectService.EXPECT().GetFullURL(suite.context, "some-short-url").Return("some-long-url", nil)

	suite.controller.RedirectToFullUrl(suite.context)
	suite.Equal(http.StatusMovedPermanently, suite.recorder.Code)
}
