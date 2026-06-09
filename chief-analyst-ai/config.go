package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Telegram struct {
		Token string `mapstructure:"token"`
	} `mapstructure:"telegram"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
