package config

import (
	"testing"
	"time"
)

func TestParseSelectorString(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
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
				Frontend: FrontendConfig{
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
