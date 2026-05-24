package config_test

import (
	"os"
	"path/filepath"
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveConfigPath(t *testing.T) {
	t.Run("returns default when it exists", func(t *testing.T) {
		dir := t.TempDir()
		defaultPath := filepath.Join(dir, config.DefaultConfigFile)
		require.NoError(t, os.WriteFile(defaultPath, []byte("settings: {}"), 0600))

		assert.Equal(t, defaultPath, config.ResolveConfigPath(defaultPath))
	})

	t.Run("falls back to legacy when default is missing", func(t *testing.T) {
		dir := t.TempDir()
		legacyPath := filepath.Join(dir, config.LegacyConfigFile)
		require.NoError(t, os.WriteFile(legacyPath, []byte("settings: {}"), 0600))

		defaultPath := filepath.Join(dir, config.DefaultConfigFile)
		assert.Equal(t, legacyPath, config.ResolveConfigPath(defaultPath))
	})

	t.Run("returns requested path when neither file exists", func(t *testing.T) {
		dir := t.TempDir()
		defaultPath := filepath.Join(dir, config.DefaultConfigFile)
		assert.Equal(t, defaultPath, config.ResolveConfigPath(defaultPath))
	})

	t.Run("does not fall back for custom paths", func(t *testing.T) {
		dir := t.TempDir()
		legacyPath := filepath.Join(dir, config.LegacyConfigFile)
		require.NoError(t, os.WriteFile(legacyPath, []byte("settings: {}"), 0600))

		customPath := filepath.Join(dir, "custom.yml")
		assert.Equal(t, customPath, config.ResolveConfigPath(customPath))
	})
}
