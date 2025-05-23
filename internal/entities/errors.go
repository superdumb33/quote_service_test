package entities

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal server error")
	ErrDuplicate = errors.New("duplicate")
	
)