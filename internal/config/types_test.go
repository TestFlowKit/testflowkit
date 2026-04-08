package config

import (
	"testing"
	"time"
)

var parseSelectorStringtestCases = []struct {
	name     string
	input    string
	expected Selector
}{
	{
		name:  "CSS selector without prefix",
		input: "button[type='submit']",
		expected: Selector{
			Type:  SelectorTypeCSS,
			Value: "button[type='submit']",
		},
	},
	{
		name:  "CSS selector with class",
		input: ".spinner",
		expected: Selector{
			Type:  SelectorTypeCSS,
			Value: ".spinner",
		},
	},
	{
		name:  "XPath selector with prefix",
		input: "xpath://div[contains(@class, 'loading')]",
		expected: Selector{
			Type:  SelectorTypeXPath,
			Value: "//div[contains(@class, 'loading')]",
		},
	},
	{
		name:  "XPath selector with button",
		input: "xpath://button[@type='submit']",
		expected: Selector{
			Type:  SelectorTypeXPath,
			Value: "//button[@type='submit']",
		},
	},
	{
		name:  "CSS selector with whitespace",
		input: "  button[type='submit']  ",
		expected: Selector{
			Type:  SelectorTypeCSS,
			Value: "button[type='submit']",
		},
	},
	{
		name:  "XPath selector with whitespace",
		input: "  xpath://div[contains(@class, 'loading')]  ",
		expected: Selector{
			Type:  SelectorTypeXPath,
			Value: "//div[contains(@class, 'loading')]",
		},
	},
}

func TestParseSelectorString(t *testing.T) {
	for _, tt := range parseSelectorStringtestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := NewSelector(tt.input)
			if result.Type != tt.expected.Type {
				t.Errorf("ParseSelectorString() Type = %v, want %v", result.Type, tt.expected.Type)
			}
			if result.Value != tt.expected.Value {
				t.Errorf("ParseSelectorString() Value = %v, want %v", result.Value, tt.expected.Value)
			}
		})
	}
}

func TestSelector_IsXPath(t *testing.T) {
	tests := []struct {
		name     string
		selector Selector
		expected bool
	}{
		{
			name: "XPath selector",
			selector: Selector{
				Type:  SelectorTypeXPath,
				Value: "//div[@class='test']",
			},
			expected: true,
		},
		{
			name: "CSS selector",
			selector: Selector{
				Type:  SelectorTypeCSS,
				Value: ".test",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.selector.IsXPath()
			if result != tt.expected {
				t.Errorf("Selector.IsXPath() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSelector_IsCSS(t *testing.T) {
	tests := []struct {
		name     string
		selector Selector
		expected bool
	}{
		{
			name: "CSS selector",
			selector: Selector{
				Type:  SelectorTypeCSS,
				Value: ".test",
			},
			expected: true,
		},
		{
			name: "XPath selector",
			selector: Selector{
				Type:  SelectorTypeXPath,
				Value: "//div[@class='test']",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.selector.IsCSS()
			if result != tt.expected {
				t.Errorf("Selector.IsCSS() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSelector_String(t *testing.T) {
	selector := Selector{
		Type:  SelectorTypeCSS,
		Value: "button[type='submit']",
	}

	expected := "button[type='submit']"
	result := selector.String()

	if result != expected {
		t.Errorf("Selector.String() = %v, want %v", result, expected)
	}
}

func TestConfig_GetTimeout(t *testing.T) {
	tests := []struct {
		name           string
		defaultTimeout int
		expectedMs     int64
	}{
		{
			name:           "10 seconds timeout",
			defaultTimeout: 10000,
			expectedMs:     10000,
		},
		{
			name:           "30 seconds timeout",
			defaultTimeout: 30000,
			expectedMs:     30000,
		},
		{
			name:           "1 minute timeout",
			defaultTimeout: 60000,
			expectedMs:     60000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				Frontend: &FrontendConfig{
					DefaultTimeout: tt.defaultTimeout,
				},
			}

			result := config.GetTimeout()
			expected := time.Duration(tt.expectedMs) * time.Millisecond

			if result != expected {
				t.Errorf("Config.GetTimeout() = %v, want %v", result, expected)
			}
		})
	}
}

func TestConfigGetAPIRequestTimeout(t *testing.T) {
	const apiBaseURL = "https://api.example.com"
	intPtr := func(v int) *int { return &v }

	tests := []struct {
		name        string
		config      *Config
		apiName     string
		requestName string
		expectedMs  int
	}{
		{
			name: "REST endpoint timeout overrides API and default timeouts",
			config: &Config{
				APIs: &APIsConfig{
					DefaultTimeout: 30000,
					Definitions: map[string]APIDefinition{
						"users": {
							Type:    APITypeREST,
							BaseURL: "https://api.example.com",
							Timeout: intPtr(20000),
							Endpoints: map[string]Endpoint{
								"get_user": {
									Method:      "GET",
									Path:        "/users/1",
									Description: "Get user",
									Timeout:     intPtr(5000),
								},
							},
						},
					},
				},
			},
			apiName:     "users",
			requestName: "get_user",
			expectedMs:  5000,
		},
		{
			name: "GraphQL operation timeout overrides API and default timeouts",
			config: &Config{
				APIs: &APIsConfig{
					DefaultTimeout: 30000,
					Definitions: map[string]APIDefinition{
						"graphql_api": {
							Type:     APITypeGraphQL,
							Endpoint: "https://api.example.com/graphql",
							Timeout:  intPtr(15000),
							Operations: map[string]GraphQLOperation{
								"get_user": {
									Type:        "query",
									Operation:   "query { user { id } }",
									Description: "Get user",
									Timeout:     intPtr(7000),
								},
							},
						},
					},
				},
			},
			apiName:     "graphql_api",
			requestName: "get_user",
			expectedMs:  7000,
		},
		{
			name: "API timeout overrides default timeout when request timeout is absent",
			config: &Config{
				APIs: &APIsConfig{
					DefaultTimeout: 30000,
					Definitions: map[string]APIDefinition{
						"users": {
							Type:    APITypeREST,
							BaseURL: apiBaseURL,
							Timeout: intPtr(12000),
							Endpoints: map[string]Endpoint{
								"list_users": {
									Method:      "GET",
									Path:        "/users",
									Description: "List users",
								},
							},
						},
					},
				},
			},
			apiName:     "users",
			requestName: "list_users",
			expectedMs:  12000,
		},
		{
			name: "Default timeout is used when request and API timeouts are absent",
			config: &Config{
				APIs: &APIsConfig{
					DefaultTimeout: 25000,
					Definitions: map[string]APIDefinition{
						"users": {
							Type:    APITypeREST,
							BaseURL: apiBaseURL,
							Endpoints: map[string]Endpoint{
								"list_users": {
									Method:      "GET",
									Path:        "/users",
									Description: "List users",
								},
							},
						},
					},
				},
			},
			apiName:     "users",
			requestName: "list_users",
			expectedMs:  25000,
		},
		{
			name:        "Built-in fallback is used when APIs config is missing",
			config:      &Config{},
			apiName:     "missing",
			requestName: "missing_request",
			expectedMs:  30000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.GetAPIRequestTimeout(tt.apiName, tt.requestName)
			if result != tt.expectedMs {
				t.Errorf("Config.GetAPIRequestTimeout() = %v, want %v", result, tt.expectedMs)
			}
		})
	}
}
