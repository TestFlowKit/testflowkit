package formatter_test

import (
	"strings"
	"testing"

	"testflowkit/pkg/formatter"
)

// ────────────────────────────────────────────────────────────────────────────
// FormatJSON
// ────────────────────────────────────────────────────────────────────────────

func TestFormatJSON_ValidObject(t *testing.T) {
	raw := []byte(`{"a":1,"b":"hello"}`)
	got, err := formatter.FormatJSON(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "{\n  \"a\": 1,\n  \"b\": \"hello\"\n}"
	if string(got) != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatJSON_ValidArray(t *testing.T) {
	raw := []byte(`[1,2,3]`)
	got, err := formatter.FormatJSON(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(got), "\n") {
		t.Error("expected newlines in formatted JSON array")
	}
}

func TestFormatJSON_InvalidJSON(t *testing.T) {
	raw := []byte(`not json at all`)
	got, err := formatter.FormatJSON(raw)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	// Must return original bytes on failure.
	if string(got) != string(raw) {
		t.Errorf("expected raw bytes back, got: %s", got)
	}
}

// ────────────────────────────────────────────────────────────────────────────
// FormatXML
// ────────────────────────────────────────────────────────────────────────────

func TestFormatXML_Valid(t *testing.T) {
	raw := []byte(`<root><child attr="v">text</child></root>`)
	got, err := formatter.FormatXML(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := string(got)
	if !strings.Contains(out, "\n") {
		t.Error("expected newlines in formatted XML")
	}
	if !strings.Contains(out, "child") {
		t.Error("expected element content preserved")
	}
}

func TestFormatXML_InvalidXML(t *testing.T) {
	raw := []byte(`<unclosed>`)
	got, err := formatter.FormatXML(raw)
	if err == nil {
		t.Fatal("expected error for invalid XML")
	}
	if string(got) != string(raw) {
		t.Errorf("expected raw bytes back, got: %s", got)
	}
}

// ────────────────────────────────────────────────────────────────────────────
// IsBinary
// ────────────────────────────────────────────────────────────────────────────

func TestIsBinary(t *testing.T) {
	cases := []struct {
		ct   string
		want bool
	}{
		{"image/png", true},
		{"image/jpeg", true},
		{"audio/mpeg", true},
		{"video/mp4", true},
		{"application/octet-stream", true},
		{"application/pdf", true},
		{"application/zip", true},
		{"application/json", false},
		{"text/xml", false},
		{"text/plain", false},
		{"", false},
	}
	for _, tc := range cases {
		if got := formatter.IsBinary(tc.ct); got != tc.want {
			t.Errorf("IsBinary(%q) = %v, want %v", tc.ct, got, tc.want)
		}
	}
}

// ────────────────────────────────────────────────────────────────────────────
// NeedsFormatting
// ────────────────────────────────────────────────────────────────────────────

func TestNeedsFormatting(t *testing.T) {
	cases := []struct {
		ct   string
		want bool
	}{
		{"application/json", true},
		{"application/json; charset=utf-8", true},
		{"text/json", true},
		{"application/xml", true},
		{"text/xml", true},
		{"text/plain", false},
		{"application/octet-stream", false},
		{"", false},
	}
	for _, tc := range cases {
		if got := formatter.NeedsFormatting(tc.ct); got != tc.want {
			t.Errorf("NeedsFormatting(%q) = %v, want %v", tc.ct, got, tc.want)
		}
	}
}

// ────────────────────────────────────────────────────────────────────────────
// Format (integration)
// ────────────────────────────────────────────────────────────────────────────

func TestFormat_EmptyBody(t *testing.T) {
	got := formatter.Format("application/json", []byte{}, 0)
	if got != "(empty body)" {
		t.Errorf("got %q", got)
	}
}

func TestFormat_OversizedBody(t *testing.T) {
	big := make([]byte, 10)
	got := formatter.Format("application/json", big, 5)
	if !strings.Contains(got, "exceeds limit") {
		t.Errorf("expected truncation notice, got: %s", got)
	}
}

func TestFormat_BinaryContent(t *testing.T) {
	got := formatter.Format("image/png", []byte{0x89, 0x50, 0x4e, 0x47}, 0)
	if !strings.Contains(got, "binary content") {
		t.Errorf("expected binary placeholder, got: %s", got)
	}
}

func TestFormat_ValidJSON(t *testing.T) {
	raw := []byte(`{"x":1}`)
	got := formatter.Format("application/json", raw, 0)
	if !strings.Contains(got, "\n") {
		t.Errorf("expected indented JSON, got: %s", got)
	}
}

func TestFormat_ValidJSONWithMIMEParams(t *testing.T) {
	raw := []byte(`{"x":1}`)
	got := formatter.Format("application/json; charset=utf-8", raw, 0)
	if !strings.Contains(got, "\n") {
		t.Errorf("expected indented JSON, got: %s", got)
	}
}

func TestFormat_InvalidJSON_FallsBackToRaw(t *testing.T) {
	raw := []byte(`not-json`)
	got := formatter.Format("application/json", raw, 0)
	if got != "not-json" {
		t.Errorf("expected raw fallback, got: %s", got)
	}
}

func TestFormat_ValidXML(t *testing.T) {
	raw := []byte(`<r><a>1</a></r>`)
	got := formatter.Format("application/xml", raw, 0)
	if !strings.Contains(got, "\n") {
		t.Errorf("expected indented XML, got: %s", got)
	}
}

func TestFormat_InvalidXML_FallsBackToRaw(t *testing.T) {
	raw := []byte(`<unclosed>`)
	got := formatter.Format("text/xml", raw, 0)
	if got != "<unclosed>" {
		t.Errorf("expected raw fallback, got: %s", got)
	}
}

func TestFormat_PlainText(t *testing.T) {
	raw := []byte(`hello world`)
	got := formatter.Format("text/plain", raw, 0)
	if got != "hello world" {
		t.Errorf("got: %s", got)
	}
}

func TestFormat_NoMaxSize_LargeBody(t *testing.T) {
	// maxSize = 0 means no limit.
	large := []byte(`{"key":"` + strings.Repeat("v", 2_000_000) + `"}`)
	got := formatter.Format("application/json", large, 0)
	// Should still attempt formatting; starts with indented brace.
	if !strings.HasPrefix(got, "{") {
		t.Errorf("unexpected output start: %s", got[:50])
	}
}
