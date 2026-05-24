package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/gherkinparser"
)

// validateJSONBodyEquals validates that the entire response body matches expected JSON or XML.
func (steps) validateJSONBodyEquals() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response body should be:`},
		func(ctx context.Context, expectedBody string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			expectedBody = scenario.ReplaceVariablesInString(scenarioCtx, expectedBody)

			responseBody := backend.GetResponseBody()
			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			err := validation.ValidateBodyEquals(responseBody, expectedBody)
			if err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the response body matches the expected JSON or XML exactly.",
			Variables: []stepbuilder.DocVariable{
				{Name: "body", Description: "Expected JSON or XML body", Type: stepbuilder.VarTypeString},
			},
			Example: gherkinparser.DocStringExample(
				"Then the response body should be:",
				"{\"status\":\"success\",\"data\":{\"id\":1}}",
			),
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}
