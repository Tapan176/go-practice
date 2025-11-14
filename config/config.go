package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Http struct {
	Port int    `env:"httpPort" env-required:"true"`
	Host string `env:"httpHost" env-required:"true"`
}

type Database struct {
	Host     string `env:"databaseHost" env-required:"true"`
	Port     int    `env:"databasePort" env-required:"true"`
	User     string `env:"databaseUser" env-required:"true"`
	Password string `env:"databasePassword" env-required:"true"`
	Name     string `env:"databaseName" env-required:"true"`
}

type Config struct {
	Env      string   `env:"env" env-required:"true"`
	Http     Http     `env:"http"`
	Database Database `env:"database"`
}

func GetConfig() *Config {
	var config Config

	err := cleanenv.ReadConfig(".config/.env", &config)
	if err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}

	return &config
}
