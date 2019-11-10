package handlers

import (
	"net/http"

	"todo/domain"
)

type authResponse struct {
	User  *domain.User     `json:"user"`
	Token *domain.JWTToken `json:"token"`
}

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		user, err := s.domain.Register(payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		// generate jwt token
		token, err := user.GenToken()
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, &authResponse{
			User:  user,
			Token: token,
		}, http.StatusCreated)
	}, &payload)
}
