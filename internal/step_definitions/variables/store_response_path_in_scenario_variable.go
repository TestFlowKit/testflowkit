package variables

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
)

func (steps) storeJSONPathIntoScenarioVariable() stepbuilder.Step {
	finalDescription := "The response path to extract the value from (GJSON for JSON, XPath for XML)"
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I store the response path {string} from the response into {string} variable`,
			`I store the JSON path {string} from the response into {string} variable`,
		},
		func(ctx context.Context, responsePath, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if backend == nil || !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			value, err := helpers.GetResponsePathValue(backend.GetResponseBody(), responsePath)
			if err != nil {
				return ctx, fmt.Errorf("failed to extract response path '%s': %w", responsePath, err)
			}

			scenarioCtx.SetVariable(varName, value)
			logger.InfoFf("Stored response path '%s' value '%v' into variable '%s'", responsePath, value, varName)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a value from a response using a response path into a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "jsonPath", Description: finalDescription, Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the variable to store the value in", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I store the response path "data.user.id" from the response into "user_id" variable`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
		},
	)
}
