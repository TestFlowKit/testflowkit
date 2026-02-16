package restapi

import (
	"context"
	"errors"
	"maps"
	"strings"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) setPathParams() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the following path parameters:`},
		func(ctx context.Context, paramsTable *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			// Parse table to map
			params, err := assistdog.NewDefault().ParseMap(paramsTable)
			if err != nil {
				return ctx, errors.New("failed to parse path parameters map: " + err.Error())
			}

			// Store in endpoint enricher
			endpoint := backend.GetEndpoint()
			if endpoint == nil {
				return ctx, errors.New("no endpoint configured in backend context")
			}
			if endpoint.PathParams == nil {
				endpoint.PathParams = make(map[string]string)
			}
			maps.Copy(endpoint.PathParams, params)

			logger.InfoFf("Path parameters set: %v", params)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: strings.Join(setPathParamsDoc, " "),
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "parameters",
					Description: "Table with parameter names and values",
					Type:        stepbuilder.VarTypeTable,
				},
			},
			Example: `Given I set the following path parameters:
  | userId | {{userId}} |
  | postId | 123        |`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}

func (steps) setPathParam() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I set the path parameter {string} to {string}`,
		},
		func(ctx context.Context, name, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			endpoint := backend.GetEndpoint()
			if endpoint == nil {
				return ctx, errors.New("no endpoint configured in backend context")
			}
			if endpoint.PathParams == nil {
				endpoint.PathParams = make(map[string]string)
			}
			endpoint.PathParams[name] = value

			logger.InfoFf("Path parameter set: %s=%s", name, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a single URL path parameter for the REST API request.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The path parameter name", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "The path parameter value", Type: stepbuilder.VarTypeString},
			},
			Example:    `Given I set the path parameter "userId" to "123"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}

var setPathParamsDoc = []string{
	"Sets URL path parameters for the REST API request.",
	" Path parameters replace placeholders in the URL template.",
}
