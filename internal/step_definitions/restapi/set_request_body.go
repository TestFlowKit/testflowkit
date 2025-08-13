package restapi

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) setRequestBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the request body to:`},
		func(ctx context.Context, body string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := scenarioCtx.SetRequestBody([]byte(body))
			if err != nil {
				return ctx, fmt.Errorf("failed to set request body: %w", err)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the request body content for any content type (JSON, XML, form data, etc.).",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "The request body content (JSON, XML, form data, or plain text)",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `And I set the request body to:
  """
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  """`,
			Category: stepbuilder.RESTAPI,
		},
	)
}

func (steps) setJSONRequestBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the JSON request body to:`},
		func(ctx context.Context, body string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := scenarioCtx.SetRequestBodyAsJSON([]byte(body))
			if err != nil {
				return ctx, fmt.Errorf("invalid JSON in request body: %w", err)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the request body content with JSON validation.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "The JSON request body content",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `And I set the JSON request body to:
  """
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  """`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
