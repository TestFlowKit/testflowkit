package actioninit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	configPath := configFile
	_, statErr := os.Stat(configPath)
	assert.False(t, os.IsNotExist(statErr), "config.yml was not created")

	content, readErr := os.ReadFile(configPath)
	require.NoError(t, readErr, "Failed to read config.yml")

	configStr := string(content)

	expectedElements := []string{
		"frontend_base_url: \"https://testflowkit.dreamsfollowers.me\"",
		"gherkin_location: \"./features\"",
		"default_timeout: 10000",
		"pages:",
		"home: \"/\"",
		"get_started: \"/get-started\"",
		"sentences: \"/sentences\"",
	}

	for _, expected := range expectedElements {
		assert.Contains(t, configStr, expected)
	}
}

func verifyDirectoryStructure(t *testing.T) {
	t.Helper()
	_, statErr := os.Stat(featuresDir)
	assert.False(t, os.IsNotExist(statErr), "features directory was not created")

	info, statErr := os.Stat(featuresDir)
	require.NoError(t, statErr, "Failed to get directory info")

	assert.Equal(t, os.FileMode(0755), info.Mode().Perm(), "Expected directory permissions 0755")
}

func verifySampleFeatureFile(t *testing.T) {
	t.Helper()
	samplePath := filepath.Join(featuresDir, "sample.feature")
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

	configInfo, statErr := os.Stat(configFile)
	require.NoError(t, statErr, "Failed to get config file info")
	assert.Equal(t, os.FileMode(0600), configInfo.Mode().Perm(), "Expected config file permissions 0600")

	sampleInfo, statErr := os.Stat(filepath.Join(featuresDir, "sample.feature"))
	require.NoError(t, statErr, "Failed to get sample feature file info")
	assert.Equal(t, os.FileMode(0600), sampleInfo.Mode().Perm(), "Expected sample feature file permissions 0600")
}

func TestInitializationErrorScenarios(t *testing.T) {
	t.Run("ExistingConfigFileError", testExistingConfigFileError)
	t.Run("ExistingFeaturesDirectoryWarning", testExistingFeaturesDirectoryWarning)
	t.Run("ExistingSampleFeatureSkip", testExistingSampleFeatureSkip)
	t.Run("CleanupMechanism", testCleanupMechanism)
}

func testExistingConfigFileError(t *testing.T) {
	t.Helper()
	cleanup := setupTestDir(t)
	defer cleanup()

	writeErr := os.WriteFile(configFile, []byte("existing config"), 0600)
	require.NoError(t, writeErr, "Failed to create existing config file")

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createConfigFile(state)
	require.Error(t, err, "Expected error when config file already exists")
	assert.Contains(t, err.Error(), "config.yml already exists", "Expected specific error message")
}

func testExistingFeaturesDirectoryWarning(t *testing.T) {
	t.Helper()
	cleanup := setupTestDir(t)
	defer cleanup()

	mkdirErr := os.MkdirAll(featuresDir, 0755)
	require.NoError(t, mkdirErr, "Failed to create existing features directory")

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createDirectoryStructure(state)
	require.NoError(t, err, "Expected no error when features directory exists")
	assert.Empty(t, state.createdDirs, "Should not track existing directory")
}

func testExistingSampleFeatureSkip(t *testing.T) {
	t.Helper()
	cleanup := setupTestDir(t)
	defer cleanup()

	mkdirErr := os.MkdirAll(featuresDir, 0755)
	require.NoError(t, mkdirErr, "Failed to create features directory")

	samplePath := filepath.Join(featuresDir, "sample.feature")
	writeErr := os.WriteFile(samplePath, []byte("existing sample"), 0600)
	require.NoError(t, writeErr, "Failed to create existing sample file")

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createSampleFeature(state)
	require.NoError(t, err, "Expected no error when sample feature exists")
	assert.Empty(t, state.createdFiles, "Should not track existing file")
}

func testCleanupMechanism(t *testing.T) {
	t.Helper()
	cleanup := setupTestDir(t)
	defer cleanup()

	testDir := "test_cleanup_dir"
	testFile := "test_cleanup_file.txt"

	mkdirErr := os.MkdirAll(testDir, 0755)
	require.NoError(t, mkdirErr, "Failed to create test directory")

	writeErr := os.WriteFile(testFile, []byte("test content"), 0600)
	require.NoError(t, writeErr, "Failed to create test file")

	state := &InitializationState{
		createdFiles: []string{testFile},
		createdDirs:  []string{testDir},
	}

	_, statErr := os.Stat(testFile)
	assert.False(t, os.IsNotExist(statErr), "Test file should exist before cleanup")
	_, statErr = os.Stat(testDir)
	assert.False(t, os.IsNotExist(statErr), "Test directory should exist before cleanup")

	state.cleanup()

	_, statErr = os.Stat(testFile)
	assert.True(t, os.IsNotExist(statErr), "Test file should be removed after cleanup")
	_, statErr = os.Stat(testDir)
	assert.True(t, os.IsNotExist(statErr), "Test directory should be removed after cleanup")
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
	content, readErr := os.ReadFile(configFile)
	require.NoError(t, readErr, "Failed to read config.yml")

	configStr := string(content)

	requiredSections := []string{
		"active_environment:",
		"settings:",
		"environments:",
		"frontend:",
	}

	for _, section := range requiredSections {
		assert.Contains(t, configStr, section, "Config missing required section: %s", section)
	}

	assert.Contains(t, configStr, "concurrency: 1", "Config should have concurrency set to 1")
	assert.Contains(t, configStr, "report_format: \"html\"", "Config should have HTML report format")
	assert.Contains(t, configStr, "headless: false", "Config should have headless set to false for demo purposes")
}

func testSampleFeatureStructure(t *testing.T) {
	t.Helper()
	content, readErr := os.ReadFile(filepath.Join(featuresDir, "sample.feature"))
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
		configFile,
		featuresDir,
		filepath.Join(featuresDir, "sample.feature"),
	}

	for _, path := range expectedPaths {
		_, statErr := os.Stat(path)
		assert.False(t, os.IsNotExist(statErr), "Expected path does not exist: %s", path)
	}

	entries, readErr := os.ReadDir(".")
	require.NoError(t, readErr, "Failed to read directory")

	expectedFiles := map[string]bool{
		configFile:  true,
		featuresDir: true,
	}

	for _, entry := range entries {
		assert.True(t, expectedFiles[entry.Name()], "Unexpected file/directory created: %s", entry.Name())
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
	content, readErr := os.ReadFile(filepath.Join(featuresDir, "sample.feature"))
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

	configContent, readErr := os.ReadFile(configFile)
	require.NoError(t, readErr, "Failed to read config.yml")

	featureContent, readErr := os.ReadFile(filepath.Join(featuresDir, "sample.feature"))
	require.NoError(t, readErr, "Failed to read sample.feature")

	configStr := string(configContent)
	featureStr := string(featureContent)

	configPages := []string{"home", "get_started", "sentences", "configuration"}
	for _, page := range configPages {
		pagePattern := "\"" + page + "\" page"
		if strings.Contains(featureStr, pagePattern) {
			pageDefPattern := page + ":"
			assert.Contains(t, configStr, pageDefPattern, "Feature uses page '%s' but it's not defined in config", page)
		}
	}

	configElements := []string{"get started", "main content", "sentence filter", "code block", "footer"}
	for _, element := range configElements {
		elementPattern := "\"" + element + "\""
		if strings.Contains(featureStr, elementPattern) {
			elementKey := strings.ReplaceAll(element, " ", "_")
			errMsg := fmt.Sprintf("Feature uses element '%s' but it's not clearly defined in config", element)
			assert.Contains(t, configStr, elementKey, errMsg)
		}
	}
}

func testTestFlowKitSentenceCompatibility(t *testing.T) {
	t.Helper()
	content, readErr := os.ReadFile(filepath.Join(featuresDir, "sample.feature"))
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
				failureMsg := "Step uses 'I' instead of 'the user'"

				otherMsg := fmt.Sprintf("Step at line %d uses 'I' instead of 'the user': %s", i+1, line)
				assert.Fail(t, failureMsg, otherMsg)
			}
		}
	}
}
