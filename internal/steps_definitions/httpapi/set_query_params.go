package httpapi

import (
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
)

func (st *steps) setQueryParams() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I have the following query parameters:$`},
		func(ctx *scenario.Context) func(*godog.Table) error {
			return func(table *godog.Table) error {
				if ctx.HttpContext.Method == "" {
					return fmt.Errorf("request has not been prepared. Please use 'I prepare a ... request' first")
				}
				for _, row := range table.Rows[1:] {
					ctx.HttpContext.QueryParams[row.Cells[0].Value] = row.Cells[1].Value
				}
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets query parameters for the request using a table format.",
			Variables: []stepbuilder.DocVariable{
				{Name: "queryParams", Description: "Table with parameter name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When I have the following query parameters:\n  | page | 1 |",
			Category: stepbuilder.Form,
		},
	)
}
