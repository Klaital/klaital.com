package sessionstore

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	pb "github.com/klaital/klaital.com/protobufs/gen/go"
)

type Repository struct {
	db                *redis.Client
	sessionExpiration time.Duration
}

type Config struct {
	RedisHost      string        `env:"SESSIONS_HOST" envDefault:"sessionstore-klaital:6379"`
	RedisPassword  string        `env:"SESSIONS_PASSWORD" envDefault:""`
	RedisDb        int           `env:"SESSIONS_DB" envDefault:"0"`
	SessionTimeout time.Duration `env:"SESSIONS_TIMEOUT" envDefault:"168h"`
}

func New() *Repository {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		slog.Error("Failed to parse Session Store config from env", "err", err)
		os.Exit(1)
	}
	// TODO: load config from env
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDb,
	})

	return &Repository{
		db:                rdb,
		sessionExpiration: cfg.SessionTimeout,
	}
}

func (r *Repository) Get(ctx context.Context, sessionId string) (*pb.User, error) {
	var u pb.User
	err := r.db.Get(ctx, sessionId).Scan(&u)
	if err != nil {
		return nil, fmt.Errorf("fetch user session: %w", err)
	}

	return &u, nil
}

func (r *Repository) Set(ctx context.Context, user *pb.User) (sessionId string, err error) {
	// Generate random session ID
	sessionId = uuid.NewString()
	// Save the user data in redis against the session ID
	_, err = r.db.SetNX(ctx, sessionId, user, r.sessionExpiration).Result()
	if err != nil {
		return "", fmt.Errorf("saving user session: %w", err)
	}
	// Success!
	return sessionId, nil
}
