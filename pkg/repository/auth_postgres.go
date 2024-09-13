package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	restJournal "rest_journal"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user restJournal.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, surname, email, password, role) values ($1, $2, $3, $4, $5) RETURNING id`, usersTable)
	err := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Password, user.Role).Scan(&id)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23514" {
				return 0, errors.New("invalid role: must be one of 'student', 'teacher', 'headTeacher'")
			}
		}
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (restJournal.User, error) {
	var user restJournal.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE email = $1 AND password = $2`, usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
