package storage

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) setLocalStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I set localStorage {string} to {string}`},
		func(ctx context.Context, key, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := setStorageValue(scenarioCtx, "localStorage", key, value)
			if err != nil {
				return ctx, fmt.Errorf("failed to set localStorage item: %w", err)
			}
			logger.InfoFf("Set localStorage['%s'] = '%s'", key, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a value in the browser's localStorage.",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The localStorage key", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "The value to store", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I set localStorage "user_preference" to "dark_mode"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
