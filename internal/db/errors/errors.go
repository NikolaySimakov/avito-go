package errors

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrSegmentsNotExist = errors.New("no segments")
	ErrDatabase         = errors.New("database errors")
)
