package config

import (
	"github.com/spf13/viper"
)

// Config represents all hightower configuration variables
type Config struct {
	EntryPoint string
}

// Parse loads a config file
func Parse() (Config, error) {
	var config Config

	viper.SetConfigName("hightower")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("testdata")

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	config.EntryPoint = viper.GetString("entrypoint")
	return config, nil
}
