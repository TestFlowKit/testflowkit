package commonbackendsteps

import (
	"context"
	"fmt"
	"strconv"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

func (s steps) validatePathCount() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the response path "{string}" should have {int} elements`},
		func(ctx context.Context, expectedCountRaw, path string) (context.Context, error) {
			expectedCount, err := strconv.Atoi(expectedCountRaw)
			if err != nil {
				return ctx, fmt.Errorf("invalid expected count '%s': %w", expectedCountRaw, err)
			}

			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()
			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			path = scenario.ReplaceVariablesInString(scenarioCtx, path)

			engine, err := newResponseEngine(backend)
			if err != nil {
				return ctx, err
			}

			return ctx, validation.ValidatePathCount(engine, path, expectedCount)
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates the number of elements matched by a response path for both JSON and XML responses.",
			Variables: []stepbuilder.DocVariable{
				responsePathDocVariable("Response path to count (GJSON for JSON, XPath for XML)"),
				{Name: "count", Description: "Expected number of matched elements", Type: stepbuilder.VarTypeInt},
			},
			Example:    `Then the response path "//tags/tag" should have 2 elements`,
			Categories: stepbuilder.Backend,
		},
	)
}
