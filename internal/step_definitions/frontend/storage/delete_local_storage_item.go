package storage

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) deleteLocalStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I delete localStorage {string}`},
		func(ctx context.Context, key string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := deleteStorageItem(scenarioCtx, "localStorage", key)
			if err != nil {
				return ctx, err
			}

			logger.InfoFf("Deleted localStorage['%s']", key)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Removes a specific item from localStorage.",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The localStorage key to delete", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I delete localStorage "user_preference"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
