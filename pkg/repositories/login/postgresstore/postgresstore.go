package login_postgresstore

import (
	"context"
	"embed"
	"fmt"

	"github.com/caarlos0/env/v10"

	kerrors "github.com/klaital/klaital.com/pkg/errors"
	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
	"github.com/klaital/klaital.com/pkg/repositories/login/postgresstore/queries"
	"github.com/klaital/klaital.com/pkg/repositories/utilities"
)

//go:embed migrations/*.sql
var MigrationsFS embed.FS

type Repository struct {
	Queries *queries.Queries
}

type Config struct {
	DbName              string `env:"LOGIN_DB_NAME" envDefault:"webcomics"`
	ConnectionStringFmt string `env:"LOGIN_DB_CONNECT_FMT" envDefault:"postgres://klaital:rootpw@postgres-klaital:5432/%s?sslmode=disable"`
}

func New() (*Repository, error) {
	// Load settings from environment variables
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parsing login_postgresstore env: %w", err)
	}

	// Connect to the database
	db, err := utilities.ConnectAndMigratePostgres(cfg.ConnectionStringFmt, cfg.DbName, MigrationsFS)
	if err != nil {
		return nil, fmt.Errorf("connecting to db for comics_postgresstore: %w", err)
	}

	// Initialize Repository
	repo := &Repository{
		Queries: queries.New(db),
	}

	return repo, nil
}

func (r *Repository) AddUser(ctx context.Context, username, email, passwordDigest string) (uint64, error) {

	id, err := r.Queries.CreateUser(ctx, queries.CreateUserParams{Username: username, Email: email, PasswordDigest: passwordDigest})
	if err != nil {
		return 0, fmt.Errorf("create user query: %w", err)
	}
	return uint64(id), err
}

func (r *Repository) GetUser(ctx context.Context, userId uint64) (*login_repository.User, error) {
	u, err := r.Queries.GetUserById(ctx, int32(userId))
	if err != nil {
		return nil, fmt.Errorf("get user by id query: %w", err)
	}
	return &login_repository.User{
		ID:             userId,
		Username:       u.Username,
		Email:          u.Email,
		PasswordDigest: u.PasswordDigest,
	}, nil
}
func (r *Repository) GetUserByEmail(ctx context.Context, email, passwd string) (*login_repository.User, error) {
	return nil, kerrors.ErrNotImplemented
}
func (r *Repository) UpdateName(ctx context.Context, id uint64, newUsername string) error {
	return kerrors.ErrNotImplemented
}
func (r *Repository) DeleteUser(ctx context.Context, id uint64) error {
	return kerrors.ErrNotImplemented
}
