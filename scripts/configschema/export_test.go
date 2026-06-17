package configschema

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExportJSON_ValidJSON(t *testing.T) {
	data, err := ExportJSON("1.0.0-test")
	require.NoError(t, err)

	var export Export
	require.NoError(t, json.Unmarshal(data, &export))

	assert.Equal(t, "Config", export.Root)
	assert.Equal(t, "1.0.0-test", export.Version)
	require.NotNil(t, export.Schema)
	assert.Equal(t, schemaTypeObject, export.Schema.Type)
}

func TestExportJSON_KeyPathsExist(t *testing.T) {
	data, err := ExportJSON("test")
	require.NoError(t, err)

	var export Export
	require.NoError(t, json.Unmarshal(data, &export))

	paths := []struct {
		path         string
		required     *bool
		containsDesc string
	}{
		{"settings.gherkin_location", nil, "Gherkin"},
		{"apis.definitions", nil, "API"},
		{"agent.default_tags_for_draft", boolPtr(false), "draft"},
		{"security_schemes", nil, "auth"},
		{"agent", boolPtr(false), "MCP"},
	}

	for _, tc := range paths {
		t.Run(tc.path, func(t *testing.T) {
			node := findNodeByPath(export.Schema, tc.path)
			require.NotNil(t, node, "path %q not found", tc.path)
			if tc.required != nil {
				assert.Equal(t, *tc.required, node.Required, "required mismatch for %q", tc.path)
			}
			if tc.containsDesc != "" {
				assert.True(t,
					strings.Contains(strings.ToLower(node.Description), strings.ToLower(tc.containsDesc)) ||
						strings.Contains(strings.ToLower(node.TypeDescription), strings.ToLower(tc.containsDesc)),
					"expected description mentioning %q for path %q, got description=%q type_description=%q",
					tc.containsDesc, tc.path, node.Description, node.TypeDescription,
				)
			}
		})
	}
}

func TestExportJSON_NoEnabledField(t *testing.T) {
	data, err := ExportJSON("test")
	require.NoError(t, err)

	var export Export
	require.NoError(t, json.Unmarshal(data, &export))

	debug := findNodeByPath(export.Schema, "settings.debug")
	require.NotNil(t, debug)

	_, hasEnabled := debug.Properties["enabled"]
	assert.False(t, hasEnabled, "CLI-only Enabled field must not appear in schema")
}

func TestExportJSON_NoEmptyDescriptions(t *testing.T) {
	data, err := ExportJSON("test")
	require.NoError(t, err)

	var export Export
	require.NoError(t, json.Unmarshal(data, &export))

	missing := CollectEmptyDescriptions(export.Schema, "")
	assert.Empty(t, missing, "fields missing godoc synopsis: %v", missing)
}

func TestExportJSON_MultiLineGodocExported(t *testing.T) {
	data, err := ExportJSON("test")
	require.NoError(t, err)

	var export Export
	require.NoError(t, json.Unmarshal(data, &export))

	debug := findNodeByPath(export.Schema, "settings.debug")
	require.NotNil(t, debug)
	assert.NotEmpty(t, debug.TypeDescription)
	assert.Contains(t, debug.TypeDescription, "debug")

	prettyPrint := findNodeByPath(export.Schema, "settings.debug.pretty_print")
	require.NotNil(t, prettyPrint)
	assert.NotEmpty(t, prettyPrint.Description)
}

func TestExportJSON_ReportFormatEnum(t *testing.T) {
	data, err := ExportJSON("test")
	require.NoError(t, err)

	var export Export
	require.NoError(t, json.Unmarshal(data, &export))

	reportFormat := findNodeByPath(export.Schema, "settings.report_format")
	require.NotNil(t, reportFormat)
	assert.Equal(t, []string{"html", "json", "junit"}, reportFormat.Enum)
}

func TestSplitDocComment(t *testing.T) {
	synopsis, details := splitDocComment("Short summary. More detail here.\n\nSecond paragraph.")
	assert.Equal(t, "Short summary.", synopsis)
	assert.Contains(t, details, "More detail here.")
}

func boolPtr(v bool) *bool { return &v }
