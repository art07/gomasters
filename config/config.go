package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

var appConfig *AppConfig

type AppConfig struct {
	Server struct {
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host" envconfig:"DB_HOST"`
		Port     string `yaml:"port" envconfig:"DB_PORT"`
		Name     string `yaml:"name" envconfig:"DB_NAME"`
		User     string `yaml:"user" envconfig:"DB_USER"`
		Password string `yaml:"password" envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

func GetAppConfig() (*AppConfig, error) {
	if appConfig == nil {
		cfgFile, err := os.Open("./config/config.yml")
		if err != nil {
			return nil, err
		}
		//goland:noinspection GoUnhandledErrorResult
		defer cfgFile.Close()

		appConfig = &AppConfig{}
		decoder := yaml.NewDecoder(cfgFile)
		if err = decoder.Decode(appConfig); err != nil {
			return nil, err
		}

		if err = envconfig.Process("", appConfig); err != nil {
			return nil, err
		}
	}
	return appConfig, nil
}
