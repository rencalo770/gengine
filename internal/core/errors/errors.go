package errors

import (
	"errors"
	"fmt"
)

func New(text string) error{
	return errors.New(text)
}

func Errorf(format string, i... interface{}) error {
	return  errors.New(fmt.Sprintf(format, i))
}



