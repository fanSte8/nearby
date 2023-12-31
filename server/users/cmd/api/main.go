package main

import (
	"log/slog"
	"os"

	"nearby/common/clients"
	"nearby/common/httperrors"
	"nearby/common/middleware"
	"nearby/common/storage"
	"nearby/users/internal/data"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type application struct {
	config           config
	logger           slog.Logger
	models           data.Models
	httpErrors       httperrors.HttpErrors
	commonMiddleware middleware.CommonMiddleware
	mailerClient     clients.IMailerClient
	storage          storage.Storage
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := newConfig()
	if err != nil {
		log.Error("Error loading configuration", "error", err)
		return
	}

	db, err := connectDB(cfg.Dsn)
	if err != nil {
		log.Error("Error connecting to the database", "error", err)
		return
	}

	defer db.Close()

	httpErrors := httperrors.NewHttpErrors(log)
	commonMiddleware := middleware.NewCommonMiddleware(httpErrors)

	mailerClient, err := clients.NewMailerClient(cfg.MailerClientUrl)
	if err != nil {
		log.Error("Error connecting to the mailer service", "error", err)
	}

	storage := storage.NewS3Storage(storage.S3Config{
		BucketName:      cfg.S3BucketName,
		Region:          cfg.AWSRegion,
		AccessKeyID:     cfg.AWSAccessKeyID,
		AccessKeySecret: cfg.AWSAccessKeySecret,
	})

	app := &application{
		config:           *cfg,
		logger:           *log,
		models:           data.NewModels(db),
		httpErrors:       httpErrors,
		commonMiddleware: commonMiddleware,
		mailerClient:     *mailerClient,
		storage:          storage,
	}

	err = app.serve()
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
}
