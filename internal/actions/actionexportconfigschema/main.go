package actionexportconfigschema

import (
	"encoding/json"
	"fmt"
	"os"
	"testflowkit/pkg/logger"
	"testflowkit/scripts/configschema"
)

func Execute(format, version string) {
	if format != "json" {
		fmt.Fprintf(os.Stderr, "error: unsupported format %q — only \"json\" is supported\n", format)
		os.Exit(1)
	}

	jsonData, err := configschema.ExportJSON(version)
	if err != nil {
		logger.Fatal("Error exporting config schema", err)
	}

	if !json.Valid(jsonData) {
		logger.Fatal("Exported config schema is not valid JSON", nil)
	}

	if _, writeErr := os.Stdout.Write(jsonData); writeErr != nil {
		logger.Fatal("Error writing config schema to stdout", writeErr)
	}
}
