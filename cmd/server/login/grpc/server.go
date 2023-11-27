package login_grpc

import (
	"fmt"
	"log/slog"
	"net"
	"time"

	service "github.com/klaital/klaital.com/pkg/service/login"
	pb "github.com/klaital/klaital.com/protobufs/gen/go"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedLoginServer

	Svc        *service.Service
	ListenAddr string

	srv *grpc.Server
	lis net.Listener
}

func New(svc *service.Service, listenAddr string) *Server {
	return &Server{
		Svc:        svc,
		ListenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	s.srv = grpc.NewServer() // add interceptors here
	pb.RegisterLoginServer(s.srv, s)
	var err error
	s.lis, err = net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("start grpc server: %w", err)
	}
	return s.srv.Serve(s.lis)
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
