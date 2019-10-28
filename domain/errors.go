package domain

import "errors"

var (
	ErrNoResult = errors.New("no result")
	ErrUserWithEmailAlreadyExist = errors.New("user with email already exist")
	ErrUserWithUsernameAlreadyExist = errors.New("user with username already exist")
)