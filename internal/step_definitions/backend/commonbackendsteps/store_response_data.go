package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
)

func (steps) storeResponseData() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store the value from "{string}" as "{string}"`},
		func(ctx context.Context, responsePath, variableName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			responseBody := backend.GetResponseBody()
			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			value, err := helpers.GetResponsePathValue(responseBody, responsePath)
			if err != nil {
				return ctx, fmt.Errorf("failed to extract value from path '%s': %w", responsePath, err)
			}

			backend.SetGraphQLVariable(variableName, value)
			scenarioCtx.SetVariable(variableName, value)

			logger.InfoFf("Stored value from '%s' as '%s': %v", responsePath, variableName, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Extracts a value from the response using a response path and stores it as a variable.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "Response path to extract value from (GJSON for JSON, XPath for XML)",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "variableName",
					Description: "Name of the variable to store the value in",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `And I store the value from "data.user.id" as "userId"
And I store the value from "token" as "authToken"`,
			Categories: stepbuilder.Backend,
		},
	)
}
