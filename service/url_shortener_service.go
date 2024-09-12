package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"shorted/configuration"
	shortedErr "shorted/error"
	"shorted/loggingUtil"
	"shorted/model"
	"shorted/store"
	"time"
)

type URLShortenerService interface {
	GetShortenedURL(ctx *gin.Context, url string) (model.ShortUrlResponse, *shortedErr.ShortedError)
	GetFullURL(ctx *gin.Context, url string) (string, *shortedErr.ShortedError)
}

type urlShortenerService struct {
	store      store.Store
	configData *configuration.ConfigData
}

func NewURLShortenerService(store store.Store,
	configData *configuration.ConfigData) URLShortenerService {
	return urlShortenerService{store: store, configData: configData}
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

	shortUrl = generateRandomString(10) //make length configurable
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

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(length)
		result = append(result, charset[randomIndex])
	}

	return string(result)

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
