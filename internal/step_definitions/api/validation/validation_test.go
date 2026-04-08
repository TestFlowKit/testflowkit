package validation

import (
	"testing"

	"github.com/stretchr/testify/require"

	"testflowkit/pkg/queryable"
)

func TestValidatePathType_SupportsAliases(t *testing.T) {
	engine, err := queryable.NewEngine([]byte(`{"title":"hello","tags":["alpha","beta"]}`), queryable.FormatJSON)
	require.NoError(t, err)

	require.NoError(t, ValidatePathType(engine, "title", "text", true))
	require.NoError(t, ValidatePathType(engine, "tags", "list", true))
}

func TestValidatePathCount_XML(t *testing.T) {
	engine, err := queryable.NewEngine(
		[]byte(`<root><tags><tag priority="high">alpha</tag><tag priority="low">beta</tag></tags></root>`),
		queryable.FormatXML,
	)
	require.NoError(t, err)

	require.NoError(t, ValidatePathCount(engine, "//tags/tag", 2))
	require.NoError(t, ValidatePathType(engine, "//tags/tag", "list", true))
	require.NoError(t, ValidatePathValue(engine, "//tags/tag[@priority='high']", "alpha", true))
}

func TestValidatePathExists_PropagatesSyntaxGuard(t *testing.T) {
	engine, err := queryable.NewEngine([]byte(`{"status":"ok"}`), queryable.FormatJSON)
	require.NoError(t, err)

	err = ValidatePathExists(engine, "//status", true)
	require.EqualError(t, err, "XPath syntax detected on JSON content; please use GJSON syntax")
}

func TestValidateBodyEquals_XML(t *testing.T) {
	actual := []byte(`<root><status>ok</status><user id="1">John</user></root>`)
	expected := `<root>
		<status>ok</status>
		<user id="1">John</user>
	</root>`

	require.NoError(t, ValidateBodyEquals(actual, expected))
}
