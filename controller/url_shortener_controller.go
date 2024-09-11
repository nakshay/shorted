package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/logger"
	"shorted/model"
)

type URLShortenerController interface {
	GetShortenedURL(ctx *gin.Context)
}

type urlShortenerController struct {
}

func NewURLShortenerController() URLShortenerController {
	return urlShortenerController{}
}

func (c urlShortenerController) GetShortenedURL(ctx *gin.Context) {
	log := logger.New(ctx).WithFields("Controller", "urlShortenerController").WithFields("Method", "GetShortenedURL")
	log.Info("URL shortening started")
	var request model.ShortURLRequest
	err := ctx.Bind(&request)
	if err != nil {
		log.Error("Error while binding request ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("URL shortening request received")
}
