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
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			headers, parseErr := assistdog.NewDefault().ParseMap(table)
			if parseErr != nil {
				return ctx, fmt.Errorf("failed to parse headers map: %w", parseErr)
			}

			for name, value := range headers {
				backend.SetHeader(name, value)
			}

			logger.InfoFf("Headers set: %v", headers)
			return ctx, nil
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
			Category: stepbuilder.Backend,
		},
	)
}
