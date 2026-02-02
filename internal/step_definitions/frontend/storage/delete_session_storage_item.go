package storage

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) deleteSessionStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I delete sessionStorage {string}`},
		func(ctx context.Context, key string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := deleteStorageItem(scenarioCtx, "sessionStorage", key)
			if err != nil {
				return ctx, fmt.Errorf("failed to delete sessionStorage item: %w", err)
			}

			logger.InfoFf("Deleted sessionStorage['%s']", key)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Removes a specific item from sessionStorage.",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The sessionStorage key to delete", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I delete sessionStorage "temp_token"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
