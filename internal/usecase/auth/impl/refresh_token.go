package impl

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"system_management/commons/utils"
	"system_management/internal/shared/dto"
)

func (u useCase) RefreshToken(ctx context.Context, param dto.RefreshTokenParam) (dto.LoginResponse, error) {
	var (
		result dto.LoginResponse
		err    error
	)
	if err := param.Validate(); err != nil {
		return result, err
	}

	token, err := utils.VerifyToken(param.RefreshToken, u.cfg.Jwt.SigningKey)
	if err != nil {
		return result, err
	}
	claims := token.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	return result, err
}
