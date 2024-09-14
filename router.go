package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/configuration"
	"shorted/controller"
	urlShortenerService "shorted/service"
	shortedErr "shorted/shorted_error"
	"shorted/storage"
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

	shortenerService := urlShortenerService.NewURLShortenerService(dbStore, config)
	urlShortenerController := controller.NewURLShortenerController(shortenerService, errorResponseInterceptor)

	redirectController := controller.NewRedirectController(shortenerService, errorResponseInterceptor)

	routes := r.Group("/api")
	{
		routes.POST("/v1/short-it", urlShortenerController.GetShortenedURL)
	}

	r.GET("/:shortURL", redirectController.RedirectToFullUrl)

	return r
}
