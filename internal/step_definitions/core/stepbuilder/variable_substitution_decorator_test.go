package stepbuilder

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testContextKey struct{}

func newVarContextMock(scCtx *scenario.Context) context.Context {
	return context.WithValue(context.TODO(), scenario.ContextKey{}, scCtx)
}

func TestVariableSubstitutionDecorator_ReplaceStringVariable(t *testing.T) {
	step := NewWithOneVariable(
		[]string{"sentence"},
		func(ctx context.Context, value string) (context.Context, error) {
			return context.WithValue(ctx, testContextKey{}, value), nil
		},
		nil,
		DocParams{
			Description: "description",
			Variables:   []DocVariable{},
		},
	)
	decorator := VariableSubstitutionDecorator{step: step}
	definition, ok := decorator.GetDefinition().(func(context.Context, string) (context.Context, error))
	require.True(t, ok, "Failed to cast definition to definitionType")

	scCtx := scenario.NewContext(&config.Config{}, make(map[string]any))
	const value = "setted variable"

	const key = "key"
	scCtx.SetVariable(key, value)

	defCtx := newVarContextMock(scCtx)
	ctx, _ := definition(defCtx, fmt.Sprintf("{{ %s }}", key))
	assert.Equal(t, value, ctx.Value(testContextKey{}))
}

func TestVariableSubstitutionDecorator_ReplaceMultipleStringVariables(t *testing.T) {
	step := NewWithTwoVariables(
		[]string{"sentence"},
		func(ctx context.Context, value1, value2 string) (context.Context, error) {
			return context.WithValue(ctx, testContextKey{}, []string{value1, value2}), nil
		},
		nil,
		DocParams{
			Description: "description",
			Variables:   []DocVariable{},
		},
	)
	decorator := VariableSubstitutionDecorator{step: step}
	definition, ok := decorator.GetDefinition().(func(context.Context, string, string) (context.Context, error))
	require.True(t, ok, "Failed to cast definition to definitionType")

	scCtx := scenario.NewContext(&config.Config{}, make(map[string]any))
	const value = "setted variable"
	const value2 = "setted variable 2"
	const key1 = "key1"
	const key2 = "key2"
	scCtx.SetVariable(key1, value)
	scCtx.SetVariable(key2, value2)

	defCtx := newVarContextMock(scCtx)
	ctx, _ := definition(defCtx, fmt.Sprintf("{{ %s }}", key1), fmt.Sprintf("{{ %s }}", key2))
	assert.Equal(t, []string{value, value2}, ctx.Value(testContextKey{}))
}

func TestVariableSubstitutionDecorator_RemovesQuotes(t *testing.T) {
	step := NewWithOneVariable(
		[]string{"sentence"},
		func(ctx context.Context, value string) (context.Context, error) {
			return context.WithValue(ctx, testContextKey{}, value), nil
		},
		nil,
		DocParams{
			Description: "description",
			Variables:   []DocVariable{},
		},
	)
	decorator := VariableSubstitutionDecorator{step: step}
	definition, ok := decorator.GetDefinition().(func(context.Context, string) (context.Context, error))
	require.True(t, ok, "Failed to cast definition to definitionType")

	scCtx := scenario.NewContext(&config.Config{}, make(map[string]any))
	const value = "setted variable"

	const key = "key"
	scCtx.SetVariable(key, value)

	defCtx := newVarContextMock(scCtx)
	ctx, err := definition(defCtx, fmt.Sprintf("{{ %s }}", key))
	require.NoError(t, err)
	assert.Equal(t, value, ctx.Value(testContextKey{}))
}

func TestVariableSubstitutionDecorator_ReplaceJSONVariable(t *testing.T) {
	step := NewWithOneVariable(
		[]string{"sentence"},
		func(ctx context.Context, value string) (context.Context, error) {
			return context.WithValue(ctx, testContextKey{}, value), nil
		},
		nil,
		DocParams{},
	)

	decorator := VariableSubstitutionDecorator{step: step}
	definition, ok := decorator.GetDefinition().(func(context.Context, string) (context.Context, error))
	require.True(t, ok, "Failed to cast definition to definitionType")

	scCtx := scenario.NewContext(&config.Config{}, make(map[string]any))
	const value = "setted variable"
	const key1 = "name"
	const key2 = "age"
	const value2 = "20"
	scCtx.SetVariable(key1, value)
	scCtx.SetVariable(key2, value2)

	jsonString := `{"name": "{{ name }}", "age": {{ age }}}`
	jsonString = scenario.ReplaceVariablesInString(scCtx, jsonString)

	const expectedJSONString = `{"name": "setted variable", "age": 20}`
	defCtx := newVarContextMock(scCtx)
	ctx, err := definition(defCtx, jsonString)
	actualJSON, ok := ctx.Value(testContextKey{}).(string)
	assert.True(t, ok)
	require.NoError(t, err)
	require.JSONEq(t, expectedJSONString, actualJSON)
}
