package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var appConfig *AppConfig

type AppConfig struct {
	// Server
	AppAddr string `envconfig:"APP_ADDR" required:"true"`

	// Postgres
	PgHost     string `envconfig:"PG_HOST" required:"true"`
	PgPort     string `envconfig:"PG_PORT" default:"5432"`
	PgDb       string `envconfig:"PG_DB" required:"true"`
	PgUser     string `envconfig:"PG_USER" required:"true"`
	PgPassword string `envconfig:"PG_PASSWORD" required:"true"`
}

func GetAppConfig() (*AppConfig, error) {
	if appConfig == nil {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}

		appConfig = &AppConfig{}
		if err := envconfig.Process("", appConfig); err != nil {
			return nil, err
		}
	}
	return appConfig, nil
}

func (c *AppConfig) GetDbString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		c.PgUser, c.PgPassword, c.PgHost, c.PgPort, c.PgDb)
}
