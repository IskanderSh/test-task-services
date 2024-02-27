package handlers

import (
	"context"
	"log/slog"

	authv1 "github.com/IskanderSh/test-task-protos/gen/go/auth-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthProvider interface {
	SignUp(ctx context.Context, name, surname, password, email string) (string, error)
	SignIn(ctx context.Context, email, password string) (string, error)
}

type AuthHandler struct {
	authv1.UnimplementedAuthServer
	log  *slog.Logger
	auth AuthProvider
}

func Register(grpc *grpc.Server, log *slog.Logger, provider AuthProvider) {
	authv1.RegisterAuthServer(grpc, &AuthHandler{log: log, auth: provider})
}

func (h *AuthHandler) SignUp(ctx context.Context, req *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {
	if err := validateSignUpRequest(req); err != nil {
		return nil, err
	}

	h.log.Debug("successfully validate sign up request")

	uuid, err := h.auth.SignUp(ctx, req.GetName(), req.GetSurname(), req.GetPassword(), req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	h.log.Debug("successfully sign up user and get uuid")

	return &authv1.SignUpResponse{
		Uuid: uuid,
	}, nil
}

func (h *AuthHandler) SignIn(ctx context.Context, req *authv1.SignInRequest) (*authv1.SignInResponse, error) {
	if err := validateSignInRequest(req); err != nil {
		return nil, err
	}

	h.log.Debug("successfully validate sign in request")

	uuid, err := h.auth.SignIn(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	h.log.Debug("successfully sign in user")

	return &authv1.SignInResponse{
		Uuid: uuid,
	}, nil
}

func validateSignUpRequest(req *authv1.SignUpRequest) error {
	if req.GetName() == "" || req.GetSurname() == "" || req.GetPassword() == "" || req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "all fields are necessary")
	}

	return nil
}

func validateSignInRequest(req *authv1.SignInRequest) error {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "email and password are required")
	}

	return nil
}
