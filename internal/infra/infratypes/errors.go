package infratypes

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
	"github.com/rmatsuoka/times_rmatsuoka/internal/types"
)

func WrapError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return &Error{
			Equal:  types.ErrNotFound,
			Origin: err,
		}
	}

	sqlite3Err, ok := err.(sqlite3.Error)
	if !ok {
		return err
	}

	var equalErr error
	switch sqlite3Err.ExtendedCode {
	case sqlite3.ErrConstraintUnique:
		equalErr = types.ErrExist
	default:
		return err
	}

	return &Error{
		Equal:  equalErr,
		Origin: sqlite3Err,
	}
}

type Error struct {
	Equal  error
	Origin error
}

func (e *Error) Error() string {
	return e.Origin.Error()
}

func (e *Error) Is(err error) bool {
	return errors.Is(err, e.Equal) || errors.Is(err, e.Origin)
}

func (e *Error) Unwrap() error {
	return e.Origin
}
