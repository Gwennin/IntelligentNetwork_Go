package errors

import (
	"fmt"
)

type INError struct {
	Code    uint16
	Message string
	Fatal   bool
}

func NewError(code uint16, message string) *INError {
	var err INError
	err.Code = code
	err.Message = message
	err.Fatal = false

	return &err
}

func FatalError(code uint16, message string) *INError {
	var err INError
	err.Code = code
	err.Message = message
	err.Fatal = true

	return &err
}

func (e *INError) Error() string {
	prefix := ""
	if e.Fatal {
		prefix = "FATAL"
	}
	return fmt.Sprintf("%s Intelligent Network error n°%d: %s", prefix, e.Code, e.Message)
}
