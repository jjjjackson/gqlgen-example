package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DB_URL"`
}

func LoadConfig() *Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err.Error())
	}

	return &config
}
