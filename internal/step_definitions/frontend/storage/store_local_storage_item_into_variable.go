package storage

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeLocalStorageItemIntoVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store localStorage {string} into {string} variable`},
		func(ctx context.Context, key, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			value, err := getStorageValue(scenarioCtx, "localStorage", key)
			if err != nil {
				return ctx, err
			}

			scenarioCtx.SetVariable(varName, value)
			logger.InfoFf("Stored localStorage['%s'] = '%s' into variable '%s'", key, value, varName)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Retrieves a value from localStorage and stores it in a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The localStorage key to retrieve", Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the variable to store the value in", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I store localStorage "user_preference" into "preference" variable`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
