package model_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/mvrilo/go-example-repo-mock/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositorySuite struct {
	suite.Suite
	db   *sqlx.DB
	mock *UserMock
}

// test constructor
func (s *UserRepositorySuite) SetupSuite() {
	db, dbmock, err := sqlmock.New()
	assert.Nil(s.T(), err)
	dbx := sqlx.NewDb(db, "mysql")
	s.mock = NewUserMock(dbmock)
	s.db = dbx
}

func (s *UserRepositorySuite) AfterTest() {
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepositorySuite) TestCreateUser() {
	type testCase struct {
		name    string
		input   *model.User
		err     error
		prepare func(*testCase)
	}

	for _, tt := range []testCase{
		{
			name:  "happy path",
			input: &model.User{Name: "Bender"},
			err:   nil,
			prepare: func(tt *testCase) {
				s.mock.CreateUser(tt.input, int64(1))
			},
		},
		{
			name:  "query error",
			input: &model.User{Name: "Bender"},
			err:   errors.New("random error"),
			prepare: func(tt *testCase) {
				s.mock.CreateUserError(tt.input, tt.err)
			},
		},
	} {
		s.Run(tt.name, func() {
			tt.prepare(&tt)

			user := tt.input
			err := user.CreateUser(s.db)
			if err != nil {
				assert.Error(s.T(), tt.err, err)
				return
			}

			assert.NoError(s.T(), err)
			assert.NotZero(s.T(), tt.input.ID)
		})
	}
}

func (s *UserRepositorySuite) TestGetUserByName() {
	type testCase struct {
		name    string
		input   *model.User
		err     error
		prepare func(*testCase)
	}

	for _, tt := range []*testCase{
		{
			name:  "happy path",
			input: &model.User{Name: "Bender"},
			err:   nil,
			prepare: func(tt *testCase) {
				s.mock.GetUserByName(tt.input)
			},
		},
		{
			name:  "user not found error",
			input: &model.User{},
			err:   model.ErrUserNotFound,
			prepare: func(tt *testCase) {
				s.mock.GetUserByNameError(tt.input, tt.err)
			},
		},
		{
			name:  "query error",
			input: &model.User{},
			err:   errors.New("random error"),
			prepare: func(tt *testCase) {
				s.mock.GetUserByNameError(tt.input, tt.err)
			},
		},
	} {
		s.Run(tt.name, func() {
			tt.prepare(tt)

			user := &model.User{}
			err := user.GetUserByName(s.db, tt.input.Name)
			if err != nil {
				assert.Error(s.T(), tt.err, err)
				return
			}

			assert.NoError(s.T(), err)
			assert.NotZero(s.T(), user.ID)
			assert.Equal(s.T(), tt.input.Name, user.Name)
		})
	}
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
