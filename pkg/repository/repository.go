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

type Repository struct {
	Authorization
	Students
	Teachers
	Groups
	Specialties
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Students:      NewStudentsPostgres(db),
		Teachers:      NewTeachersPostgres(db),
		Groups:        NewGroupsPostgres(db),
		Specialties:   NewSpecialtiesPostgres(db),
	}
}
