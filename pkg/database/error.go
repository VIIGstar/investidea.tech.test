package database

import (
	"errors"
)

var (
	InvalidRequestError = errors.New("invalid request")
	NotFoundError       = errors.New("not found")
	NoChanges           = errors.New("no changes")
)

func IsNotFoundRecord(err error) bool {
	return err.Error() == NotFoundError.Error()
}
