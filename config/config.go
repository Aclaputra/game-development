package config

import (
	"log"

	"github.com/spf13/viper"
)

func ExecConfig() {

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error get config: %s", err)
	}
}
