package commonbackendsteps

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/api/protocol"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) prepareRequest() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{
			`I prepare a request to {string}`,
		},
		func(ctx context.Context, name string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			// Auto-detect protocol based on step text or use existing protocol
			var adapter scenario.APIProtocol
			if backend.GetProtocol() != nil {
				adapter = backend.GetProtocol()
			} else {
				adapter = protocol.NewRESTAPIAdapter()
				logger.InfoFf("Using REST API protocol for endpoint: %s", name)
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
				{Name: "name", Description: "The name of the endpoint (REST)", Type: stepbuilder.VarTypeString},
			},
			Example:  `Given I prepare a request to "getUser"`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
