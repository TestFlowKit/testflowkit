package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlattenMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		prefix   string
		expected map[string]string
	}{
		{
			name: "simple flat map",
			input: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			prefix: "",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "nested map",
			input: map[string]any{
				"api": map[string]any{
					"url": "http://localhost",
					"key": "secret",
				},
			},
			prefix: "",
			expected: map[string]string{
				"api.url": "http://localhost",
				"api.key": "secret",
			},
		},
		{
			name: "deeply nested map",
			input: map[string]any{
				"database": map[string]any{
					"primary": map[string]any{
						"host": "localhost",
						"port": 5432,
					},
					"replica": map[string]any{
						"host": "replica.example.com",
						"port": 5433,
					},
				},
			},
			prefix: "",
			expected: map[string]string{
				"database.primary.host": "localhost",
				"database.primary.port": "5432",
				"database.replica.host": "replica.example.com",
				"database.replica.port": "5433",
			},
		},
		{
			name: "with prefix",
			input: map[string]any{
				"host": "localhost",
				"port": 8080,
			},
			prefix: "server",
			expected: map[string]string{
				"server.host": "localhost",
				"server.port": "8080",
			},
		},
		{
			name: "array values",
			input: map[string]any{
				"tags": []any{"tag1", "tag2", "tag3"},
			},
			prefix: "",
			expected: map[string]string{
				"tags": "[tag1 tag2 tag3]",
			},
		},
		{
			name: "empty array",
			input: map[string]any{
				"empty": []any{},
			},
			prefix: "",
			expected: map[string]string{
				"empty": "",
			},
		},
		{
			name: "nil value",
			input: map[string]any{
				"nullable": nil,
			},
			prefix: "",
			expected: map[string]string{
				"nullable": "",
			},
		},
		{
			name: "map[any]any type",
			input: map[string]any{
				"config": map[any]any{
					"enabled": true,
					123:       "numeric-key",
				},
			},
			prefix: "",
			expected: map[string]string{
				"config.enabled": "true",
				"config.123":     "numeric-key",
			},
		},
		{
			name: "mixed types",
			input: map[string]any{
				"string":  "text",
				"number":  42,
				"float":   3.14,
				"boolean": true,
				"array":   []any{1, 2, 3},
				"nested": map[string]any{
					"key": "value",
				},
			},
			prefix: "",
			expected: map[string]string{
				"string":     "text",
				"number":     "42",
				"float":      "3.14",
				"boolean":    "true",
				"array":      "[1 2 3]",
				"nested.key": "value",
			},
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			prefix:   "",
			expected: map[string]string{},
		},
		{
			name: "complex nested structure",
			input: map[string]any{
				"frontend": map[string]any{
					"url": "http://localhost:3000",
					"auth": map[string]any{
						"enabled": true,
						"methods": []any{"oauth", "jwt"},
					},
				},
				"backend": map[string]any{
					"api": map[string]any{
						"url":     "http://localhost:8080",
						"timeout": "30s",
					},
					"database": map[string]any{
						"host": "localhost",
						"port": 5432,
					},
				},
			},
			prefix: "",
			expected: map[string]string{
				"frontend.url":          "http://localhost:3000",
				"frontend.auth.enabled": "true",
				"frontend.auth.methods": "[oauth jwt]",
				"backend.api.url":       "http://localhost:8080",
				"backend.api.timeout":   "30s",
				"backend.database.host": "localhost",
				"backend.database.port": "5432",
			},
		},
		{
			name: "array of maps",
			input: map[string]any{
				"servers": []any{
					map[string]any{"host": "server1", "port": 8080},
					map[string]any{"host": "server2", "port": 8081},
				},
			},
			prefix: "",
			expected: map[string]string{
				"servers": "[map[host:server1 port:8080] map[host:server2 port:8081]]",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FlattenMap(tt.input, tt.prefix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLoadEnvFile(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    map[string]string
		expectError bool
	}{
		{
			name: "simple env file",
			fileContent: `
api_key: "secret123"
api_url: "http://localhost:8080"
`,
			expected: map[string]string{
				"api_key": "secret123",
				"api_url": "http://localhost:8080",
			},
			expectError: false,
		},
		{
			name: "nested env file",
			fileContent: `
database:
  host: localhost
  port: 5432
  credentials:
    username: admin
    password: secret
`,
			expected: map[string]string{
				"database.host":                 "localhost",
				"database.port":                 "5432",
				"database.credentials.username": "admin",
				"database.credentials.password": "secret",
			},
			expectError: false,
		},
		{
			name: "env file with arrays",
			fileContent: `
allowed_origins:
  - http://localhost:3000
  - http://localhost:3001
tags:
  - development
  - testing
`,
			expected: map[string]string{
				"allowed_origins": "[http://localhost:3000 http://localhost:3001]",
				"tags":            "[development testing]",
			},
			expectError: false,
		},
		{
			name: "mixed types env file",
			fileContent: `
frontend_base_url: "http://localhost:3000"
api:
  base_url: "http://localhost:8080"
  timeout: 30
  retry: true
features:
  - authentication
  - authorization
`,
			expected: map[string]string{
				"frontend_base_url": "http://localhost:3000",
				"api.base_url":      "http://localhost:8080",
				"api.timeout":       "30",
				"api.retry":         "true",
				"features":          "[authentication authorization]",
			},
			expectError: false,
		},
		{
			name:        "invalid yaml",
			fileContent: `invalid: yaml: content:`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "empty file",
			fileContent: ``,
			expected:    map[string]string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "test.env.yml")
			err := os.WriteFile(tmpFile, []byte(tt.fileContent), 0644)
			require.NoError(t, err)

			// Test LoadEnvFile
			result, err := LoadEnvFile(tmpFile)

			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestLoadEnvFile_FileNotFound(t *testing.T) {
	result, err := LoadEnvFile("/non/existent/file.yml")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to read env file")
}

func TestLoadEnvFile_RealWorldExample(t *testing.T) {
	fileContent := `
# Production environment configuration
frontend_base_url: "https://example.com"
rest_api_base_url: "https://api.example.com"
graphql_endpoint: "https://api.example.com/graphql"

# API Configuration
api:
  key: "prod-api-key-12345"
  timeout: "60s"
  rate_limit: 1000
  features:
    - caching
    - compression

# Database Configuration  
database:
  primary:
    host: "db-primary.example.com"
    port: 5432
    name: "production_db"
  replica:
    host: "db-replica.example.com"
    port: 5432
    name: "production_db"

# Feature Flags
features:
  new_ui: true
  beta_api: false
  analytics: true
`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "prod.env.yml")
	err := os.WriteFile(tmpFile, []byte(fileContent), 0644)
	require.NoError(t, err)

	result, err := LoadEnvFile(tmpFile)
	require.NoError(t, err)

	// Verify some key values
	assert.Equal(t, "https://example.com", result["frontend_base_url"])
	assert.Equal(t, "https://api.example.com", result["rest_api_base_url"])
	assert.Equal(t, "prod-api-key-12345", result["api.key"])
	assert.Equal(t, "60s", result["api.timeout"])
	assert.Equal(t, "1000", result["api.rate_limit"])
	assert.Equal(t, "[caching compression]", result["api.features"])
	assert.Equal(t, "db-primary.example.com", result["database.primary.host"])
	assert.Equal(t, "5432", result["database.primary.port"])
	assert.Equal(t, "true", result["features.new_ui"])
	assert.Equal(t, "false", result["features.beta_api"])
}
