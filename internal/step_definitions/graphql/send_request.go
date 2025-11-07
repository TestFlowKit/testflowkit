package graphql

import (
	"context"
	"fmt"
	"time"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/graphql"
	"testflowkit/pkg/logger"
)

func (steps) sendGraphQLRequest() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I send the GraphQL request`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			cfg := scenarioCtx.GetConfig()

			// Get GraphQL endpoint
			endpoint, err := cfg.GetGraphQLEndpoint()
			if err != nil {
				return ctx, fmt.Errorf("failed to get GraphQL endpoint: %w", err)
			}

			// Prepare headers by combining config headers with context headers
			headers := cfg.GetGraphQLHeaders()
			for k, v := range scenarioCtx.GetGraphQLHeaders() {
				headers[k] = v
			}

			// Create GraphQL client
			var options []graphql.ClientOption
			if len(headers) > 0 {
				options = append(options, graphql.WithHeaders(headers))
			}
			client := graphql.NewClient(endpoint, options...)
			scenarioCtx.SetGraphQLClient(client)

			// Get the prepared request
			request := scenarioCtx.GetGraphQLRequest()
			if request == nil {
				return ctx, fmt.Errorf("no GraphQL request prepared - use 'I prepare a GraphQL request' step first")
			}

			// Ensure variables are set in the request
			request.Variables = scenarioCtx.GetGraphQLVariables()

			logger.InfoFf("Sending GraphQL request to: %s", endpoint)
			logger.InfoFf("Query: %s", request.Query)
			if len(request.Variables) > 0 {
				logger.InfoFf("Variables: %v", request.Variables)
			}

			startTime := time.Now()

			// Execute the GraphQL request
			response, err := client.Execute(ctx, *request)
			duration := time.Since(startTime)

			if err != nil {
				return ctx, fmt.Errorf("failed to send GraphQL request: %w", err)
			}

			// Store the response in scenario context
			scenarioCtx.SetGraphQLResponse(response)

			logger.InfoFf("GraphQL request completed - Status: %d, Duration: %v, Errors: %d",
				response.StatusCode, duration, len(response.Errors))

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sends the prepared GraphQL request and stores the response.",
			Example:     `When I send the GraphQL request`,
			Category:    stepbuilder.GraphQL,
		},
	)
}
