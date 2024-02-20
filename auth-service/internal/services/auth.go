package services

import (
	"context"
	"log/slog"

	authv1 "github.com/IskanderSh/test-task-protos/gen/go/auth-service"
)

type AuthService struct {
	log *slog.Logger
}

func NewAuthService(log *slog.Logger) *AuthService {
	return &AuthService{log: log}
}

func (s *AuthService) SignUp(ctx context.Context, request *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {
	return nil, nil
}

func (s *AuthService) SignIn(ctx context.Context, request *authv1.SignInRequest) (*authv1.SignInResponse, error) {
	return nil, nil
}
