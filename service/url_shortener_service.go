package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shorted/configuration"
	"shorted/loggingUtil"
	"shorted/model"
	shortedErr "shorted/shorted_error"
	"shorted/storage"
	"shorted/util"
)

//go:generate mockgen -source=./url_shortener_service.go -destination=../mocks/mock_url_shortener_service.go -package=mocks

type URLShortenerService interface {
	GetShortenedURL(ctx *gin.Context, url string) (model.ShortUrlResponse, *shortedErr.ShortedError)
	GetFullURL(ctx *gin.Context, url string) (string, *shortedErr.ShortedError)
}

type urlShortenerService struct {
	store                 storage.Store
	configData            *configuration.ConfigData
	randomStringGenerator util.RandomStringGenerator
}

func NewURLShortenerService(store storage.Store,
	configData *configuration.ConfigData,
	randomStringGenerator util.RandomStringGenerator) URLShortenerService {
	return urlShortenerService{store: store, configData: configData, randomStringGenerator: randomStringGenerator}
}

func (service urlShortenerService) GetShortenedURL(ctx *gin.Context, fullURL string) (model.ShortUrlResponse, *shortedErr.ShortedError) {
	logger := loggingUtil.GetLogger(ctx).
		WithFields("File", "urlShortenerService").
		WithFields("Method", "GetShortenedURL")

	logger.Debugf("Checking if short url exist for full url %v", fullURL)
	shortUrl, found := service.store.IsShortURLExists(fullURL)
	if found {
		logger.Info("Short url found")
		return service.buildResponse(shortUrl), nil
	}

	shortUrl = service.randomStringGenerator.GenerateRandomString(service.configData.RandomCharacterLength)
	err := service.store.SaveShortURL(shortUrl, fullURL)
	if err != nil {
		logger.Error("Error while saving short url ", err)
		return model.ShortUrlResponse{}, shortedErr.InternalServerErrorWithMessage(err.Error())
	}
	return service.buildResponse(shortUrl), nil

}

func (service urlShortenerService) buildResponse(shortUrl string) model.ShortUrlResponse {
	shortUrl = fmt.Sprintf("%v/%v", service.configData.ServiceDomain, shortUrl)
	return model.ShortUrlResponse{ShortUrl: shortUrl}
}

func (service urlShortenerService) GetFullURL(ctx *gin.Context, shortUrl string) (string, *shortedErr.ShortedError) {
	logger := loggingUtil.GetLogger(ctx).
		WithFields("File", "urlShortenerService").
		WithFields("Method", "GetFullURL")

	logger.Infof("Started finding full url for a short url %v", shortUrl)
	fullURL, found := service.store.FindFullURL(shortUrl)
	if !found {
		return "", shortedErr.URLNotFoundErr
	}

	return fullURL, nil
}
