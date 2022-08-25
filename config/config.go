package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	HTTPPort   int64  `env:"PINKIE_HTTP_PORT" envDefault:"8080"`
	DBHost     string `env:"PINKIE_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int64  `env:"PINKIE_DB_PORT" envDefault:"3306"`
	DBUser     string `env:"PINKIE_DB_USER" envDefault:"pinkie"`
	DBPassword string `env:"PINKIE_DB_PASSWORD" envDefault:"pinkie"`
	DBName     string `env:"PINKIE_DB_NAME" envDefault:"pinkie"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
