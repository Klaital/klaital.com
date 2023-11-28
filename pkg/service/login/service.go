package login

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
	"github.com/klaital/klaital.com/pkg/repositories/login/sessionstore"
	pb "github.com/klaital/klaital.com/protobufs/gen/go"
)

type Service struct {
	Repo     login_repository.Repository
	Sessions *sessionstore.Repository
}

func New(repo login_repository.Repository, sessions *sessionstore.Repository) *Service {
	return &Service{
		Repo:     repo,
		Sessions: sessions,
	}
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
	u := &pb.User{
		Id:       newId,
		Username: name,
		Email:    email,
	}
	sessionToken, err := s.Sessions.Set(ctx, u)
	if err != nil {
		slog.Error("Failed to log in newly-registered user", "err", err)
	}
	resp := &pb.RegisterResponse{
		User:         u,
		SessionToken: sessionToken,
	}
	return resp, nil
}

func (s *Service) Login(ctx context.Context, email, rawPassword string) (*pb.LoginResponse, error) {
	// Pull user data from database
	u, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("login fetching user data: %w", err)
	}

	if !login_repository.CheckPassword([]byte(u.PasswordDigest), []byte(rawPassword)) {
		return nil, login_repository.ErrPasswordMismatch
	}

	pbUser := &pb.User{
		Id:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
	sessionToken, err := s.Sessions.Set(ctx, pbUser)
	if err != nil {
		slog.Error("Failed to log in user", "err", err)
		return nil, err
	}
	resp := &pb.LoginResponse{
		User:         pbUser,
		SessionToken: sessionToken,
	}
	return resp, nil
}
