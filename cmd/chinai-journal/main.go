package main

import (
	"chinai-journal/internal/config"
	"chinai-journal/internal/lib/logger/sl"
	"chinai-journal/internal/storage/sqlite"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	mwLogger "chinai-journal/internal/http-server/middleware/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	// DONE: init config

	cfg := config.MustLoad()
	fmt.Println(cfg)

	// DONE: init logger

	log := setupLogger(cfg.Env)

	log.Info("starting url_shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// DONE: init storage

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1) // we can use `return` instead
	}

	_ = storage
	// DONE: init router

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// http routes

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.HTML(w, r, "pong")
	})

	// TODO: run server

	log.Info("starting server", slog.String("Addr", cfg.Address))
	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.TimeOut,
		WriteTimeout: cfg.HTTPServer.TimeOut,
		IdleTimeout:  cfg.HTTPServer.IdleTimeOut,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server", sl.Err(err))
	}

	log.Error("server stopped")
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {

	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
