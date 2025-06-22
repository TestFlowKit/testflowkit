package httpapi

import (
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
)

func (st steps) setPathParams() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I set the following path params:$`},
		func(ctx *scenario.Context) func(*godog.Table) error {
			return func(table *godog.Table) error {
				if ctx.HttpContext.EndpointName == "" {
					return fmt.Errorf("request has not been prepared. Please use 'I prepare a ... request' first")
				}
				endpointPath := ctx.HttpContext.EndpointConfigs[strings.ToLower(ctx.HttpContext.EndpointName)]
				for _, row := range table.Rows[1:] {
					param := row.Cells[0].Value
					value := row.Cells[1].Value
					endpointPath = strings.ReplaceAll(endpointPath, "{"+param+"}", value)
				}
				ctx.HttpContext.EndpointConfigs[strings.ToLower(ctx.HttpContext.EndpointName)] = endpointPath
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets path parameters for the endpoint by replacing placeholders in the endpoint path.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path params", Description: "Table with param name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When I set the following path params:\n  | id | 123 |",
			Category: stepbuilder.Form,
		},
	)
}
