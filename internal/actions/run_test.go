package actions

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testflowkit/internal/config"
	"testing"
	"time"
)

func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "simple step",
			expected: "simple_step",
		},
		{
			input:    "step with special chars!@#$%",
			expected: "step_with_special_chars",
		},
		{
			input:    "step with spaces and dots...",
			expected: "step_with_spaces_and_dots",
		},
		{
			input:    "step with multiple___underscores",
			expected: "step_with_multiple_underscores",
		},
		{
			input:    "_step with leading underscore",
			expected: "step_with_leading_underscore",
		},
		{
			input:    "step with trailing underscore_",
			expected: "step_with_trailing_underscore",
		},
	}

	for _, tt := range tests {
		result := sanitizeFilename(tt.input)
		if result != tt.expected {
			t.Errorf("sanitizeFilename(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestScreenshotOnFailureConfig(t *testing.T) {
	// Test with screenshot on failure enabled
	config1 := &config.Config{
		Settings: config.GlobalSettings{
			ScreenshotOnFailure: true,
		},
	}

	if !config1.IsScreenshotOnFailureEnabled() {
		t.Error("Expected screenshot on failure to be enabled")
	}

	// Test with screenshot on failure disabled
	config2 := &config.Config{
		Settings: config.GlobalSettings{
			ScreenshotOnFailure: false,
		},
	}

	if config2.IsScreenshotOnFailureEnabled() {
		t.Error("Expected screenshot on failure to be disabled")
	}
}

func TestScreenshotDirectoryCreation(t *testing.T) {
	testDir := "test_screenshots"

	defer os.RemoveAll(testDir)

	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	if _, existsErr := os.Stat(testDir); os.IsNotExist(existsErr) {
		t.Error("Screenshot directory was not created")
	}
}

func TestScreenshotPathGeneration(t *testing.T) {
	stepText := "the user clicks on the submit button"

	safeStepName := sanitizeFilename(stepText)
	generatedPath := filepath.Join(screenshotDir, safeStepName+".png")

	// Check that the path contains the expected base filename
	expectedBaseName := "the_user_clicks_on_the_submit_button"
	if !strings.Contains(generatedPath, expectedBaseName) {
		t.Errorf("Expected path to contain %s, got %s", expectedBaseName, generatedPath)
	}

	// Check that the path ends with .png
	if !strings.HasSuffix(generatedPath, ".png") {
		t.Errorf("Expected path to end with .png, got %s", generatedPath)
	}
}

func TestScreenshotFilenameWithTimestamp(t *testing.T) {
	stepText := "the user clicks on the submit button"

	// Generate filename with timestamp (simulating the actual implementation)
	safeStepName := sanitizeFilename(stepText)
	timestamp := time.Now().Format("20060102_150405_000")
	filename := safeStepName + "_" + timestamp + ".png"

	// Check that the filename contains the expected components
	expectedBaseName := "the_user_clicks_on_the_submit_button"
	if !strings.Contains(filename, expectedBaseName) {
		t.Errorf("Expected filename to contain %s, got %s", expectedBaseName, filename)
	}

	// Check that the filename contains a timestamp pattern (YYYYMMDD_HHMMSS_mmm)
	timestampPattern := regexp.MustCompile(`\d{8}_\d{6}_\d{3}`)
	if !timestampPattern.MatchString(filename) {
		t.Errorf("Expected filename to contain timestamp pattern, got %s", filename)
	}

	// Check that the filename ends with .png
	if !strings.HasSuffix(filename, ".png") {
		t.Errorf("Expected filename to end with .png, got %s", filename)
	}
}
