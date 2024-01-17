package impl

import (
	"system_management/internal/repository/db/management_system/used_money"
	"system_management/internal/repository/db/management_system/used_money/impl"
)

func (m managementSystemRepo) GetUsedMoneyRepository() used_money.Repository {
	db := m.dbExecutor
	if m.dbExecutor == nil {
		db = m.db
	}
	return impl.NewUsedMoneyRepo(db)
}
