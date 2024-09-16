package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type StudentsService struct {
	repo repository.Students
}

func NewStudentsService(repo repository.Students) *StudentsService {
	return &StudentsService{repo: repo}
}

func (s *StudentsService) GetAll() ([]restJournal.UserParse, error) {
	return s.repo.GetAll()
}

func (s *StudentsService) GetById(studentId int) (restJournal.UserParse, error) {
	return s.repo.GetById(studentId)
}
