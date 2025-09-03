package assertions

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) checkResponseStatusCode() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response status code should be {number}`},
		func(ctx context.Context, expectedStatusCode string) (context.Context, error) {
			expectedStatusCodeInt, err := strconv.Atoi(expectedStatusCode)
			if err != nil {
				return ctx, fmt.Errorf("invalid status code: %s", expectedStatusCode)
			}

			scenarioCtx := scenario.MustFromContext(ctx)

			response := scenarioCtx.GetResponse()
			if response == nil {
				return ctx, errors.New("no response available. Please send a request first")
			}

			actualStatusCode := response.StatusCode
			if actualStatusCode != expectedStatusCodeInt {
				const errFormat = "expected status code %d, but got %d with response\n%s"
				return ctx, fmt.Errorf(errFormat, expectedStatusCodeInt, actualStatusCode, string(response.Body))
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
