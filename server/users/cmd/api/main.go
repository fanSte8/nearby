package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"time"

	"nearby/common/clients"
	"nearby/common/httperrors"
	"nearby/common/middleware"
	"nearby/users/internal/data"

	"github.com/caarlos0/env/v9"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	Environment     string `env:"ENVIRONMENT" envDefault:"development"`
	Port            int    `env:"PORT" envDefault:"3000"`
	Dsn             string `env:"DSN"`
	JWTSecret       string `env:"JWT_SECRET"`
	MailerClientUrl string `env:"MAILER_SERVICE"`
}

type application struct {
	config           config
	logger           slog.Logger
	models           data.Models
	httpErrors       httperrors.HttpErrors
	commonMiddleware middleware.CommonMiddleware
	mailerClient     clients.MailerClient
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

	mailerClient, err := clients.NewMailerClient(cfg.MailerClientUrl)
	if err != nil {
		log.Error("Error connecting to the mailer service", "error", err)
	}

	app := &application{
		config:           *cfg,
		logger:           *log,
		models:           data.NewModels(db),
		httpErrors:       httpErrors,
		commonMiddleware: commonMiddleware,
		mailerClient:     *mailerClient,
	}

	err = app.serve()
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}

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

func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	} else {
		return err
	}
}
