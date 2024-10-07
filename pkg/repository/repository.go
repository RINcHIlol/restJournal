package repository

import (
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type Authorization interface {
	CreateUser(user restJournal.User) (int, error)
	GetUser(email, password string) (restJournal.User, error)
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
	GetBySpec(id int) ([]restJournal.Subjects, error)
}

type JournalGroups interface {
	GetAll(id int, subject int) ([]restJournal.StudentGrade, error)
}

// в put проверяй id юзера на id учителя тип ток он может менять чо поставил
type JournalStudents interface {
	GetById(studentId, teacherId int) ([]restJournal.StudentGrade, error)
	PutById(gradeId, teacherId, grade int) error
	Post(grade restJournal.Grade) error
	DeleteById(gradeId int) error
}

type Repository struct {
	Authorization
	Students
	Teachers
	Groups
	Specialties
	Subjects
	JournalGroups
	JournalStudents
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		Students:        NewStudentsPostgres(db),
		Teachers:        NewTeachersPostgres(db),
		Groups:          NewGroupsPostgres(db),
		Specialties:     NewSpecialtiesPostgres(db),
		Subjects:        NewSubjectsPostgres(db),
		JournalGroups:   NewJournalGroupsPostgres(db),
		JournalStudents: NewJournalStudentsPostgres(db),
	}
}
