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

	cfg, err := config.GetAppConfig(".env")
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

	server := &http.Server{
		Addr:    cfg.AppAddr,
		Handler: r,
	}
	logger.Info("Start http server", zap.String("server", cfg.AppAddr))
	logger.Fatal("fatal server error", zap.Error(server.ListenAndServe()))

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//<-c
	//
	//cctx, cancel := context.WithTimeout(context.Background(), wait)
	//defer cancel()
	//err = server.Shutdown(cctx)
	//if err != nil {
	//	logger.Error("shutdown error", zap.Error(err))
	//}
	//logger.Error("shutting down")
	//os.Exit(0)
}
