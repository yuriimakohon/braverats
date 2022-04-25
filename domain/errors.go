package domain

import "errors"

var (
	ErrFirstPlayerNoSuchCard  = errors.New("first player has no such card")
	ErrSecondPlayerNoSuchCard = errors.New("second player has no such card")
)
