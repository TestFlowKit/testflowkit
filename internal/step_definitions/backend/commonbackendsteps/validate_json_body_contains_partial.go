package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

// validateJSONBodyContainsPartial validates that the response body contains all fields
// from the expected partial JSON (deep subset check — extra fields are ignored).
func (steps) validateJSONBodyContainsPartial() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response body should contain:`},
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

			err := validation.ValidateBodyContainsPartial(responseBody, expectedBody)
			if err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the response body contains all fields from the expected partial JSON." +
				" Extra fields in the actual response are ignored.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "Partial JSON body that must be present in the response",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `Then the response body should contain:
"""
{"status":"success","data":{"id":1}}
"""`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI, stepbuilder.GraphQL},
		},
	)
}
