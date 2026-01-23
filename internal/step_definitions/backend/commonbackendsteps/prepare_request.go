package commonbackendsteps

import (
	"context"
	"fmt"
	"strings"

	"testflowkit/internal/config"
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
		func(ctx context.Context, fullName string) (context.Context, error) {
			const partsNb = 2

			parts := strings.SplitN(fullName, ".", partsNb)
			if len(parts) != partsNb {
				return ctx, fmt.Errorf("invalid endpoint reference '%s', expected format 'api_name.endpoint_name'", fullName)
			}

			apiName := parts[0]
			endpointName := parts[1]

			scenarioCtx := scenario.MustFromContext(ctx)
			cfg := scenarioCtx.GetConfig()

			apiDef, err := cfg.GetAPI(apiName)
			if err != nil {
				return ctx, fmt.Errorf("failed to get API '%s': %w", apiName, err)
			}

			var adapter protocol.APIProtocol
			switch apiDef.Type {
			case config.APITypeGraphQL:
				adapter = protocol.NewGraphQLAdapter()
			case config.APITypeREST:
				adapter = protocol.NewRESTAPIAdapter()
			default:
				return ctx, fmt.Errorf("unsupported API type: %s", apiDef.Type)
			}

			// Prepare the request using the protocol adapter
			ctx, err = adapter.PrepareRequest(ctx, apiName, endpointName)
			if err != nil {
				return ctx, fmt.Errorf("failed to prepare request: %w", err)
			}

			logger.InfoFf("Request prepared: %s - %s.%s", adapter.GetProtocolName(), apiName, endpointName)
			return ctx, nil
		},
		nil,
		(func() stepbuilder.DocParams {
			desc := "The API and endpoint name in format 'api_name.endpoint_name'"
			return stepbuilder.DocParams{
				Description: "Prepares a request to a named API endpoint or operation",
				Variables: []stepbuilder.DocVariable{
					{Name: "request name", Description: desc, Type: stepbuilder.VarTypeString},
				},
				Example:    `Given I prepare a request to "users_api.getUser"`,
				Categories: stepbuilder.Backend,
			}
		})(),
	)
}
