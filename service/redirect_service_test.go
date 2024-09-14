package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"shorted/configuration"
	"shorted/mocks"
	"shorted/shorted_error"
	"testing"
)

type RedirectServiceTestSuite struct {
	suite.Suite
	context         *gin.Context
	configData      *configuration.ConfigData
	mockCtrl        *gomock.Controller
	mockStore       *mocks.MockStore
	redirectService RedirectService
}

func TestRedirectServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RedirectServiceTestSuite))
}

func (suite *RedirectServiceTestSuite) SetupTest() {
	suite.configData = &configuration.ConfigData{
		ServiceDomain:         "http://localhost:8080",
		RandomCharacterLength: 15,
	}
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockStore = mocks.NewMockStore(suite.mockCtrl)
	suite.redirectService = NewRedirectService(suite.mockStore)
}

func (suite *RedirectServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *RedirectServiceTestSuite) TestGetFullURLShouldReturnFullURLForGivenShortURLIfExists() {
	shortURL := "short-url"
	fullURL := "https://google.com/iamfellinggood"
	suite.mockStore.EXPECT().FindFullURL(shortURL).Return(fullURL, true)
	response, err := suite.redirectService.GetFullURL(suite.context, shortURL)
	suite.Nil(err)
	suite.Equal(fullURL, response)
}

func (suite *RedirectServiceTestSuite) TestGetFullURLShouldReturnErrorIfURLDoesNotExists() {
	shortURL := "short-url"
	expectedErr := shorted_error.URLNotFoundErr
	suite.mockStore.EXPECT().FindFullURL(shortURL).Return("", false)
	response, err := suite.redirectService.GetFullURL(suite.context, shortURL)
	suite.Equal(expectedErr, err)
	suite.Empty(response)
}
