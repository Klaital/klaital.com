package main

import (
	"log/slog"
	"os"

	comics_repo "github.com/klaital/klaital.com/pkg/repositories/comics"
	"github.com/klaital/klaital.com/pkg/repositories/comics/postgresstore"
	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
	login_postgresstore "github.com/klaital/klaital.com/pkg/repositories/login/postgresstore"
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

	// TODO: initialize the grpc server
	// TODO: initialize the grpc-gateway server
}
