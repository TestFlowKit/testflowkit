package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTestsTimeOutShouldBeOverridenByArgs(t *testing.T) {
	args := argsConfig{
		Run: &runCmd{
			Timeout: 15 * time.Second,
		},
	}

	file := cliConfig{
		Timeout: "10s",
	}

	assert.Equal(t, "15s", initAppConfig(args, file).Timeout)
}

func TestTestsTimeOutShouldBeTheSameThanFile(t *testing.T) {
	args := argsConfig{}

	file := cliConfig{
		Timeout: "10s",
	}

	assert.Equal(t, "10s", initAppConfig(args, file).Timeout)
}

func TestGherkinLocationShouldBeOverridenByCLIFileConfig(t *testing.T) {
	args := argsConfig{
		Run: &runCmd{
			GherkinLocation: "new_path",
		},
	}

	fileConfig := cliConfig{
		GherkinLocation: "old_path",
	}

	assert.Equal(t, "new_path", initAppConfig(args, fileConfig).GherkinLocation)
}
