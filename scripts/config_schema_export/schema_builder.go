package main

import (
	"fmt"
	"reflect"
	"strings"

	"testflowkit/internal/config"
)

const (
	agentDocDescription = "IDE agent configuration consumed by the testflowkit-mcp server. Never read by the tkit CLI."
	envDescription      = "Key-value map of project environment variables. Accessible anywhere in config " +
		"and in Gherkin via {{ env.key }}. Keys are case-sensitive."
	defaultSecurityDescription = "Name of a security_scheme applied to all API requests when no API-level " +
		"or endpoint-level security_ref is set. Use \"none\" to disable inherited auth."
	securitySchemesDescription = "Named, reusable authentication credentials. Reference from an API or " +
		"endpoint via security_ref.name. Use \"none\" as the name to disable all inherited auth."
	versionDescription = "Schema version. Only 1 is currently valid. The MCP server rejects unsupported " +
		"versions with a clear error."
	projectConfigDescription = "Path to the testflowkit.yml (or legacy config.yml) runtime config. " +
		"Resolved relative to this file."
	stepCatalogDescription = "Controls how the MCP server resolves the step definitions catalog. Omit all " +
		"sub-fields to auto-fetch from the GitHub Release matching the installed tkit version."
	stepCatalogFileDescription = "Local path to a pre-generated step-definitions.json. Takes precedence over " +
		"url and release fetch. Useful for contributors or air-gapped environments."
	stepCatalogURLDescription = "Direct HTTPS URL to download the catalog JSON. Takes precedence over release " +
		"fetch. Supports {{ env.VAR }} interpolation."
	runCommandDescription = "CLI command used to run draft tests. Informational for Phase 3 run loop; not " +
		"executed by the MCP server in Phase 1."
)

func buildSchema(version string, comments map[string]map[string]string) Schema {
	version = strings.TrimPrefix(version, "v")
	return Schema{
		SchemaVersion: "1",
		GeneratedBy:   fmt.Sprintf("testflowkit v%s", version),
		Files: map[string]FileDoc{
			"testflowkit.yml":       buildTestflowkitYmlDoc(comments),
			"testflowkit.agent.yml": buildAgentYmlDoc(),
		},
	}
}

func buildTestflowkitYmlDoc(comments map[string]map[string]string) FileDoc {
	return FileDoc{
		Description:   "Runtime test configuration consumed by the tkit CLI. Never read by the MCP server.",
		Interpolation: "{{ env.VAR_NAME }} can be used in any string value to inject an OS or env-section variable.",
		Sections: []Section{
			structSection("settings", reflect.TypeOf(config.GlobalSettings{}), true,
				"Global test runner settings.", comments),
			{
				Key:         "env",
				Required:    false,
				Type:        "map<string, any>",
				Description: envDescription,
			},
			structSection("frontend", reflect.TypeOf(config.FrontendConfig{}), false,
				"Browser/UI test configuration. Omit entirely if the project has no frontend scenarios.", comments),
			securitySchemesSection(comments),
			{
				Key:         "default_security",
				Required:    false,
				Type:        schemaTypeString,
				Description: defaultSecurityDescription,
			},
			structSection("apis", reflect.TypeOf(config.APIsConfig{}), false,
				"REST and GraphQL API definitions. Omit entirely if the project has no API scenarios.", comments),
		},
	}
}

// structSection reflects all yaml-tagged fields of t into a Section.
func structSection(
	key string,
	t reflect.Type,
	required bool,
	description string,
	comments map[string]map[string]string,
) Section {
	return Section{
		Key:         key,
		Required:    required,
		Description: description,
		Fields:      reflectStructFields(t, comments),
	}
}

// securitySchemesSection builds the discriminator-variant section for
// security_schemes by reflecting SecurityScheme fields from the Go source.
// Field metadata (types, descriptions, constraints) is fully derived from the
// struct; the variant grouping is the only manually declared domain knowledge.
func securitySchemesSection(comments map[string]map[string]string) Section {
	all := reflectStructFields(reflect.TypeOf(config.SecurityScheme{}), comments)
	fm := fieldMapByKey(all)

	typeField := fm["type"]
	typeField.Required = true

	variantKeys := map[string][]string{
		"bearer": {"token"},
		"basic":  {"username", "password"},
		"apikey": {"key", "placement", "header_name", "query_param"},
		"oauth2": {
			"client_id",
			"client_secret",
			"token_url",
			"token_endpoint_auth_method",
			"scopes",
			"audience",
			"duration",
			"retry_on_401",
			"proxy_url",
		},
		"certificate": {"cert_file", "key_file"},
	}
	variantOrder := []string{"bearer", "basic", "apikey", "oauth2", "certificate"}

	var variants []Variant
	for _, typeName := range variantOrder {
		fields := []Field{withEnum(typeField, []string{typeName})}
		for _, k := range variantKeys[typeName] {
			if f, ok := fm[k]; ok {
				fields = append(fields, f)
			}
		}
		variants = append(variants, Variant{When: "type: " + typeName, Fields: fields})
	}

	return Section{
		Key:           "security_schemes",
		Required:      false,
		Type:          "map<string, SecurityScheme>",
		Description:   securitySchemesDescription,
		Discriminator: "type",
		Variants:      variants,
	}
}

func withEnum(f Field, enum []string) Field {
	f.Enum = enum
	return f
}

func buildAgentYmlDoc() FileDoc {
	return FileDoc{
		Description: agentDocDescription,
		Sections: []Section{
			{
				Key:         "version",
				Required:    true,
				Type:        schemaTypeInt,
				Description: versionDescription,
				Fields: []Field{
					{Key: "version", Type: schemaTypeInt, Required: true, Enum: []string{"1"}, Description: "Must be 1."},
				},
			},
			{
				Key:         "project",
				Required:    false,
				Description: "Paths used by the MCP server to locate runtime config and feature files.",
				Fields: []Field{
					{Key: "test_config", Type: schemaTypeString, Default: "./testflowkit.yml",
						Description: projectConfigDescription},
					{Key: "features_glob", Type: schemaTypeString, Default: "features/**/*.feature",
						Description: "Glob pattern for feature files. Used by list_features, read_feature, and write_feature."},
				},
			},
			{
				Key:         "step_catalog",
				Required:    false,
				Description: stepCatalogDescription,
				Fields: []Field{
					{Key: "file", Type: schemaTypeString,
						Description: stepCatalogFileDescription},
					{Key: "url", Type: schemaTypeString, Interpolation: true,
						Description: stepCatalogURLDescription},
				},
			},
			{
				Key:         "agent",
				Required:    false,
				Description: "Agent behaviour configuration.",
				Fields: []Field{
					{Key: "default_tags_for_draft", Type: schemaTypeString, Default: "@wip @ai-generated",
						Description: "Space-separated Gherkin tags added to every scenario drafted by the agent."},
					{Key: "run_command", Type: schemaTypeString, Default: "tkit run -c testflowkit.yml --tags @wip",
						Description: runCommandDescription},
				},
			},
		},
	}
}
