// comics_service provides the business logic for the Comics Reader application
package comics_service

import (
	comics_repo "github.com/klaital/klaital.com/pkg/repositories/comics"
	login_repository "github.com/klaital/klaital.com/pkg/repositories/login"
)

type Service struct {
	ComicsRepo comics_repo.Repository
	LoginRepo  login_repository.Repository
}

func New(comicsRepo comics_repo.Repository, loginRepo login_repository.Repository) *Service {
	return &Service{
		ComicsRepo: comicsRepo,
		LoginRepo:  loginRepo,
	}
}
