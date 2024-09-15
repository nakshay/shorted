package service

import (
	"github.com/gin-gonic/gin"
	"shorted/loggingUtil"
	"shorted/model"
	"shorted/storage"
)

type MetricsService interface {
	GetMetrics(ctx *gin.Context) model.MetricsResponse
}
type metricsService struct {
	store storage.Store
}

func NewMetricsService(store storage.Store) MetricsService {
	return metricsService{store: store}
}

func (service metricsService) GetMetrics(ctx *gin.Context) model.MetricsResponse {
	logger := loggingUtil.GetLogger(ctx).WithFields("File", "metricsService").WithFields("Method", "GetMetrics")
	logger.Infof("Getting top 3 hits ")
	response := service.store.GetMetricsForTopDomain(3)
	logger.Debug("hits retrieved successfully")
	return response
}
