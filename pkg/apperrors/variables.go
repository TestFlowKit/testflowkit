package apperrors

import "fmt"

// VariableNotFoundError is returned when a scenario variable does not exist.
type VariableNotFoundError struct {
	Name string
}

func (e *VariableNotFoundError) Error() string {
	return fmt.Sprintf("variable '%s' not found", e.Name)
}

// VariableNotStringError is returned when a variable exists but its value is not a string.
type VariableNotStringError struct {
	Name string
}

func (e *VariableNotStringError) Error() string {
	return fmt.Sprintf("variable '%s' is not a string", e.Name)
}
