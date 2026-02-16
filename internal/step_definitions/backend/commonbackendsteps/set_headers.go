package commonbackendsteps

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) setHeaders() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the following headers:`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			headers, parseErr := assistdog.NewDefault().ParseMap(table)
			if parseErr != nil {
				return ctx, fmt.Errorf("failed to parse headers map: %w", parseErr)
			}

			err := setHeadersHelper(ctx, headers)
			if err == nil {
				logger.InfoFf("Headers set: %v", headers)
			}
			return ctx, err
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets HTTP headers for the request.",
			Variables: []stepbuilder.DocVariable{
				{Name: "headers", Description: "Table with header name and value pairs.", Type: stepbuilder.VarTypeTable},
			},
			Example: `And I set the following headers:
  | Authorization | Bearer {{token}}      |
  | Content-Type  | application/json      |
  | X-Request-ID  | {{requestId}}         |`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (steps) setHeader() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I set the header {string} to {string}`,
		},
		func(ctx context.Context, name, value string) (context.Context, error) {
			err := setHeadersHelper(ctx, map[string]string{name: value})
			if err == nil {
				logger.InfoFf("Header set: %s=%s", name, value)
			}
			return ctx, err
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a single HTTP header for the request.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The header name", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "The header value", Type: stepbuilder.VarTypeString},
			},
			Example:    `And I set the header "Authorization" to "Bearer {{token}}"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func setHeadersHelper(ctx context.Context, headers map[string]string) error {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	for name, value := range headers {
		backend.SetGraphQLHeader(name, value)
	}

	return nil
}
