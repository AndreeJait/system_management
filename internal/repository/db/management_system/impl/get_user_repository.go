package impl

import (
	"system_management/internal/repository/db/management_system/user"
	"system_management/internal/repository/db/management_system/user/impl"
)

func (m managementSystemRepo) GetUserRepository() user.Repository {
	db := m.dbExecutor
	if m.dbExecutor == nil {
		db = m.db
	}
	return impl.NewUserRepository(db)
}
