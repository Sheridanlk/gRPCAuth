package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/logger"
	"syscall"
)

func main() {
	cfg := config.Load()

	log := logger.Setup(cfg.Env)

	log.Info("Statrting application!",
		slog.String("env", cfg.Env),
	)

	application := app.New(log, cfg.GRPC.Port, cfg.Postrgres.User, cfg.Postrgres.Password, cfg.Postrgres.Name, cfg.Postrgres.Port, cfg.TokenTTL)

	go application.GRPCSrv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCSrv.Stop()

	log.Info("Application stopped")
}
