package handlers

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"

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

func (s *Server) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := domain.ParseToken(r)

		if err != nil {
			unauthorizedResponse(w)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int64(claims["id"].(float64))

			user, err := s.domain.GetUserByID(userID)
			if err != nil {
				unauthorizedResponse(w)
				return
			}

			ctx := context.WithValue(r.Context(), "currentUser", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			unauthorizedResponse(w)
			return
		}
	})
}
