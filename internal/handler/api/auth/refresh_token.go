package auth

import (
	"github.com/labstack/echo/v4"
	"system_management/internal/shared/dto"
)

func (h handler) refreshToken(c echo.Context) error {
	var param dto.RefreshTokenParam
	if err := c.Bind(&param); err != nil {
		return err
	}
	return nil
}
