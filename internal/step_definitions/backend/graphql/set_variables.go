package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

// setGraphQLVariables sets GraphQL variables from a table using the unified parser.
func (steps) setGraphQLVariables() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the following GraphQL variables:`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			// Convert table to map[string]string
			varsMap, err := assistdog.NewDefault().ParseMap(table)
			if err != nil {
				return ctx, fmt.Errorf("failed to parse variables table: %w", err)
			}

			if errSetVars := backend.SetVariablesFromStrings(varsMap); errSetVars != nil {
				return ctx, fmt.Errorf("failed to set GraphQL variables: %w", errSetVars)
			}

			logger.InfoFf("GraphQL variables set: %v", backend.GetVariables())
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets GraphQL variables using the unified type-aware parser. Supports primitives, arrays, and objects.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "variables",
					Description: "Table with variable names and values (supports JSON arrays/objects)",
					Type:        stepbuilder.VarTypeTable,
				},
			},
			Example: `Given I set the following GraphQL variables:
  | name    | value                           |
  | userId  | 123                             |
  | tags    | ["go", "testing"]               |
  | filter  | {"status": "active", "limit": 5}|`,
			Category: stepbuilder.GraphQL,
		},
	)
}
