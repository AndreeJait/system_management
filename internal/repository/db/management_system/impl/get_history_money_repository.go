package impl

import (
	"system_management/internal/repository/db/management_system/history_money"
	"system_management/internal/repository/db/management_system/history_money/impl"
)

func (m managementSystemRepo) GetHistoryMoneyRepository() history_money.Repository {
	db := m.dbExecutor
	if m.dbExecutor == nil {
		db = m.db
	}
	return impl.NewHistoryMoneyRepository(db)
}
