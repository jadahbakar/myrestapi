package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App *Application
		Db  *Database
	}

	Application struct {
		Name string
		Port int
	}
	Database struct {
		Url      string
		Timezone string
	}
)

func New(mode string, configPath string) (config *Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigFile("app")
	if mode == "devel" {
		viper.SetConfigName("app_test")
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	configuration := Config{
		App: &Application{
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetInt("APP_PORT"),
		},
		Db: &Database{
			Url:      viper.GetString("DATABASE_URL"),
			Timezone: viper.GetString("DATABASE_TIMEZONE"),
		},
	}
	return &configuration, err
}
