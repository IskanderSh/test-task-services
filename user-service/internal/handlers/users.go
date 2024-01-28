package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	userv1 "github.com/IskanderSh/test-task-protos/gen/go"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/domain/models"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler interface {
	Create(ctx context.Context, name, surname, email, role string) (string, error)
	Get(ctx context.Context, uuid string) (models.User, error)
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

func (s *serverAPI) CreateRequest(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	s.log.Debug(fmt.Sprintf("successfully validate create request with params: %s, %s, %s, %s"),
		req.GetName(), req.GetSurname(), req.GetEmail(), req.GetRole())

	uuid, err := s.user.Create(ctx, req.GetName(), req.GetSurname(), req.GetEmail(), req.GetRole())
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			return nil, status.Error(codes.InvalidArgument, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.log.Debug("return create user request successfully")

	return &userv1.CreateUserResponse{
		Uuid: uuid,
	}, nil
}

func (s *serverAPI) GetRequest(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	if err := validateGetRequest(req); err != nil {
		return nil, err
	}

	s.log.Debug(fmt.Sprintf("successfully validate get request with params: %s"), req.GetUuid())

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

func (s *serverAPI) UpdateRequest(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	s.log.Debug("successfully validate update request")

	ok, err := s.user.Update(ctx, req.GetUuid(), req.GetName(), req.GetSurname(), req.GetEmail(), req.GetRole())
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "user with such uuid not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.log.Debug("return update user request successfully")

	return &userv1.UpdateUserResponse{
		Ok: ok,
	}, nil
}

func (s *serverAPI) DeleteRequest(ctx context.Context, req *userv1.DeleteUserRequest) (*userv1.DeleteUserResponse, error) {
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

func validateCreateRequest(req *userv1.CreateUserRequest) error {
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

func validateDeleteRequest(req *userv1.DeleteUserRequest) error {
	if req.GetUuid() == "" {
		return status.Error(codes.InvalidArgument, "uuid is required")
	}

	return nil
}
