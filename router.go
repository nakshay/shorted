package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/configuration"
	"shorted/controller"
	"shorted/service"
	shortedErr "shorted/shorted_error"
	"shorted/storage"
	"shorted/util"
)

func setupRouter(config *configuration.ConfigData) *gin.Engine {

	r := gin.Default()
	// Health endpoint
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// initialization
	dbStore := storage.Init()

	errorResponseInterceptor := shortedErr.NewErrorResponseInterceptor()
	randomStringGenerator := util.NewRandomStringGenerator()

	shortenerService := service.NewURLShortenerService(dbStore, config, randomStringGenerator)
	redirectService := service.NewRedirectService(dbStore)
	metricService := service.NewMetricsService(dbStore)

	urlShortenerController := controller.NewURLShortenerController(shortenerService, errorResponseInterceptor)
	redirectController := controller.NewRedirectController(redirectService, errorResponseInterceptor)
	metricsController := controller.NewMetricController(metricService, config)

	routes := r.Group("/api")
	{
		routes.POST("/v1/short-it", urlShortenerController.GetShortenedURL)
		routes.GET("/v1/metrics", metricsController.GetMetrics)
	}

	r.GET("/:shortURL", redirectController.RedirectToFullUrl)

	return r
}
