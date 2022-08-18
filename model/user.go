package model

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID   int64
	Name string
}

func (u *User) CreateUser(db *sqlx.DB) error {
	res, err := db.Exec("INSERT INTO users (name) VALUES (?)", u.Name)
	if err != nil {
		return err
	}
	u.ID, err = res.LastInsertId()
	return err
}

func (u *User) GetUserByName(db *sqlx.DB, name string) error {
	err := db.Get(u, "SELECT id,name FROM users WHERE name = ?", name)
	if err == sql.ErrNoRows {
		return ErrUserNotFound
	}
	return err
}
