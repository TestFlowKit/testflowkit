package restapi

import (
	"context"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

// setJSONBody sets a JSON request body with validation.
func (steps) setJSONBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the JSON request body to:`},
		func(ctx context.Context, jsonBody string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			backend.SetRESTRequestBody([]byte(jsonBody))
			logger.InfoFf("JSON request body set and validated (%d bytes)", len(jsonBody))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a JSON request body for the REST API request with JSON validation. The body must be valid JSON.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "json",
					Description: "JSON body content",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `Given I set the JSON request body to:
"""
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 30,
  "tags": ["developer", "golang"]
}
"""`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
