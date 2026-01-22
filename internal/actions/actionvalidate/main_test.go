package actionvalidate

import (
	"os"
	"path/filepath"
	"testflowkit/pkg/gherkinparser"
	"testflowkit/pkg/variables"
	"testing"
)

type envVarsTestCase struct {
	name              string
	gherkinContents   []string
	configContent     string
	envVars           map[string]string
	shouldExit        bool
	expectedUndefined []string
}

// collectUndefinedVars replicates the logic from validateEnvVars.
func collectUndefinedVars(parsedFeatures []*gherkinparser.Feature, configFilePath string) map[string]bool {
	undefinedEnvs := make(map[string]bool)

	// Check Gherkin files
	for _, f := range parsedFeatures {
		missing := variables.FindUndefinedEnvReferences(string(f.Contents))
		for _, m := range missing {
			undefinedEnvs[m] = true
		}
	}

	// Check config file
	if configFilePath != "" {
		configContent, err := os.ReadFile(configFilePath)
		if err == nil {
			missing := variables.FindUndefinedEnvReferences(string(configContent))
			for _, m := range missing {
				undefinedEnvs[m] = true
			}
		}
	}

	return undefinedEnvs
}

func verifyUndefinedVars(t *testing.T, undefinedEnvs map[string]bool, expected []string) {
	t.Helper()

	if len(undefinedEnvs) == 0 {
		t.Error("Expected undefined env vars but got none")
		return
	}

	// Check that all expected undefined vars are found
	for _, expectedVar := range expected {
		if !undefinedEnvs[expectedVar] {
			t.Errorf("Expected undefined var %q not found", expectedVar)
		}
	}

	// Check no unexpected vars
	if len(undefinedEnvs) != len(expected) {
		t.Errorf("Expected %d undefined vars, got %d", len(expected), len(undefinedEnvs))
	}
}

func createTempConfigFile(t *testing.T, content string) string {
	t.Helper()

	if content == "" {
		return ""
	}

	tmpFile, err := os.CreateTemp("", "config-*.yml")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	t.Cleanup(func() { os.Remove(tmpFile.Name()) })

	if _, errWriteTmpfFile := tmpFile.WriteString(content); errWriteTmpfFile != nil {
		t.Fatalf("Failed to write config file: %v", errWriteTmpfFile)
	}
	tmpFile.Close()

	return tmpFile.Name()
}

func TestValidateEnvVars(t *testing.T) {
	// Setup: Initialize environment variables for testing
	originalEnvStore := variables.GetAllEnvVariables()
	defer func() {
		variables.SetEnvVariables(originalEnvStore)
	}()

	tests := []envVarsTestCase{
		{
			name: "all env vars defined",
			gherkinContents: []string{
				`Given I set "{{ env.api_url }}" as base URL`,
			},
			configContent: `
settings:
  base_url: "{{ env.api_url }}"
`,
			envVars: map[string]string{
				"api_url": "http://localhost:3000",
			},
			shouldExit:        false,
			expectedUndefined: nil,
		},
		{
			name: "undefined env var in gherkin",
			gherkinContents: []string{
				`Given I set "{{ env.undefined_var }}" as value`,
			},
			configContent: `settings:
  timeout: 1000
`,
			envVars: map[string]string{
				"api_url": "http://localhost:3000",
			},
			shouldExit:        true,
			expectedUndefined: []string{"undefined_var"},
		},
		{
			name:            "undefined env var in config",
			gherkinContents: []string{},
			configContent: `
settings:
  base_url: "{{ env.missing_url }}"
`,
			envVars: map[string]string{
				"api_url": "http://localhost:3000",
			},
			shouldExit:        true,
			expectedUndefined: []string{"missing_url"},
		},
		{
			name: "multiple undefined env vars",
			gherkinContents: []string{
				`Given I set "{{ env.var1 }}" as value`,
				`And I use "{{ env.var2 }}" for testing`,
			},
			configContent: `
settings:
  url: "{{ env.var3 }}"
`,
			envVars: map[string]string{
				"defined_var": "value",
			},
			shouldExit:        true,
			expectedUndefined: []string{"var1", "var2", "var3"},
		},
		{
			name: "duplicate undefined env vars",
			gherkinContents: []string{
				`Given I set "{{ env.same_var }}" as value`,
				`And I use "{{ env.same_var }}" again`,
			},
			configContent: `
settings:
  url: "{{ env.same_var }}"
`,
			envVars:           map[string]string{},
			shouldExit:        true,
			expectedUndefined: []string{"same_var"},
		},
		{
			name:            "empty config file path",
			gherkinContents: []string{},
			configContent:   "",
			envVars: map[string]string{
				"api_url": "http://localhost:3000",
			},
			shouldExit:        false,
			expectedUndefined: nil,
		},
		{
			name: "nested env vars",
			gherkinContents: []string{
				`Given I set "{{ env.database.host }}" as host`,
			},
			configContent: `
settings:
  db: "{{ env.database.port }}"
`,
			envVars: map[string]string{
				"database.host": "localhost",
				"database.port": "5432",
			},
			shouldExit:        false,
			expectedUndefined: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			variables.SetEnvVariables(tt.envVars)

			// Create parsed features
			var parsedFeatures []*gherkinparser.Feature
			for i, content := range tt.gherkinContents {
				parsedFeatures = append(parsedFeatures, &gherkinparser.Feature{
					Name:     "Test Feature " + string(rune(i)),
					Contents: []byte(content),
				})
			}

			configFilePath := createTempConfigFile(t, tt.configContent)
			undefinedEnvs := collectUndefinedVars(parsedFeatures, configFilePath)

			if tt.shouldExit {
				verifyUndefinedVars(t, undefinedEnvs, tt.expectedUndefined)
				return
			}
			if len(undefinedEnvs) > 0 {
				t.Errorf("Expected no undefined env vars but got: %v", undefinedEnvs)
			}
		})
	}
}

func TestValidateEnvVarsWithMissingConfigFile(t *testing.T) {
	variables.SetEnvVariables(map[string]string{})
	defer variables.ResetEnvVariables()

	nonExistentPath := filepath.Join(os.TempDir(), "non-existent-config-file.yml")
	undefinedEnvs := collectUndefinedVars([]*gherkinparser.Feature{}, nonExistentPath)

	if len(undefinedEnvs) != 0 {
		t.Errorf("Expected no undefined vars with missing config file, got: %v", undefinedEnvs)
	}
}

func TestValidateEnvVarsEmptyContent(t *testing.T) {
	variables.SetEnvVariables(map[string]string{})
	defer variables.ResetEnvVariables()

	parsedFeatures := []*gherkinparser.Feature{
		{Name: "Empty Feature", Contents: []byte("")},
	}

	configFilePath := createTempConfigFile(t, "")
	undefinedEnvs := collectUndefinedVars(parsedFeatures, configFilePath)

	if len(undefinedEnvs) != 0 {
		t.Errorf("Expected no undefined vars with empty content, got: %v", undefinedEnvs)
	}
}

func TestValidateEnvVarsComplexScenario(t *testing.T) {
	variables.SetEnvVariables(map[string]string{
		"frontend_base_url": "http://localhost:3000",
		"api_key":           "test-key",
	})
	defer variables.ResetEnvVariables()

	gherkinContent := `
Feature: Complex test
  Scenario: Test with env vars
    Given I navigate to "{{ env.frontend_base_url }}"
    And I use API key "{{ env.api_key }}"
    And I connect to "{{ env.database_url }}"
    When I send request to "{{ env.api_endpoint }}"
`

	configContent := `
settings:
  base_url: "{{ env.frontend_base_url }}"
  timeout: 1000

backend:
  endpoint: "{{ env.backend_service }}"
  auth: "{{ env.api_key }}"
`

	parsedFeatures := []*gherkinparser.Feature{
		{Name: "Complex Feature", Contents: []byte(gherkinContent)},
	}

	configFilePath := createTempConfigFile(t, configContent)
	undefinedEnvs := collectUndefinedVars(parsedFeatures, configFilePath)

	// Should find: database_url, api_endpoint, backend_service
	expectedUndefined := map[string]bool{
		"database_url":    true,
		"api_endpoint":    true,
		"backend_service": true,
	}

	if len(undefinedEnvs) != len(expectedUndefined) {
		t.Errorf("Expected %d undefined vars, got %d: %v", len(expectedUndefined), len(undefinedEnvs), undefinedEnvs)
	}

	for expectedVar := range expectedUndefined {
		if !undefinedEnvs[expectedVar] {
			t.Errorf("Expected undefined var %q not found", expectedVar)
		}
	}

	// Verify defined vars are not in undefined list
	definedVars := []string{"frontend_base_url", "api_key"}
	for _, definedVar := range definedVars {
		if undefinedEnvs[definedVar] {
			t.Errorf("Defined var %q should not be in undefined list", definedVar)
		}
	}
}
