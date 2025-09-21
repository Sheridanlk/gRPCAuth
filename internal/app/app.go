package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storageConnect string, tokenTTL time.Duration) *App {
	// TODO: инициализация хранилища

	// TODO: инициализация сервисного слоя (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
