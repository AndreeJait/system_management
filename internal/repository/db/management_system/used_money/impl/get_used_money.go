package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"system_management/commons/utils"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
	"system_management/internal/shared/dto"
)

func buildQueryGetUsedMoney(query *bun.SelectQuery, param dto.GetUsedMoneyParam) {
	query.Where("user_id = ?", param.UserID)

	if param.Day != 0 {
		query.Where("EXTRACT(DAY FROM created_at) = ?", param.Day)
	}

	if param.Month != 0 {
		query.Where("EXTRACT(MONTH FROM created_at) = ?", param.Month)
	}

	if param.Month != 0 {
		query.Where("EXTRACT(YEAR FROM created_at) = ?", param.Year)
	}

	query.Where("purpose ILIKE ? OR description ILIKE ?",
		fmt.Sprintf("%%%s%%", param.Key),
		fmt.Sprintf("%%%s%%", param.Key))
}

func (r repository) GetUsedMoney(ctx context.Context, param dto.GetUsedMoneyParam) ([]model.UsedMoney, error) {
	var usedMoneys = make([]model.UsedMoney, 0)
	query := r.db.NewSelect().Model(&usedMoneys)

	buildQueryGetUsedMoney(query, param)
	utils.BuildPagination(param.LastID, param.Limit, param.OrderType, query)
	err := query.Scan(ctx)
	if err != nil {
		return usedMoneys, errors.Wrap(err, constant.FailedToFetchData)
	}
	return usedMoneys, nil
}
