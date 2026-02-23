package validation

import (
	"strings"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/pkg/logger"
)

// ValidateJSONPathValue validates that a JSON path contains the expected value
// Works for both GraphQL and REST API responses.
func ValidateJSONPathValue(jsonBody []byte, jsonPath, expectedValue string) error {
	actualValue, err := jsonhelpers.GetPathValueAsString(jsonBody, jsonPath)
	if err != nil {
		return ErrPathNotFound
	}

	if actualValue != expectedValue {
		return &ValueMismatchError{
			Path:     jsonPath,
			Expected: expectedValue,
			Actual:   actualValue,
		}
	}

	logger.InfoFf("response validation passed: %s = %s", jsonPath, expectedValue)
	return nil
}

// ValidateJSONPathExists validates that a JSON path exists in the response.
func ValidateJSONPathExists(jsonBody []byte, jsonPath string) error {
	if jsonBody == nil {
		return ErrNoResponse
	}

	if !jsonhelpers.PathExists(jsonBody, jsonPath) {
		return ErrPathNotFound
	}

	logger.InfoFf("response validation passed: path '%s' exists", jsonPath)
	return nil
}

// ValidateBodyContains validates that the response body contains the expected substring.
func ValidateBodyContains(body []byte, expectedSubstring string) error {
	if body == nil {
		return ErrNoResponse
	}

	bodyStr := string(body)
	if !strings.Contains(bodyStr, expectedSubstring) {
		return ErrBodyNotContains
	}

	logger.InfoFf("response validation passed: body contains '%s'", expectedSubstring)
	return nil
}

// ValidateJSONBodyEquals validates that the response body matches the expected JSON.
func ValidateJSONBodyEquals(actual []byte, expected string) error {
	if actual == nil {
		return ErrNoResponse
	}

	if err := jsonhelpers.CompareJSON([]byte(expected), actual); err != nil {
		return ErrJSONNotEqual
	}

	logger.InfoFf("response validation passed: JSON matches expected content")
	return nil
}
