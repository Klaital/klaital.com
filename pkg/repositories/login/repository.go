package login_repository

import (
	"context"
	"errors"
)

type Repository interface {
	AddUser(ctx context.Context, username, email, password string) (uint64, error)
	GetUser(ctx context.Context, userId uint64) (*User, error)
	GetUserByEmail(ctx context.Context, email, passwd string) (*User, error)
	UpdateName(ctx context.Context, id uint64, newUsername string) error
	DeleteUser(ctx context.Context, id uint64) error
}

type User struct {
	ID             uint64
	Username       string
	Email          string
	PasswordDigest string
}

var ErrPasswordMismatch = errors.New("incorrect password")
var ErrEmailInUse = errors.New("email in use")
