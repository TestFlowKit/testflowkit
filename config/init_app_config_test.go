package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliPathDefaultValue(t *testing.T) {
	appConfig := initAppConfig(appArgsConfig{}, appFileConfig{})
	assert.Equal(t, defaultCliConfigPath, appConfig.GherkinLocation)
}

func TestReportFormatDefaultValue(t *testing.T) {
	appConfig := initAppConfig(appArgsConfig{}, appFileConfig{})
	assert.Equal(t, defaultReportFormat, appConfig.ReportFormat)
}

func TestTimeoutDefaultValue(t *testing.T) {
	appConfig := initAppConfig(appArgsConfig{}, appFileConfig{})
	assert.Equal(t, defaultTimeout, appConfig.Timeout)
}
