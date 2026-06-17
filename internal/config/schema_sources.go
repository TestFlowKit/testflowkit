package config

import "embed"

// Schema source files embedded for config schema export godoc extraction.
//
//go:embed main_type.go types.go security_types.go agent_types.go
var SchemaSourceFiles embed.FS

// SchemaSourceFileNames lists config type source files used for godoc extraction.
var SchemaSourceFileNames = []string{"main_type.go", "types.go", "security_types.go", "agent_types.go"}
