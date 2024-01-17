package middleware

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"system_management/commons/ierr"
	"system_management/commons/response"
	"system_management/commons/utils"
	"system_management/internal/shared/constant"
)

func MustLoggedIn(signingKey, encryptionKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := utils.VerifyTokenFromRequest(c, signingKey)
			if err != nil {
				return response.HTTPError(err, http.StatusUnauthorized, ierr.ErrUnauthorized.Code, ierr.ErrUnauthorized.Message)
			}

			claims := token.Claims.(jwt.MapClaims)

			var tokenType string
			if val, ok := claims["token_type"].(string); ok {
				tokenType = val
			}

			if tokenType != constant.TokenTypeAccess {
				return response.ErrUnauthorized(ierr.ErrUnauthorized)
			}

			claims["full_name"], err = utils.Decrypt(claims["full_name"].(string), encryptionKey)
			if err != nil {
				return response.ErrInternalServerError(err)
			}

			if claims["username"] != nil {
				claims["username"], err = utils.Decrypt(claims["username"].(string), encryptionKey)
				if err != nil {
					return response.ErrInternalServerError(err)
				}
			}

			ctx := context.WithValue(c.Request().Context(), utils.ContextKeyUser, token)
			r := c.Request().WithContext(ctx)
			c.SetRequest(r)

			if ctx.Value(utils.ContextKeyUser) != nil {
				return next(c)
			}

			return response.ErrUnauthorized(err)
		}
	}
}
