package impl

import (
	"context"
	"github.com/uptrace/bun"
	"system_management/internal/repository/db/management_system"
)

func (m managementSystemRepo) Transaction(ctx context.Context, txFunc management_system.TransactionCallback) (out interface{}, err error) {

	var tx bun.Tx

	tx, err = m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			rErr := tx.Rollback() // err is non-nil; don't change it
			if rErr != nil {
				err = rErr
			}
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	registry := managementSystemRepo{
		db:         m.db,
		dbExecutor: tx,
	}
	out, err = txFunc(ctx, registry)

	return
}
