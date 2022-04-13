package repository

import (
	"Kahoot"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user Kahoot.User) (string, error)
	GetUser(username, password string) (Kahoot.User, error)
}

type Repository struct {
	Authorisation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorisation: NewAuthRepository(db)}
}
