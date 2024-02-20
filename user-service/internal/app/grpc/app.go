package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/IskanderSh/test-task-services/user-services/internal/handlers"
	"github.com/IskanderSh/test-task-services/user-services/internal/services"
	"google.golang.org/grpc"
)

type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	userService *services.UserService,
	port int,
) *GRPCApp {
	gRPCServer := grpc.NewServer()

	handlers.Register(gRPCServer, userService, log)

	return &GRPCApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *GRPCApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *GRPCApp) Run() error {
	a.log.Info("running grpc app")
	const op = "grpcGRPCApp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return err
	}

	a.log.Info("grpc server is running",
		slog.String("addr", l.Addr().String()),
		slog.String("op", op))

	return nil
}

func (a *GRPCApp) Stop() {
	const op = "grpcapp.Stop"

	a.log.Info("stopping gRPC server", slog.String("op", op))

	a.gRPCServer.GracefulStop()
}
