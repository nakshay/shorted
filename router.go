package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/configuration"
	shortContrller "shorted/controller"
	shortedErr "shorted/error"
	urlShortenerService "shorted/service"
	"shorted/store"
)

func setupRouter(config *configuration.ConfigData) *gin.Engine {

	r := gin.Default()
	// Health endpoint
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// initialization
	s := store.Init()

	errorResponseInterceptor := shortedErr.NewErrorResponseInterceptor()

	shortenerService := urlShortenerService.NewURLShortenerService(s, config)
	urlShortenerController := shortContrller.NewURLShortenerController(shortenerService, errorResponseInterceptor)

	routes := r.Group("/api")
	{
		routes.POST("/v1/short-it", urlShortenerController.GetShortenedURL)
	}

	return r
}
