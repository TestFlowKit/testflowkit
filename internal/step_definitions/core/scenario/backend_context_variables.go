package scenario

import (
	"maps"
)

// GetVariable gets a variable by name.
func (bc *BackendContext) GetVariable(name string) (any, bool) {
	value, exists := bc.Variables[name]
	return value, exists
}

// SetVariable sets a variable with a pre-parsed value.
func (bc *BackendContext) SetVariable(name string, value any) {
	bc.Variables[name] = value
}

// SetVariablesFromStrings sets multiple variables by parsing string values.
func (bc *BackendContext) SetVariablesFromStrings(variables map[string]string) error {
	parsedVariables, err := bc.parser.ParseVariables(variables)
	if err != nil {
		return err
	}
	maps.Copy(bc.Variables, parsedVariables)
	return nil
}

// GetVariables returns all variables.
func (bc *BackendContext) GetVariables() map[string]any {
	return bc.Variables
}

// ClearVariables clears all variables.
func (bc *BackendContext) ClearVariables() {
	bc.Variables = make(map[string]any)
}
