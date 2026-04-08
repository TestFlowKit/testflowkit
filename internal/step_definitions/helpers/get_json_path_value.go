package helpers

import (
	"fmt"

	"testflowkit/pkg/queryable"
)

func GetResponsePathValue(responseBody []byte, path string) (any, error) {
	return getPathValue(responseBody, path, queryable.FormatAuto)
}

func GetJSONPathValue(jsonBody []byte, path string) (any, error) {
	return getPathValue(jsonBody, path, queryable.FormatJSON)
}

func getPathValue(body []byte, path string, format queryable.Format) (any, error) {
	engine, err := queryable.NewEngine(body, format)
	if err != nil {
		return nil, err
	}

	result, err := engine.Get(path)
	if err != nil {
		return nil, err
	}
	if !result.Exists {
		return nil, fmt.Errorf("response path '%s' does not exist in response body", path)
	}

	return result.Value, nil
}
