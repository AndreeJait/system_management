package target_used_money

import (
	"context"
	"system_management/internal/model"
	"system_management/internal/shared/dto"
)

type Repository interface {
	GetTargetUsedMoneyByUserID(ctx context.Context, param dto.GetTargetUsedMoneyByIDParam) ([]model.TargetUsedMoney, error)
	CountTargetUsedMoneyByUserID(ctx context.Context, param dto.GetTargetUsedMoneyByIDParam) (int64, error)
	Insert(ctx context.Context, targetUsedMoney model.TargetUsedMoney) (int64, error)
}
