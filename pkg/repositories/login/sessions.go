package login_repository

import (
	"context"

	pb "github.com/klaital/klaital.com/protobufs/gen/go"
)

type SessionStore interface {
	Create(ctx context.Context, user *pb.User) (string, error)
	Read(ctx context.Context, sessionId string) (*pb.User, error)
	Delete(ctx context.Context, sessionId string) error
}
