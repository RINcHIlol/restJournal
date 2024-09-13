package repository

import (
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type Authorization interface {
	CreateUser(user restJournal.User) (int, error)
	GetUser(email, password string) (restJournal.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
