package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

// validateJSONPathExists validates that a JSON path exists in the response.
func (s steps) validateJSONPathExists() stepbuilder.Step {
	return s.newJSONPathExistsStep(
		`the response should have field {string}`,
		true,
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

// validateJSONPathNotExists validates that a JSON path does not exist in the response.
func (s steps) validateJSONPathNotExists() stepbuilder.Step {
	return s.newJSONPathExistsStep(
		`the response should not have field {string}`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path does not exist in the response.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path that should not exist", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response should not have field "user.password"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) newJSONPathExistsStep(sentence string, shouldExist bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{sentence},
		func(ctx context.Context, jsonPath string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			err := validation.ValidateJSONPathExists(responseBody, jsonPath)
			if shouldExist && err != nil {
				return ctx, err
			}

			if !shouldExist && err == nil {
				return ctx, fmt.Errorf("path '%s' exists in response but should not", jsonPath)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
