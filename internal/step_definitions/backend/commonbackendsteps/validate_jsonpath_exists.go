package commonbackendsteps

import (
	"context"

	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathExists validates that a JSON path exists in the response.
func (steps) validateJSONPathExists() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response should have field {string}`},
		func(ctx context.Context, jsonPath string) (context.Context, error) {
			err := commonJSONPathHandler(ctx, jsonPath) // Reuse the common handler to check if the path exists
			return ctx, err
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path exists in the response.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to check", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response should have field "user.name"`,
			Categories: stepbuilder.Backend,
		},
	)
}
