package shorted_error

import (
	"github.com/gin-gonic/gin"
	"shorted/loggingUtil"
)

//go:generate mockgen -source=./error_response_interceptor.go -destination=../mocks/mock_error_response_interceptor.go -package=mocks

type ErrorResponseInterceptor interface {
	HandleBadRequest(ctx *gin.Context, bindErr error)
	HandleServiceErr(ctx *gin.Context, bindErr *ShortedError)
}

type errorResponseInterceptor struct {
}

func NewErrorResponseInterceptor() ErrorResponseInterceptor {
	return errorResponseInterceptor{}
}

func (errorResponseInterceptor) HandleBadRequest(ctx *gin.Context, bindErr error) {
	log := loggingUtil.GetLogger(ctx).WithFields("File", "errorResponseInterceptor").WithFields("Method", "HandleBadRequest")
	log.Errorf("Bad request received %v", bindErr)
	err := BadRequestError()
	ctx.JSON(err.httpStatusCode, err.errorMessage)
}

func (errorResponseInterceptor) HandleServiceErr(ctx *gin.Context, serviceErr *ShortedError) {
	log := loggingUtil.GetLogger(ctx).WithFields("File", "errorResponseInterceptor").WithFields("Method", "HandleServiceErr")
	log.Errorf("Service shorted_error %v", serviceErr)
	ctx.JSON(serviceErr.httpStatusCode, serviceErr.errorMessage)

}
