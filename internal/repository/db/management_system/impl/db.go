package impl

import (
	"github.com/uptrace/bun"
	"system_management/internal/repository/db/management_system"
)

type managementSystemRepo struct {
	db         *bun.DB
	dbExecutor management_system.DBI
}

func NewManagementSystemRepo(db *bun.DB) management_system.ManagementSystemRepo {
	return &managementSystemRepo{
		db: db,
	}
}
