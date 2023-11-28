package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/caarlos0/env/v10"

	comics_grpc "github.com/klaital/klaital.com/cmd/server/comics/grpc"
	comics_gateway "github.com/klaital/klaital.com/cmd/server/comics/grpc-gateway"
	login_grpc "github.com/klaital/klaital.com/cmd/server/login/grpc"
	login_gateway "github.com/klaital/klaital.com/cmd/server/login/grpc-gateway"
	comics_repo "github.com/klaital/klaital.com/pkg/repositories/comics"
	comics_postgresstore "github.com/klaital/klaital.com/pkg/repositories/comics/postgresstore"
	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
	login_postgresstore "github.com/klaital/klaital.com/pkg/repositories/login/postgresstore"
	"github.com/klaital/klaital.com/pkg/repositories/login/sessionstore"
	comics_service "github.com/klaital/klaital.com/pkg/service/comics"
	login_service "github.com/klaital/klaital.com/pkg/service/login"
)

type config struct {
	LoginRestAddr   string `env:"LOGIN_REST_ADDR" envDefault:":8080"`
	LoginGrpcAddr   string `env:"LOGIN_GRPC_ADDR" envDefault:":9000"`
	ComicsRestAddr  string `env:"COMICS_REST_ADDR" envDefault:":8081"`
	ComicsGrpcAddr  string `env:"COMICS_GRPC_ADDR" envDefault:":9001"`
	LibraryRestAddr string `env:"LIBRARY_REST_ADDR" envDefault:":8082"`
	LibraryGrpcAddr string `env:"LIBRARY_GRPC_ADDR" envDefault:":9002"`

	LogLevel  int    `env:"LOG_LEVEL" envDefault:"-4"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"text"`        // support "text" and "json"
	LogPretty bool   `env:"LOG_PRETTY_PRINT" envDefault:"false"` // instead of pretty-json, local logging is best done with a text format handler
	LogSource bool   `env:"LOG_SOURCE" envDefault:"true"`        // turning this off can increase performance
}

func main() {
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		slog.Error("Failed to parse config from env", "err", err)
		os.Exit(1)
	}
	logHandlerOpts := &slog.HandlerOptions{
		Level:     slog.Level(cfg.LogLevel),
		AddSource: cfg.LogSource,
	}
	logger := slog.Default()
	switch cfg.LogFormat {
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stderr, logHandlerOpts))
	case "text":
		logger = slog.New(slog.NewTextHandler(os.Stderr, logHandlerOpts))
	default:
		slog.Error("invalid LOG_FORMAT. Defaulting to JSON", "format", cfg.LogFormat)
		logger = slog.New(slog.NewJSONHandler(os.Stderr, logHandlerOpts))
	}
	slog.SetDefault(logger)

	// Load the Repositories
	var comicsRepo comics_repo.Repository
	var loginRepo login_repository.Repository

	comicsRepo, err = comics_postgresstore.New()
	if err != nil {
		slog.Error("Failed to connect to comics Repository", "err", err)
		os.Exit(1)
	}
	loginRepo, err = login_postgresstore.New()
	if err != nil {
		slog.Error("Failed to connect to login Repository", "err", err)
		os.Exit(1)
	}
	sessionStore := sessionstore.New()

	// Initialize the business layers
	comicSvc := comics_service.New(comicsRepo, loginRepo)
	loginSvc := login_service.New(loginRepo, sessionStore)

	// initialize the grpc servers
	comicsGrpcSvc := comics_grpc.New(comicSvc, cfg.ComicsGrpcAddr)
	loginGrpcSvc := login_grpc.New(loginSvc, cfg.LoginGrpcAddr)
	go comicsGrpcSvc.Start()
	go loginGrpcSvc.Start()

	// initialize the grpc-gateway server
	comicsGatewaySvc := comics_gateway.New(cfg.ComicsRestAddr, cfg.ComicsGrpcAddr)
	go comicsGatewaySvc.Start()
	loginGatewaySvc := login_gateway.New(cfg.LoginRestAddr, cfg.LoginGrpcAddr)
	go loginGatewaySvc.Start()

	// listen for the shutdown signal, and gracefully halt the servers
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// Block here until a shutdown signal is received
	<-sigChan
	slog.Info("Shutting down")
	// Shut down the servers
	loginGatewaySvc.Stop()
	comicsGatewaySvc.Stop()
	loginGrpcSvc.Stop()
	comicsGrpcSvc.Stop()
}
