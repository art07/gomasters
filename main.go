package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"net/http"
	"playground/rest-api/gomasters/config"
	"playground/rest-api/gomasters/handler"
	adminRepo "playground/rest-api/gomasters/repository/admin"
	userRepo "playground/rest-api/gomasters/repository/user"
	"playground/rest-api/gomasters/router"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("REST API app started.")

	cfg, err := config.GetAppConfig()
	if err != nil {
		logger.Fatal("Config reading error.", zap.Error(err))
	}

	// https://github.com/jackc/pgx/blob/master/stdlib/sql.go
	db, err := sql.Open("pgx", fmt.Sprintf(
		"user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		logger.Fatal("Open DB error.", zap.Error(err))
	}
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
	if err = db.Ping(); err != nil {
		logger.Fatal("Ping DB error.", zap.Error(err))
	}

	uRepo := userRepo.NewUserRepository(logger, db)
	aRepo := adminRepo.NewAdminRepository(logger, db)

	userHandler := handler.NewHandler(logger, uRepo)
	adminHandler := handler.NewHandler(logger, aRepo)

	r := router.NewRouter(userHandler, adminHandler)

	addr := fmt.Sprint(cfg.Server.Host, ":", cfg.Server.Port)
	logger.Info("Start server.", zap.String("server", addr))
	if err = http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("Server error.", zap.Error(err))
	}
}
