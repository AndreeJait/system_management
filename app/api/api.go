package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"net/http"
	"os"
	"os/signal"
	"system_management/commons/ierr"
	customMiddleware "system_management/commons/middleware"
	"system_management/commons/response"
	"system_management/commons/utils"
	"system_management/config"
	"system_management/internal/handler/api"
	"system_management/internal/shared/constant"
	"time"
)

func main() {
	Run()
}

type Api struct {
	cfg *config.Config
	db  *bun.DB
	e   *echo.Echo
}

func startServer(appMode constant.AppMode) {
	cfg := config.LoadConfig(config.GetEnvModeByAppMode(appMode))
	db := utils.ConnectToDB(*cfg)
	e := echo.New()

	a := Api{
		cfg: cfg,
		db:  db,
		e:   e,
	}

	a.runServer()
}

func (a *Api) initHandler() {
	handler := api.Handler{
		Config: a.cfg,
		Route:  a.e.Group("/api"),
		DB:     a.db,
	}
	handler.RegisterHandler()
}

func (a *Api) runServer() {
	a.e.Use(customMiddleware.Logging)

	a.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	a.e.HTTPErrorHandler = CustomHTTPErrorHandler(a.cfg)

	a.initHandler()

	a.e.Any("", func(c echo.Context) error {
		return response.ErrNotFound(ierr.ErrNotFound)
	})

	a.e.Any("/*", func(c echo.Context) error {
		return response.ErrNotFound(ierr.ErrNotFound)
	})

	go func() {
		var prefix = "https://"
		if a.cfg.Server.Host == "localhost" || a.cfg.Server.Host == "127.0.0.1" {
			prefix = "http://"
		}
		logrus.Infof("starting %s(%s) server at %s%s:%s/api",
			a.cfg.Server.Name, a.cfg.Server.Version,
			prefix, a.cfg.Server.Host, a.cfg.Server.Port)
		if err := a.e.Start(fmt.Sprintf(":%s", a.cfg.Server.Port)); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logrus.Info("graceful shutdown the server")
	if err := a.e.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
}
