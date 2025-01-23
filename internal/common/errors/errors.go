package errors

import (
	"fmt"
)

type Error struct {
	Code    int
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func New(message string) error {
	return &Error{
		Message: message,
	}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &Error{
		Message: message,
		Cause:   err,
	}
}
