package impl

import (
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/user"
)

type repository struct {
	db management_system.DBI
}

func NewUserRepository(db management_system.DBI) user.Repository {
	return &repository{
		db: db,
	}
}
