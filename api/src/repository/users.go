package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare("insert into users (name, nick, mail, password) values (?, ?, ?, ?)")
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

func (repo Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	lines, err := repo.db.Query(
		"select id, name, nick, mail, createin from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Mail,
			&user.CreateIn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo Users) SearchByID(ID uint64) (models.User, error) {
	lines, err := repo.db.Query(
		"select id, name, nick, mail, createin from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Mail,
			&user.CreateIn,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo Users) Update(ID uint64, user models.User) error {
	statement, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, mail = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Mail, ID); err != nil {
		return err
	}

	return nil
}

func (repo Users) Delete(ID uint64) error {
	statement, err := repo.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
