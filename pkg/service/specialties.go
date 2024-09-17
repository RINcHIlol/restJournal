package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type SpecialtiesService struct {
	repo repository.Specialties
}

func NewSpecialtiesService(repo repository.Specialties) *SpecialtiesService {
	return &SpecialtiesService{repo: repo}
}

func (s *SpecialtiesService) GetAll() ([]restJournal.Specialty, error) {
	return s.repo.GetAll()
}

func (s *SpecialtiesService) GetById(id int) (restJournal.Specialty, error) {
	return s.repo.GetById(id)
}
