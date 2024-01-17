package impl

import (
	"system_management/commons/utils"
	"system_management/internal/repository/db/management_system"
	"system_management/internal/repository/db/management_system/user"
	"system_management/internal/repository/db/management_system/user/impl"
)

func (m managementSystemRepo) GetUserRepository() user.Repository {
	return impl.NewUserRepository(utils.ReturnFirstIfNotNil[management_system.DBI](m.dbExecutor, m.db))
}
