package impl

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"system_management/commons/utils"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
	"system_management/internal/shared/dto"
)

func buildQueryGetTargetUsedMoneyByUserID(query *bun.SelectQuery, param dto.GetTargetUsedMoneyByIDParam) {
	query.Where("user_id = ?", param.UserID)
	if param.Month != 0 {
		query.Where("month = ?", param.Month)
	}

	if param.Year != 0 {
		query.Where("year = ?", param.Year)
	}
}

func (r repository) GetTargetUsedMoneyByUserID(ctx context.Context, param dto.GetTargetUsedMoneyByIDParam) ([]model.TargetUsedMoney, error) {
	var result = make([]model.TargetUsedMoney, 0)
	query := r.db.NewSelect().Model(&result)

	buildQueryGetTargetUsedMoneyByUserID(query, param)
	utils.BuildPagination(param.LastID, param.Limit, param.OrderType, query)

	err := query.Scan(ctx)
	if err != nil {
		return result, errors.Wrap(err, constant.FailedToFetchData)
	}
	return result, nil
}
