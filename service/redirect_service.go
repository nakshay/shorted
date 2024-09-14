package service

import (
	"github.com/gin-gonic/gin"
	"shorted/loggingUtil"
	shortedErr "shorted/shorted_error"
	"shorted/storage"
)

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
		return "", shortedErr.URLNotFoundErr
	}

	return fullURL, nil
}
