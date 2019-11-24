package handlers

import (
	"net/http"

	"todo/domain"
)

func (s *Server) createTodo() http.HandlerFunc {
	var payload domain.CreateTodoPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		currentUser := s.currentUserFromCTX(r)

		todo, err := s.domain.CreateTodo(payload, currentUser)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, todo, http.StatusCreated)
	}, &payload)
}
