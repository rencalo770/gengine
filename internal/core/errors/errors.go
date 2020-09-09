package errors

import (
	"errors"
)

func New(text string) error {
	return errors.New(text)
}
