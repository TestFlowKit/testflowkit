package gherkinparser

import (
	"testing"

	messages "github.com/cucumber/messages/go/v21"
)

type TestCase struct {
	name         string
	step         *messages.Step
	macroTitles  []string
	expectedCall *MacroCall
	expectError  bool
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
		macroTitles: []string{"user login with credentials"},
		expectedCall: &MacroCall{
			MacroName: "user login with credentials",
			Variables: map[string]string{
				"username": "oki",
				"password": "ler123",
			},
		},
		expectError: false,
	},
	{
		name: "handles macro call without table",
		step: &messages.Step{
			Text: "user logout",
		},
		macroTitles: []string{"user logout"},
		expectedCall: &MacroCall{
			MacroName: "user logout",
			Variables: map[string]string{},
		},
		expectError: false,
	},
	{
		name: "returns error for non-macro step",
		step: &messages.Step{
			Text: "some other step",
		},
		macroTitles:  []string{"user login with credentials"},
		expectedCall: nil,
		expectError:  true,
	},
}

func TestParseMacroCallWithTable(t *testing.T) {
	for _, tt := range parseMacroCallWithTableTests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseMacroCallWithTable(tt.step, tt.macroTitles)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			testParseMacroCallWithTableAssertions(t, result, tt)
		})
	}
}

func testParseMacroCallWithTableAssertions(t *testing.T, result *MacroCall, tt TestCase) {
	if result.MacroName != tt.expectedCall.MacroName {
		t.Errorf("expected macro name %s, got %s", tt.expectedCall.MacroName, result.MacroName)
	}

	if len(result.Variables) != len(tt.expectedCall.Variables) {
		t.Errorf("expected %d variables, got %d", len(tt.expectedCall.Variables), len(result.Variables))
	}

	for key, expectedValue := range tt.expectedCall.Variables {
		actualValue, exists := result.Variables[key]
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
		stepText  string
		variables map[string]string
		expected  string
	}{
		{
			name:     "substitutes single variable",
			stepText: "the user fills the username field with |username|",
			variables: map[string]string{
				"username": "oki",
			},
			expected: "the user fills the username field with oki",
		},
		{
			name:     "substitutes multiple variables",
			stepText: "the user fills the |field| field with |value|",
			variables: map[string]string{
				"field": "username",
				"value": "oki",
			},
			expected: "the user fills the username field with oki",
		},
		{
			name:      "handles no variables",
			stepText:  "the user clicks the button",
			variables: map[string]string{},
			expected:  "the user clicks the button",
		},
		{
			name:     "handles variable not in map",
			stepText: "the user fills the |username| field with |password|",
			variables: map[string]string{
				"username": "oki",
			},
			expected: "the user fills the oki field with |password|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubstituteVariables(tt.stepText, tt.variables)
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

func TestMacroCallStruct(t *testing.T) {
	// Test the MacroCall struct
	mc := MacroCall{
		MacroName: "user login with credentials",
		Variables: map[string]string{
			"username": "oki",
			"password": "ler123",
		},
	}

	if mc.MacroName != "user login with credentials" {
		t.Errorf("expected macro name 'user login with credentials', got %s", mc.MacroName)
	}

	if len(mc.Variables) != 2 {
		t.Errorf("expected 2 variables, got %d", len(mc.Variables))
	}

	if mc.Variables["username"] != "oki" {
		t.Errorf("expected username 'oki', got %s", mc.Variables["username"])
	}

	if mc.Variables["password"] != "ler123" {
		t.Errorf("expected password 'ler123', got %s", mc.Variables["password"])
	}
}
