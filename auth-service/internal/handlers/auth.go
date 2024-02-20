package handlers

import (
	"context"
	"log/slog"

	authv1 "github.com/IskanderSh/test-task-protos/gen/go/auth-service"
	"google.golang.org/grpc"
)

type AuthProvider interface {
	SignUp(ctx context.Context, request *authv1.SignUpRequest) (*authv1.SignUpResponse, error)
	SignIn(ctx context.Context, request *authv1.SignInRequest) (*authv1.SignInResponse, error)
}

type AuthHandler struct {
	authv1.UnimplementedAuthServer
	log  *slog.Logger
	auth AuthProvider
}

func Register(grpc *grpc.Server, log *slog.Logger, provider AuthProvider) {
	authv1.RegisterAuthServer(grpc, &AuthHandler{log: log, auth: provider})
}
