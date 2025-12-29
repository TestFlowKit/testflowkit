package actionrun

import (
	"os"
	"testflowkit/internal/config"
	"testing"
)

func TestScreenshotOnFailureConfig(t *testing.T) {
	// Test with screenshot on failure enabled
	config1 := &config.Config{
		Frontend: &config.FrontendConfig{
			ScreenshotOnFailure: true,
		},
	}

	if !config1.IsScreenshotOnFailureEnabled() {
		t.Error("Expected screenshot on failure to be enabled")
	}

	// Test with screenshot on failure disabled
	config2 := &config.Config{
		Frontend: &config.FrontendConfig{
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
