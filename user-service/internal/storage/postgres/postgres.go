package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/config"
	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/domain/models"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type UserStorage struct {
	db *sql.DB
}

var (
	ErrUserNotFound = errors.New("user with such uuid not found")
)

func New(dbConfig *config.Storage) (*UserStorage, error) {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.DB, dbConfig.Port)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &UserStorage{db: db}, nil
}

func (s *UserStorage) GetUser(ctx context.Context, uuid string) (*models.User, error) {}

func (s *UserStorage) UpdateUser(ctx context.Context, uuid, name, surname, email, role string) (bool, error) {
}

func (s *UserStorage) DeleteUser(ctx context.Context, uuid string) (bool, error) {}
