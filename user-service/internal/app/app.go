package app

import (
	"log/slog"

	grpcapp "github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/app/grpc"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/config"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/services"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/storage/postgres"
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
