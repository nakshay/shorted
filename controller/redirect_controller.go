package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/loggingUtil"
	urlShortenerService "shorted/service"
	shortedErr "shorted/shorted_error"
)

type RedirectController interface {
	RedirectToFullUrl(ctx *gin.Context)
}

type redirectController struct {
	urlShortenerService      urlShortenerService.URLShortenerService
	errorResponseInterceptor shortedErr.ErrorResponseInterceptor
}

func NewRedirectController(urlShortenerService urlShortenerService.URLShortenerService,
	errorResponseInterceptor shortedErr.ErrorResponseInterceptor) RedirectController {
	return redirectController{urlShortenerService: urlShortenerService, errorResponseInterceptor: errorResponseInterceptor}
}

func (c redirectController) RedirectToFullUrl(ctx *gin.Context) {
	logger := loggingUtil.GetLogger(ctx).WithFields("File", "redirectController").
		WithFields("Method", "RedirectToFullUrl")

	shortURL := ctx.Param("shortURL")
	logger.Infof("Received a short url %v", shortURL)
	fullUrl, err := c.urlShortenerService.GetFullURL(ctx, shortURL)
	if err != nil {
		logger.Errorf("Failed to get full url: %v", err)
		c.errorResponseInterceptor.HandleServiceErr(ctx, err)
		return
	}
	logger.Info("Full URL found, redirecting...")
	ctx.Redirect(http.StatusMovedPermanently, fullUrl)

}
