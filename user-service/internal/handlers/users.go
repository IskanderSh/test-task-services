package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	userv1 "github.com/IskanderSh/test-task-protos/gen/go/user-service"
	"github.com/IskanderSh/test-task-services/user-services/internal/domain/models"
	"github.com/IskanderSh/test-task-services/user-services/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler interface {
	Get(ctx context.Context, uuid string) (*models.User, error)
	Update(ctx context.Context, uuid, name, surname, email, role string) (bool, error)
	Delete(ctx context.Context, uuid string) (bool, error)
}

type serverAPI struct {
	userv1.UnimplementedUserServer
	user UserHandler
	log  *slog.Logger
}

func Register(gRPC *grpc.Server, user UserHandler, log *slog.Logger) {
	userv1.RegisterUserServer(gRPC, &serverAPI{user: user, log: log})
}

func (s *serverAPI) Get(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	if err := validateGetRequest(req); err != nil {
		return nil, err
	}

	s.log.Debug(fmt.Sprintf("successfully validate get request with params: %s", req.GetUuid()))

	user, err := s.user.Get(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.log.Debug("return get user request successfully")

	return &userv1.GetUserResponse{
		Uuid:    user.Uuid,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Role:    user.Role,
	}, nil
}

func (s *serverAPI) Update(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	s.log.Debug("successfully validate update request")

	ok, err := s.user.Update(ctx, req.GetUuid(), req.GetName(), req.GetSurname(), req.GetEmail(), req.GetRole())
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "user with such uuid not found")
		} else if errors.Is(err, services.ErrInvalidCreds) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.log.Debug("return update user request successfully")

	return &userv1.UpdateUserResponse{
		Ok: ok,
	}, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *userv1.DeleteUserRequest) (*userv1.DeleteUserResponse, error) {
	if err := validateDeleteRequest(req); err != nil {
		return nil, err
	}

	s.log.Debug("successfully validate delete user request")

	ok, err := s.user.Delete(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "user with such uuid not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.log.Debug("return delete user request successfully")

	return &userv1.DeleteUserResponse{
		Ok: ok,
	}, nil
}

func validateGetRequest(req *userv1.GetUserRequest) error {
	if req.GetUuid() == "" {
		return status.Error(codes.InvalidArgument, "uuid is required")
	}

	return nil
}

func validateUpdateRequest(req *userv1.UpdateUserRequest) error {
	if req.GetUuid() == "" {
		return status.Error(codes.InvalidArgument, "uuid is required")
	}
	if req.GetEmail() == "" && req.GetName() == "" && req.GetSurname() == "" && req.GetRole() == "" {
		return status.Error(codes.InvalidArgument, "one of the update params should be non empty")
	}

	return nil
}

func validateDeleteRequest(req *userv1.DeleteUserRequest) error {
	if req.GetUuid() == "" {
		return status.Error(codes.InvalidArgument, "uuid is required")
	}

	return nil
}
