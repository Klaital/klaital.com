package grpc

import comics_service "github.com/klaital/klaital.com/pkg/service/comics"

type Server struct {
	Svc *comics_service.Service
}

func New(svc *comics_service.Service) *Server {
	return &Server{Svc: svc}
}
