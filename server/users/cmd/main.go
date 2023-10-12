package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Port        int    `env:"PORT" envDefault:"3000"`
	Dsn         string `env:"DSN"`
}

type application struct {
	config config
	logger slog.Logger
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
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

	app := &application{
		config: *cfg,
		logger: *log,
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
