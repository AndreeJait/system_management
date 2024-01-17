package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"system_management/commons/ierr"

	"github.com/pkg/errors"
)

// Data is an alias for map
type Data map[string]interface{}

func buildResponseMsg(defaultMsg string, msg ...string) string {
	if len(msg) == 0 {
		return defaultMsg
	}
	var response string
	for i, item := range msg {
		response += item
		if len(msg)-1 != i {
			response += ", "
		}
	}
	return response
}

// Response struct
type Response struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// Success responses with JSON format-responseMsg
func Success(c echo.Context, code int, data interface{}, msg ...string) error {

	responseMsg := buildResponseMsg("Success", msg...)

	if data == nil {
		data = map[string]interface{}{}
	}

	res := Response{
		Success: true,
		Message: responseMsg,
		Data:    data,
	}
	return c.JSON(code, res)
}

// SuccessOK returns code 200
func SuccessOK(c echo.Context, data interface{}, msg ...string) error {
	return Success(c, http.StatusOK, data, msg...)
}

// SuccessCreated returns code 201
func SuccessCreated(c echo.Context, data interface{}, msg ...string) error {
	return Success(c, http.StatusCreated, data, msg...)
}

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	HTTPCode  int    `json:"-"`
	Success   bool   `json:"success" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code,omitempty"`
	RequestID string `json:"request_id"`
	Internal  error  `json:"-"`
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by CustomHTTPErrorHandler
func (e ErrorResponse) StatusCode() int {
	return e.HTTPCode
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrInternalServerError creates a new error response representing an internal server error (HTTP 500)
func ErrInternalServerError(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	var val ierr.Error = ierr.ErrInternal
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusUnauthorized,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

func ErrUnauthorized(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	var val ierr.Error = ierr.ErrUnauthorized
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusUnauthorized,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrForbidden creates a new error response representing an authorization failure (HTTP 403)
func ErrForbidden(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	var val ierr.Error = ierr.ErrForbidden
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusForbidden,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrNotFound creates a new error response representing a resource not found (HTTP 404)
func ErrNotFound(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	var val ierr.Error = ierr.ErrNotFound
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusNotFound,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrBadRequest creates a new error response representing a bad request (HTTP 400)
func ErrBadRequest(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	var val ierr.Error = ierr.ErrorBadRequest
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusBadRequest,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

func HTTPError(err error, statusCode int, errorCode string, message string) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	return ErrorResponse{
		HTTPCode:  statusCode,
		Message:   message,
		ErrorCode: errorCode,
		Internal:  err,
	}
}
