package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"system_management/config"
)

func ConnectToDB(cfg config.Config) *bun.DB {
	connectionConfig := cfg.DB
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connectionConfig.Host,
		connectionConfig.Port,
		connectionConfig.Username,
		connectionConfig.Password,
		connectionConfig.Name,
	)
	hsqldb, err := sql.Open("postgres", connection)
	if err != nil {
		logrus.Fatal(err)
	}

	// Create a Bun db on top of it.
	db := bun.NewDB(hsqldb, pgdialect.New())
	if cfg.Server.Mode == config.Development {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	logrus.Info("connected to database")
	return db
}
