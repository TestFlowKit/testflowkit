package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFeaturesFilesPathsLocationOverriding(t *testing.T) {
	args := appArgs{GherkinLocation: "new_path"}
	noOverrideArgs := appArgs{GherkinLocation: ""}

	cliConfig := ClI{
		testingConfig: testingConfig{
			GherkinLocation: "old_path",
		},
	}

	assert.Equal(t, "new_path", overrideClIConfig(args, cliConfig).GherkinLocation)
	assert.Equal(t, "old_path", overrideClIConfig(noOverrideArgs, cliConfig).GherkinLocation)
}

func TestTimeoutOverriding(t *testing.T) {
	args := appArgs{Timeout: 15 * time.Second}
	noOverrideArgs := appArgs{Timeout: 0 * time.Millisecond}

	cliConfig := ClI{
		testingConfig: testingConfig{
			Timeout: "10s",
		},
	}

	assert.Equal(t, "15s", overrideClIConfig(args, cliConfig).Timeout)
	assert.Equal(t, "10s", overrideClIConfig(noOverrideArgs, cliConfig).Timeout)
}

func TestTagsOverriding(t *testing.T) {
	args := appArgs{Tags: "@OK && ~@NOK"}
	noOverrideArgs := appArgs{Tags: ""}

	cliConfig := ClI{
		testingConfig: testingConfig{
			Tags: "@DEFAULT",
		},
	}

	assert.Equal(t, "@OK && ~@NOK", overrideClIConfig(args, cliConfig).Tags)
	assert.Equal(t, "@DEFAULT", overrideClIConfig(noOverrideArgs, cliConfig).Tags)
}

func TestParallelArgsShouldOverrideAlwaysCLIConfig(t *testing.T) {
	args := appArgs{Parallel: 0}
	secondArgs := appArgs{Parallel: 10}
	third := appArgs{Parallel: -1}

	cliConfig := ClI{
		testingConfig: testingConfig{
			Parallel: 5,
		},
	}

	assert.Equal(t, 0, overrideClIConfig(args, cliConfig).Parallel)
	assert.Equal(t, 10, overrideClIConfig(secondArgs, cliConfig).Parallel)
	assert.Equal(t, 5, overrideClIConfig(third, cliConfig).Parallel)
}

func TestDisplayBrowserArgsShouldOverrideAlwaysCLIConfig(t *testing.T) {
	args := appArgs{DisplayBrowser: true}
	secondArgs := appArgs{DisplayBrowser: false}

	cliConfig := ClI{}

	assert.True(t, overrideClIConfig(args, cliConfig).displayBrowser)
	assert.False(t, overrideClIConfig(secondArgs, cliConfig).displayBrowser)
}
