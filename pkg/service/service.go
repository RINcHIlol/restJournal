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

type Teachers interface {
	GetAll() ([]restJournal.UserParse, error)
	GetById(teacherId int) (restJournal.UserParse, error)
}

type Groups interface {
	GetAll() ([]restJournal.Group, error)
	GetById(groupId int) (restJournal.Group, error)
	GetAllStudents(groupId int) ([]restJournal.UserParse, error)
}

type Specialties interface {
	GetAll() ([]restJournal.Specialty, error)
	GetById(id int) (restJournal.Specialty, error)
}

type Subjects interface {
	GetAll() ([]restJournal.Subjects, error)
	GetById(id int) (restJournal.Subjects, error)
	GetSubjectsBySpecialty(id int) ([]restJournal.Subjects, error)
}

type JournalGroups interface {
	GetAll(id int, subject int) ([]restJournal.StudentGrade, error)
}

// в put проверяй id юзера на id учителя тип ток он может менять чо поставил
type JournalStudents interface {
	GetStudentGrades(studentId, teacherId int) ([]restJournal.StudentGrade, error)
	PutStudentGrade(gradeId, teacherId, grade int) error
	PostStudentGrade(grade restJournal.Grade) error
	DeleteStudentGrade(gradeId int) error
}

//more

//more

type Service struct {
	Authorization
	Students
	Teachers
	Groups
	Specialties
	Subjects
	JournalGroups
	JournalStudents
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repos.Authorization),
		Students:        NewStudentsService(repos.Students),
		Teachers:        NewTeachersService(repos.Teachers),
		Groups:          NewGroupsService(repos.Groups),
		Specialties:     NewSpecialtiesService(repos.Specialties),
		Subjects:        NewSubjectsService(repos.Subjects),
		JournalGroups:   NewJournalGroupsService(repos.JournalGroups),
		JournalStudents: NewJournalStudentsService(repos.JournalStudents),
	}
}
