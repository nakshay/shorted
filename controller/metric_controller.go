package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/loggingUtil"
	"shorted/service"
)

type MetricController interface {
	GetMetrics(ctx *gin.Context)
}

type metricController struct {
	metricService service.MetricsService
}

func NewMetricController(metricService service.MetricsService) MetricController {
	return metricController{metricService: metricService}
}

func (c metricController) GetMetrics(ctx *gin.Context) {
	logger := loggingUtil.GetLogger(ctx).WithFields("File", "metricController").WithFields("Method", "GetMetrics")
	logger.Infof("Getting top hits ")
	response := c.metricService.GetMetrics(ctx)
	logger.Debugf("Read metrics successful")
	ctx.JSON(http.StatusOK, response)

}
