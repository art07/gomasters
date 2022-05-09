package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"net/http"
	"playground/rest-api/gomasters/config"
	"playground/rest-api/gomasters/router"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Golang REST API started")

	cfg, err := config.GetAppConfig()
	if err != nil {
		logger.Fatal("config reading error", zap.Error(err))
	}
	logger.Info("Config OK")

	// https://github.com/jackc/pgx/blob/master/stdlib/sql.go
	db, err := sql.Open("pgx", cfg.GetDbString())
	if err != nil {
		logger.Fatal("open db error", zap.Error(err))
	}
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
	if err = db.Ping(); err != nil {
		logger.Fatal("ping db error", zap.Error(err))
	}
	logger.Info("Db OK")

	r := router.NewRouter(db, logger)
	logger.Info("Start http server", zap.String("server", cfg.GetDbString()))
	logger.Fatal("fatal server error", zap.Error(http.ListenAndServe(cfg.GetServerString(), r)))
}
