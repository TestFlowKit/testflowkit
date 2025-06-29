package restapi

import (
	"context"
	"errors"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
)

func (steps) setQueryParams() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I have the following query parameters:$`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			endpoint := scenarioCtx.GetEndpoint()
			if endpoint.Path == "" {
				return ctx, errors.New("request has not been prepared. Please use 'I prepare a ... request' first")
			}
			for _, row := range table.Rows[1:] {
				scenarioCtx.AddQueryParam(row.Cells[0].Value, row.Cells[1].Value)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets query parameters for the request using a table format.",
			Variables: []stepbuilder.DocVariable{
				{Name: "queryParams", Description: "Table with parameter name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When I have the following query parameters:\n  | page | 1 |",
			Category: stepbuilder.RESTAPI,
		},
	)
}
