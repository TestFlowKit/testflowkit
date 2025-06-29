package restapi

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) checkResponseStatusCode() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response status code should be {number}$`},
		func(ctx context.Context, expectedStatusCode int) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			response := scenarioCtx.GetResponse()
			if response == nil {
				return ctx, errors.New("no response available. Please send a request first")
			}

			actualStatusCode := response.StatusCode
			if actualStatusCode != expectedStatusCode {
				return ctx, fmt.Errorf("expected status code %d, but got %d", expectedStatusCode, actualStatusCode)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the HTTP response has the expected status code.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "statusCode",
					Description: "The expected HTTP status code (e.g., 200, 201, 404, 500)",
					Type:        stepbuilder.VarTypeInt,
				},
			},
			Example:  `Then the response status code should be 200`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
