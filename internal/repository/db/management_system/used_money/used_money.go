package used_money

import (
	"context"
	"system_management/internal/model"
	"system_management/internal/shared/dto"
)

type Repository interface {
	Insert(ctx context.Context, usedMoney model.UsedMoney) (int64, error)

	GetUsedMoney(ctx context.Context, param dto.GetUsedMoneyParam) ([]model.UsedMoney, error)
	CountUsedMoney(ctx context.Context, param dto.GetUsedMoneyParam) (int64, error)
}
