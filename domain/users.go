package domain

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type JWTToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func (u *User) GenToken() (*JWTToken, error) {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))

	expiresAt := time.Now().Add(time.Hour * 24 * 7) // 1 week

	jwtToken.Claims = jwt.MapClaims{
		"id":  u.ID,
		"exp": expiresAt.Unix(),
	}

	accessToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}, nil
}

func (d *Domain) GetUserByID(id int64) (*User, error) {
	user, err := d.DB.UserRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
