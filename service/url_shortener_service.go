package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shorted/configuration"
	"shorted/loggingUtil"
	"shorted/model"
	"shorted/storage"
	"shorted/util"
)

//go:generate mockgen -source=./url_shortener_service.go -destination=../mocks/mock_url_shortener_service.go -package=mocks

type URLShortenerService interface {
	GetShortenedURL(ctx *gin.Context, url string) model.ShortUrlResponse
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

func (service urlShortenerService) GetShortenedURL(ctx *gin.Context, fullURL string) model.ShortUrlResponse {
	logger := loggingUtil.GetLogger(ctx).
		WithFields("File", "urlShortenerService").
		WithFields("Method", "GetShortenedURL")

	logger.Debugf("Checking if short url exist for full url %v", fullURL)
	shortUrl, found := service.store.IsShortURLExists(fullURL)
	if found {
		logger.Info("Short url found")
		return service.buildResponse(shortUrl)
	}
	logger.Info("Short URL not found, generating a new one")
	shortUrl = service.randomStringGenerator.GenerateRandomString(service.configData.RandomCharacterLength)
	service.store.SaveShortURL(shortUrl, fullURL)
	return service.buildResponse(shortUrl)

}

func (service urlShortenerService) buildResponse(shortUrl string) model.ShortUrlResponse {
	shortUrl = fmt.Sprintf("%v/%v", service.configData.ServiceDomain, shortUrl)
	return model.ShortUrlResponse{ShortUrl: shortUrl}
}
