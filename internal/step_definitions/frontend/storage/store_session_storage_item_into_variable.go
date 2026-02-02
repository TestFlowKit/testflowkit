package storage

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeSessionStorageItemIntoVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store sessionStorage {string} into {string} variable`},
		func(ctx context.Context, key, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			result, err := getStorageValue(scenarioCtx, "sessionStorage", key)
			if err != nil {
				return ctx, err
			}

			var value string
			if result != "" {
				value = result
			}

			scenarioCtx.SetVariable(varName, value)
			logger.InfoFf("Stored sessionStorage['%s'] = '%s' into variable '%s'", key, value, varName)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Retrieves a value from sessionStorage and stores it in a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The sessionStorage key to retrieve", Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the variable to store the value in", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I store sessionStorage "temp_token" into "token" variable`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
