package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/orenvadi/auth-grpc/internal/app/grpc"
	"github.com/orenvadi/auth-grpc/internal/config"
	"github.com/orenvadi/auth-grpc/internal/services/auth"
	"github.com/orenvadi/auth-grpc/internal/storage/surrdb"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, cfg *config.Config, tokenTTL time.Duration) *App {
	// DONE init storage

	// storage, err := sqlite.New(storagePath)
	// storage, err := postgres.New(fmt.Sprintf("postgres://%s?sslmode=disable", storageDSN))

	storage, err := surrdb.New(cfg.Storage.Host, cfg.Storage.User, cfg.Storage.Password, cfg.Storage.DbName, cfg.Storage.DbNameSpace)
	if err != nil {
		panic(err)
	}

	// DONE init auth server (auth)

	authService := auth.New(log, storage, storage, storage, storage, storage, tokenTTL, cfg.JwtSecret)

	grpcApp := grpcapp.New(log, authService, authService, storage, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
