package actioninit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDirectoryStructure(t *testing.T) {
	defer os.RemoveAll(featuresDir)
	os.RemoveAll(featuresDir)

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createDirectoryStructure(state)
	require.NoError(t, err, "Expected no error")

	_, statErr := os.Stat(featuresDir)
	assert.False(t, os.IsNotExist(statErr), "Features directory was not created")

	assert.Len(t, state.createdDirs, 1, "Expected state to track one created directory")
	assert.Equal(t, featuresDir, state.createdDirs[0], "Expected state to track the features directory")

	state2 := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}
	err = createDirectoryStructure(state2)
	require.NoError(t, err, "Expected no error when directory exists")

	assert.Empty(t, state2.createdDirs, "Expected state not to track existing directory")
}

func TestCreateDirectoryStructurePathValidation(t *testing.T) {
	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createDirectoryStructure(state)
	require.NoError(t, err, "Expected no error for valid path")

	os.RemoveAll(featuresDir)
}

func TestCreateSampleFeature(t *testing.T) {
	defer os.RemoveAll(featuresDir)
	os.RemoveAll(featuresDir)

	err := os.MkdirAll(featuresDir, 0755)
	require.NoError(t, err, "Failed to create features directory")

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err = createSampleFeature(state)
	require.NoError(t, err, "Expected no error")

	sampleFeaturePath := filepath.Join(featuresDir, "sample.feature")
	_, statErr := os.Stat(sampleFeaturePath)
	assert.False(t, os.IsNotExist(statErr), "Sample feature file was not created")

	assert.Len(t, state.createdFiles, 1, "Expected state to track one created file")
	assert.Equal(t, sampleFeaturePath, state.createdFiles[0], "Expected state to track the sample feature file")

	content, readErr := os.ReadFile(sampleFeaturePath)
	require.NoError(t, readErr, "Failed to read sample feature file")

	contentStr := string(content)
	assert.Contains(t, contentStr, "@SAMPLE", "Sample feature file should contain @SAMPLE tag")
	assert.Contains(t, contentStr, "Feature: TestFlowKit Documentation Site Sample Test")
	assert.Contains(t, contentStr, "the user opens a new browser tab")

	state2 := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}
	err = createSampleFeature(state2)
	require.NoError(t, err, "Expected no error when file exists")

	assert.Empty(t, state2.createdFiles, "Expected state not to track existing file")
}

func TestCreateSampleFeatureWithoutFeaturesDirectory(t *testing.T) {
	defer os.RemoveAll(featuresDir)
	os.RemoveAll(featuresDir)

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createSampleFeature(state)
	require.Error(t, err, "Expected error when features directory doesn't exist")
}

func TestInitializationStateCleanup(t *testing.T) {
	defer os.RemoveAll("test_cleanup")
	os.RemoveAll("test_cleanup")

	testDir := "test_cleanup"
	testFile := "test_cleanup/test.txt"

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

func TestCreateConfigFile(t *testing.T) {
	defer os.Remove(configFile)
	os.Remove(configFile)

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	err := createConfigFile(state)
	require.NoError(t, err, "Expected no error")

	_, statErr := os.Stat(configFile)
	assert.False(t, os.IsNotExist(statErr), "Config file was not created")

	assert.Len(t, state.createdFiles, 1, "Expected state to track one created file")
	assert.Equal(t, configFile, state.createdFiles[0], "Expected state to track the config file")

	state2 := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}
	err = createConfigFile(state2)
	require.Error(t, err, "Expected error when config file already exists")

	assert.Empty(t, state2.createdFiles, "Expected state not to track existing file")
}
