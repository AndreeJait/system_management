package management_system

import (
	"context"
	"github.com/uptrace/bun"
	"system_management/internal/repository/db/management_system/history_money"
	"system_management/internal/repository/db/management_system/info_saving"
	"system_management/internal/repository/db/management_system/target_used_money"
	"system_management/internal/repository/db/management_system/used_money"
	"system_management/internal/repository/db/management_system/user"
)

type TransactionCallback func(ctx context.Context, repo ManagementSystemRepo) (interface{}, error)

type ManagementSystemRepo interface {
	Transaction(ctx context.Context, txFunc TransactionCallback) (out interface{}, err error)
	GetUserRepository() user.Repository
	GetUsedMoneyRepository() used_money.Repository
	GetInfoSavingRepository() info_saving.Repository
	GetTargetUsedMoneyRepository() target_used_money.Repository
	GetHistoryMoneyRepository() history_money.Repository
}

// DBI is a DB interface implemented by *DB and *Tx.
type DBI interface {
	NewValues(model interface{}) *bun.ValuesQuery
	NewSelect() *bun.SelectQuery
	NewInsert() *bun.InsertQuery
	NewUpdate() *bun.UpdateQuery
	NewDelete() *bun.DeleteQuery
	NewCreateTable() *bun.CreateTableQuery
	NewDropTable() *bun.DropTableQuery
	NewCreateIndex() *bun.CreateIndexQuery
	NewDropIndex() *bun.DropIndexQuery
	NewTruncateTable() *bun.TruncateTableQuery
	NewAddColumn() *bun.AddColumnQuery
	NewDropColumn() *bun.DropColumnQuery
	NewRaw(query string, args ...interface{}) *bun.RawQuery
}
