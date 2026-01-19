package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeResponseData() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store the value from "{string}" as "{string}"`},
		func(ctx context.Context, jsonPath, variableName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to extract data from")
			}

			// Get the appropriate response body based on protocol
			responseBody := backend.GetResponseBody()
			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			value, err := jsonhelpers.GetPathValue(responseBody, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to extract value from path '%s': %w", jsonPath, err)
			}

			// Store in both backend context and global context
			backend.SetGraphQLVariable(variableName, value)
			scenarioCtx.SetVariable(variableName, value)

			logger.InfoFf("Stored value from '%s' as '%s': %v", jsonPath, variableName, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Extracts a value from the response using a JSON path and stores it as a variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "jsonPath", Description: "JSON path to extract value from", Type: stepbuilder.VarTypeString},
				{Name: "variableName", Description: "Name of the variable to store the value in", Type: stepbuilder.VarTypeString},
			},
			Example: `And I store the value from "data.user.id" as "userId"
And I store the value from "token" as "authToken"`,
			Categories: stepbuilder.Backend,
		},
	)
}
