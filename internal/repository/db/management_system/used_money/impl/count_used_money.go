package impl

import (
	"context"
	"github.com/pkg/errors"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
	"system_management/internal/shared/dto"
)

func (r repository) CountUsedMoney(ctx context.Context, param dto.GetUsedMoneyParam) (int64, error) {
	var countModel = model.CountModel{}
	query := r.db.NewSelect().
		Model((*model.UsedMoney)(nil))

	query.ColumnExpr("count(*) as total_data")

	buildQueryGetUsedMoney(query, param)

	err := query.Scan(ctx, &countModel)
	if err != nil {
		return countModel.TotalData, errors.Wrap(err, constant.FailedToCountData)
	}

	return countModel.TotalData, nil
}
