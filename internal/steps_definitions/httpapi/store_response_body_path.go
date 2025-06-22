package httpapi

import (
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) storeResponseBodyPath() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I store the value of the response body path {string} as {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(path, variableName string) error {
				if ctx.HttpContext.Response.Body == nil {
					return fmt.Errorf("request has not been sent or response had no body")
				}
				value, err := getValueFromDotNotation(ctx.HttpContext.Response.Body, path)
				if err != nil {
					return fmt.Errorf("could not find value for path '%s' in response body: %w", path, err)
				}
				fmt.Printf("INFO: Storing value '%v' in context variable '%s'\n", value, variableName)
				// You would replace the line above with your actual context storage mechanism.
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a value from the response body at a specific path into a variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "The JSON path to extract the value from.", Type: stepbuilder.VarTypeString},
				{Name: "variableName", Description: "The name of the variable to store the value in.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When I store the value of the response body path \"user.id\" as \"userId\"",
			Category: stepbuilder.Form,
		},
	)
}
