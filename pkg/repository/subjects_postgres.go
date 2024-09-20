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
	var subject restJournal.Subjects

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, subjectsTable)
	err := r.db.Get(&subject, query, id)
	if err != nil {
		return subject, err
	}
	return subject, nil
}

func (r *SubjectsPostgres) GetBySpec(id int) ([]restJournal.Subjects, error) {
	var subjects []restJournal.Subjects

	query := fmt.Sprintf(`SELECT s.name, s.description FROM %s ss JOIN %s s ON ss.subject_id = s.id
                                 WHERE ss.specialty_id = $1;`, specialtySubjectsTable, subjectsTable)
	err := r.db.Select(&subjects, query, id)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}
