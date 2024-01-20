package impl

import (
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/history_money"
)

type repository struct {
	db management_system.DBI
}

func NewHistoryMoneyRepository(db management_system.DBI) history_money.Repository {
	return &repository{
		db: db,
	}
}
