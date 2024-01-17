package impl

import (
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/used_money"
)

type repository struct {
	db management_system.DBI
}

func NewUsedMoneyRepo(db management_system.DBI) used_money.Repository {
	return &repository{
		db: db,
	}
}
