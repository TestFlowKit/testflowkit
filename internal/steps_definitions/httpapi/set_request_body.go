package httpapi

import (
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) setRequestBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I set the request body to:$`},
		func(ctx *scenario.Context) func(string) error {
			return func(body string) error {
				if ctx.HttpContext.Method == "" {
					return fmt.Errorf("request has not been prepared. Please use 'I prepare a ... request' first")
				}
				ctx.HttpContext.RequestBody = []byte(body)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the request body content.",
			Variables: []stepbuilder.DocVariable{
				{Name: "body", Description: "The request body content.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When I set the request body to:\n  \"\"\"\n  {\"name\": \"John\"}\n  \"\"\"",
			Category: stepbuilder.Form,
		},
	)
}
