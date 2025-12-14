package validation

import (
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/pkg/logger"
)

// ValidateJSONPathValue validates that a JSON path contains the expected value
// Works for both GraphQL and REST API responses.
func ValidateJSONPathValue(jsonBody []byte, jsonPath, expectedValue string) error {
	if jsonBody == nil {
		return errors.New("no response available - send a request first")
	}

	actualValue, err := jsonhelpers.GetPathValueAsString(jsonBody, jsonPath)
	if err != nil {
		return fmt.Errorf("failed to get value at path '%s': %w", jsonPath, err)
	}

	if actualValue != expectedValue {
		return fmt.Errorf("expected value '%s' at path '%s', but got '%s'", expectedValue, jsonPath, actualValue)
	}

	logger.InfoFf("response validation passed: %s = %s", jsonPath, expectedValue)
	return nil
}

// ValidateJSONPathExists validates that a JSON path exists in the response.
func ValidateJSONPathExists(jsonBody []byte, jsonPath string) error {
	if jsonBody == nil {
		return errors.New("no response available - send a request first")
	}

	if !jsonhelpers.PathExists(jsonBody, jsonPath) {
		return fmt.Errorf("path '%s' does not exist in response", jsonPath)
	}

	logger.InfoFf("response validation passed: path '%s' exists", jsonPath)
	return nil
}

// ValidateBodyContains validates that the response body contains the expected substring.
func ValidateBodyContains(body []byte, expectedSubstring string) error {
	if body == nil {
		return errors.New("no response available - send a request first")
	}

	bodyStr := string(body)
	if !strings.Contains(bodyStr, expectedSubstring) {
		return fmt.Errorf("response body does not contain expected substring '%s'", expectedSubstring)
	}

	logger.InfoFf("response validation passed: body contains '%s'", expectedSubstring)
	return nil
}

// ValidateJSONBodyEquals validates that the response body matches the expected JSON.
func ValidateJSONBodyEquals(actual []byte, expected string) error {
	if actual == nil {
		return errors.New("no response available - send a request first")
	}

	if err := jsonhelpers.CompareJSON([]byte(expected), actual); err != nil {
		return fmt.Errorf("response JSON validation failed: %w", err)
	}

	logger.InfoFf("response validation passed: JSON matches expected content")
	return nil
}
