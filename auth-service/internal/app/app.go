package app

import (
	"log/slog"

	grpcapp "github.com/IskanderSh/test-task-services/auth-service/internal/app/grpc"
	"github.com/IskanderSh/test-task-services/auth-service/internal/config"
	"github.com/IskanderSh/test-task-services/auth-service/internal/services"
)

type App struct {
	GRPCServer *grpcapp.GRPCApp
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	service, err := services.NewAuthService(log, cfg.UserService)
	if err != nil {
		panic(err)
	}

	grpcApp := grpcapp.NewGRPCApp(log, service, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
