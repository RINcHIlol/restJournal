package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type Authorization interface {
	CreateUser(user restJournal.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, string, error)
}

type Students interface {
	GetAll() ([]restJournal.UserParse, error)
	GetById(studentId int) (restJournal.UserParse, error)
}

//more

//more

type Service struct {
	Authorization
	Students
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Students:      NewStudentsService(repos.Students),
	}
}
