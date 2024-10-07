package service

import (
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
)

type JournalStudentsService struct {
	repo repository.JournalStudents
}

func NewJournalStudentsService(repo repository.JournalStudents) *JournalStudentsService {
	return &JournalStudentsService{repo: repo}
}

func (s *JournalStudentsService) GetStudentGrades(studentId, teacherId int) ([]restJournal.StudentGrade, error) {
	return s.repo.GetById(studentId, teacherId)
}

func (s *JournalStudentsService) PutStudentGrade(gradeId, teacherId, grade int) error {
	return s.repo.PutById(gradeId, teacherId, grade)
}

func (s *JournalStudentsService) PostStudentGrade(grade restJournal.Grade) error {
	return s.repo.Post(grade)
}

func (s *JournalStudentsService) DeleteStudentGrade(gradeId int) error {
	return s.repo.DeleteById(gradeId)
}
