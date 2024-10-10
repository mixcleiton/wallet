package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func LoadConfig(path string) (*DatabaseConfig, error) {
	var config *DatabaseConfig
	log.Println(path)
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error unmarshaling config: %s", err)
	}

	return config, nil
}
