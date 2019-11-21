package domain

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
}

func (r *RegisterPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	v.MustBeNotEmpty("email", r.Email)
	v.MustBeValidEmail("email", r.Email)

	v.MustBeNotEmpty("password", r.Password)
	v.MustBeLongerThan("password", r.Password, 6)

	v.MustBeNotEmpty("confirmPassword", r.ConfirmPassword)
	v.MustMatch(
		ElementMatcher{
			field: "confirmPassword",
			value: r.ConfirmPassword,
		},
		ElementMatcher{
			field: "password",
			value: r.Password,
		},
	)

	v.MustBeNotEmpty("username", r.Username)
	v.MustBeLongerThan("username", r.Username, 3)

	return v.IsValid(), v.errors
}

func (d *Domain) Register(payload RegisterPayload) (*User, error) {
	userExist, _ := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExist != nil {
		return nil, ErrUserWithEmailAlreadyExist
	}

	userExist, _ = d.DB.UserRepo.GetByUsername(payload.Username)
	if userExist != nil {
		return nil, ErrUserWithUsernameAlreadyExist
	}

	password, err := d.setPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	data := &User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: *password,
	}

	user, err := d.DB.UserRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) setPassword(password string) (*string, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	password = string(passwordHash)

	return &password, nil
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
}

func ParseToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, authExtractor, func(t *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("JWT_SECRET"))
		return b, nil
	})

	return token, err
}
