package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mvrilo/go-example-repo-mock/entity"
	"github.com/mvrilo/go-example-repo-mock/repository"
)

// ensure it's implemented
var _ repository.UserRepository = (*UserMysqlRepository)(nil)

// implements UserRepository
type UserMysqlRepository struct {
	db *sqlx.DB
}

func NewUserMysqlRepository(db *sqlx.DB) *UserMysqlRepository {
	return &UserMysqlRepository{db}
}

func (r *UserMysqlRepository) CreateUser(ctx context.Context, user *entity.User) error {
	res, err := r.db.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		return err
	}
	user.ID, err = res.LastInsertId()
	return err
}

func (r *UserMysqlRepository) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User
	err := r.db.GetContext(ctx, &user, "SELECT id,name FROM users WHERE name = ?", name)
	if err == sql.ErrNoRows {
		return nil, entity.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
