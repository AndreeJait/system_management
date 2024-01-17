package auth

import (
	"github.com/labstack/echo/v4"
	"system_management/commons/response"
	"system_management/commons/utils"
)

func (h handler) infoUser(c echo.Context) error {
	ctx := c.Request().Context()

	user := utils.GetLoggedInUser(ctx)

	return response.SuccessOK(c, user, "success login")
}
