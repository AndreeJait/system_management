package impl

import (
	"system_management/internal/repository/db/management_system/info_saving"
	"system_management/internal/repository/db/management_system/info_saving/impl"
)

func (m managementSystemRepo) GetInfoSavingRepository() info_saving.Repository {
	db := m.dbExecutor
	if m.dbExecutor == nil {
		db = m.db
	}
	return impl.NewInfoSavingRepository(db)
}
