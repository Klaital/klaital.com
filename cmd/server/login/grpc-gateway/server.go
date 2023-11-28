package login_gateway

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
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
	err := pb.RegisterLoginHandlerFromEndpoint(context.Background(), mux, s.GrpcListenAddr, opts)
	if err != nil {
		return fmt.Errorf("starting gateway server: %w", err)
	}

	withCors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*", "https://login.klaital.com"},
		// AllowOriginFunc: func(origin string) bool { return true },
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
		Logger:           cors.Logger(log.New(os.Stderr, "", 0)),
	}).Handler(mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	s.srv = &http.Server{Addr: s.RestListenAddr, Handler: withCors}
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := s.srv.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to gracefully shut down gateway server", "err", err)
	} else {
		slog.Info("Shut down login gateway server")
	}
}
