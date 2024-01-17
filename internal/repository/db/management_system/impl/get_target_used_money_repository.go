package impl

import (
	"system_management/internal/repository/db/management_system/target_used_money"
	"system_management/internal/repository/db/management_system/target_used_money/impl"
)

func (m managementSystemRepo) GetTargetUsedMoneyRepository() target_used_money.Repository {
	db := m.dbExecutor
	if m.dbExecutor == nil {
		db = m.db
	}
	return impl.NewTargetUsedMoneyRepo(db)
}
