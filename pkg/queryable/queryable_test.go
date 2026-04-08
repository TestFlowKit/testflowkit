package queryable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectFormatFromContentType(t *testing.T) {
	require.Equal(t, FormatJSON, DetectFormatFromContentType("application/json; charset=utf-8"))
	require.Equal(t, FormatXML, DetectFormatFromContentType("application/soap+xml; charset=utf-8"))
	require.Equal(t, FormatAuto, DetectFormatFromContentType(""))
}

func TestNewEngine_JSONSupportsQueries(t *testing.T) {
	body := []byte(
		`{"status":"ok","enabled":true,"count":2,"info":{"email":"contact@example.com"},` +
			`"tags":["alpha","beta"]}`,
	)

	engine, err := NewEngine(body, FormatJSON)
	require.NoError(t, err)

	result, err := engine.Get("info.email")
	require.NoError(t, err)
	require.True(t, result.Exists)
	require.Equal(t, "contact@example.com", result.Raw)
	require.Equal(t, "string", result.Kind)

	exists, err := engine.Exists("status")
	require.NoError(t, err)
	require.True(t, exists)

	results, err := engine.GetAll("tags")
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Equal(t, "alpha", results[0].Raw)
	require.Equal(t, "beta", results[1].Raw)
}

func TestNewEngine_XMLSupportsXPathQueries(t *testing.T) {
	body := []byte(
		`<root><status>ok</status><enabled>true</enabled><count>2</count>` +
			`<info><email>contact@example.com</email></info>` +
			`<tags><tag priority="high">alpha</tag><tag priority="low">beta</tag></tags>` +
			`<orders><order><id>1001</id><total>42.50</total></order>` +
			`<order><id>1002</id><total>10.00</total></order></orders></root>`,
	)

	engine, err := NewEngine(body, FormatXML)
	require.NoError(t, err)

	result, err := engine.Get("//info/email")
	require.NoError(t, err)
	require.True(t, result.Exists)
	require.Equal(t, "contact@example.com", result.Raw)
	require.Equal(t, "string", result.Kind)

	boolResult, err := engine.Get("//enabled")
	require.NoError(t, err)
	require.Equal(t, "boolean", boolResult.Kind)

	countResult, err := engine.Get("//count")
	require.NoError(t, err)
	require.Equal(t, "integer", countResult.Kind)

	attributeResult, err := engine.Get("//tags/tag[@priority='high']/@priority")
	require.NoError(t, err)
	require.True(t, attributeResult.Exists)
	require.Equal(t, "high", attributeResult.Raw)

	predicateResult, err := engine.Get("//orders/order[id='1002']/total")
	require.NoError(t, err)
	require.True(t, predicateResult.Exists)
	require.Equal(t, "10.00", predicateResult.Raw)

	results, err := engine.GetAll("//tags/tag")
	require.NoError(t, err)
	require.Len(t, results, 2)
}

func TestNewEngine_AutoDetectsXMLFromBody(t *testing.T) {
	engine, err := NewEngine([]byte(`<root><name>John</name></root>`), FormatAuto)
	require.NoError(t, err)

	result, err := engine.Get("//name")
	require.NoError(t, err)
	require.True(t, result.Exists)
	require.Equal(t, "John", result.Raw)
}

func TestNewEngine_RejectsXPathSyntaxForJSON(t *testing.T) {
	engine, err := NewEngine([]byte(`{"status":"ok"}`), FormatJSON)
	require.NoError(t, err)

	_, err = engine.Get("//status")
	require.EqualError(t, err, "XPath syntax detected on JSON content; please use GJSON syntax")
}

func TestNewEngine_RejectsGJSONSyntaxForXML(t *testing.T) {
	engine, err := NewEngine([]byte(`<root><user><id>123</id></user></root>`), FormatXML)
	require.NoError(t, err)

	_, err = engine.Get(`users.#(id==123).name`)
	require.EqualError(t, err, "JSONPath/GJSON syntax detected on XML content; please use XPath")
}
