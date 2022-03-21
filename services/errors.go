package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amount not be zero")
	ErrRepository = errors.New("repository error")
)
