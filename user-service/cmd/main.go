package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/app"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"

	debugLvl = "DEBUG"
	infoLvl  = "INFO"
	warnLvl  = "WARN"
	errorLvl = "ERROR"
)

func main() {
	// load config file
	cfg := config.MustLoad()

	// init logger
	log := setupLogger(cfg)
	log.Info("logger initialized successfully")

	// init app
	application := app.New(log, cfg)
	log.Info("app created successfully")

	// start app
	go application.GRPCServer.MustRun()

	// graceful shutdown
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	}

	return log
}

func getLogLevel(lvl string) slog.Level {
	var res slog.Level

	switch strings.ToUpper(lvl) {
	case debugLvl:
		res = slog.LevelDebug
	case infoLvl:
		res = slog.LevelInfo
	case warnLvl:
		res = slog.LevelWarn
	case errorLvl:
		res = slog.LevelError
	}

	return res
}