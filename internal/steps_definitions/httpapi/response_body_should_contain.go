package httpapi

import (
	"fmt"
	"strings"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) responseBodyShouldContain() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response body should contain {string}$`},
		func(ctx *scenario.Context) func(string) error {
			return func(text string) error {
				if ctx.HttpContext.Response.Body == nil {
					return fmt.Errorf("request has not been sent or response had no body")
				}
				if !strings.Contains(string(ctx.HttpContext.Response.Body), text) {
					return fmt.Errorf("expected response body to contain '%s', but it did not. Body: %s", text, string(ctx.HttpContext.Response.Body))
				}
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the response body contains the specified text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "The text that should be contained in the response body.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the response body should contain \"success\"",
			Category: stepbuilder.Visual,
		},
	)
}
