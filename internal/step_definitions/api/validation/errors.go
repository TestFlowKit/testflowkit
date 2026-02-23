package validation

import (
	"errors"
	"fmt"
)

var (
	ErrNoResponse      = errors.New("no response available - send a request first")
	ErrPathNotFound    = errors.New("json path not found")
	ErrValueMismatch   = errors.New("json value mismatch")
	ErrBodyNotContains = errors.New("body does not contain expected substring")
	ErrJSONNotEqual    = errors.New("response JSON validation failed")
)

// ValueMismatchError provides detailed mismatch data while supporting errors.Is with ErrValueMismatch.
type ValueMismatchError struct {
	Path     string
	Expected string
	Actual   string
}

func (e *ValueMismatchError) Error() string {
	return fmt.Sprintf("expected value '%s' at path '%s', but got '%s'", e.Expected, e.Path, e.Actual)
}

func (e *ValueMismatchError) Unwrap() error {
	return ErrValueMismatch
}
