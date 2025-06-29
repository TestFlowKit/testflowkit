package restapi

import (
	"context"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
)

func (steps) setHeaders() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I have the following headers:$`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			for _, row := range table.Rows[1:] {
				scenarioCtx.AddHeader(row.Cells[0].Value, row.Cells[1].Value)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets HTTP headers for the request using a table format.",
			Variables: []stepbuilder.DocVariable{
				{Name: "headers", Description: "Table with header name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When I have the following headers:\n  | Content-Type | application/json |",
			Category: stepbuilder.RESTAPI,
		},
	)
}
