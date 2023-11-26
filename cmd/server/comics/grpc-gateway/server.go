package grpcgateway

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/klaital/klaital.com/protobufs/gen/go"
)

type Server struct {
	RestListenAddr string
	GrpcListenAddr string

	srv *http.Server
}

func New(restListenAddr string, grpcListenAddr string) *Server {
	return &Server{
		RestListenAddr: restListenAddr,
		GrpcListenAddr: grpcListenAddr,
	}
}

func (s *Server) Start() error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterComicsHandlerFromEndpoint(context.Background(), mux, s.GrpcListenAddr, opts)
	if err != nil {
		return fmt.Errorf("starting gateway server: %w", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	s.srv = &http.Server{Addr: s.RestListenAddr}
	s.srv.Handler = mux
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := s.srv.Shutdown(ctx)
	slog.Error("Failed to gracefully shut down gateway server", "err", err)
}
