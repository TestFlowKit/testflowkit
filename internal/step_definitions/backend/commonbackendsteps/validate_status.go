package commonbackendsteps

import (
	"context"
	"fmt"
	"strconv"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

func (s steps) validateStatusCode() stepbuilder.Step {
	return s.newStatusCodeStep(
		`the response status code should be {int}`,
		true,
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

func (s steps) validateStatusCodeNot() stepbuilder.Step {
	return s.newStatusCodeStep(
		`the response status code should not be {int}`,
		false,
		stepbuilder.DocParams{
			Description: "Validates the HTTP response status code is not a specific value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "statusCode", Description: "Status code that should not match", Type: stepbuilder.VarTypeInt},
			},
			Example:    `Then the response status code should not be 500`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) newStatusCodeStep(sentence string, shouldEqual bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{sentence},
		func(ctx context.Context, codeStr string) (context.Context, error) {
			expectedCode, err := strconv.Atoi(codeStr)
			if err != nil {
				return ctx, fmt.Errorf("invalid status code: %s", codeStr)
			}
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			actualCode := backend.GetStatusCode()
			if shouldEqual && actualCode != expectedCode {
				return ctx, fmt.Errorf("status code mismatch: expected %d, got %d", expectedCode, actualCode)
			}

			if !shouldEqual && actualCode == expectedCode {
				return ctx, fmt.Errorf("status code should not be %d", expectedCode)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
