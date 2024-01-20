package user

import (
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"system_management/commons/middleware"
	"system_management/config"
	"system_management/internal/usecase/user"
)

type handler struct {
	route  *echo.Group
	db     *bun.DB
	cfg    *config.Config
	ucUser user.UseCase
}

func RegisterUserApi(route *echo.Group, db *bun.DB, cfg *config.Config) {

	//repoManagementSystem := repoImpl.NewManagementSystemRepo(db)
	//ucUser := authUcImpl.NewAuthUseCase(cfg, repoManagementSystem)

	h := handler{
		db:    db,
		cfg:   cfg,
		route: route,
	}

	route.Use(middleware.MustLoggedIn(cfg.Jwt.SigningKey, cfg.Jwt.EncryptionKey))
	route.GET("/me", h.infoUser)
}
