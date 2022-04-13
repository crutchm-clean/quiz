package repository

import (
	"Kahoot"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func (a AuthRepository) CreateUser(user Kahoot.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s(id, login, username, email, password) VALUES($1,$2,$3,$4,$5) RETURNING id", usersTable)
	row := a.db.QueryRow(query, uuid.New().String()[:8], user.Login, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (a AuthRepository) GetUser(username, password string) (Kahoot.User, error) {
	var user Kahoot.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := a.db.Get(&user, query, username, password)
	return user, err
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
