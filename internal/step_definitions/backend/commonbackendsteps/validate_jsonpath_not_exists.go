package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathNotExists validates that a JSON path does not exist in the response.
func (steps) validateJSONPathNotExists() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response should not have field {string}`},
		func(ctx context.Context, jsonPath string) (context.Context, error) {
			err := commonJSONPathHandler(ctx, jsonPath) // Reuse the common handler to check if the path exists
			if errors.Is(err, validation.ErrPathNotFound) {
				// Path does not exist, which is the expected outcome
				return ctx, nil
			}

			if err != nil {
				return ctx, err
			}

			return ctx, errors.New("validation failed: JSON path " + jsonPath + " exists when it should not")
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path does not exist in the response.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to check", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response should not have field "user.name"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func commonJSONPathHandler(ctx context.Context, jsonPath string) error {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	if !backend.HasResponse() {
		return errors.New("no response available to validate")
	}

	jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)

	responseBody := backend.GetResponseBody()

	if responseBody == nil {
		return errors.New("response body is empty")
	}

	return validation.ValidateJSONPathExists(responseBody, jsonPath)
}
