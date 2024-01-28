package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/domain/models"
)

type UserService struct {
	log         *slog.Logger
	usrProvider UserProvider
}

type UserProvider interface {
	Create()
	Get()
	Update()
	Delete()
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
)

func (p *UserService) Create(ctx context.Context, name, surname, email, role string) (string, error) {

}

func (p *UserService) Get(ctx context.Context, uuid string) (models.User, error) {

}

func (p *UserService) Update(ctx context.Context, uuid, name, surname, email, role string) (bool, error) {

}

func (p *UserService) Delete(ctx context.Context, uuid string) (bool, error) {

}
