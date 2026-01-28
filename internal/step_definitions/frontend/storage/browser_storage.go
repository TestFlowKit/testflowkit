package storage

import (
	"context"
	"errors"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

type storageSteps struct{}

func GetSteps() []stepbuilder.Step {
	s := storageSteps{}
	return []stepbuilder.Step{
		s.setLocalStorageItem(),
		s.setSessionStorageItem(),
		s.storeLocalStorageItemIntoVariable(),
		s.storeSessionStorageItemIntoVariable(),
		s.deleteLocalStorageItem(),
		s.deleteSessionStorageItem(),
		s.clearLocalStorage(),
		s.clearSessionStorage(),
	}
}

func (storageSteps) setLocalStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I set localStorage {string} to {string}`},
		func(ctx context.Context, key, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			key = scenario.ReplaceVariablesInString(scenarioCtx, key)
			value = scenario.ReplaceVariablesInString(scenarioCtx, value)

			script := fmt.Sprintf(`localStorage.setItem('%s', '%s')`,
				escapeJSString(key), escapeJSString(value))
			page.ExecuteJS(script)

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

func (storageSteps) setSessionStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I set sessionStorage {string} to {string}`},
		func(ctx context.Context, key, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			key = scenario.ReplaceVariablesInString(scenarioCtx, key)
			value = scenario.ReplaceVariablesInString(scenarioCtx, value)

			script := fmt.Sprintf(`sessionStorage.setItem('%s', '%s')`,
				escapeJSString(key), escapeJSString(value))
			page.ExecuteJS(script)

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

func (storageSteps) storeLocalStorageItemIntoVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store localStorage {string} into {string} variable`},
		func(ctx context.Context, key, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			key = scenario.ReplaceVariablesInString(scenarioCtx, key)

			script := fmt.Sprintf(`localStorage.getItem('%s')`, escapeJSString(key))
			result := page.ExecuteJS(script)

			var value string
			if result != "" {
				value = result
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

func (storageSteps) storeSessionStorageItemIntoVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store sessionStorage {string} into {string} variable`},
		func(ctx context.Context, key, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			key = scenario.ReplaceVariablesInString(scenarioCtx, key)

			script := fmt.Sprintf(`sessionStorage.getItem('%s')`, escapeJSString(key))
			result := page.ExecuteJS(script)

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

func (storageSteps) deleteLocalStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I delete localStorage {string}`},
		func(ctx context.Context, key string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			key = scenario.ReplaceVariablesInString(scenarioCtx, key)

			script := fmt.Sprintf(`localStorage.removeItem('%s')`, escapeJSString(key))
			page.ExecuteJS(script)

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

func (storageSteps) deleteSessionStorageItem() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I delete sessionStorage {string}`},
		func(ctx context.Context, key string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			key = scenario.ReplaceVariablesInString(scenarioCtx, key)

			script := fmt.Sprintf(`sessionStorage.removeItem('%s')`, escapeJSString(key))
			page.ExecuteJS(script)

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

func (storageSteps) clearLocalStorage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I clear localStorage`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			page.ExecuteJS(`localStorage.clear()`)

			logger.InfoFf("Cleared all localStorage items")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Removes all items from localStorage.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I clear localStorage`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}

func (storageSteps) clearSessionStorage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I clear sessionStorage`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, err := scenarioCtx.GetCurrentPageOnly()
			if err != nil {
				return ctx, errors.New("no current page available")
			}

			page.ExecuteJS(`sessionStorage.clear()`)

			logger.InfoFf("Cleared all sessionStorage items")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Removes all items from sessionStorage.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I clear sessionStorage`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}

// escapeJSString escapes single quotes in JavaScript strings
func escapeJSString(s string) string {
	result := ""
	for _, char := range s {
		switch char {
		case '\'':
			result += "\\'"
		case '\\':
			result += "\\\\"
		case '\n':
			result += "\\n"
		case '\r':
			result += "\\r"
		default:
			result += string(char)
		}
	}
	return result
}
