package entity

import "errors"

var (
	ErrInvalidStringInput  = errors.New("invalid string input, please try again")
	ErrInvalidIntegerInput = errors.New("invalid integer input, please try again")
)
