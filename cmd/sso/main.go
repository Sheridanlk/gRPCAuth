package main

import (
	"log/slog"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/logger"
)

func main() {
	cfg := config.Load()

	log := logger.Setup(cfg.Env)

	log.Info("Statrting application!",
		slog.String("env", cfg.Env),
	)

	connect := ""

	application := app.New(log, cfg.GRPC.Port, connect, cfg.TokenTTL)

	application.GRPCSrv.Run()
}
