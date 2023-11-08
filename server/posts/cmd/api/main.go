package main

import (
	"log/slog"
	"os"

	"nearby/common/clients"
	"nearby/common/httperrors"
	"nearby/common/middleware"
	"nearby/common/storage"
	"nearby/posts/internal/data"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type application struct {
	config              config
	logger              slog.Logger
	models              data.Models
	httpErrors          httperrors.HttpErrors
	commonMiddleware    middleware.CommonMiddleware
	storage             storage.Storage
	usersClient         clients.UsersClient
	notificationsClient clients.NotificationsClient
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

	err = migrateDB(db)
	if err != nil {
		log.Error("Error running database migrations", "error", err)
		return
	}

	httpErrors := httperrors.NewHttpErrors(log)
	commonMiddleware := middleware.NewCommonMiddleware(httpErrors)

	storage := storage.NewS3Storage(storage.S3Config{
		BucketName:      cfg.S3BucketName,
		Region:          cfg.AWSRegion,
		AccessKeyID:     cfg.AWSAccessKeyID,
		AccessKeySecret: cfg.AWSAccessKeySecret,
	})

	usersClient, err := clients.NewUsersClient(cfg.UsersClientUrl)
	if err != nil {
		log.Error("Error creating users client", "error", err)
		return
	}

	notificationsClient, err := clients.NewNotificationsClient(cfg.NotificationsClientUrl)
	if err != nil {
		log.Error("Error creating notifications client", "error", err)
		return
	}

	app := &application{
		config:              *cfg,
		logger:              *log,
		models:              data.NewModels(db),
		httpErrors:          httpErrors,
		commonMiddleware:    commonMiddleware,
		storage:             storage,
		usersClient:         *usersClient,
		notificationsClient: *notificationsClient,
	}

	err = app.serve()
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
}
