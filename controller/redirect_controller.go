package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorted/loggingUtil"
	"shorted/service"
	shortedErr "shorted/shorted_error"
	"strings"
)

type RedirectController interface {
	RedirectToFullUrl(ctx *gin.Context)
}

type redirectController struct {
	redirectService          service.RedirectService
	errorResponseInterceptor shortedErr.ErrorResponseInterceptor
}

func NewRedirectController(urlShortenerService service.RedirectService,
	errorResponseInterceptor shortedErr.ErrorResponseInterceptor) RedirectController {
	return redirectController{redirectService: urlShortenerService, errorResponseInterceptor: errorResponseInterceptor}
}

func (c redirectController) RedirectToFullUrl(ctx *gin.Context) {
	logger := loggingUtil.GetLogger(ctx).WithFields("File", "redirectController").
		WithFields("Method", "RedirectToFullUrl")

	shortURL := ctx.Param("shortURL")
	logger.Infof("Received a short url %v", shortURL)
	if len(strings.TrimSpace(shortURL)) == 0 {
		logger.Infof("Short url is missing in query params")
		err := shortedErr.BadRequestErrorWithErrorMessage("short url is missing")
		c.errorResponseInterceptor.HandleBadRequest(ctx, err)
		return
	}
	logger.Info("Short url present, getting full URL")
	fullUrl, err := c.redirectService.GetFullURL(ctx, shortURL)
	if err != nil {
		logger.Errorf("Failed to get full url: %v", err)
		c.errorResponseInterceptor.HandleServiceErr(ctx, err)
		return
	}
	logger.Info("Full URL found, redirecting...")
	ctx.Redirect(http.StatusMovedPermanently, fullUrl)

}
