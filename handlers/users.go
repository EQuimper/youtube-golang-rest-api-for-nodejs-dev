package handlers

import (
	"fmt"
	"net/http"

	"todo/domain"
)

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("payload", payload)
	}, &payload)
}
