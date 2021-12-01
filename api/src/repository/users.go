package repository

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db sql.DB
}

func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, err) {
	return 0, nil
}
