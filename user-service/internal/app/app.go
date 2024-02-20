package app

import (
	"log/slog"

	grpcapp "github.com/IskanderSh/test-task-services/user-services/internal/app/grpc"
	"github.com/IskanderSh/test-task-services/user-services/internal/config"
	"github.com/IskanderSh/test-task-services/user-services/internal/services"
	"github.com/IskanderSh/test-task-services/user-services/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.GRPCApp
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	postgres, err := storage.New(&cfg.Storage)
	if err != nil {
		panic(err)
	}

	userService := services.New(log, postgres)

	grpcApp := grpcapp.New(log, userService, cfg.GRPC.Port)

	return &App{GRPCServer: grpcApp}
}
