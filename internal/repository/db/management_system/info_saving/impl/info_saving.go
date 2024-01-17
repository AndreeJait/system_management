package impl

import (
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/info_saving"
)

type repository struct {
	db management_system.DBI
}

func NewInfoSavingRepository(db management_system.DBI) info_saving.Repository {
	return &repository{
		db: db,
	}
}
