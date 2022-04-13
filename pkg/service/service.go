package service

import (
	"Kahoot"
	"Kahoot/pkg/repository"
)

type Authorisation interface {
	CreateUser(user Kahoot.User) (string, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}
type Service struct {
	Authorisation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
	}
}
