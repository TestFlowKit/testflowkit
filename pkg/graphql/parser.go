package graphql

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

// ResponseParser provides utilities for parsing GraphQL responses.
type ResponseParser struct {
	response *Response
}

// NewResponseParser creates a new response parser.
func NewResponseParser(response *Response) *ResponseParser {
	return &ResponseParser{response: response}
}

// GetDataAtPath extracts data at the specified JSON path.
func (rp *ResponseParser) GetDataAtPath(path string) (gjson.Result, error) {
	if rp.response.Data == nil {
		// Provide detailed error context including any GraphQL errors
		errorContext := map[string]interface{}{
			"path": path,
		}
		if rp.response.HasErrors() {
			errorContext["graphql_errors"] = rp.response.GetErrorMessages()
		}
		return gjson.Result{}, errors.New("response contains no data")
	}

	result := gjson.GetBytes(rp.response.Data, path)
	if !result.Exists() {
		// Provide context about available paths for debugging
		availablePaths := rp.GetAllPaths()
		return gjson.Result{}, fmt.Errorf("path '%s' not found in response data (available paths: %v)", path, availablePaths)
	}

	return result, nil
}

// GetDataAsString extracts data at path as string.
func (rp *ResponseParser) GetDataAsString(path string) (string, error) {
	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

// GetDataAsInt extracts data at path as int64.
func (rp *ResponseParser) GetDataAsInt(path string) (int64, error) {
	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return 0, err
	}
	return result.Int(), nil
}

// GetDataAsFloat extracts data at path as float64.
func (rp *ResponseParser) GetDataAsFloat(path string) (float64, error) {
	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return 0, err
	}
	return result.Float(), nil
}

// GetDataAsBool extracts data at path as bool.
func (rp *ResponseParser) GetDataAsBool(path string) (bool, error) {
	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return false, err
	}
	return result.Bool(), nil
}

// GetDataAsArray extracts data at path as array.
func (rp *ResponseParser) GetDataAsArray(path string) ([]gjson.Result, error) {
	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return nil, err
	}

	if !result.IsArray() {
		return nil, fmt.Errorf("path '%s' does not contain an array", path)
	}

	return result.Array(), nil
}

// GetDataAsInterface extracts data at path as interface{}.
func (rp *ResponseParser) GetDataAsInterface(path string) (interface{}, error) {
	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return nil, err
	}

	var value interface{}
	if errDecode := json.Unmarshal([]byte(result.Raw), &value); errDecode != nil {
		return nil, fmt.Errorf("failed to unmarshal data at path '%s': %w", path, errDecode)
	}

	return value, nil
}

// PathExists checks if a path exists in the response data.
func (rp *ResponseParser) PathExists(path string) bool {
	if rp.response.Data == nil {
		return false
	}

	result := gjson.GetBytes(rp.response.Data, path)
	return result.Exists()
}

// GetAllPaths returns all available paths in the response data (for debugging).
func (rp *ResponseParser) GetAllPaths() []string {
	if rp.response.Data == nil {
		return []string{}
	}

	var paths []string
	gjson.GetBytes(rp.response.Data, "@this").ForEach(func(key, _ gjson.Result) bool {
		paths = append(paths, key.String())
		return true
	})

	return paths
}

// GetDataWithErrorContext extracts data at path and includes error context if extraction fails.
func (rp *ResponseParser) GetDataWithErrorContext(path string) (gjson.Result, map[string]interface{}, error) {
	errorContext := map[string]interface{}{
		"path":        path,
		"has_data":    rp.response.HasData(),
		"has_errors":  rp.response.HasErrors(),
		"status_code": rp.response.StatusCode,
	}

	if rp.response.HasErrors() {
		errorContext["error_summary"] = rp.response.GetErrorSummary()
		errorContext["error_messages"] = rp.response.GetErrorMessages()
	}

	if rp.response.Data == nil {
		return gjson.Result{}, errorContext, errors.New("response contains no data")
	}

	result := gjson.GetBytes(rp.response.Data, path)
	if !result.Exists() {
		errorContext["available_paths"] = rp.GetAllPaths()
		return gjson.Result{}, errorContext, fmt.Errorf("path '%s' not found in response data", path)
	}

	return result, errorContext, nil
}

// ValidateDataPath checks if a path exists and optionally validates the data type.
func (rp *ResponseParser) ValidateDataPath(path string, expectedType string) error {
	if !rp.PathExists(path) {
		return fmt.Errorf("path '%s' does not exist in response data", path)
	}

	if expectedType == "" {
		return nil // No type validation requested
	}

	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return err
	}

	// Validate data type
	switch expectedType {
	case "string":
		if result.Type != gjson.String {
			return fmt.Errorf("path '%s' expected to be string but got %s", path, result.Type.String())
		}
	case "number":
		if result.Type != gjson.Number {
			return fmt.Errorf("path '%s' expected to be number but got %s", path, result.Type.String())
		}
	case "boolean":
		if result.Type != gjson.True && result.Type != gjson.False {
			return fmt.Errorf("path '%s' expected to be boolean but got %s", path, result.Type.String())
		}
	case "array":
		if !result.IsArray() {
			return fmt.Errorf("path '%s' expected to be array but got %s", path, result.Type.String())
		}
	case "object":
		if !result.IsObject() {
			return fmt.Errorf("path '%s' expected to be object but got %s", path, result.Type.String())
		}
	default:
		return fmt.Errorf("unknown expected type: %s", expectedType)
	}

	return nil
}

// GetErrorDetails returns detailed error information for validation steps.
func (rp *ResponseParser) GetErrorDetails() map[string]interface{} {
	if !rp.response.HasErrors() {
		return map[string]interface{}{
			"has_errors": false,
			"count":      0,
		}
	}

	summary := rp.response.GetErrorSummary()
	details := rp.response.GetDetailedErrorInfo()

	return map[string]interface{}{
		"has_errors":       true,
		"count":            summary.TotalErrors,
		"summary":          summary,
		"details":          details,
		"has_critical":     rp.response.HasCriticalErrors(),
		"first_error":      rp.response.GetFirstError(),
		"messages":         rp.response.GetErrorMessages(),
		"errors_as_string": rp.response.GetErrorsAsString(),
	}
}

// GetResponseMetadata returns metadata about the response for context preservation.
func (rp *ResponseParser) GetResponseMetadata() map[string]interface{} {
	return map[string]interface{}{
		"has_data":            rp.response.HasData(),
		"has_errors":          rp.response.HasErrors(),
		"has_extensions":      len(rp.response.Extensions) > 0,
		"status_code":         rp.response.StatusCode,
		"is_successful":       rp.response.IsSuccessful(),
		"has_critical_errors": rp.response.HasCriticalErrors(),
		"error_count":         len(rp.response.Errors),
		"available_paths":     rp.GetAllPaths(),
	}
}

// ExtractVariableData extracts data suitable for storing in scenario context variables.
func (rp *ResponseParser) ExtractVariableData(path string) (interface{}, error) {
	result, errorContext, err := rp.GetDataWithErrorContext(path)
	if err != nil {
		// Include error context for better debugging
		return nil, fmt.Errorf("failed to extract variable data: %w (context: %+v)", err, errorContext)
	}

	// Return the appropriate Go type based on the JSON type
	switch result.Type {
	case gjson.String:
		return result.String(), nil
	case gjson.Number:
		return result.Float(), nil
	case gjson.True:
		return true, nil
	case gjson.False:
		return false, nil
	case gjson.JSON:
		// For objects and arrays, return as interface{}
		var value interface{}
		if errJSONDecode := json.Unmarshal([]byte(result.Raw), &value); errJSONDecode != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON data at path '%s': %w", path, errJSONDecode)
		}
		return value, nil
	case gjson.Null:
		return result.Value(), nil
	default:
		return result.Value(), nil
	}
}
