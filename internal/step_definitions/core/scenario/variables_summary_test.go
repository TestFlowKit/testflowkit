package scenario

import (
	"testflowkit/internal/config"
	"testing"
)

func TestGenerateVariablesSummary(t *testing.T) {
	cfg := &config.Config{}
	vars := map[string]any{
		"user":     map[string]any{"id": 123, "name": "alice"},
		"password": "supersecret",
		"count":    42,
	}

	ctx := NewContext(cfg, vars, nil)

	summary := ctx.GenerateVariablesSummary()

	if summary == "(no variables)" {
		t.Fatalf("expected variables summary, got none")
	}

	// Expect redaction of sensitive key
	if !contains(summary, "password") || !contains(summary, "[REDACTED]") {
		t.Fatalf("expected password to be redacted in summary: %s", summary)
	}

	// Expect user object to be present
	if !contains(summary, "user") || !contains(summary, "id") || !contains(summary, "alice") {
		t.Fatalf("expected user JSON in summary: %s", summary)
	}

	if !contains(summary, "count") || !contains(summary, "42") {
		t.Fatalf("expected count in summary: %s", summary)
	}
}

func contains(s, sub string) bool {
	return stringsContains(s, sub)
}

// avoid importing strings in test helper to keep assertions explicit.
func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
