package storage

import (
	"errors"
	"fmt"
	"slices"
	"testflowkit/internal/step_definitions/core/scenario"
)

// escapeJSString escapes single quotes and other special characters in JavaScript strings.
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

func getStorageValue(scenarioCtx *scenario.Context, storageType, key string) (string, error) {
	validateStorageType(storageType)

	page, err := scenarioCtx.GetCurrentPageOnly()
	if err != nil {
		return "", errors.New("no current page available")
	}

	key = scenario.ReplaceVariablesInString(scenarioCtx, key)

	script := fmt.Sprintf(`%s.getItem('%s')`, storageType, escapeJSString(key))
	value, errExJS := page.ExecuteJS(script)
	if errExJS != nil {
		return "", errExJS
	}

	return value, nil
}

func setStorageValue(scenarioCtx *scenario.Context, storageType, key, value string) error {
	validateStorageType(storageType)

	page, err := scenarioCtx.GetCurrentPageOnly()
	if err != nil {
		return errors.New("no current page available")
	}

	key = scenario.ReplaceVariablesInString(scenarioCtx, key)
	value = scenario.ReplaceVariablesInString(scenarioCtx, value)

	script := fmt.Sprintf(`%s.setItem('%s', '%s')`,
		storageType, escapeJSString(key), escapeJSString(value))
	_, err = page.ExecuteJS(script)

	return err
}

func deleteStorageItem(scenarioCtx *scenario.Context, storageType, key string) error {
	validateStorageType(storageType)
	page, err := scenarioCtx.GetCurrentPageOnly()
	if err != nil {
		return errors.New("no current page available")
	}

	key = scenario.ReplaceVariablesInString(scenarioCtx, key)

	script := fmt.Sprintf(`%s.removeItem('%s')`, storageType, escapeJSString(key))
	_, err = page.ExecuteJS(script)

	return err
}

func clearStorage(scenarioCtx *scenario.Context, storageType string) error {
	validateStorageType(storageType)

	page, err := scenarioCtx.GetCurrentPageOnly()
	if err != nil {
		return errors.New("no current page available")
	}

	script := fmt.Sprintf(`%s.clear()`, storageType)
	_, err = page.ExecuteJS(script)

	return err
}

func validateStorageType(storageType string) {
	if !slices.Contains([]string{
		"localStorage",
		"sessionStorage",
	}, storageType) {
		panic(fmt.Errorf("invalid storage type: %s", storageType))
	}
}
