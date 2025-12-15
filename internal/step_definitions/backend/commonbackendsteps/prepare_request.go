package commonbackendsteps

import (
	"context"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/api/protocol"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) prepareRequest() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I prepare a (?i)(graphql|rest) request to {string}`,
		},
		func(ctx context.Context, protocolType string, name string) (context.Context, error) {
			var adapter protocol.APIProtocol
			switch strings.ToLower(protocolType) {
			case "graphql":
				adapter = protocol.NewGraphQLAdapter()
			case "rest":
				adapter = protocol.NewRESTAPIAdapter()
			default:
				return ctx, fmt.Errorf("unsupported protocol type: %s", protocolType)
			}

			// Prepare the request using the protocol adapter
			ctx, err := adapter.PrepareRequest(ctx, name)
			if err != nil {
				return ctx, fmt.Errorf("failed to prepare request: %w", err)
			}

			logger.InfoFf("Request prepared: %s - %s", adapter.GetProtocolName(), name)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Prepares a request",
			Variables: []stepbuilder.DocVariable{
				{Name: "protocolType", Description: "The API protocol type (graphql or REST)", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The name of the operation or endpoint", Type: stepbuilder.VarTypeString},
			},
			Example:  `Given I prepare a REST request to "getUser"`,
			Category: stepbuilder.Backend,
		},
	)
}
