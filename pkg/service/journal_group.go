package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type JournalGroupsService struct {
	repo repository.JournalGroups
}

func NewJournalGroupsService(repo repository.JournalGroups) *JournalGroupsService {
	return &JournalGroupsService{repo: repo}
}

func (s *JournalGroupsService) GetAll(id int, subject int) ([]restJournal.StudentGrade, error) {
	return s.repo.GetAll(id, subject)
}
