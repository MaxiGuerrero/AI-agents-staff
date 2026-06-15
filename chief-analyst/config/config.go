package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	WhiteUserId   int64  `mapstructure:"white_user_id"`
	WhiteUserName string `mapstructure:"white_user_name"`
	LlmBaseUrl    string `mapstructure:"llm_base_url"`
}

func LoadConfig() *Config {
	viper.SetConfigFile(getConfigPath())

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../config.yaml"
	}
	return configPath
}
