package jsonhelpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsJSON_ExactMatch(t *testing.T) {
	expected := []byte(`{"status":"success","id":1}`)
	actual := []byte(`{"status":"success","id":1}`)
	require.NoError(t, ContainsJSON(expected, actual))
}

func TestContainsJSON_SubsetMatch(t *testing.T) {
	expected := []byte(`{"status":"success"}`)
	actual := []byte(`{"status":"success","data":{"id":1},"extra":"ignored"}`)
	require.NoError(t, ContainsJSON(expected, actual))
}

func TestContainsJSON_NestedSubsetMatch(t *testing.T) {
	expected := []byte(`{"data":{"id":1}}`)
	actual := []byte(`{"data":{"id":1,"name":"John"},"meta":"extra"}`)
	require.NoError(t, ContainsJSON(expected, actual))
}

func TestContainsJSON_ArrayMatch(t *testing.T) {
	expected := []byte(`{"tags":["alpha","beta"]}`)
	actual := []byte(`{"tags":["alpha","beta"],"other":true}`)
	require.NoError(t, ContainsJSON(expected, actual))
}

func TestContainsJSON_MissingKey(t *testing.T) {
	expected := []byte(`{"status":"success","missing":"field"}`)
	actual := []byte(`{"status":"success"}`)
	err := ContainsJSON(expected, actual)
	require.Error(t, err)
	require.Contains(t, err.Error(), "missing")
}

func TestContainsJSON_WrongValue(t *testing.T) {
	expected := []byte(`{"status":"success"}`)
	actual := []byte(`{"status":"error"}`)
	err := ContainsJSON(expected, actual)
	require.Error(t, err)
	require.Contains(t, err.Error(), "status")
}

func TestContainsJSON_NestedWrongValue(t *testing.T) {
	expected := []byte(`{"data":{"id":2}}`)
	actual := []byte(`{"data":{"id":1}}`)
	err := ContainsJSON(expected, actual)
	require.Error(t, err)
	require.Contains(t, err.Error(), "data.id")
}

func TestContainsJSON_ArrayLengthMismatch(t *testing.T) {
	expected := []byte(`{"tags":["alpha","beta","gamma"]}`)
	actual := []byte(`{"tags":["alpha","beta"]}`)
	err := ContainsJSON(expected, actual)
	require.Error(t, err)
	require.Contains(t, err.Error(), "length")
}

func TestContainsJSON_ArrayElementMismatch(t *testing.T) {
	expected := []byte(`{"tags":["alpha","delta"]}`)
	actual := []byte(`{"tags":["alpha","beta"]}`)
	err := ContainsJSON(expected, actual)
	require.Error(t, err)
}

func TestContainsJSON_TypeMismatch(t *testing.T) {
	expected := []byte(`{"id":1}`)
	actual := []byte(`{"id":"1"}`)
	err := ContainsJSON(expected, actual)
	require.Error(t, err)
	require.Contains(t, err.Error(), "id")
}

func TestContainsJSON_InvalidExpected(t *testing.T) {
	err := ContainsJSON([]byte(`not-json`), []byte(`{"id":1}`))
	require.Error(t, err)
	require.Contains(t, err.Error(), "unmarshal expected")
}

func TestContainsJSON_InvalidActual(t *testing.T) {
	err := ContainsJSON([]byte(`{"id":1}`), []byte(`not-json`))
	require.Error(t, err)
	require.Contains(t, err.Error(), "unmarshal actual")
}
