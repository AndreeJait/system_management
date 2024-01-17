package api

import (
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"system_management/config"
	"system_management/internal/handler/api/auth"
)

type Handler struct {
	Config *config.Config
	Route  *echo.Group
	DB     *bun.DB
}

func (h Handler) RegisterHandler() {
	// register Auth
	auth.RegisterAuthApi(h.Route.Group("/auth"), h.DB, h.Config)
}
