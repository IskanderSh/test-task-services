package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Suplab-Team/test-task-go/tree/IskanderSh/user-service/internal/config"
	_ "github.com/lib/pq"
)

type UserStorage struct {
	db *sql.DB
}

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

func (s *UserStorage) CreateUser() {

}

func (s *UserStorage) GetUser() {}

func (s *UserStorage) UpdateUser() {}

func (s *UserStorage) DeleteUser() {}
