package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"PORT"`
}

func loadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal("Cannot Read Environment Variables")
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("Cannot Parse Environment Variables")
	}

	return
}
