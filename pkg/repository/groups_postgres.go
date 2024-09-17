package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type GroupsPostgres struct {
	db *sqlx.DB
}

func NewGroupsPostgres(db *sqlx.DB) *GroupsPostgres {
	return &GroupsPostgres{db: db}
}

func (r *GroupsPostgres) GetAll() ([]restJournal.Group, error) {
	var groups []restJournal.Group
	query := fmt.Sprintf(`SELECT * FROM %s`, groupsTable)
	err := r.db.Select(&groups, query)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *GroupsPostgres) GetById(teacherId int) (restJournal.Group, error) {
	var group restJournal.Group
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, groupsTable)
	err := r.db.Get(&group, query, teacherId)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (r *GroupsPostgres) GetAllStudents(groupId int) ([]restJournal.UserParse, error) {
	var users []restJournal.UserParse
	query := fmt.Sprintf(`SELECT u.name, u.surname, u.email, u.role, s.name AS specialty FROM %s u JOIN %s gu 
        ON u.id = gu.user_id  JOIN %s g ON g.id = gu.group_id JOIN %s s ON g.specialty_id = s.id 
        WHERE g.id = $1 AND u.role = 'student'`, usersTable, groupUsersTable, groupsTable, specialtiesTable)
	err := r.db.Select(&users, query, groupId)
	if err != nil {
		return nil, err
	}
	return users, nil
}
