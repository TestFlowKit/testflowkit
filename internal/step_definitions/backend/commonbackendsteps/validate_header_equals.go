package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/apperrors"
)

func (s steps) validateResponseHeaderEquals() stepbuilder.Step {
	return s.newResponseHeaderEqualsStep(
		`the response header {string} should equal {string}`,
		true,
		stepbuilder.DocParams{
			Description: "Validates that a specific response header has the expected value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "header", Description: "Name of the header to validate", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "Expected value of the header", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response header "Content-Type" should equal "application/json"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) validateResponseHeaderNotEquals() stepbuilder.Step {
	return s.newResponseHeaderEqualsStep(
		`the response header {string} should not equal {string}`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that a specific response header does not have a given value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "header", Description: "Name of the header to validate", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "Value that should not be present", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response header "Cache-Control" should not equal "no-store"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) newResponseHeaderEqualsStep(
	sentence string, shouldEqual bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, headerName, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			headerName = scenario.ReplaceVariablesInString(scenarioCtx, headerName)
			value = scenario.ReplaceVariablesInString(scenarioCtx, value)

			response := backend.GetResponse()
			if response == nil || response.Headers == nil {
				return ctx, errors.New("response headers are not available")
			}

			actualValue, found := helpers.GetHeaderCaseInsensitive(response.Headers, headerName)
			if !found {
				return ctx, fmt.Errorf("response header '%s' not found", headerName)
			}

			if shouldEqual && actualValue != value {
				return ctx, fmt.Errorf("header '%s' has value '%s', expected '%s'", headerName, actualValue, value)
			}

			if !shouldEqual && actualValue == value {
				return ctx, fmt.Errorf("header '%s' has forbidden value '%s'", headerName, actualValue)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
