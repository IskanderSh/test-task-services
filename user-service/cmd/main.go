package main

import (
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IskanderSh/test-task-services/user-services/internal/app"
	"github.com/IskanderSh/test-task-services/user-services/internal/config"
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
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))
	application.GRPCServer.Stop()

	log.Info("application stopped")
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
