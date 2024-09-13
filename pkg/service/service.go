package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type Authorization interface {
	CreateUser(user restJournal.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

//more

//more

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
