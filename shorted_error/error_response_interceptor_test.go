package shorted_error

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ErrorResponseInterceptorTest struct {
	suite.Suite
	mockCtrl                 *gomock.Controller
	recorder                 *httptest.ResponseRecorder
	context                  *gin.Context
	errorResponseInterceptor ErrorResponseInterceptor
}

func TestNewErrorResponseInterceptor(t *testing.T) {
	suite.Run(t, new(ErrorResponseInterceptorTest))
}

func (suite *ErrorResponseInterceptorTest) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.context.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	suite.errorResponseInterceptor = NewErrorResponseInterceptor()

}

func (suite *ErrorResponseInterceptorTest) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *ErrorResponseInterceptorTest) TestHandleBadRequest() {
	errorMsg := "Bad Request"
	suite.errorResponseInterceptor.HandleBadRequest(suite.context, errors.New("bad request"))

	suite.Equal(errorMsg, "Bad Request")
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)

}

func (suite *ErrorResponseInterceptorTest) TestHandleServiceErr() {
	expectedErr := InternalServerError()
	suite.errorResponseInterceptor.HandleServiceErr(suite.context, InternalServerError())
	bytes, _ := io.ReadAll(suite.recorder.Body)
	var response string
	json.Unmarshal(bytes, &response)

	suite.Equal(expectedErr.errorMessage, response)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)

}
