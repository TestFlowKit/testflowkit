package graphql

import (
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

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

// ValidateDataPath checks if a path exists and optionally validates the data type.
func (rp *ResponseParser) ValidateDataPath(path string, expectedType string) error {
	if !rp.PathExists(path) {
		return fmt.Errorf("path '%s' does not exist in response data", path)
	}

	if expectedType == "" {
		return nil
	}

	result, err := rp.GetDataAtPath(path)
	if err != nil {
		return err
	}

	validator, exists := rp.getTypeValidator(expectedType)
	if !exists {
		return fmt.Errorf("unknown expected type: %s", expectedType)
	}

	if !validator(result) {
		return fmt.Errorf("path '%s' expected to be %s but got %s", path, expectedType, result.Type.String())
	}

	return nil
}

func (rp *ResponseParser) getTypeValidator(expectedType string) (func(gjson.Result) bool, bool) {
	typeValidators := map[string]func(gjson.Result) bool{
		"string": func(r gjson.Result) bool {
			return r.Type == gjson.String
		},
		"number": func(r gjson.Result) bool {
			return r.Type == gjson.Number
		},
		"boolean": func(r gjson.Result) bool {
			return r.Type == gjson.True || r.Type == gjson.False
		},
		"array": func(r gjson.Result) bool {
			return r.IsArray()
		},
		"object": func(r gjson.Result) bool {
			return r.IsObject()
		},
	}

	validator, exists := typeValidators[expectedType]
	return validator, exists
}
