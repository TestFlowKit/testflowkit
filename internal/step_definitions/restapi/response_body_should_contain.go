package restapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) responseBodyShouldContain() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response body should contain {string}`},
		func(ctx context.Context, expectedContent string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			response := scenarioCtx.GetResponse()
			if response == nil {
				return ctx, errors.New("no response available. Please send a request first")
			}

			responseBody := string(response.Body)
			if strings.Contains(responseBody, expectedContent) {
				return ctx, nil
			}

			intro := "response body does not contain expected content"
			msg := fmt.Sprintf("%s '%s'. Response: %s", intro, expectedContent, responseBody)
			return ctx, errors.New(msg)
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the response body contains the specified content.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "content",
					Description: "The text content that should be present in the response body",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the response body should contain "success"`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
