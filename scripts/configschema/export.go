package configschema

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testflowkit/internal/config"
)

// ExportJSON builds the config schema tree from the Config struct and its
// godoc comments, then returns indented JSON suitable for AI agent consumption.
func ExportJSON(cliVersion string) ([]byte, error) {
	docs, err := loadDocIndex()
	if err != nil {
		return nil, fmt.Errorf("load doc index: %w", err)
	}

	schema := buildSchemaNode(reflect.TypeOf(config.Config{}), "Config", "", true, docs, "")

	export := Export{
		Root:    "Config",
		Version: cliVersion,
		Schema:  schema,
	}

	return json.MarshalIndent(export, "", "  ")
}
