package impl

import (
	"system_management/commons/utils"
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/used_money"
	"system_management/internal/repository/db/management_system/used_money/impl"
)

func (m managementSystemRepo) GetUsedMoneyRepository() used_money.Repository {
	return impl.NewUsedMoneyRepo(utils.ReturnFirstIfNotNil[management_system.DBI](m.dbExecutor, m.db))
}
