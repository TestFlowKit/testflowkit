package restapi

import (
	"context"
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) setRequestBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the request body to:`},
		func(ctx context.Context, body string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := scenarioCtx.SetRequestBody([]byte(body))
			if err != nil {
				return ctx, fmt.Errorf("invalid JSON in request body: %w", err)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the request body content with automatic JSON validation.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "The request body content (JSON, XML, or plain text)",
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
