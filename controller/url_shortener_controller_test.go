package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"shorted/mocks"
	"shorted/model"
	"testing"
)

type UrlShortenerControllerTestSuite struct {
	suite.Suite
	context                      *gin.Context
	mockCtrl                     *gomock.Controller
	recorder                     *httptest.ResponseRecorder
	mockErrorResponseInterceptor *mocks.MockErrorResponseInterceptor
	mockURLShortenerService      *mocks.MockURLShortenerService
	controller                   URLShortenerController
}

func TestUrlShortenerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UrlShortenerControllerTestSuite))
}

func (suite *UrlShortenerControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.context.Request = httptest.NewRequest("POST", "/api/v1/short-it", nil)
	suite.mockErrorResponseInterceptor = mocks.NewMockErrorResponseInterceptor(suite.mockCtrl)
	suite.mockURLShortenerService = mocks.NewMockURLShortenerService(suite.mockCtrl)
	suite.controller = NewURLShortenerController(suite.mockURLShortenerService, suite.mockErrorResponseInterceptor)
}

func (suite *UrlShortenerControllerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *UrlShortenerControllerTestSuite) TestURLShortenerShouldReturnShortURLSuccessfully() {
	longURL := "long-url"
	request := model.ShortURLRequest{URL: longURL}
	expectedResponse := model.ShortUrlResponse{ShortUrl: "short-url"}
	requestBytes, _ := json.Marshal(request)
	suite.context.Request = httptest.NewRequest("POST", "/api/v1/short-it",
		bytes.NewBufferString(string(requestBytes)))
	suite.mockURLShortenerService.EXPECT().GetShortenedURL(suite.context, longURL).Return(expectedResponse)

	suite.controller.GetShortenedURL(suite.context)
	responseBytes, _ := io.ReadAll(suite.recorder.Body)
	var actualResponse model.ShortUrlResponse
	_ = json.Unmarshal(responseBytes, &actualResponse)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Equal(expectedResponse, actualResponse)

}

func (suite *UrlShortenerControllerTestSuite) TestURLShortenerShouldReturnBadRequestErrorWhenRequestIsInvalid() {

	responseBytes, _ := io.ReadAll(suite.recorder.Body)
	var actualResponse model.ShortUrlResponse
	_ = json.Unmarshal(responseBytes, &actualResponse)
	suite.mockErrorResponseInterceptor.EXPECT().HandleBadRequest(suite.context, gomock.Any())
	suite.controller.GetShortenedURL(suite.context)
	suite.Empty(actualResponse)

}
