package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
}

func LoadAppConfig(path string) (AppConfig, error) {
	if path == "" {
		return AppConfig{}, errors.New("config path is empty")
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return AppConfig{}, errors.New("config file not found")
		}

		return AppConfig{}, fmt.Errorf("failed to read config file: %s", err.Error())
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, fmt.Errorf("failed to unmarshal config: %s", err.Error())
	}

	return config, nil
}
