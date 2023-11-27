package login_grpc

import (
	"context"
	"log/slog"

	pb "github.com/klaital/klaital.com/protobufs/gen/go"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp, err := s.Svc.Register(ctx, req.Username, req.Email, req.Password)

	if err != nil {
		slog.Error("Failed to register new user", "err", err)
		return nil, err // TODO: convert this into a proper grpc error
	}

	return resp, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := s.Svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		slog.Error("Failed to log user in", "err", err)
		return nil, err // TODO: convert this into a proper grpc error
	}

	return resp, nil
}
