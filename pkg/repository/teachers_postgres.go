package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type TeachersPostgres struct {
	db *sqlx.DB
}

func NewTeachersPostgres(db *sqlx.DB) *TeachersPostgres {
	return &TeachersPostgres{db: db}
}

func (r *TeachersPostgres) GetAll() ([]restJournal.UserParse, error) {
	var teachers []restJournal.UserParse
	query := fmt.Sprintf(`SELECT u.name, u.surname, u.email, u.role, s.name AS
    specialty FROM %s u JOIN %s s ON u.specialty_id = s.id WHERE u.role = 'teacher'`, usersTable, specialtiesTable)
	err := r.db.Select(&teachers, query)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *TeachersPostgres) GetById(teacherId int) (restJournal.UserParse, error) {
	var teacher restJournal.UserParse
	query := fmt.Sprintf(`SELECT u.name, u.surname, u.email, u.role, s.name AS
    specialty FROM %s u JOIN %s s ON u.specialty_id = s.id WHERE u.role = 'teacher' and u.id = $1`, usersTable, specialtiesTable)
	err := r.db.Get(&teacher, query, teacherId)
	if err != nil {
		return restJournal.UserParse{}, err
	}
	return teacher, nil
}
