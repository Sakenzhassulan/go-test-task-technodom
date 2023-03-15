package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbName       string `mapstructure:"DB_NAME"`
	DbCollection string `mapstructure:"DB_COLLECTION"`
	DbUri        string `mapstructure:"DB_URI"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
}

func NewConfig() (config Config, err error) {
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.Unmarshal(&config)
	return
}
