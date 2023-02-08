package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig(name string) Config {
	viper.AddConfigPath("./internal/config/")
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Unable to load config", err)
	}

	config := Config{}
	err = viper.Unmarshal(&config)

	if err != nil {
		log.Panic("Unable to parse config", err)
	}

	return config

}
