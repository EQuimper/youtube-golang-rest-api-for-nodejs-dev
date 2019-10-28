package handlers

import "net/http"

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := s.domain.Register()
	}
}
