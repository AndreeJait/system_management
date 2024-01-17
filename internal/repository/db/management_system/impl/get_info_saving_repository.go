package impl

import (
	"system_management/commons/utils"
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/info_saving"
	"system_management/internal/repository/db/management_system/info_saving/impl"
)

func (m managementSystemRepo) GetInfoSavingRepository() info_saving.Repository {
	return impl.NewInfoSavingRepository(utils.ReturnFirstIfNotNil[management_system.DBI](m.dbExecutor, m.db))
}
