package repository

import (
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type JournalGroupsPostgres struct {
	db *sqlx.DB
}

func NewJournalGroupsPostgres(db *sqlx.DB) *JournalGroupsPostgres {
	return &JournalGroupsPostgres{db: db}
}

func (r *JournalGroupsPostgres) GetAll(id int, subject int) ([]restJournal.StudentGrade, error) {
	var students []restJournal.StudentGrade

	query := `SELECT u.name AS student_name, u.surname AS student_surname, g.grade, s.name AS subject_name, g.created_at
        FROM grades g JOIN users u ON g.user_id = u.id JOIN subjects s ON g.subject_id = s.id
        WHERE g.group_id = $1 AND g.subject_id = $2;`

	err := r.db.Select(&students, query, id, subject)
	if err != nil {
		return nil, err
	}

	return students, nil
}
