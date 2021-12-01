package repository

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare("insert into users (name, nick, mail, password) values (?, ?. ?. ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Mail, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil

}
