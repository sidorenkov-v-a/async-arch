package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrTaskNotFound = errors.New("task not found")
)
