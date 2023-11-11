package main

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type config struct {
	Environment            string `env:"ENVIRONMENT" envDefault:"development"`
	Port                   int    `env:"PORT" envDefault:"3001"`
	Dsn                    string `env:"DSN"`
	JWTSecret              string `env:"JWT_SECRET"`
	UsersClientUrl         string `env:"USERS_SERVICE"`
	NotificationsClientUrl string `env:"NOTIFICATIONS_SERVICE"`
	S3BucketName           string `env:"S3_BUCKET_NAME"`
	AWSRegion              string `env:"AWS_REGION"`
	AWSAccessKeyID         string `env:"AWS_ACCESS_KEY_ID"`
	AWSAccessKeySecret     string `env:"AWS_ACCESS_KEY_SECRET"`
	testing                bool
}

func newConfig() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &config{}

	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
