package config

import (
	"os"
	"path/filepath"
)

const (
	// DefaultConfigFile is the primary configuration file name.
	DefaultConfigFile = "testflowkit.yml"
	// LegacyConfigFile is supported for backward compatibility when DefaultConfigFile is absent.
	LegacyConfigFile = "config.yml"
)

// ResolveConfigPath returns path when it exists. When path names DefaultConfigFile and
// that file is missing, LegacyConfigFile in the same directory is used if present.
func ResolveConfigPath(path string) string {
	if _, err := os.Stat(path); err == nil {
		return path
	}
	if filepath.Base(path) != DefaultConfigFile {
		return path
	}

	legacyPath := filepath.Join(filepath.Dir(path), LegacyConfigFile)
	if _, err := os.Stat(legacyPath); err == nil {
		return legacyPath
	}

	return path
}
