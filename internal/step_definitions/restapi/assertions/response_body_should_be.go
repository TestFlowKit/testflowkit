package assertions

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) responseBodyShouldBe() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response body should be:`},
		func(ctx context.Context, expectedBody string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			response := scenarioCtx.GetResponse()
			if response == nil {
				return ctx, errors.New("no response available. Please send a request first")
			}

			responseBody := string(response.Body)
			trimmedExpected := strings.TrimSpace(expectedBody)
			trimmedActual := strings.TrimSpace(responseBody)

			if trimmedActual == trimmedExpected {
				return ctx, nil
			}

			intro := "response body does not match expected content exactly"
			msg := fmt.Sprintf("%s. Expected: '%s', Actual: '%s'", intro, trimmedExpected, trimmedActual)
			return ctx, errors.New(msg)
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the complete response body exactly matches the specified content.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "expectedBody",
					Description: "The exact response body content that should match",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the response body should be '{"status": "success", "message": "User created"}'`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
