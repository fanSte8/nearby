package main

import (
	"log/slog"
	"os"

	"nearby/common/httperrors"
	"nearby/common/middleware"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type config struct {
	ServerPort   int    `env:"PORT"  envDefault:"3003"`
	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     int    `env:"SMTP_PORT"`
	SmtpUsername string `env:"SMTP_USERNAME"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
	SmtpSender   string `env:"SMTP_SENDER"`
}

type application struct {
	config           config
	logger           slog.Logger
	mailer           imailer
	httpErrors       httperrors.HttpErrors
	commonMiddleware middleware.CommonMiddleware
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Error("Error loading .env file", "error", err)
			return
		}
	}

	cfg := &config{}

	err := env.Parse(cfg)
	if err != nil {
		log.Error("Error reading config", "error", err)
		return
	}

	mailer := newMailer(cfg)

	httpErrors := httperrors.NewHttpErrors(log)
	commonMiddleware := middleware.NewCommonMiddleware(httpErrors)

	app := application{
		config:           *cfg,
		logger:           *log,
		mailer:           mailer,
		httpErrors:       httpErrors,
		commonMiddleware: commonMiddleware,
	}

	err = app.serve()
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
}
