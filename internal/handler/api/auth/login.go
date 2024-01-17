package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"system_management/commons/ierr"
	"system_management/commons/response"
	"system_management/internal/shared/dto"
)

func (h handler) login(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.LoginParam

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.ucAuth.Login(ctx, req)
	if err != nil {
		switch errors.Cause(err) {
		case ierr.ErrInvalidCred, ierr.ErrUserNotActive:
			return response.ErrBadRequest(err)
		default:
			return err
		}
	}

	return response.SuccessOK(c, res, "success login")
}
