package impl

import (
	"context"
	"system_management/internal/shared/dto"
)

func (u useCase) Login(ctx context.Context, param dto.LoginParam) (dto.LoginResponse, error) {
	return dto.LoginResponse{
		AccessToken: "Andree Accessed",
	}, nil
}
