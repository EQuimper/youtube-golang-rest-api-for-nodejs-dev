package handlers

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"todo/domain"
)

type Server struct {
	domain *domain.Domain
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
}

func NewServer(domain *domain.Domain) *Server {
	return &Server{domain: domain}
}

func SetupRouter(domain *domain.Domain) *chi.Mux {
	server := NewServer(domain)

	r := chi.NewRouter()

	setupMiddleware(r)

	server.setupEndpoints(r)

	return r
}
