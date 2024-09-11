package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	shortContrller "shorted/controller"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	// Health endpoint
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// initialization

	urlShortnerController := shortContrller.NewURLShortenerController()

	routes := r.Group("/api")
	{
		routes.POST("/v1/short-it", urlShortnerController.GetShortenedURL)
	}

	return r
}
