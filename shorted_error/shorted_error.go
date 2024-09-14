package shorted_error

import "fmt"

type ShortedError struct {
	httpStatusCode int
	errorCode      string
	errorMessage   string
}

func (err ShortedError) Error() string {
	return fmt.Sprintf("Error code : %v, Error Message: %v", err.errorCode, err.errorMessage)
}

var (
	URLNotFoundErr = &ShortedError{
		httpStatusCode: 404,
		errorCode:      "URL_NOT_FOUND",
		errorMessage:   "Invalid short URL OR URL expired",
	}
)

func BadRequestError() *ShortedError {
	return &ShortedError{httpStatusCode: 400, errorCode: "SHORTED_BAD_REQUEST", errorMessage: "Bad Request"}
}
func InternalServerError() *ShortedError {
	return &ShortedError{
		httpStatusCode: 500,
		errorCode:      "SHORTED_INTERNAL_SERVER_ERROR",
		errorMessage:   "Something went wrong",
	}
}

func BadRequestErrorWithErrorMessage(errorMessage string) *ShortedError {
	badRequestErr := BadRequestError()
	badRequestErr.errorMessage = errorMessage
	return badRequestErr
}

func InternalServerErrorWithMessage(errorMessage string) *ShortedError {
	internalServerErr := InternalServerError()
	internalServerErr.errorMessage = errorMessage
	return internalServerErr
}
