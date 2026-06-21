package variables

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/variables"
)

func (steps) storeJSONPathIntoGlobalVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I save the response path {string} as global variable {string}`,
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

			variables.SetGlobalVariable(varName, value)
			logger.Infof("Stored response path '%s' value '%v' into global variable '%s'", responsePath, value, varName)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a value from a response using a response path into a global variable.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "responsePath",
					Description: "The response path to extract (GJSON for JSON, XPath for XML)",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        stepbuilder.DocVarVarName,
					Description: "The name of the global variable",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
			Example:    `Then I save the response path "token" as global variable "AUTH_TOKEN"`,
		},
	)
}
