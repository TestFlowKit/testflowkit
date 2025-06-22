package httpapi

import (
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) checkResponseStatusCode() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response status code should be {number}$`},
		func(ctx *scenario.Context) func(int) error {
			return func(expectedCode int) error {
				if ctx.HttpContext.Response.Body == nil {
					return fmt.Errorf("request has not been sent yet")
				}
				if ctx.HttpContext.Response.StatusCode != expectedCode {
					return fmt.Errorf("expected status code %d, but got %d. Body: %s", expectedCode, ctx.HttpContext.Response.StatusCode, string(ctx.HttpContext.Response.Body))
				}
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the response has the expected status code.",
			Variables: []stepbuilder.DocVariable{
				{Name: "statusCode", Description: "The expected HTTP status code.", Type: stepbuilder.VarTypeInt},
			},
			Example:  "Then the response status code should be 200",
			Category: stepbuilder.Visual,
		},
	)
}
