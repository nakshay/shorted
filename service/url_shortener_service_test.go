package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"shorted/configuration"
	"shorted/mocks"
	"shorted/model"
	"testing"
)

type UrlShortenerServiceTestSuite struct {
	suite.Suite
	context                   *gin.Context
	configData                *configuration.ConfigData
	mockCtrl                  *gomock.Controller
	mockStore                 *mocks.MockStore
	mockRandomStringGenerator *mocks.MockRandomStringGenerator
	urlShortenerService       URLShortenerService
}

func TestUrlShortenerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UrlShortenerServiceTestSuite))
}

func (suite *UrlShortenerServiceTestSuite) SetupTest() {
	suite.configData = &configuration.ConfigData{
		ServiceDomain:         "http://localhost:8080",
		RandomCharacterLength: 15,
	}
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockStore = mocks.NewMockStore(suite.mockCtrl)
	suite.mockRandomStringGenerator = mocks.NewMockRandomStringGenerator(suite.mockCtrl)
	suite.urlShortenerService = NewURLShortenerService(suite.mockStore, suite.configData, suite.mockRandomStringGenerator)
}

func (suite *UrlShortenerServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *UrlShortenerServiceTestSuite) TestNewUrlShortenerServiceShouldGenerateNewShortURLIfAlreadyNotPresent() {
	longURL := "long-url"
	randomString := "ajsidpfidncjdur"
	expectedResponse := model.ShortUrlResponse{ShortUrl: suite.configData.ServiceDomain + "/" + randomString}
	suite.mockStore.EXPECT().IsShortURLExists(longURL).Return("", false)
	suite.mockRandomStringGenerator.EXPECT().GenerateRandomString(suite.configData.RandomCharacterLength).Return(randomString)
	suite.mockStore.EXPECT().SaveShortURL(randomString, longURL)
	response := suite.urlShortenerService.GetShortenedURL(suite.context, longURL)
	suite.Equal(expectedResponse, response)
}

func (suite *UrlShortenerServiceTestSuite) TestNewUrlShortenerServiceShouldReturnResponseIfShortURLAlreadyExist() {
	longURL := "long-url"
	randomString := "ajsidpfidncjdur"
	expectedResponse := model.ShortUrlResponse{ShortUrl: suite.configData.ServiceDomain + "/" + randomString}
	suite.mockStore.EXPECT().IsShortURLExists(longURL).Return(randomString, true)
	response := suite.urlShortenerService.GetShortenedURL(suite.context, longURL)
	suite.Equal(expectedResponse, response)
}
