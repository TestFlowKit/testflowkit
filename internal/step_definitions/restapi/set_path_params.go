package restapi

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) setPathParams() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I set the following path params:$`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			endpoint := scenarioCtx.GetEndpoint()
			if endpoint.Path == "" {
				return ctx, errors.New("request has not been prepared. Please use 'I prepare a ... request' first")
			}
			params, err := assistdog.NewDefault().ParseMap(table)
			if err != nil {
				return ctx, err
			}

			parsedParams, err := scenario.ReplaceVariablesInMap(scenarioCtx, params)
			if err != nil {
				return ctx, err
			}

			var unknownParams []string

			for param := range parsedParams {
				isKnowParam := strings.Contains(endpoint.Path, "{"+param+"}")

				if isKnowParam {
					paramValue := parsedParams[param]
					scenarioCtx.AddPathParam(param, paramValue)
					continue
				}

				unknownParams = append(unknownParams, param)
			}

			if len(unknownParams) > 0 {
				return ctx, fmt.Errorf("unknown path parameters: %v", unknownParams)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets path parameters for the endpoint by replacing placeholders in the endpoint path.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path params", Description: "Table with param name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When I set the following path params:\n  | id | 123 |",
			Category: stepbuilder.RESTAPI,
		},
	)
}
