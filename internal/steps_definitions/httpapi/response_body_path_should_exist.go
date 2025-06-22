package httpapi

import (
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) responseBodyPathShouldExist() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response body path {string} should exist$`},
		func(ctx *scenario.Context) func(string) error {
			return func(path string) error {
				if ctx.HttpContext.Response.Body == nil {
					return fmt.Errorf("request has not been sent or response had no body")
				}
				_, err := getValueFromDotNotation(ctx.HttpContext.Response.Body, path)
				if err != nil {
					return fmt.Errorf("path '%s' did not resolve to a value in the response body: %w", path, err)
				}
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that a specific path exists in the JSON response body.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "The JSON path to check (e.g., 'user.name').", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the response body path \"user.name\" should exist",
			Category: stepbuilder.Visual,
		},
	)
}
