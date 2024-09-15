package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"shorted/mocks"
	"shorted/model"
	"testing"
)

type MetricsServiceTestSuite struct {
	suite.Suite
	context        *gin.Context
	mockCtrl       *gomock.Controller
	mockStore      *mocks.MockStore
	metricsService MetricsService
}

func TestNewMetricsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MetricsServiceTestSuite))
}

func (suite *MetricsServiceTestSuite) SetupTest() {
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockStore = mocks.NewMockStore(suite.mockCtrl)
	suite.metricsService = NewMetricsService(suite.mockStore)
}
func (suite *MetricsServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *MetricsServiceTestSuite) TestGetMetricsShouldReturnMetricsSuccessfully() {
	expectedResponse := model.MetricsResponse{TopHits: []model.TopHit{{
		URL:  "domain1.com",
		Hits: 4}, {
		URL:  "domain2.com",
		Hits: 3}, {
		URL:  "domain3.com",
		Hits: 1,
	}}}
	suite.mockStore.EXPECT().GetMetricsForTopDomain(3).Return(expectedResponse)
	response := suite.metricsService.GetMetrics(suite.context, 3)
	suite.Equal(expectedResponse, response)
}
