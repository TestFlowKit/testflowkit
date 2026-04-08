package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetJSONPathValue(t *testing.T) {
	jsonBody := []byte(`{"name": "John", "age": 30}`)
	value, err := GetJSONPathValue(jsonBody, "name")
	require.NoError(t, err)
	require.Equal(t, "John", value)
}

func TestGetUndefineJSONPathValue(t *testing.T) {
	jsonBody := []byte(`{"name": "John", "age": 30}`)
	value, err := GetJSONPathValue(jsonBody, "status")
	require.Error(t, err)
	require.Nil(t, value)
}

func TestGetResponsePathValue_XML(t *testing.T) {
	xmlBody := []byte(`<root><user><name>John</name><active>true</active></user></root>`)

	value, err := GetResponsePathValue(xmlBody, "//user/name")
	require.NoError(t, err)
	require.Equal(t, "John", value)
}
