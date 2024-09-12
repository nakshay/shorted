package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"shorted/configuration"
	shortedErr "shorted/error"
	"shorted/logger"
	"shorted/model"
	"shorted/store"
	"time"
)

type URLShortenerService interface {
	GetShortenedURL(ctx *gin.Context, url string) (model.ShortUrlResponse, *shortedErr.ShortedError)
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
	log := logger.New(ctx).
		WithFields("Service", "urlShortenerService").
		WithFields("Method", "GetShortenedURL")

	shortUrl := generateRandomString(10) //make length configurable
	err := service.store.Save(shortUrl, fullURL)
	if err != nil {
		log.Error("Error while saving short url ", err)
		return model.ShortUrlResponse{}, shortedErr.InternalServerErrorWithMessage(err.Error())
	}
	shortUrl = fmt.Sprintf("%v/%v", service.configData.ServiceDomain, shortUrl)

	return model.ShortUrlResponse{ShortUrl: shortUrl}, nil

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
