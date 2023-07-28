package model

import "errors"

var (
	// ErrBadRequest :nodoc:
	ErrBadRequest = errors.New("bad request")

	// ErrIncorrectEmailOrPassword :nodoc:
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")

	// ErrNotFound :nodoc:
	ErrNotFound = errors.New("record not found")

	// ErrUnauthorized :nodoc:
	ErrUnauthorized = errors.New("unauthorized")

	// ErrIncorrectRole :nodoc:
	ErrIncorrectRole = errors.New("incorrect role")

	// ErrInsufficientStock :nodoc:
	ErrInsufficientStock = errors.New("insufficient stock")
)
