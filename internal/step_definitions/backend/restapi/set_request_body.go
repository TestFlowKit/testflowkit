package restapi

import (
	"context"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) setRequestBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{
			`I set the request body to:`,
		},
		func(ctx context.Context, body string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			// Check if it looks like a file path (second sentence pattern)
			// File paths would be passed directly, multiline strings would have newlines

			backend.SetRESTRequestBody([]byte(body))
			logger.InfoFf("Request body set (%d bytes)", len(body))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the raw request body for the REST API request.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "Raw body content or file path",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `Given I set the request body to:
"""
{"name": "John", "email": "john@example.com"}
"""
`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}
