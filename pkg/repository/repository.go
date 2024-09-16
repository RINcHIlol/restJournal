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

type Repository struct {
	Authorization
	Students
	Teachers
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Students:      NewStudentsPostgres(db),
		Teachers:      NewTeachersPostgres(db),
	}
}
