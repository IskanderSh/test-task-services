package handlers

import (
	"context"
	"log/slog"

	userv1 "github.com/IskanderSh/test-task-protos/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler interface {
	Create()
	Get()
	Update()
	Delete()
}

type serverAPI struct {
	userv1.UnimplementedUserServer
	user UserHandler
	log  *slog.Logger
}

func Register(gRPC *grpc.Server, user UserHandler, log *slog.Logger) {
	userv1.RegisterUserServer(gRPC, &serverAPI{user: user, log: log})
}

func (s *serverAPI) Create(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	return nil, nil
}

func (s *serverAPI) Get(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	return nil, nil
}

func (s *serverAPI) Update(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	return nil, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *userv1.DeleteUserRequest) (*userv1.DeleteUserResponse, error) {
	return nil, nil
}

func validateCreate(req *userv1.CreateUserRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}
	if req.GetSurname() == "" {
		return status.Error(codes.InvalidArgument, "surname is required")
	}
	if req.GetRole() == "" {
		return status.Error(codes.InvalidArgument, "role is required")
	}

	return nil
}
