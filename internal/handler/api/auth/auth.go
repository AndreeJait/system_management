package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"system_management/config"
	repoImpl "system_management/internal/repository/db/management_system/impl"
	"system_management/internal/usecase/auth"
	authUcImpl "system_management/internal/usecase/auth/impl"
)

type handler struct {
	route  *echo.Group
	db     *bun.DB
	cfg    *config.Config
	ucAuth auth.UseCase
}

func RegisterAuthApi(route *echo.Group, db *bun.DB, cfg *config.Config) {

	repoManagementSystem := repoImpl.NewManagementSystemRepo(db)
	ucAuth := authUcImpl.NewAuthUseCase(cfg, repoManagementSystem)

	h := handler{
		db:     db,
		cfg:    cfg,
		route:  route,
		ucAuth: ucAuth,
	}

	route.POST("/login", h.login)
}
