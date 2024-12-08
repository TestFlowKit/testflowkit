package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTestsTimeOutShouldBeOverridenByArgs(t *testing.T) {
	args := appArgsConfig{
		Timeout: 15 * time.Second,
	}

	file := appFileConfig{
		Timeout: "10s",
	}

	assert.Equal(t, "15s", InitAppConfig(args, file).Timeout)
}

func TestTestsTimeOutShouldBeTheSameThanFile(t *testing.T) {
	args := appArgsConfig{}

	file := appFileConfig{
		Timeout: "10s",
	}

	assert.Equal(t, "10s", InitAppConfig(args, file).Timeout)
}

func TestGherkinLocationShouldBeOverridenByCLIFileConfig(t *testing.T) {
	args := appArgsConfig{GherkinLocation: "new_path"}

	fileConfig := appFileConfig{
		GherkinLocation: "old_path",
	}

	assert.Equal(t, "new_path", InitAppConfig(args, fileConfig).GherkinLocation)
}
