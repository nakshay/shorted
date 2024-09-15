package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"shorted/configuration"
	"shorted/mocks"
	"shorted/model"
	"testing"
)

type MetricControllerTestSuite struct {
	suite.Suite
	context           *gin.Context
	mockCtrl          *gomock.Controller
	recorder          *httptest.ResponseRecorder
	mockMetricService *mocks.MockMetricsService
	configData        *configuration.ConfigData
	controller        MetricController
}

func TestNewMetricControllerTestSuite(t *testing.T) {
	suite.Run(t, new(MetricControllerTestSuite))
}

func (suite *MetricControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.context.Request = httptest.NewRequest("POST", "/api/v1/short-it", nil)
	suite.mockMetricService = mocks.NewMockMetricsService(suite.mockCtrl)
	suite.configData = &configuration.ConfigData{
		MetricDefaultSize: 3,
	}
	suite.controller = NewMetricController(suite.mockMetricService, suite.configData)
}

func (suite *MetricControllerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *MetricControllerTestSuite) TestGetMetricsShouldCallServiceMethodWithValueInConfig() {
	suite.context.Request = httptest.NewRequest("GET", "/api/v1/metric", nil)

	expectedResponse := model.MetricsResponse{TopHits: []model.TopHit{{
		URL:  "domain1.com",
		Hits: 4}, {
		URL:  "domain2.com",
		Hits: 3}, {
		URL:  "domain3.com",
		Hits: 1,
	}}}
	suite.mockMetricService.EXPECT().GetMetrics(suite.context, 3).Return(expectedResponse)
	suite.controller.GetMetrics(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)

}

func (suite *MetricControllerTestSuite) TestGetMetricsShouldCallServiceMethodByOverridingMetricSizeFromQueryParams() {
	suite.context.Request = httptest.NewRequest("GET", "/api/v1/metric?limit=5", nil)

	expectedResponse := model.MetricsResponse{TopHits: []model.TopHit{{
		URL:  "domain1.com",
		Hits: 4}, {
		URL:  "domain2.com",
		Hits: 3}, {
		URL:  "domain3.com",
		Hits: 1,
	}}}
	suite.mockMetricService.EXPECT().GetMetrics(suite.context, 5).Return(expectedResponse)
	suite.controller.GetMetrics(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)

}
