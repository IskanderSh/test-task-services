package services

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/domain/models"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/storage/postgres"
	"github.com/pkg/errors"
)

type UserService struct {
	log         *slog.Logger
	usrProvider UserProvider
}

type UserProvider interface {
	GetUser(ctx context.Context, uuid string) (*models.User, error)
	UpdateUser(ctx context.Context, uuid, name, surname, email, role string) (bool, error)
	DeleteUser(ctx context.Context, uuid string) (bool, error)
}

func New(
	log *slog.Logger,
	usrProvider UserProvider,
) *UserService {
	return &UserService{log: log, usrProvider: usrProvider}
}

var (
	ErrUserNotFound      = errors.New("user with such uuid not found")
	ErrUserAlreadyExists = errors.New("user with this credentials already exists")
	ErrInvalidCreds      = errors.New("invalid credentials")

	nameRegexp = regexp.MustCompile(`[A-Z][a-z]{5,30}`).MatchString

	emailRegexp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString
)

type funcNameString func(string) bool

func (p *UserService) Get(ctx context.Context, uuid string) (*models.User, error) {
	user, err := p.usrProvider.GetUser(ctx, uuid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			p.log.Warn("user not found", uuid)
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (p *UserService) Update(ctx context.Context, uuid, name, surname, email, role string) (bool, error) {
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
			return false, errors.Wrap(ErrInvalidCreds, args.message)
		}
	}

	ok, err := p.usrProvider.UpdateUser(ctx, uuid, name, surname, email, role)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			p.log.Warn("user not found", uuid)
			return false, ErrUserNotFound
		}
		return false, err
	}

	return ok, nil
}

func (p *UserService) Delete(ctx context.Context, uuid string) (bool, error) {
	ok, err := p.usrProvider.DeleteUser(ctx, uuid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			p.log.Warn("user not found", uuid)
			return false, ErrUserNotFound
		}

		return false, err
	}

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
