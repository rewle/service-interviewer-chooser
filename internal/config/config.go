package config

import (
	"github.com/rewle/service-select-participants/internal/utils"
	"github.com/spf13/viper"
)

type Config struct {
	Addr       string
	DbUser     string
	DbPassword string
	DbAddr     string
	Db         string
}

func Init() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	utils.PanicOnErr(viper.ReadInConfig())
	var cfg Config
	utils.PanicOnErr(viper.Unmarshal(&cfg))

	return &cfg
}
