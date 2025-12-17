package fileutils_test

import (
	"path/filepath"
	"testflowkit/internal/utils/fileutils"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	featuresDir = "features"
	configFile  = "config.yml"
)

func TestValidatePath(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		allowedPrefixes []string
		expectError     bool
	}{
		{
			name:        "valid simple path",
			path:        featuresDir,
			expectError: false,
		},
		{
			name:        "directory traversal attempt",
			path:        "../etc/passwd",
			expectError: true,
		},
		{
			name:        "nested directory traversal",
			path:        featuresDir + "/../../../etc/passwd",
			expectError: true,
		},
		{
			name:            "valid path with allowed prefix",
			path:            filepath.Join(featuresDir, "sample.feature"),
			allowedPrefixes: []string{featuresDir},
			expectError:     false,
		},
		{
			name:            "invalid path outside allowed prefix",
			path:            configFile,
			allowedPrefixes: []string{featuresDir},
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fileutils.ValidatePath(tt.path, tt.allowedPrefixes...)
			if tt.expectError {
				require.Error(t, err, "Expected error for path %s", tt.path)
			} else {
				require.NoError(t, err, "Expected no error for path %s", tt.path)
			}
		})
	}
}

func TestPathValidationSecurity(t *testing.T) {
	t.Run("DirectoryTraversalPrevention", testDirectoryTraversalPrevention)
	t.Run("AllowedPathValidation", testAllowedPathValidation)
	t.Run("PrefixRestrictionValidation", testPrefixRestrictionValidation)
}

func testDirectoryTraversalPrevention(t *testing.T) {
	t.Helper()
	maliciousPaths := []string{
		"../etc/passwd",
		"../../etc/passwd",
		"features/../../../etc/passwd",
		"features/../../config.yml",
	}

	for _, path := range maliciousPaths {
		err := fileutils.ValidatePath(path)
		require.Error(t, err, "Expected error for malicious path: %s", path)
	}
}

func testAllowedPathValidation(t *testing.T) {
	t.Helper()
	validPaths := []string{
		featuresDir,
		configFile,
		filepath.Join(featuresDir, "sample.feature"),
		filepath.Join(featuresDir, "test.feature"),
	}

	for _, path := range validPaths {
		err := fileutils.ValidatePath(path)
		require.NoError(t, err, "Expected no error for valid path: %s", path)
	}
}

func testPrefixRestrictionValidation(t *testing.T) {
	t.Helper()

	err := fileutils.ValidatePath(filepath.Join(featuresDir, "sample.feature"), featuresDir)
	require.NoError(t, err, "Expected no error for path within allowed prefix")

	err = fileutils.ValidatePath(configFile, featuresDir)
	require.Error(t, err, "Expected error for path outside allowed prefix")
}
