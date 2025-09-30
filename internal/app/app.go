package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/services/auth"
	"sso/internal/storage/postgres"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
	Storage *postgres.Storage
}

func New(log *slog.Logger, grpcPort int, storageUser string, storagePassword string, storageName string, storagePort int, tokenTTL time.Duration) *App {
	storage, err := postgres.New(storageUser, storagePassword, storageName, storagePort)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
		Storage: storage,
	}
}
