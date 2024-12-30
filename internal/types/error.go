package types

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("already exists")
)
