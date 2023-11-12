package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type config struct {
	ServerPort           int    `env:"PORT" envDefault:"3004"`
	UsersService         string `env:"USERS_SERVICE" envDefault:"http://localhost:3000/v1/"`
	PostsService         string `env:"POSTS_SERVICE" envDefault:"http://localhost:3001/v1/"`
	NotificationsService string `env:"NOTIFICATIONS_SERVICE" envDefault:"http://localhost:3002/v1/"`
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

	r := mux.NewRouter()

	r.Path("/users").Handler(reverseProxy(cfg.UsersService, "/v1"))
	r.Path("/users/{rest:.*}").Handler(reverseProxy(cfg.UsersService, "/v1"))

	r.Path("/posts").Handler(reverseProxy(cfg.PostsService, "/v1"))
	r.Path("/posts/{rest:.*}").Handler(reverseProxy(cfg.PostsService, "/v1"))

	r.Path("/notifications").Handler(reverseProxy(cfg.NotificationsService, "/v1"))
	r.Path("/notifications/{rest:.*}").Handler(reverseProxy(cfg.NotificationsService, "/v1"))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:      r,
		ErrorLog:     slog.NewLogLogger(log.Handler(), 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Info(fmt.Sprintf("Starting server on port %d", cfg.ServerPort))

		err := srv.ListenAndServe()
		if err != nil {
			log.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	sig := <-c
	log.Info("Got signal", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}

func reverseProxy(targetURL, prefix string) http.Handler {
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal("Error parsing target URL:", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Director = func(r *http.Request) {
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.URL.Path = prefix + r.URL.Path
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = target.Host
	}

	return proxy
}
