package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type StudentsPostgres struct {
	db *sqlx.DB
}

func NewStudentsPostgres(db *sqlx.DB) *StudentsPostgres {
	return &StudentsPostgres{db: db}
}

func (r *StudentsPostgres) GetAll() ([]restJournal.UserParse, error) {
	var students []restJournal.UserParse
	query := fmt.Sprintf(`SELECT u.name, u.surname, u.email, u.role, s.name AS
    specialty FROM %s u JOIN %s s ON u.specialty_id = s.id WHERE u.role = 'student'`, usersTable, specialtiesTable)
	err := r.db.Select(&students, query)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentsPostgres) GetById(studentId int) (restJournal.UserParse, error) {
	var student restJournal.UserParse
	query := fmt.Sprintf(`SELECT u.name, u.surname, u.email, u.role, s.name AS
    specialty FROM %s u JOIN %s s ON u.specialty_id = s.id WHERE u.role = 'student' and u.id = $1`, usersTable, specialtiesTable)
	err := r.db.Get(&student, query, studentId)
	if err != nil {
		return restJournal.UserParse{}, err
	}
	return student, nil
}
