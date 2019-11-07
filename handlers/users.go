package handlers

import (
	"net/http"

	"todo/domain"
)

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		user, err := s.domain.Register(payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		// generate jwt token

		jsonResponse(w, user, http.StatusCreated)
	}, &payload)
}
