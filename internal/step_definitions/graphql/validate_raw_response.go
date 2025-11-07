package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) validateGraphQLResponseDataShouldBe() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the GraphQL response data should be`},
		func(ctx context.Context, expectedResponse string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			response := scenarioCtx.GetGraphQLResponse()

			if response == nil {
				return ctx, fmt.Errorf("no GraphQL response available - send a GraphQL request first")
			}

			if response.Data == nil {
				return ctx, fmt.Errorf("GraphQL response contains no data")
			}

			// Use only the data part of the response (like in Bruno/GraphQL clients)
			actualDataStr := string(response.Data)

			// Normalize JSON strings for comparison (remove extra whitespace)
			var expectedJSON, actualJSON interface{}

			if err := json.Unmarshal([]byte(expectedResponse), &expectedJSON); err != nil {
				return ctx, fmt.Errorf("expected response is not valid JSON: %w", err)
			}

			if err := json.Unmarshal(response.Data, &actualJSON); err != nil {
				return ctx, fmt.Errorf("actual response data is not valid JSON: %w", err)
			}

			// Marshal both back to normalized JSON for comparison
			expectedNormalized, _ := json.Marshal(expectedJSON)
			actualNormalized, _ := json.Marshal(actualJSON)

			if string(expectedNormalized) != string(actualNormalized) {
				return ctx, fmt.Errorf("GraphQL response data does not match expected response.\nExpected: %s\nActual: %s", string(expectedNormalized), actualDataStr)
			}

			logger.InfoFf("GraphQL response data validation passed - matches expected JSON")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response data exactly matches the expected JSON using multiline string format. This compares only the data portion of the response.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "expectedResponse",
					Description: "The expected JSON data that the GraphQL response should match (supports multiline strings)",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `Then the GraphQL response data should be
"""
{
  "user": {
    "id": "123",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
"""`,
			Category: stepbuilder.GraphQL,
		},
	)
}
