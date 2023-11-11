package main

import (
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type config struct {
	Environment    string `env:"ENVIRONMENT" envDefault:"development"`
	Port           int    `env:"PORT" envDefault:"3002"`
	Dsn            string `env:"DSN"`
	JWTSecret      string `env:"JWT_SECRET"`
	UsersClientUrl string `env:"USERS_SERVICE"`
}

func newConfig() (*config, error) {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	cfg := &config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
