package config_test

import (
	"cucumber/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapAppDetailsConfig(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.InitByFilePath("cli.test.yml")

	assert.Equal(t, "MonApplication", cliConfig.AppName)
	assert.Equal(t, "ma chic app", cliConfig.AppDescription)
	assert.Equal(t, "1.0.0", cliConfig.AppVersion)
}

func TestReportingConfig(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.InitByFilePath("cli.test.yml")

	assert.Equal(t, "html", cliConfig.ReportFormat)
	assert.True(t, cliConfig.ReportEnabled)
}

func TestTestingConfig(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.InitByFilePath("cli.test.yml")

	cliConfig.SetDisplayBrowser(true)

	assert.False(t, cliConfig.IsHeadlessModeEnabled())
	assert.Equal(t, "60s", cliConfig.Timeout)
	assert.Equal(t, 4, cliConfig.Parallel)
	assert.Equal(t, "2s", cliConfig.SlowMotion)
	assert.Equal(t, 2*time.Second, cliConfig.GetSlowMotion())
	assert.Equal(t, "./features", cliConfig.GherkinLocation)
}

func TestSLowMotionMustBeZeroWhenHeadlessModeEnabled(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.InitByFilePath("cli.test.yml")

	const fileContent = `
configuration:
  display_browser: false
  timeout: 60s
  slowMotion: 2s
  concurrency: 4
  gherkin_location: "./features"  # Chemin vers les fichiers Gherkin
application:
  app_name: "MonApplication"
  app_description: "ma chic app"
  app_version: "1.0.0"
reporting:
  enable: true
  report_format: "html"  # Options possibles : "html", "json", "xml"


`
	cliConfig.InitByFileContent(fileContent)

	assert.True(t, cliConfig.IsHeadlessModeEnabled())
	assert.Equal(t, time.Duration(0), cliConfig.GetSlowMotion())
}

func TestSetConcurrencyTo1WhenEnableBrowserIsEnabled(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.SetDisplayBrowser(true)
	cliConfig.Parallel = 10

	assert.Equal(t, 0, cliConfig.GetConcurrency())
}

func TestGetConcurrencyValueWhenEnableBrowserIsDisabled(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.SetDisplayBrowser(false)
	cliConfig.Parallel = 10

	assert.Equal(t, 10, cliConfig.GetConcurrency())
}
