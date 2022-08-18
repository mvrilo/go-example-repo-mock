package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/mvrilo/go-example-repo-mock/entity"
	mock "github.com/mvrilo/go-example-repo-mock/mock/mysql"
	"github.com/mvrilo/go-example-repo-mock/repository"
	mysqlrepo "github.com/mvrilo/go-example-repo-mock/repository/mysql"
)

type UserRepositorySuite struct {
	suite.Suite
	db             *sqlx.DB
	mock           *mock.UserMysqlRepositoryMock
	userRepository repository.UserRepository
}

// test constructor
func (s *UserRepositorySuite) SetupSuite() {
	db, dbmock, err := sqlmock.New()
	assert.Nil(s.T(), err)
	dbx := sqlx.NewDb(db, "mysql")
	s.db = dbx
	s.mock = mock.NewUserMysqlRepositoryMock(dbmock)
	s.userRepository = mysqlrepo.NewUserMysqlRepository(dbx)
}

func (s *UserRepositorySuite) AfterTest() {
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepositorySuite) TestCreateUser() {
	type testCase struct {
		name    string
		input   *entity.User
		err     error
		prepare func(*testCase)
	}

	for _, tt := range []testCase{
		{
			name:  "happy path",
			input: &entity.User{Name: "Bender"},
			err:   nil,
			prepare: func(tt *testCase) {
				s.mock.CreateUser(tt.input, int64(1))
			},
		},
		{
			name:  "query error",
			input: &entity.User{Name: "Bender"},
			err:   errors.New("random error"),
			prepare: func(tt *testCase) {
				s.mock.CreateUserError(tt.input, tt.err)
			},
		},
	} {
		s.Run(tt.name, func() {
			tt.prepare(&tt)

			err := s.userRepository.CreateUser(context.Background(), tt.input)
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
		input   *entity.User
		err     error
		prepare func(*testCase)
	}

	for _, tt := range []*testCase{
		{
			name:  "happy path",
			input: &entity.User{Name: "Bender"},
			err:   nil,
			prepare: func(tt *testCase) {
				s.mock.GetUserByName(tt.input)
			},
		},
		{
			name:  "user not found error",
			input: &entity.User{},
			err:   entity.ErrUserNotFound,
			prepare: func(tt *testCase) {
				s.mock.GetUserByNameError(tt.input, tt.err)
			},
		},
		{
			name:  "query error",
			input: &entity.User{},
			err:   errors.New("random error"),
			prepare: func(tt *testCase) {
				s.mock.GetUserByNameError(tt.input, tt.err)
			},
		},
	} {
		s.Run(tt.name, func() {
			tt.prepare(tt)

			user, err := s.userRepository.GetUserByName(context.Background(), tt.input.Name)
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
