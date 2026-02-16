package restapi

import (
	"context"
	"errors"
	"maps"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) setQueryParams() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the following query parameters:`},
		func(ctx context.Context, paramsTable *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			params, err := assistdog.NewDefault().ParseMap(paramsTable)
			if err != nil {
				return ctx, errors.New("failed to parse query parameters map: " + err.Error())
			}

			// Store in endpoint enricher
			endpoint := backend.GetEndpoint()
			if endpoint == nil {
				return ctx, errors.New("no endpoint configured in backend context")
			}

			if endpoint.QueryParams == nil {
				endpoint.QueryParams = make(map[string]string)
			}
			maps.Copy(endpoint.QueryParams, params)

			logger.InfoFf("Query parameters set: %v", params)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets URL query parameters for the REST API request.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "parameters",
					Description: "Table with parameter names and values",
					Type:        stepbuilder.VarTypeTable,
				},
			},
			Example: `Given I set the following query parameters:
  | name   | value        |
  | page   | 1            |
  | limit  | 10           |
  | filter | {{category}} |`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}

func (steps) setQueryParam() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I set query parameter {string} to {string}`,
		},
		func(ctx context.Context, name, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			endpoint := backend.GetEndpoint()
			if endpoint == nil {
				return ctx, errors.New("no endpoint configured in backend context")
			}
			if endpoint.QueryParams == nil {
				endpoint.QueryParams = make(map[string]string)
			}
			endpoint.QueryParams[name] = value

			logger.InfoFf("Query parameter set: %s=%s", name, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a single URL query parameter for the REST API request.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The query parameter name", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "The query parameter value", Type: stepbuilder.VarTypeString},
			},
			Example:    `Given I set query parameter "page" to "1"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}
