package login

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
	pb "github.com/klaital/klaital.com/protobufs/gen/go"
)

type Service struct {
	Repo login_repository.Repository
}

func New(repo login_repository.Repository) *Service {
	return &Service{Repo: repo}
}

var ErrPasswordTooWeak = errors.New("password too weak")
var ErrEmailInUse = errors.New("email already in use")

// Register creates a new user account. Errors are returned if the email is already in use, or if the password is not good enough
func (s *Service) Register(ctx context.Context, name, email, passwordRaw string) (*pb.RegisterResponse, error) {
	// TODO: add more checks for password strength
	if len(passwordRaw) < 4 {
		return nil, ErrPasswordTooWeak
	}

	hashedPw, err := login_repository.HashPassword(passwordRaw, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hashing user password: %w", err)
	}
	newId, err := s.Repo.AddUser(ctx, name, email, hashedPw)
	if err != nil {
		return nil, fmt.Errorf("adding user record: %w", err)
	}

	// TODO: use the session store to create a new user session

	resp := &pb.RegisterResponse{
		User: &pb.User{
			Id:       newId,
			Username: name,
			Email:    email,
		},
		SessionToken: "",
	}
	return resp, nil
}
