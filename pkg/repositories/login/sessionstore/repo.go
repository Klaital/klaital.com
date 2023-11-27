package sessionstore

import (
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rdb *redis.Client
}

type Config struct {
}

func New() *Repository {
	// TODO: load config from env

}
