package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type TeachersService struct {
	repo repository.Students
}

func NewTeachersService(repo repository.Teachers) *TeachersService {
	return &TeachersService{repo: repo}
}

func (s *TeachersService) GetAll() ([]restJournal.UserParse, error) {
	return s.repo.GetAll()
}

func (s *TeachersService) GetById(studentId int) (restJournal.UserParse, error) {
	return s.repo.GetById(studentId)
}
