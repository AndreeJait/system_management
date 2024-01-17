package impl

import (
	"system_management/commons/utils"
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/target_used_money"
	"system_management/internal/repository/db/management_system/target_used_money/impl"
)

func (m managementSystemRepo) GetTargetUsedMoneyRepository() target_used_money.Repository {
	return impl.NewTargetUsedMoneyRepo(
		utils.ReturnFirstIfNotNil[management_system.DBI](m.dbExecutor, m.db))
}
