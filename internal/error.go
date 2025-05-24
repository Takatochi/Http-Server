package internal

import "fmt"

type Error struct {
	Code    int
	Message string
	Err     error
}

func NewError(_code int, _message string, _err error) *Error {
	return &Error{
		Code:    _code,
		Message: _message,
		Err:     _err,
	}
}
func (e Error) Error() string {
	return fmt.Sprintf("Error: %d - %s - %s", e.Code, e.Message, e.Err)
}
