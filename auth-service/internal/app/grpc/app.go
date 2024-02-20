package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/IskanderSh/test-task-services/auth-service/internal/handlers"
	"github.com/IskanderSh/test-task-services/auth-service/internal/services"
	"google.golang.org/grpc"
)

type GRPCApp struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	port       int
}

func NewGRPCApp(
	log *slog.Logger,
	service *services.AuthService,
	port int,
) *GRPCApp {
	grpcServer := grpc.NewServer()

	handlers.Register(grpcServer, log, service)

	return &GRPCApp{
		log:        log,
		GRPCServer: grpcServer,
		port:       port,
	}
}

func (a *GRPCApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(fmt.Sprintf("couldn't run grpc application on port: %d", a.port))
	}
}

func (a *GRPCApp) Run() error {
	const op = "grpcApp.Run"

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	if err := a.GRPCServer.Serve(lis); err != nil {
		return err
	}

	a.log.Info("grpc server is running",
		slog.String("addr", lis.Addr().String()),
		slog.String("op", op))

	return nil
}

func (a *GRPCApp) Stop() {

}
