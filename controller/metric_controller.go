package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/configuration"
	"shorted/loggingUtil"
	"shorted/service"
	"strconv"
)

type MetricController interface {
	GetMetrics(ctx *gin.Context)
}

type metricController struct {
	metricService service.MetricsService
	configData    *configuration.ConfigData
}

func NewMetricController(metricService service.MetricsService, configData *configuration.ConfigData) MetricController {
	return metricController{metricService: metricService, configData: configData}
}

func (controller metricController) GetMetrics(ctx *gin.Context) {
	logger := loggingUtil.GetLogger(ctx).WithFields("File", "metricController").WithFields("Method", "GetMetrics")
	logger.Infof("Getting top hits ")
	topNDomains, _ := strconv.Atoi(ctx.DefaultQuery("limit", strconv.Itoa(controller.configData.MetricDefaultSize)))
	response := controller.metricService.GetMetrics(ctx, topNDomains)
	logger.Debugf("Read metrics successful")
	ctx.JSON(http.StatusOK, response)

}
