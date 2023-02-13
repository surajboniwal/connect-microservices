package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress      string `mapstructure:"SERVER_ADDRESS"`
	AuthServiceAddress string `mapstructure:"AUTHENTICATION_SERVICE_ADDRESS"`
}

func LoadConfig(name string) (Config, error) {
	viper.AddConfigPath("./internal/config/")
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = viper.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil

}
