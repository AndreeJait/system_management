package impl

import (
	"context"
	"database/sql"
	errors2 "errors"
	"github.com/pkg/errors"
	"system_management/commons/ierr"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
)

func (r repository) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	err = r.db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		if errors2.Is(err, sql.ErrNoRows) {
			return user, ierr.ErrNotFound
		}
		return user, errors.Wrap(err, constant.FailedToFetchData)
	}
	return
}
