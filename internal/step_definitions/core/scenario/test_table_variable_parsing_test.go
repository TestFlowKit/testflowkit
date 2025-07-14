package scenario

import (
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableVariableParsing_MapData(t *testing.T) {
	// Create a scenario context with variables
	ctx := NewContext(&config.Config{})
	ctx.SetVariable("userName", "John Doe")
	ctx.SetVariable("userAge", "30")

	// Create test data that simulates what would come from a parsed table
	data := map[string]any{
		"Name": "{{ userName }}",
		"Age":  "{{ userAge }}",
	}

	// Parse variables in the table data
	parsedData, err := ReplaceVariablesInMap(ctx, data)
	require.NoError(t, err)

	// Verify the variables were replaced
	assert.Equal(t, "John Doe", parsedData["Name"])
	assert.Equal(t, "30", parsedData["Age"])
}

func TestTableVariableParsing_ArrayData(t *testing.T) {
	// Create a scenario context with variables
	ctx := NewContext(&config.Config{})
	ctx.SetVariable("userName", "Jane Smith")
	ctx.SetVariable("userAge", "25")

	// Create test data that simulates what would come from a parsed table
	data := []any{
		map[string]any{
			"Name": "{{ userName }}",
			"Age":  "{{ userAge }}",
		},
	}

	// Parse variables in the table data
	parsedData, err := ReplaceVariablesInArray(ctx, data)
	require.NoError(t, err)

	// Verify the variables were replaced
	rowMap, ok := parsedData[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "Jane Smith", rowMap["Name"])
	assert.Equal(t, "25", rowMap["Age"])
}

func TestTableVariableParsing_UndefinedVariables(t *testing.T) {
	// Create a scenario context without variables
	ctx := NewContext(&config.Config{})

	// Create test data with undefined variables
	data := map[string]any{
		"Name": "{{ undefinedVar }}",
		"Age":  "{{ anotherUndefinedVar }}",
	}

	// Parse variables in the table data
	parsedData, err := ReplaceVariablesInMap(ctx, data)
	require.NoError(t, err)

	// Verify undefined variables remain unchanged
	assert.Equal(t, "{{ undefinedVar }}", parsedData["Name"])
	assert.Equal(t, "{{ anotherUndefinedVar }}", parsedData["Age"])
}

func TestTableVariableParsing_MixedData(t *testing.T) {
	// Create a scenario context with variables
	ctx := NewContext(&config.Config{})
	ctx.SetVariable("userName", "Alice")
	ctx.SetVariable("userAge", "28")

	// Create test data with mixed defined and undefined variables
	data := map[string]any{
		"Name":     "{{ userName }}",
		"Age":      "{{ userAge }}",
		"Location": "{{ undefinedLocation }}",
		"Status":   "Active",
	}

	// Parse variables in the table data
	parsedData, err := ReplaceVariablesInMap(ctx, data)
	require.NoError(t, err)

	// Verify the variables were replaced correctly
	assert.Equal(t, "Alice", parsedData["Name"])
	assert.Equal(t, "28", parsedData["Age"])
	assert.Equal(t, "{{ undefinedLocation }}", parsedData["Location"])
	assert.Equal(t, "Active", parsedData["Status"])
}

func TestShouldReplaceVariableOccurenceByThisValue(t *testing.T) {
	ctx := NewContext(&config.Config{})
	ctx.SetVariable("testVar", "testValue")
	ctx.SetVariable("anotherVar", "anotherValue")

	const sentence = "This is a test sentence with a variable: {{ testVar }} and another one: {{anotherVar}}."
	replacedSentence := ReplaceVariablesInString(ctx, sentence)
	assert.Equal(t, "This is a test sentence with a variable: testValue and another one: anotherValue.", replacedSentence)
}

func TestShouldNotReplaceUndefinedVariable(t *testing.T) {
	ctx := NewContext(&config.Config{})
	ctx.SetVariable("anotherVar", "anotherValue")

	expectedSentence := "This is a test sentence with a variable: {{ testVar }} and another one: anotherValue."
	const sentence = "This is a test sentence with a variable: {{ testVar }} and another one: {{anotherVar}}."
	replacedSentence := ReplaceVariablesInString(ctx, sentence)
	assert.Equal(t, expectedSentence, replacedSentence)
}

func TestVariableSubstitutionDecorator_ReplaceVariableInATrueSentence(t *testing.T) {
	scCtx := NewContext(&config.Config{})
	const key, value = "postTitle", "My Post Title"
	const key2, value2 = "postField", "My Post Field"
	scCtx.SetVariable(key, value)
	scCtx.SetVariable(key2, value2)

	const sentence = "the user enters {{ postTitle }} into the {{ postField }} field"

	replacedSentence := ReplaceVariablesInString(scCtx, sentence)

	assert.Equal(t, "the user enters My Post Title into the My Post Field field", replacedSentence)
}
