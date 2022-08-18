package model_test

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mvrilo/go-example-repo-mock/model"
)

type UserMock struct {
	sqlmock.Sqlmock
}

func NewUserMock(mock sqlmock.Sqlmock) *UserMock {
	return &UserMock{mock}
}

func (m *UserMock) CreateUser(user *model.User, id int64) {
	m.
		ExpectExec(regexp.QuoteMeta("INSERT INTO users (name) VALUES (?)")).
		WithArgs(user.Name).
		WillReturnResult(sqlmock.NewResult(id, 1))
}

func (m *UserMock) CreateUserError(user *model.User, err error) {
	m.
		ExpectExec(regexp.QuoteMeta("INSERT INTO users (name) VALUES (?)")).
		WithArgs(user.Name).
		WillReturnError(err)
}

func (m *UserMock) GetUserByName(user *model.User) {
	m.
		ExpectQuery("SELECT id,name FROM users WHERE name = ?").
		WithArgs(user.Name).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(int64(1), user.Name),
		)
}

func (m *UserMock) GetUserByNameError(user *model.User, err error) {
	m.
		ExpectQuery("SELECT id,name FROM users WHERE name = ?").
		WithArgs(user.Name).
		WillReturnError(err)
}
