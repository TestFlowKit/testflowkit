package helpers

import "strings"

func GetHeaderCaseInsensitive(headers map[string]string, name string) (string, bool) {
	// First try exact match
	if value, exists := headers[name]; exists {
		return value, true
	}

	// Fall back to case-insensitive search
	nameLower := strings.ToLower(name)
	for key, value := range headers {
		if strings.ToLower(key) == nameLower {
			return value, true
		}
	}

	return "", false
}
