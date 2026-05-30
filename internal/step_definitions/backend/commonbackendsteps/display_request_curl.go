package commonbackendsteps

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
)

func (steps) displayRequestCURL() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{
			`I display the request cURL`,
		},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if backend.GetProtocol() == nil {
				return ctx, apperrors.ErrNoRequestPrepared
			}

			if err := backend.SubstituteVariables(scenarioCtx); err != nil {
				return ctx, fmt.Errorf("failed to substitute variables: %w", err)
			}

			curlCommand, err := backend.GetProtocol().GetCURLCommand(ctx)
			if err != nil {
				return ctx, fmt.Errorf("failed to build request cURL: %w", err)
			}

			logger.Debug("=== REQUEST CURL ===")
			logger.Debug(curlCommand)
			logger.Debug("=== END REQUEST CURL ===")

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Displays the current prepared request as a cURL command for REST and GraphQL contexts.",
			Example:     `And I display the request cURL`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.RESTAPI, stepbuilder.GraphQL, stepbuilder.Debug},
		},
	)
}
