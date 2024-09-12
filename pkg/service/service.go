package service

import "rest_journal/pkg/repository"

type Authorization interface{}

//more

//more

type Service struct {
	Authorization Authorization
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
