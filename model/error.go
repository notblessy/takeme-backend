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

	// ErrMaxTotalReached :nodoc:
	ErrMaxTotalReached = errors.New("maximum total swipe reached")

	// ErrEmailAlreadyRegistered :nodoc:
	ErrEmailAlreadyRegistered = errors.New("email already registered")
)
