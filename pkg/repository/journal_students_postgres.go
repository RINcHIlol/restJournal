package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type JournalStudentsPostgres struct {
	db *sqlx.DB
}

func NewJournalStudentsPostgres(db *sqlx.DB) *JournalStudentsPostgres {
	return &JournalStudentsPostgres{db: db}
}

func (r *JournalStudentsPostgres) GetById(studentId int, teacherId int) ([]restJournal.StudentGrade, error) {
	var grades []restJournal.StudentGrade
	var gradeId int

	queryForId := `SELECT teacher_id FROM grades WHERE user_id = $1 LIMIT 1;` // Извлекаем первый id оценки
	err := r.db.QueryRow(queryForId, studentId).Scan(&gradeId)
	if err != nil {
		return nil, fmt.Errorf("failed to get grade ID: %w", err)
	}

	if gradeId != teacherId {
		return nil, fmt.Errorf("you can't get that")
	}

	query := fmt.Sprintf(`SELECT u.name AS student_name, u.surname AS student_surname, g.grade, s.name
    AS subject_name, g.created_at FROM %s g JOIN %s u ON g.user_id = u.id JOIN %s s ON g.subject_id = s.id
    WHERE u.id = $1;`, gradesTable, usersTable, subjectsTable)

	err = r.db.Select(&grades, query, studentId)
	if err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *JournalStudentsPostgres) PutById(gradeId int) error {
	return nil
}

func (r *JournalStudentsPostgres) Post(grade restJournal.Grade) error {
	return nil
}

func (r *JournalStudentsPostgres) DeleteById(gradeId int) error {
	return nil
}
