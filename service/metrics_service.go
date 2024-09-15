package service

import (
	"github.com/gin-gonic/gin"
	"shorted/loggingUtil"
	"shorted/model"
	"shorted/storage"
)

//go:generate mockgen -source=./metrics_service.go -destination=../mocks/mock_metric_service.go -package=mocks

type MetricsService interface {
	GetMetrics(ctx *gin.Context, topNDomains int) model.MetricsResponse
}
type metricsService struct {
	store storage.Store
}

func NewMetricsService(store storage.Store) MetricsService {
	return metricsService{store: store}
}

func (service metricsService) GetMetrics(ctx *gin.Context, topNDomains int) model.MetricsResponse {
	logger := loggingUtil.GetLogger(ctx).WithFields("File", "metricsService").WithFields("Method", "GetMetrics")
	logger.Infof("Getting top 3 hits ")
	response := service.store.GetMetricsForTopDomain(topNDomains)
	logger.Debug("hits retrieved successfully")
	return response
}
