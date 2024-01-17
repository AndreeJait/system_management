package user

import (
	"context"
	"system_management/internal/model"
)

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (user model.User, err error)
}
