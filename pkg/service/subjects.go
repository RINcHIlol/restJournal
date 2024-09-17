package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type SubjectsService struct {
	repo repository.Subjects
}

func NewSubjectsService(repo repository.Subjects) *SubjectsService {
	return &SubjectsService{repo: repo}
}

func (s *SubjectsService) GetAll() ([]restJournal.Subjects, error) {
	return s.repo.GetAll()
}

func (s *SubjectsService) GetById(id int) (restJournal.Subjects, error) {
	return s.repo.GetById(id)
}

func (s *SubjectsService) GetSubjectsBySpecialty(id int) ([]restJournal.Subjects, error) {
	return s.repo.GetBySpec(id)
}
