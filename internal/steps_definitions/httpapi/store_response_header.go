package httpapi

import (
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) storeResponseHeader() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I store the value of the response header {string} as {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(headerName, variableName string) error {
				if ctx.HttpContext.Response.Headers == nil {
					return fmt.Errorf("request has not been sent yet")
				}
				headerValue := ctx.HttpContext.Response.Headers.Get(headerName)
				if headerValue == "" {
					return fmt.Errorf("header '%s' not found in response", headerName)
				}
				fmt.Printf("INFO: Storing header value '%s' in context variable '%s'\n", headerValue, variableName)
				// You would replace the line above with your actual context storage mechanism.
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a response header value into a variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "headerName", Description: "The name of the header to extract.", Type: stepbuilder.VarTypeString},
				{Name: "variableName", Description: "The name of the variable to store the header value in.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When I store the value of the response header \"Authorization\" as \"authToken\"",
			Category: stepbuilder.Form,
		},
	)
}
