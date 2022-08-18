package mock

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mvrilo/go-example-repo-mock/entity"
)

type UserMysqlRepositoryMock struct {
	sqlmock.Sqlmock
}

func NewUserMysqlRepositoryMock(mock sqlmock.Sqlmock) *UserMysqlRepositoryMock {
	return &UserMysqlRepositoryMock{mock}
}

func (m *UserMysqlRepositoryMock) CreateUser(user *entity.User, id int64) {
	m.
		ExpectExec(regexp.QuoteMeta("INSERT INTO users (name) VALUES (?)")).
		WithArgs(user.Name).
		WillReturnResult(sqlmock.NewResult(id, 1))
}

func (m *UserMysqlRepositoryMock) CreateUserError(user *entity.User, err error) {
	m.
		ExpectExec(regexp.QuoteMeta("INSERT INTO users (name) VALUES (?)")).
		WithArgs(user.Name).
		WillReturnError(err)
}

func (m *UserMysqlRepositoryMock) GetUserByName(user *entity.User) {
	m.
		ExpectQuery("SELECT id,name FROM users WHERE name = ?").
		WithArgs(user.Name).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(int64(1), user.Name),
		)
}

func (m *UserMysqlRepositoryMock) GetUserByNameError(user *entity.User, err error) {
	m.
		ExpectQuery("SELECT id,name FROM users WHERE name = ?").
		WithArgs(user.Name).
		WillReturnError(err)
}
