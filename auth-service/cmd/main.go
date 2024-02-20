package main

import (
	"log/slog"
	"os"

	"github.com/IskanderSh/test-task-services/auth-service/internal/app"
	"github.com/IskanderSh/test-task-services/auth-service/internal/config"
)

func main() {
	// parse config
	cfg := config.MustLoad()

	// init logger
	log := setupLogger(cfg)
	log.Info("logger initialized successfully")

	// init app
	application := app.New(log, cfg)

	// start app
	go application.GRPCServer.MustRun()

	// graceful shutdown
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	}

	return log
}

func getLogLevel(logLevel string) slog.Level {
	switch logLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "error":
		return slog.LevelError
	}

	return slog.LevelDebug
}
