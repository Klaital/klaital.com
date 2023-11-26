package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"time"

	comics_service "github.com/klaital/klaital.com/pkg/service/comics"
	pb "github.com/klaital/klaital.com/protobufs/gen/go"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedComicsServer

	Svc        *comics_service.Service
	ListenAddr string

	srv *grpc.Server
	lis net.Listener
}

func New(svc *comics_service.Service, listenAddr string) *Server {
	return &Server{
		Svc:        svc,
		ListenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	s.srv = grpc.NewServer() // add interceptors here
	pb.RegisterComicsServer(s.srv, s)
	var err error
	s.lis, err = net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("start grpc server: %w", err)
	}
	s.srv.Serve(s.lis)
}

func (s *Server) Stop() {
	timeout := 2 * time.Second
	halted := make(chan struct{}, 1)
	go func() {
		s.srv.GracefulStop()
		halted <- struct{}{}
	}()

	select {
	case <-time.After(timeout):
		slog.Error("Failed to gracefully halt gRPC server.")
		s.srv.Stop()
	case <-halted:
		slog.Info("gRPC server gracefully closed")
	}

	s.lis.Close()
}
