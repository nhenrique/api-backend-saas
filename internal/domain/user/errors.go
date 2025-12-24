package user

import "errors"

var (
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidCompany  = errors.New("invalid company")
	ErrInvalidRole     = errors.New("invalid role")
)
