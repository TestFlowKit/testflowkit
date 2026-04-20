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

func TestValidateBodyContainsPartial_SubsetMatch(t *testing.T) {
	actual := []byte(`{"status":"success","data":{"id":1,"name":"John"},"meta":"extra"}`)
	require.NoError(t, ValidateBodyContainsPartial(actual, `{"status":"success","data":{"id":1}}`))
}

func TestValidateBodyContainsPartial_ExactMatch(t *testing.T) {
	actual := []byte(`{"status":"ok"}`)
	require.NoError(t, ValidateBodyContainsPartial(actual, `{"status":"ok"}`))
}

func TestValidateBodyContainsPartial_MissingKey(t *testing.T) {
	actual := []byte(`{"status":"success"}`)
	err := ValidateBodyContainsPartial(actual, `{"status":"success","missing":"field"}`)
	require.Error(t, err)
	require.Contains(t, err.Error(), "missing")
}

func TestValidateBodyContainsPartial_WrongValue(t *testing.T) {
	actual := []byte(`{"status":"error"}`)
	err := ValidateBodyContainsPartial(actual, `{"status":"success"}`)
	require.Error(t, err)
}

func TestValidateBodyContainsPartial_NilBody(t *testing.T) {
	err := ValidateBodyContainsPartial(nil, `{"status":"ok"}`)
	require.Error(t, err)
}

func TestValidateBodyContainsPartial_EmptyBody(t *testing.T) {
	err := ValidateBodyContainsPartial([]byte(``), `{"status":"ok"}`)
	require.Error(t, err)
}

func TestValidateBodyContainsPartial_EmptyExpected(t *testing.T) {
	err := ValidateBodyContainsPartial([]byte(`{"status":"ok"}`), ``)
	require.Error(t, err)
}

func TestValidateBodyContainsPartial_ComplexJSON(t *testing.T) {
	actual := []byte(`{
		"status": "success",
		"pagination": {"page": 1, "total": 100, "perPage": 10},
		"data": {
			"user": {
				"id": 42,
				"name": "Jane Doe",
				"email": "jane@example.com",
				"roles": ["admin", "editor"],
				"address": {
					"city": "Paris",
					"country": "France",
					"zip": "75001"
				}
			},
			"posts": [
				{"id": 1, "title": "Hello World", "published": true},
				{"id": 2, "title": "Second Post", "published": false}
			]
		},
		"meta": {"requestId": "abc-123", "duration": 42}
	}`)

	t.Run("top-level subset", func(t *testing.T) {
		require.NoError(t, ValidateBodyContainsPartial(actual, `{"status":"success"}`))
	})

	t.Run("nested object subset", func(t *testing.T) {
		require.NoError(t, ValidateBodyContainsPartial(actual, `{"data":{"user":{"id":42,"name":"Jane Doe"}}}`))
	})

	t.Run("deeply nested subset", func(t *testing.T) {
		partial := `{"data":{"user":{"address":{"city":"Paris","country":"France"}}}}`
		require.NoError(t, ValidateBodyContainsPartial(actual, partial))
	})

	t.Run("array field subset", func(t *testing.T) {
		require.NoError(t, ValidateBodyContainsPartial(actual, `{"data":{"user":{"roles":["admin","editor"]}}}`))
	})

	t.Run("nested array of objects subset", func(t *testing.T) {
		posts := `[{"id":1,"title":"Hello World","published":true},{"id":2,"title":"Second Post","published":false}]`
		partial := `{"data":{"posts":` + posts + `}}`
		require.NoError(t, ValidateBodyContainsPartial(actual, partial))
	})

	t.Run("pagination subset", func(t *testing.T) {
		require.NoError(t, ValidateBodyContainsPartial(actual, `{"pagination":{"page":1,"total":100}}`))
	})

	t.Run("wrong nested value fails", func(t *testing.T) {
		err := ValidateBodyContainsPartial(actual, `{"data":{"user":{"id":99}}}`)
		require.Error(t, err)
		require.Contains(t, err.Error(), "id")
	})

	t.Run("missing nested key fails", func(t *testing.T) {
		err := ValidateBodyContainsPartial(actual, `{"data":{"user":{"address":{"city":"Paris","postalCode":"75001"}}}}`)
		require.Error(t, err)
		require.Contains(t, err.Error(), "postalCode")
	})

	t.Run("wrong array element fails", func(t *testing.T) {
		err := ValidateBodyContainsPartial(actual, `{"data":{"user":{"roles":["admin","viewer"]}}}`)
		require.Error(t, err)
	})
}
