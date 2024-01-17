package impl

import (
	"context"
	"github.com/pkg/errors"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
)

func (r repository) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	err = r.db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return user, errors.Wrap(err, constant.FailedToFetchData)
	}
	return
}
