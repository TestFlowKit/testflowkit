package storage

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) setSessionStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I set sessionStorage {string} to {string}`},
		func(ctx context.Context, key, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := setStorageValue(scenarioCtx, "sessionStorage", key, value)
			if err != nil {
				return ctx, err
			}

			logger.InfoFf("Set sessionStorage['%s'] = '%s'", key, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a value in the browser's sessionStorage.",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The sessionStorage key", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "The value to store", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I set sessionStorage "temp_token" to "xyz123"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
