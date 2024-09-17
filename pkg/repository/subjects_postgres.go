package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type SubjectsPostgres struct {
	db *sqlx.DB
}

func NewSubjectsPostgres(db *sqlx.DB) *SubjectsPostgres {
	return &SubjectsPostgres{db: db}
}

func (r *SubjectsPostgres) GetAll() ([]restJournal.Subjects, error) {
	var subjects []restJournal.Subjects

	query := fmt.Sprintf(`SELECT * FROM %s`, subjectsTable)
	err := r.db.Select(&subjects, query)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func (r *SubjectsPostgres) GetById(id int) (restJournal.Subjects, error) {
	//todo
}

func (r *SubjectsPostgres) GetBySpec(id int) ([]restJournal.Subjects, error) {
	//todo
}
