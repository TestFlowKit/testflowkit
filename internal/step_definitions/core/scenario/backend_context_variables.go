package scenario

import (
	"maps"
)

func (bc *BackendContext) GetGraphQLVariable(name string) (any, bool) {
	value, exists := bc.GraphQL.Variables[name]
	return value, exists
}

func (bc *BackendContext) SetGraphQLVariable(name string, value any) {
	bc.GraphQL.Variables[name] = value
}

// SetVariablesFromStrings sets multiple variables by parsing string values.
func (bc *BackendContext) SetVariablesFromStrings(variables map[string]string) error {
	parsedVariables, err := bc.parser.ParseVariables(variables)
	if err != nil {
		return err
	}
	maps.Copy(bc.GraphQL.Variables, parsedVariables)
	return nil
}

func (bc *BackendContext) GetGraphQLVariables() map[string]any {
	return bc.GraphQL.Variables
}
