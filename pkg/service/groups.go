package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type GroupsService struct {
	repo repository.Groups
}

func NewGroupsService(repo repository.Groups) *GroupsService {
	return &GroupsService{repo: repo}
}

func (s *GroupsService) GetAll() ([]restJournal.Group, error) {
	return s.repo.GetAll()
}

func (s *GroupsService) GetById(groupId int) (restJournal.Group, error) {
	return s.repo.GetById(groupId)
}

func (s *GroupsService) GetAllStudents(groupId int) ([]restJournal.UserParse, error) {
	return s.repo.GetAllStudents(groupId)
}
