package validationerror

import "errors"

var (
	ErrInvalidAccount          = errors.New("invalid account")
	ErrAccountNotFound         = errors.New("account not found")
	ErrReceiverAccountNotFound = errors.New("receiver account not found")
	ErrInsufficientBalance     = errors.New("insufficient balanced")
)
