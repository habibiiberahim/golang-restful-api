package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigType("env")
	config.SetConfigName(".env")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return config
}