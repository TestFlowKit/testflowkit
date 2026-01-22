package config

import (
	"fmt"
	"maps"
	"os"

	"github.com/goccy/go-yaml"
)

// LoadEnvFile loads a YAML file and flattens it into a map of strings.
// Nested YAML structures are converted to dot notation keys.
// Example: api: { url: "x" } becomes map["api.url"] = "x"
// Returns an error if the file cannot be read or parsed.
func LoadEnvFile(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read env file %s: %w", path, err)
	}

	var rawEnv map[string]any
	if errMarshall := yaml.Unmarshal(data, &rawEnv); errMarshall != nil {
		return nil, fmt.Errorf("failed to parse env file %s: %w", path, errMarshall)
	}

	return FlattenMap(rawEnv, ""), nil
}

// FlattenMap converts a nested map into a flat map with dot notation.
// Supports nested maps and arrays. Arrays are converted to comma-separated strings.
// Example inputs:
//
//	map["api"]["url"] = "x" -> map["api.url"] = "x"
//	map["tags"] = ["a", "b"] -> map["tags"] = "[a b]"
func FlattenMap(data map[string]any, prefix string) map[string]string {
	result := make(map[string]string)

	for k, v := range data {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}

		switch val := v.(type) {
		case map[string]any:
			nested := FlattenMap(val, key)
			maps.Copy(result, nested)
		case map[any]any:
			// Handle map[any]any which yaml might produce
			converted := make(map[string]any)
			for mk, mv := range val {
				converted[fmt.Sprintf("%v", mk)] = mv
			}
			nested := FlattenMap(converted, key)
			for nk, nv := range nested {
				result[nk] = nv
			}
		case []any:
			// Handle arrays - convert to comma-separated string
			var items []string
			for _, item := range val {
				items = append(items, fmt.Sprintf("%v", item))
			}
			if len(items) > 0 {
				result[key] = fmt.Sprintf("%v", items)
			} else {
				result[key] = ""
			}
		case nil:
			result[key] = ""
		default:
			result[key] = fmt.Sprintf("%v", val)
		}
	}

	return result
}
