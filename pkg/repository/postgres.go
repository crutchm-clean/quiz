package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

func NewPostgresDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres:22334455@localhost:5432/kahoot?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
