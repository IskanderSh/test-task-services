package services

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/IskanderSh/test-task-services/user-services/internal/domain/models"
	"github.com/IskanderSh/test-task-services/user-services/internal/lib/error/wrapper"
	"github.com/IskanderSh/test-task-services/user-services/internal/storage/postgres"
	"github.com/pkg/errors"
)

type UserService struct {
	log         *slog.Logger
	usrProvider UserProvider
}

type UserProvider interface {
	GetUser(ctx context.Context, uuid string) (*models.User, error)
	UpdateUser(ctx context.Context, uuid string, updateUser models.UpdateUser) (bool, error)
	DeleteUser(ctx context.Context, uuid string) (bool, error)
}

func New(
	log *slog.Logger,
	usrProvider UserProvider,
) *UserService {
	return &UserService{log: log, usrProvider: usrProvider}
}

var (
	ErrUserNotFound = errors.New("user with such uuid not found")
	ErrInvalidCreds = errors.New("invalid credentials")

	nameRegexp = regexp.MustCompile(`[A-Z][a-z]{5,30}`).MatchString

	emailRegexp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString
)

type funcNameString func(string) bool

func (p *UserService) Get(ctx context.Context, uuid string) (*models.User, error) {
	const op = "services.Get"

	log := p.log.With(
		slog.String("op", op),
		slog.String("uuid", uuid))

	log.Info("getting user")

	user, err := p.usrProvider.GetUser(ctx, uuid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", uuid)

			return nil, wrapper.Wrap(op, ErrUserNotFound)
		}
		log.Error("failed to get user", err)

		return nil, wrapper.Wrap(op, err)
	}

	log.Info("user successfully get")

	return user, nil
}

func (p *UserService) Update(ctx context.Context, uuid, name, surname, email, role string) (bool, error) {
	const op = "service.Update"

	log := p.log.With(
		slog.String("op", op),
		slog.String("uuid", uuid))

	log.Info("getting user")

	_, err := p.usrProvider.GetUser(ctx, uuid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", "uuid", uuid)

			return false, wrapper.Wrap(op, ErrUserNotFound)
		}
		return false, wrapper.Wrap(op, err)
	}

	log.Info("updating user")

	validateArgs := []struct {
		funcNameString
		value   string
		message string
	}{
		{validateName, name, "invalid name"},
		{validateName, surname, "invalid surname"},
		{validateEmail, email, "invalid email"},
		{validateRole, role, "invalid role"},
	}

	for _, args := range validateArgs {
		if !args.funcNameString(args.value) {
			log.Warn("invalid credentials", args.message, args.value)

			return false, errors.Wrap(ErrInvalidCreds, args.message)
		}
	}

	updateUser := models.UpdateUser{
		Name:    &name,
		Surname: &surname,
		Email:   &email,
		Role:    &role,
	}

	ok, err := p.usrProvider.UpdateUser(ctx, uuid, updateUser)
	if err != nil {
		return false, wrapper.Wrap(op, err)
	}

	log.Info("user successfully updated")

	return ok, nil
}

func (p *UserService) Delete(ctx context.Context, uuid string) (bool, error) {
	const op = "service.Delete"

	log := p.log.With(
		slog.String("op", op),
		slog.String("uuid", uuid))

	log.Info("getting user")

	_, err := p.usrProvider.GetUser(ctx, uuid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", "uuid", uuid)

			return false, wrapper.Wrap(op, ErrUserNotFound)
		}
		return false, wrapper.Wrap(op, err)
	}

	log.Info("deleting user")

	ok, err := p.usrProvider.DeleteUser(ctx, uuid)
	if err != nil {
		return false, wrapper.Wrap(op, err)
	}

	log.Info("user successfully deleted")

	return ok, nil
}

func validateName(name string) bool {
	if !nameRegexp(name) {
		return false
	}
	return true
}

func validateEmail(email string) bool {
	if !emailRegexp(email) {
		return false
	}
	return true
}

func validateRole(role string) bool {
	if role == "user" || role == "admin" {
		return true
	}
	return false
}
