package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", s.registerUser())
		})

		r.Route("/todos", func(r chi.Router) {
			r.Use(s.withUser)
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				response := map[string]string{"hello": "world"}
				jsonResponse(w, response, 200)
			})
		})
	})
}
