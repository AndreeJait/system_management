package main

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"system_management/commons/ierr"
	"system_management/commons/response"
	"system_management/config"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// CustomHTTPErrorHandler sets error response for different type of errors and logs
func CustomHTTPErrorHandler(cfg *config.Config) echo.HTTPErrorHandler {

	return func(err error, c echo.Context) {

		// var internalErr error
		if _, ok := err.(response.ErrorResponse); !ok {
			err = response.ErrorResponse{
				Success:   false,
				HTTPCode:  http.StatusInternalServerError,
				Message:   ierr.ErrInternal.Message,
				ErrorCode: ierr.ErrInternal.Code,
				Internal:  err,
			}
		}

		// Get internal error
		internalErr := errors.Cause(err.(response.ErrorResponse).Internal)

		// handles resource not found errors
		if errors.Is(internalErr, echo.ErrNotFound) {
			err = response.HTTPError(internalErr, http.StatusNotFound, ierr.ErrNotFound.Code, "requested endpoint is not registered")
		}

		// Handles validation error
		if errors.As(internalErr, &validation.Errors{}) || errors.As(internalErr, &validation.ErrorObject{}) {
			err = response.HTTPError(internalErr, http.StatusBadRequest, ierr.ErrorBadRequest.Code, internalErr.Error())
		}

		var resp response.ErrorResponse
		if res, ok := err.(response.ErrorResponse); ok {
			resp = res
		} else {
			resp = response.ErrInternalServerError(err)
		}

		if sterr, ok := resp.Internal.(stackTracer); ok {
			if cfg.Server.Mode == config.Development {
				fmt.Printf("%+v\n", sterr.StackTrace())
			}
		}

		logrus.WithContext(c.Request().Context()).Error(resp.Internal)

		err = c.JSON(resp.HTTPCode, resp)
		if err != nil {
			logrus.Error(err)
		}
	}
}
