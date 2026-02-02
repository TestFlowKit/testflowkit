package actioninit

import (
	"os"
	"path/filepath"
	"strings"
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testConfigFile  = "config.yml"
	testFeaturesDir = "features"
)

func setupTestDir(t *testing.T) func() {
	t.Helper()
	tempDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err, "Failed to get current directory")

	err = os.Chdir(tempDir)
	require.NoError(t, err, "Failed to change to temp directory")

	cleanup := func() {
		_ = os.Chdir(originalDir)
	}

	return cleanup
}

func TestCompleteInitializationFlow(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	Execute(&config.Config{}, nil)

	t.Run("VerifyConfigFileCreation", verifyConfigFileCreation)
	t.Run("VerifyDirectoryStructure", verifyDirectoryStructure)
	t.Run("VerifySampleFeatureFile", verifySampleFeatureFile)
	t.Run("VerifyFilePermissions", verifyFilePermissions)
}

func verifyConfigFileCreation(t *testing.T) {
	t.Helper()
	_, statErr := os.Stat(testConfigFile)
	assert.False(t, os.IsNotExist(statErr), "config.yml was not created")

	content, readErr := os.ReadFile(testConfigFile)
	require.NoError(t, readErr, "Failed to read config.yml")

	configStr := string(content)

	expectedElements := []string{
		"base_url: \"https://testflowkit.github.io\"",
		"gherkin_location: \"./features\"",
		"base_url: \"{{ env.base_url }}\"",
		"default_timeout: 10000",
		"pages:",
		"home: \"/\"",
		"sentences: \"/sentences\"",
	}

	for _, expected := range expectedElements {
		assert.Contains(t, configStr, expected)
	}
}

func verifyDirectoryStructure(t *testing.T) {
	t.Helper()
	_, statErr := os.Stat(testFeaturesDir)
	assert.False(t, os.IsNotExist(statErr), "features directory was not created")

	info, statErr := os.Stat(testFeaturesDir)
	require.NoError(t, statErr, "Failed to get directory info")

	assert.Equal(t, os.FileMode(0755), info.Mode().Perm(), "Expected directory permissions 0755")
}

func verifySampleFeatureFile(t *testing.T) {
	t.Helper()
	samplePath := filepath.Join(testFeaturesDir, "sample.feature")
	_, statErr := os.Stat(samplePath)
	assert.False(t, os.IsNotExist(statErr), "sample.feature was not created")

	content, readErr := os.ReadFile(samplePath)
	require.NoError(t, readErr, "Failed to read sample.feature")

	featureStr := string(content)

	expectedElements := []string{
		"@SAMPLE",
		"Feature: TestFlowKit Documentation Site Sample Test",
		"the user opens a new browser tab",
		"the user goes to the \"home\" page",
		"the page title should be",
		"the user should see",
		"should be visible",
		"the current URL should contain",
	}

	for _, expected := range expectedElements {
		assert.Contains(t, featureStr, expected, "Sample feature file missing expected content: %s", expected)
	}

	scenarioCount := strings.Count(featureStr, "Scenario:")
	assert.GreaterOrEqual(t, scenarioCount, 3, "Expected at least 3 scenarios")
}

func verifyFilePermissions(t *testing.T) {
	t.Helper()

	configInfo, statErr := os.Stat(testConfigFile)
	require.NoError(t, statErr, "Failed to get config file info")
	assert.Equal(t, os.FileMode(0600), configInfo.Mode().Perm(), "Expected config file permissions 0600")

	sampleInfo, statErr := os.Stat(filepath.Join(testFeaturesDir, "sample.feature"))
	require.NoError(t, statErr, "Failed to get sample feature file info")
	assert.Equal(t, os.FileMode(0600), sampleInfo.Mode().Perm(), "Expected sample feature file permissions 0600")
}

func TestInitializationErrorScenarios(t *testing.T) {
	// Note: These tests verify that initialization properly detects and reports errors.
	// Since the Execute function calls logger.Fatal() which exits the process,
	// we can't easily test the actual error behavior in unit tests.
	// These tests are kept as documentation of expected behavior.

	t.Run("ExistingConfigFileError", func(t *testing.T) {
		t.Skip("Skipping: Execute calls os.Exit on error, which cannot be tested in unit tests")
	})
	t.Run("ExistingFeaturesDirectoryWarning", testExistingFeaturesDirectoryWarning)
	t.Run("ExistingSampleFeatureSkip", testExistingSampleFeatureSkip)
}

func testExistingFeaturesDirectoryWarning(t *testing.T) {
	t.Helper()
	cleanup := setupTestDir(t)
	defer cleanup()

	mkdirErr := os.MkdirAll(testFeaturesDir, 0755)
	require.NoError(t, mkdirErr, "Failed to create existing features directory")

	// Execute should handle existing directory gracefully
	Execute(&config.Config{}, nil)

	// Verify directory still exists
	_, statErr := os.Stat(testFeaturesDir)
	assert.NoError(t, statErr, "Features directory should still exist")
}

func testExistingSampleFeatureSkip(t *testing.T) {
	t.Helper()
	cleanup := setupTestDir(t)
	defer cleanup()

	mkdirErr := os.MkdirAll(testFeaturesDir, 0755)
	require.NoError(t, mkdirErr, "Failed to create features directory")

	samplePath := filepath.Join(testFeaturesDir, "sample.feature")
	writeErr := os.WriteFile(samplePath, []byte("existing sample"), 0600)
	require.NoError(t, writeErr, "Failed to create existing sample file")

	Execute(&config.Config{}, nil)

	// Verify the original sample wasn't overwritten
	content, _ := os.ReadFile(samplePath)
	assert.Equal(t, "existing sample", string(content), "Existing sample should not be overwritten")
}

func TestGeneratedFilesStructure(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	Execute(&config.Config{}, nil)

	t.Run("ConfigFileStructure", testConfigFileStructure)
	t.Run("SampleFeatureStructure", testSampleFeatureStructure)
	t.Run("DirectoryStructure", testDirectoryStructure)
}

func testConfigFileStructure(t *testing.T) {
	t.Helper()
	content, readErr := os.ReadFile(testConfigFile)
	require.NoError(t, readErr, "Failed to read config.yml")

	configStr := string(content)

	requiredSections := []string{
		"settings:",
		"env:",
		"frontend:",
	}

	for _, section := range requiredSections {
		assert.Contains(t, configStr, section, "Config missing required section: %s", section)
	}

	assert.Contains(t, configStr, "concurrency: 1", "Config should have concurrency set to 1")
	assert.Contains(t, configStr, "report_format: \"html\"", "Config should have HTML report format")
	assert.Contains(t, configStr, "headless: false", "Config should have headless set to false for demo purposes")
	assert.Contains(t, configStr, "base_url:", "Config should have base_url in env section")
	assert.Contains(t, configStr, "base_url: \"{{ env.base_url }}\"", "Config should have base_url in frontend section")
}

func testSampleFeatureStructure(t *testing.T) {
	t.Helper()
	content, readErr := os.ReadFile(filepath.Join(testFeaturesDir, "sample.feature"))
	require.NoError(t, readErr, "Failed to read sample.feature")

	featureStr := string(content)

	assert.Contains(t, featureStr, "Feature:", "Sample feature should contain Feature declaration")
	assert.Contains(t, featureStr, "Background:", "Sample feature should contain Background section")

	scenarioCount := strings.Count(featureStr, "Scenario:")
	assert.GreaterOrEqual(t, scenarioCount, 3, "Expected at least 3 scenarios for comprehensive examples")

	sentenceTypes := []string{
		"the user goes to",
		"should be visible",
		"should contain",
		"the user clicks",
		"the user enters",
	}

	for _, sentenceType := range sentenceTypes {
		assert.Contains(t, featureStr, sentenceType, "Sample feature should demonstrate %s sentence pattern", sentenceType)
	}
}

func testDirectoryStructure(t *testing.T) {
	t.Helper()

	expectedPaths := []string{
		testConfigFile,
		testFeaturesDir,
		filepath.Join(testFeaturesDir, "sample.feature"),
	}

	for _, path := range expectedPaths {
		_, statErr := os.Stat(path)
		assert.False(t, os.IsNotExist(statErr), "Expected path does not exist: %s", path)
	}
}

func TestSampleFeatureRunnability(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	Execute(&config.Config{}, nil)

	t.Run("FeatureFileValidGherkin", testFeatureFileValidGherkin)
	t.Run("ConfigurationCompatibility", testConfigurationCompatibility)
	t.Run("TestFlowKitSentenceCompatibility", testTestFlowKitSentenceCompatibility)
}

func testFeatureFileValidGherkin(t *testing.T) {
	t.Helper()
	content, readErr := os.ReadFile(filepath.Join(testFeaturesDir, "sample.feature"))
	require.NoError(t, readErr, "Failed to read sample.feature")

	featureStr := string(content)

	gherkinKeywords := []string{
		"Feature:",
		"Background:",
		"Scenario:",
		"Given ",
		"When ",
		"Then ",
		"And ",
	}

	for _, keyword := range gherkinKeywords {
		assert.Contains(t, featureStr, keyword, "Sample feature should contain Gherkin keyword: %s", keyword)
	}

	lines := strings.Split(featureStr, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Given ") || strings.HasPrefix(line, "When ") ||
			strings.HasPrefix(line, "Then ") || strings.HasPrefix(line, "And ") {
			assert.NotEmpty(t, strings.TrimSpace(line[5:]), "Empty step definition at line %d: %s", i+1, line)
		}
	}
}

func testConfigurationCompatibility(t *testing.T) {
	t.Helper()

	configContent, readErr := os.ReadFile(testConfigFile)
	require.NoError(t, readErr, "Failed to read config.yml")

	featureContent, readErr := os.ReadFile(filepath.Join(testFeaturesDir, "sample.feature"))
	require.NoError(t, readErr, "Failed to read sample.feature")

	configStr := string(configContent)
	featureStr := string(featureContent)

	configPages := []string{"home", "sentences"}
	for _, page := range configPages {
		pagePattern := "\"" + page + "\" page"
		if strings.Contains(featureStr, pagePattern) {
			pageDefPattern := page + ":"
			assert.Contains(t, configStr, pageDefPattern, "Feature uses page '%s' but it's not defined in config", page)
		}
	}
}

func testTestFlowKitSentenceCompatibility(t *testing.T) {
	t.Helper()
	content, readErr := os.ReadFile(filepath.Join(testFeaturesDir, "sample.feature"))
	require.NoError(t, readErr, "Failed to read sample.feature")

	featureStr := string(content)

	validSentencePatterns := []string{
		"the user opens a new browser tab",
		"the user goes to the",
		"the page title should be",
		"the user should see",
		"should be visible",
		"the current URL should contain",
		"the user clicks the",
		"the user enters",
		"into the",
		"field",
	}

	for _, pattern := range validSentencePatterns {
		assert.Contains(t, featureStr, pattern, "Sample feature should use TestFlowKit sentence pattern: %s", pattern)
	}

	lines := strings.Split(featureStr, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Given ") || strings.HasPrefix(line, "When ") ||
			strings.HasPrefix(line, "Then ") || strings.HasPrefix(line, "And ") {
			stepText := strings.TrimSpace(line[5:])

			if strings.Contains(stepText, "I ") && !strings.Contains(stepText, "I store") {
				assert.Fail(t,
					"Step uses 'I' instead of 'the user'",
					"Step at line %d uses 'I' instead of 'the user': %s", i+1, line)
			}
		}
	}
}
