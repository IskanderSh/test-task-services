package storage

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/IskanderSh/test-task-services/user-services/internal/config"
	"github.com/IskanderSh/test-task-services/user-services/internal/domain/models"
	"github.com/IskanderSh/test-task-services/user-services/internal/lib/error/wrapper"
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
	connectionStr := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.DB, dbConfig.Port)
	fmt.Println(connectionStr)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &UserStorage{db: db}, nil
}

func (s *UserStorage) GetUser(ctx context.Context, uuid string) (*models.User, error) {
	const op = "storage.GetUser"

	var user models.User

	row := s.db.QueryRow(GetUserQuery, uuid)

	if err := row.Scan(&user.Uuid, &user.Name, &user.Surname, &user.Email, &user.Role); err != nil {
		return nil, wrapper.Wrap(op, ErrUserNotFound)
	}

	return &user, nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, uuid string, updateUser models.UpdateUser) (bool, error) {
	const op = "storage.UpdateUser"

	var updateQuery bytes.Buffer
	updateQuery.WriteString(UpdateUserQuery)

	updateParams := make([]interface{}, 1)
	paramCount := 1

	if updateUser.Name != nil {
		updateQuery.WriteString(fmt.Sprintf(" name=$%d", paramCount))
		updateParams = append(updateParams, *updateUser.Name)
		paramCount += 1
	}

	if updateUser.Surname != nil {
		updateQuery.WriteString(fmt.Sprintf(" surname=$%d", paramCount))
		updateParams = append(updateParams, *updateUser.Surname)
		paramCount += 1
	}

	if updateUser.Email != nil {
		updateQuery.WriteString(fmt.Sprintf(" email=$%d", paramCount))
		updateParams = append(updateParams, *updateUser.Email)
		paramCount += 1
	}

	if updateUser.Role != nil {
		updateQuery.WriteString(fmt.Sprintf(" role=$%d", paramCount))
		updateParams = append(updateParams, *updateUser.Role)
		paramCount += 1
	}

	updateQuery.WriteString(fmt.Sprintf(" where uuid=$%d", paramCount))
	updateParams = append(updateParams, uuid)

	queryString := updateQuery.String()

	_, err := s.db.Exec(queryString, updateParams...)
	if err != nil {
		return false, wrapper.Wrap(op, err)
	}

	return true, nil
}

func (s *UserStorage) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	const op = "storage.DeleteUser"

	_, err := s.db.Exec(DeleteUserQuery, uuid)
	if err != nil {
		return false, wrapper.Wrap(op, err)
	}

	return true, nil
}
