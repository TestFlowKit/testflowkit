package actionrun

import (
	"os"
	"testflowkit/internal/config"
	"testing"
)

const frontendClickStep = "the user clicks the \"submit_button\" button"
const backendRequestStep = "I send the request"

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

func TestIsFrontendStepTextMatchWhenFrontendStepExists(t *testing.T) {
	stepText := frontendClickStep

	matched := IsFrontendStepTextMatch(stepText)
	if !matched {
		t.Fatal("expected frontend step detection to match")
	}
}

func TestIsFrontendStepTextMatchWhenNoFrontendStepExists(t *testing.T) {
	stepText := backendRequestStep

	matched := IsFrontendStepTextMatch(stepText)
	if matched {
		t.Fatal("expected frontend step detection to not match")
	}
}

func TestIsFrontendStepTextMatchIsCaseInsensitive(t *testing.T) {
	stepText := "THE USER CLICKS THE \"submit_button\" BUTTON"

	matched := IsFrontendStepTextMatch(stepText)
	if !matched {
		t.Fatal("expected case-insensitive frontend step detection to match")
	}
}

func TestShouldPreinitializeBrowserEngineWhenConfigIsNil(t *testing.T) {
	shouldInit := shouldPreinitializeBrowserEngine(nil, frontendClickStep)
	if shouldInit {
		t.Fatal("expected pre-initialization to be skipped when config is nil")
	}
}

func TestShouldPreinitializeBrowserEngineWhenFrontendNotDefined(t *testing.T) {
	cfg := &config.Config{}

	shouldInit := shouldPreinitializeBrowserEngine(cfg, frontendClickStep)
	if shouldInit {
		t.Fatal("expected pre-initialization to be skipped when frontend is not defined")
	}
}

func TestShouldPreinitializeBrowserEngineWhenStepIsNotFrontend(t *testing.T) {
	cfg := &config.Config{
		Frontend: &config.FrontendConfig{},
	}

	shouldInit := shouldPreinitializeBrowserEngine(cfg, backendRequestStep)
	if shouldInit {
		t.Fatal("expected pre-initialization to be skipped for non-frontend step")
	}
}

func TestShouldPreinitializeBrowserEngineWhenFrontendStepMatches(t *testing.T) {
	cfg := &config.Config{
		Frontend: &config.FrontendConfig{},
	}

	shouldInit := shouldPreinitializeBrowserEngine(cfg, frontendClickStep)
	if !shouldInit {
		t.Fatal("expected pre-initialization for frontend step")
	}
}
