package repository

import (
	"context"

	"github.com/mvrilo/go-example-repo-mock/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByName(ctx context.Context, name string) (*entity.User, error)
}
