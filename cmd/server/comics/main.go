package main

import (
	"log/slog"
	"os"

	"github.com/klaital/klaital.com/cmd/server/comics/grpc"
	grpcgateway "github.com/klaital/klaital.com/cmd/server/comics/grpc-gateway"
	comics_repo "github.com/klaital/klaital.com/pkg/repositories/comics"
	"github.com/klaital/klaital.com/pkg/repositories/comics/postgresstore"
	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
	login_postgresstore "github.com/klaital/klaital.com/pkg/repositories/login/postgresstore"
	comics_service "github.com/klaital/klaital.com/pkg/service/comics"
)

func main() {
	// Load the comics and login Repositories
	var comicsRepo comics_repo.Repository
	var loginRepo login_repository.Repository
	var err error

	comicsRepo, err = postgresstore.New()
	if err != nil {
		slog.Error("Failed to connect to comics Repository", "err", err)
		os.Exit(1)
	}
	loginRepo, err = login_postgresstore.New()
	if err != nil {
		slog.Error("Failed to connect to login Repository", "err", err)
		os.Exit(1)
	}

	// Initialize the business layer
	comicSvc := comics_service.New(comicsRepo, loginRepo)

	// initialize the grpc server
	grpcSvc := grpc.New(comicSvc, ":8080")
	grpcSvc.Start()

	// initialize the grpc-gateway server
	gatewaySvc := grpcgateway.New(":8080", ":9000")
	gatewaySvc.Start()

	// TODO: listen for the shutdown signal, and gracefully halt the servers
}
