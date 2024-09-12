package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	shortedErr "shorted/error"
	"shorted/logger"
	"shorted/model"
	"shorted/service"
)

type URLShortenerController interface {
	GetShortenedURL(ctx *gin.Context)
}

type urlShortenerController struct {
	service                  service.URLShortenerService
	errorResponseInterceptor shortedErr.ErrorResponseInterceptor
}

func NewURLShortenerController(service service.URLShortenerService,
	errorResponseInterceptor shortedErr.ErrorResponseInterceptor) URLShortenerController {
	return urlShortenerController{service: service, errorResponseInterceptor: errorResponseInterceptor}
}

func (controller urlShortenerController) GetShortenedURL(ctx *gin.Context) {
	log := logger.New(ctx).WithFields("Controller", "urlShortenerController").WithFields("Method", "GetShortenedURL")
	log.Info("URL shortening started")
	var request model.ShortURLRequest
	err := ctx.Bind(&request)
	if err != nil {
		log.Errorf("Error while binding request %v", err)
		controller.errorResponseInterceptor.HandleBadRequest(ctx, err)
		return
	}
	shortenedURL, serviceErr := controller.service.GetShortenedURL(ctx, request.URL)
	if serviceErr != nil {
		log.Errorf("Error recieved from service: Error %v", serviceErr)
		controller.errorResponseInterceptor.HandleServiceErr(ctx, serviceErr)
		return
	}
	ctx.JSON(http.StatusOK, shortenedURL)
	log.Info("URL shortening request received")
}
