package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/graphql"
	"testflowkit/pkg/logger"

	"github.com/cucumber/godog"
)

func (steps) prepareGraphQLRequest() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I prepare a GraphQL request for the {string} operation`},
		func(ctx context.Context, operationName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			cfg := scenarioCtx.GetConfig()

			operation, err := cfg.GetGraphQLOperation(operationName)
			if err != nil {
				return ctx, fmt.Errorf("failed to resolve GraphQL operation '%s': %w", operationName, err)
			}

			// Create GraphQL request with the operation
			request := &graphql.Request{
				Query:     operation.Operation,
				Variables: make(map[string]interface{}),
			}

			scenarioCtx.SetGraphQLRequest(request)

			logger.InfoFf("GraphQL request prepared for operation '%s' (%s)", operationName, operation.Description)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Prepares a GraphQL request for a configured operation.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "operationName",
					Description: "The logical operation name as defined in configuration",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Given I prepare a GraphQL request for the "get_user_profile" operation`,
			Category: stepbuilder.GraphQL,
		},
	)
}

func (steps) setGraphQLVariables() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the following GraphQL variables:`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			variables := make(map[string]string)

			for _, row := range table.Rows[1:] { // Skip header row
				key := row.Cells[0].Value
				value := row.Cells[1].Value
				variables[key] = value
			}

			err := scenarioCtx.SetGraphQLVariablesFromStrings(variables)
			if err != nil {
				return ctx, fmt.Errorf("failed to set GraphQL variables: %w", err)
			}

			// Update the current request with variables
			request := scenarioCtx.GetGraphQLRequest()
			if request != nil {
				request.Variables = scenarioCtx.GetGraphQLVariables()
			}

			logger.InfoFf("GraphQL variables set: %v", variables)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets GraphQL variables for the prepared request. Supports strings, numbers, booleans, JSON objects, and arrays.",
			Variables: []stepbuilder.DocVariable{
				{Name: "variables", Description: "Table with variable name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example: `And I set the following GraphQL variables:
  | userId    | 123                           |
  | tags      | ["frontend", "testing"]       |
  | filters   | {"status": "active"}          |
  | isActive  | true                          |`,
			Category: stepbuilder.GraphQL,
		},
	)
}

func (steps) setGraphQLArrayVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I set the GraphQL variable {string} to array {string}`},
		func(ctx context.Context, variableName, arrayValue string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := scenarioCtx.SetGraphQLArrayVariable(variableName, arrayValue)
			if err != nil {
				return ctx, fmt.Errorf("failed to set GraphQL array variable '%s': %w", variableName, err)
			}

			// Update the current request with variables
			request := scenarioCtx.GetGraphQLRequest()
			if request != nil {
				request.Variables = scenarioCtx.GetGraphQLVariables()
			}

			logger.InfoFf("GraphQL array variable '%s' set to: %s", variableName, arrayValue)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a GraphQL variable to an array value.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "variableName",
					Description: "Name of the GraphQL variable",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "arrayValue",
					Description: "JSON array value (e.g., [\"item1\", \"item2\"] or [1, 2, 3])",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I set the GraphQL variable "tags" to array ["frontend", "testing", "automation"]`,
			Category: stepbuilder.GraphQL,
		},
	)
}

func (steps) setGraphQLHeaders() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the following GraphQL headers:`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			headers := make(map[string]string)

			for _, row := range table.Rows[1:] { // Skip header row
				key := row.Cells[0].Value
				value := row.Cells[1].Value
				headers[key] = value
			}

			scenarioCtx.SetGraphQLHeaders(headers)

			logger.InfoFf("GraphQL headers set: %v", headers)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets custom headers for the GraphQL request using a table format.",
			Variables: []stepbuilder.DocVariable{
				{Name: "headers", Description: "Table with header name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example: `And I set the following GraphQL headers:
  | Authorization | Bearer token123 |
  | X-Client-ID   | test-client     |`,
			Category: stepbuilder.GraphQL,
		},
	)
}
