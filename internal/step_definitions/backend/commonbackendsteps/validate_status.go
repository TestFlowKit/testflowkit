package commonbackendsteps

import (
	"context"
	"fmt"
	"strconv"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

func (steps) validateStatusCode() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response status code should be {number}`},
		func(ctx context.Context, expectedCodeStr string) (context.Context, error) {
			expectedCode, err := strconv.Atoi(expectedCodeStr)
			if err != nil {
				return ctx, fmt.Errorf("invalid status code: %s", expectedCodeStr)
			}
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			actualCode := backend.GetStatusCode()
			if actualCode != expectedCode {
				return ctx, fmt.Errorf("status code mismatch: expected %d, got %d", expectedCode, actualCode)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates the HTTP response status code.",
			Variables: []stepbuilder.DocVariable{
				{Name: "statusCode", Description: "Expected HTTP status code", Type: stepbuilder.VarTypeInt},
			},
			Example:    `Then the response status code should be 200`,
			Categories: stepbuilder.Backend,
		},
	)
}
