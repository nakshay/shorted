package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"net/url"
	"shorted/loggingUtil"
	"shorted/model"
	"shorted/service"
	shortedErr "shorted/shorted_error"
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
	log := loggingUtil.GetLogger(ctx).WithFields("file", "urlShortenerController").WithFields("Method", "GetShortenedURL")
	log.Info("URL shortening started")
	var request model.ShortURLRequest
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		log.Errorf("Error while binding request %v", err)
		controller.errorResponseInterceptor.HandleBadRequest(ctx, err)
		return
	}
	if _, parsedErr := url.ParseRequestURI(request.URL); parsedErr != nil {
		log.Errorf("Error while parsing url %v, %v", request.URL, parsedErr)
		err := shortedErr.BadRequestErrorWithErrorMessage(parsedErr.Error())
		controller.errorResponseInterceptor.HandleBadRequest(ctx, err)
		return
	}
	shortenedURL := controller.service.GetShortenedURL(ctx, request.URL)
	log.Info("URL shortening request received")
	ctx.JSON(http.StatusOK, shortenedURL)
}
