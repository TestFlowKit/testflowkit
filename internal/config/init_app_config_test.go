package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliPathDefaultValue(t *testing.T) {
	appConfig := initAppConfig(argsConfig{}, cliConfig{})
	assert.Equal(t, defaultCliConfigPath, appConfig.GherkinLocation)
}

func TestReportFormatDefaultValue(t *testing.T) {
	appConfig := initAppConfig(argsConfig{}, cliConfig{})
	assert.Equal(t, defaultReportFormat, appConfig.ReportFormat)
}

func TestTimeoutDefaultValue(t *testing.T) {
	appConfig := initAppConfig(argsConfig{}, cliConfig{})
	assert.Equal(t, defaultTimeout, appConfig.Timeout)
}
