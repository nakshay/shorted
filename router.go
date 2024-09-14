package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/configuration"
	"shorted/controller"
	urlShortenerService "shorted/service"
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
	shortenerService := urlShortenerService.NewURLShortenerService(dbStore, config, randomStringGenerator)
	urlShortenerController := controller.NewURLShortenerController(shortenerService, errorResponseInterceptor)

	redirectController := controller.NewRedirectController(shortenerService, errorResponseInterceptor)

	routes := r.Group("/api")
	{
		routes.POST("/v1/short-it", urlShortenerController.GetShortenedURL)
	}

	r.GET("/:shortURL", redirectController.RedirectToFullUrl)

	return r
}
