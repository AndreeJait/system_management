package impl

import (
	"system_management/config"
	"system_management/internal/repository/db/management_system"
	"system_management/internal/usecase/auth"
)

type useCase struct {
	repoManagementSystem management_system.ManagementSystemRepo
	cfg                  *config.Config
}

func NewAuthUseCase(cfg *config.Config, repoManagementSystem management_system.ManagementSystemRepo) auth.UseCase {
	return &useCase{
		repoManagementSystem: repoManagementSystem,
		cfg:                  cfg,
	}
}
