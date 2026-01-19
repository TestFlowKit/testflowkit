package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) sendRequest() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{
			`I send the request`,
			`I execute the request`,
		},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if backend.GetProtocol() == nil {
				return ctx, errors.New("no request has been prepared - use 'I prepare a request' step first")
			}

			protocol := backend.GetProtocol()
			logger.InfoFf("Sending %s request...", protocol.GetProtocolName())

			// Substitute variables in the backend context before sending
			if err := backend.SubstituteVariables(scenarioCtx); err != nil {
				return ctx, fmt.Errorf("failed to substitute variables: %w", err)
			}

			// Send the request using the protocol adapter
			ctx, err := protocol.SendRequest(ctx)
			if err != nil {
				return ctx, fmt.Errorf("failed to send request: %w", err)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sends the prepared request",
			Example:     `When I send the request`,
			Categories:  stepbuilder.Backend,
		},
	)
}
