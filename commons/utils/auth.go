package utils

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"strings"
	"system_management/internal/shared/dto"
)

type ContextUser string

const ContextKeyUser ContextUser = "user"

// VerifyTokenFromRequest verifies token from the request
func VerifyTokenFromRequest(c echo.Context, signingKey string) (*jwt.Token, error) {
	tokenString := extractToken(c)
	return VerifyToken(tokenString, signingKey)
}

// VerifyToken verifies the given token
func VerifyToken(tokenString, signingKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
}

func extractToken(c echo.Context) string {

	bearToken := c.Request().Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// GetLoggedInUser returns logged in user crm
func GetLoggedInUser(ctx context.Context) dto.UserToken {

	v := ctx.Value(ContextKeyUser)
	if v == nil {
		return dto.UserToken{}
	}
	loggedInUser := v.(*jwt.Token)

	claims := loggedInUser.Claims.(jwt.MapClaims)

	var id int64 = int64(claims["id"].(float64))

	var fullName string
	if val, ok := claims["full_name"].(string); ok {
		fullName = val
	}

	var username string
	if val, ok := claims["username"].(string); ok {
		username = val
	}

	var role string
	if val, ok := claims["role"].(string); ok {
		role = val
	}

	return dto.UserToken{
		ID:       id,
		FullName: fullName,
		Username: username,
		Role:     role,
	}

}
