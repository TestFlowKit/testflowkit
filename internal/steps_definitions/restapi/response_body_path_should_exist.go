package restapi

import (
	"context"
	"errors"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) responseBodyPathShouldExist() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response body path {string} should exist$`},
		func(ctx context.Context, jsonPath string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			response := scenarioCtx.GetResponse()
			if response == nil {
				return ctx, errors.New("no response available. Please send a request first")
			}

			_, err := getValueFromDotNotation(response.Body, jsonPath)
			if err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that a specific JSON path exists in the response body.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "The JSON path to check (e.g., 'user.id', 'data.items[0].name')",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the response body path "user.id" should exist`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
