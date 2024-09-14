package service

import (
	"github.com/gin-gonic/gin"
	"shorted/loggingUtil"
	shortedErr "shorted/shorted_error"
	"shorted/storage"
)

//go:generate mockgen -source=./redirect_service.go -destination=../mocks/mock_redirect_service.go -package=mocks

type RedirectService interface {
	GetFullURL(ctx *gin.Context, url string) (string, *shortedErr.ShortedError)
}

type redirectService struct {
	store storage.Store
}

func NewRedirectService(store storage.Store) RedirectService {
	return redirectService{store: store}
}

func (service redirectService) GetFullURL(ctx *gin.Context, shortUrl string) (string, *shortedErr.ShortedError) {
	logger := loggingUtil.GetLogger(ctx).
		WithFields("File", "urlShortenerService").
		WithFields("Method", "GetFullURL")

	logger.Infof("Started finding full url for a short url %v", shortUrl)
	fullURL, found := service.store.FindFullURL(shortUrl)
	if !found {
		err := shortedErr.URLNotFoundErr
		logger.Errorf("URL not found. Error %v", err)
		return "", err
	}
	logger.Info("URL found for given short URL")
	return fullURL, nil
}
