package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNoResult                     = errors.New("no result")
	ErrUserWithEmailAlreadyExist    = errors.New("user with email already exist")
	ErrUserWithUsernameAlreadyExist = errors.New("user with username already exist")
	ErrEmailBadFormat               = errors.New("email is not valid")
)

type ErrNotLongEnough struct {
	field  string
	amount int
}

func (e ErrNotLongEnough) Error() string {
	return fmt.Sprintf("%v not long enough, %d characters is required", e.field, e.amount)
}

type ErrIsRequired struct {
	field string
}

func (e ErrIsRequired) Error() string {
	return fmt.Sprintf("%v is required", e.field)
}
