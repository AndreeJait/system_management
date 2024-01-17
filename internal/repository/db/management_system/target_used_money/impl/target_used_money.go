package impl

import (
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/target_used_money"
)

type repository struct {
	db management_system.DBI
}

func NewTargetUsedMoneyRepo(db management_system.DBI) target_used_money.Repository {
	return &repository{
		db: db,
	}
}
