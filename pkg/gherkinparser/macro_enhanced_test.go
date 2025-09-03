package gherkinparser

import (
	"fmt"
	"testing"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name              string
	step              *messages.Step
	macros            []*messages.Scenario
	expectedVariables MacroVariables
	expectError       bool
}

var parseMacroCallWithTableTests = []TestCase{
	{
		name: "parses macro call with table data",
		step: &messages.Step{
			Text: "user login with credentials",
			DataTable: &messages.DataTable{
				Rows: []*messages.TableRow{
					{
						Cells: []*messages.TableCell{
							{Value: "username"},
							{Value: "password"},
						},
					},
					{
						Cells: []*messages.TableCell{
							{Value: "oki"},
							{Value: "ler123"},
						},
					},
				},
			},
		},
		macros: []*messages.Scenario{
			{
				Name: "user login with credentials",
			},
		},
		expectedVariables: map[string]string{
			"username": "oki",
			"password": "ler123",
		},
		expectError: false,
	},
	{
		name: "handles macro call without table",
		step: &messages.Step{
			Text: "user logout",
		},
		macros: []*messages.Scenario{
			{
				Name: "user logout",
			},
		},
		expectedVariables: map[string]string{},
		expectError:       false,
	},
	{
		name: "returns error for non-macro step",
		step: &messages.Step{
			Text: "some other step",
		},
		macros: []*messages.Scenario{
			{
				Name: "user login with credentials",
			},
		},
		expectedVariables: nil,
		expectError:       true,
	},
}

func TestParseMacroCallWithTable(t *testing.T) {
	for _, tt := range parseMacroCallWithTableTests[2:3] {
		t.Run(tt.name, func(t *testing.T) {
			result := getMacroVariables(tt.step)
			testParseMacroCallWithTableAssertions(t, result, tt)
		})
	}
}

func testParseMacroCallWithTableAssertions(t *testing.T, result MacroVariables, tt TestCase) {
	if len(result) != len(tt.expectedVariables) {
		t.Errorf("expected %d variables, got %d", len(tt.expectedVariables), len(result))
	}

	for key, expectedValue := range tt.expectedVariables {
		actualValue, exists := result[key]
		if !exists {
			t.Errorf("expected variable %s to exist", key)
			continue
		}

		if actualValue != expectedValue {
			t.Errorf("expected variable %s to have value %s, got %s", key, expectedValue, actualValue)
		}
	}
}
func TestSubstituteVariables(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		variables map[string]string
		expected  string
	}{
		{
			name: "substitutes single variable",
			text: "the user fills the username field with |username|",
			variables: map[string]string{
				"username": "oki",
			},
			expected: "the user fills the username field with oki",
		},
		{
			name: "substitutes multiple variables",
			text: "the user fills the |field| field with |value|",
			variables: map[string]string{
				"field": "username",
				"value": "oki",
			},
			expected: "the user fills the username field with oki",
		},
		{
			name:      "handles no variables",
			text:      "the user clicks the button",
			variables: map[string]string{},
			expected:  "the user clicks the button",
		},
		{
			name: "handles variable not in map",
			text: "the user fills the |username| field with |password|",
			variables: map[string]string{
				"username": "oki",
			},
			expected: "the user fills the oki field with |password|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := substituteVariables(tt.text, tt.variables)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestMacroVariableStruct(t *testing.T) {
	// Test the MacroVariable struct
	mv := MacroVariable{
		Name:  "username",
		Value: "oki",
	}

	if mv.Name != "username" {
		t.Errorf("expected name 'username', got %s", mv.Name)
	}

	if mv.Value != "oki" {
		t.Errorf("expected value 'oki', got %s", mv.Value)
	}
}

func Test_GetCompleteContentForSimpleStep(t *testing.T) {
	stepDef := messages.Step{
		Text: "user type in ok input",
	}

	assert.Equal(t, "user type in ok input", getCompleteStepContentWhithoutKeyword(&stepDef))
}

func Test_GetCompleteContentForStepWithDocString(t *testing.T) {
	docString := messages.DocString{
		Content: `{
			"title": "Test Post Title",
			"body": "This is a test post body for API testing",
			"userId": 1
		}`,
		Delimiter: `"""`,
	}

	stepDef := messages.Step{
		Text:      "user send a reuest",
		DocString: &docString,
	}

	expected := fmt.Sprintf("%s\n%s\n%s\n%s", stepDef.Text, docString.Delimiter, docString.Content, docString.Delimiter)
	assert.Equal(t, expected, getCompleteStepContentWhithoutKeyword(&stepDef))
}

func Test_GetStepEndLine(t *testing.T) {
	tests := []struct {
		name           string
		featureContent []string
		stepStartLine  int
		expected       int
	}{
		{
			name: "simple step without docstring or datatable",
			featureContent: []string{
				"Given I am on the homepage",
				"When I click login",
				"Then I should see welcome message",
			},
			stepStartLine: 0,
			expected:      0,
		},
		{
			name: "step with docstring",
			featureContent: []string{
				"Given I send request with body",
				`"""`,
				`{"name": "test"}`,
				`"""`,
				"Then I should see response",
			},
			stepStartLine: 0,
			expected:      3,
		},
		{
			name: "step with datatable",
			featureContent: []string{
				"Given I have following users",
				"| name  | age |",
				"| John  | 30  |",
				"| Alice | 25  |",
				"When I query users",
			},
			stepStartLine: 0,
			expected:      3,
		},
		{
			name: "last step in feature",
			featureContent: []string{
				"Given first step",
				"When second step",
				"Then final step",
				"| data |",
				"| test |",
			},
			stepStartLine: 2,
			expected:      4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getStepEndLine(tt.stepStartLine, tt.featureContent)
			assert.Equal(t, tt.expected, result)
		})
	}
}
