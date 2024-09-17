package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type SpecialtiesPostgres struct {
	db *sqlx.DB
}

func NewSpecialtiesPostgres(db *sqlx.DB) *SpecialtiesPostgres {
	return &SpecialtiesPostgres{db: db}
}

func (s *SpecialtiesPostgres) GetAll() ([]restJournal.Specialty, error) {
	var specialties []restJournal.Specialty
	query := fmt.Sprintf("SELECT * FROM %s", specialtiesTable)
	err := s.db.Select(&specialties, query)
	if err != nil {
		return nil, err
	}
	return specialties, nil
}

func (s *SpecialtiesPostgres) GetById(id int) (restJournal.Specialty, error) {
	var specialty restJournal.Specialty
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", specialtiesTable)
	err := s.db.Get(&specialty, query, id)
	if err != nil {
		return specialty, err
	}
	return specialty, nil
}
