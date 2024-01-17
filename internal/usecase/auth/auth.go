package auth

import (
	"context"
	"system_management/internal/shared/dto"
)

type UseCase interface {
	Login(ctx context.Context, param dto.LoginParam) (dto.LoginResponse, error)
}
