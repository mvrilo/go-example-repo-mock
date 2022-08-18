package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID   int64
	Name string
}
