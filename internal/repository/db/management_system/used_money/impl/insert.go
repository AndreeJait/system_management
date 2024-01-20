package impl

import (
	"context"
	"github.com/pkg/errors"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
)

func (r repository) Insert(ctx context.Context, usedMoney model.UsedMoney) (int64, error) {
	sqlResult, err := r.db.NewInsert().Model(&usedMoney).
		Exec(ctx)
	if err != nil {
		return 0, errors.Wrap(err, constant.FailedToInsertData)
	}

	lastInsertID, err := sqlResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, err
}
