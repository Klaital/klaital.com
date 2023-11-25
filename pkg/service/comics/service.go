// comics_service provides the business logic for the Comics Reader application
package comics_service

import (
	comics_repo "github.com/klaital/klaital.com/pkg/repositories/comics"
)

type Service struct {
	Repo comics_repo.Repository
}

func New(repo comics_repo.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}
