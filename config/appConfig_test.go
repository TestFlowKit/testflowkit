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

	assert.True(t, cliConfig.DisplayBrowser)
	assert.Equal(t, "60s", cliConfig.Timeout)
	assert.Equal(t, 4, cliConfig.Concurrency)
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
  tags:
    include:
      - "login"
      - "user"
    exclude:
      - "slow"
      - "experimental"
application:
  app_name: "MonApplication"
  app_description: "ma chic app"
  app_version: "1.0.0"
reporting:
  enable: true
  report_format: "html"  # Options possibles : "html", "json", "xml"


`
	cliConfig.InitByFileContent(fileContent)

	assert.False(t, cliConfig.DisplayBrowser)
	assert.Equal(t, time.Duration(0), cliConfig.GetSlowMotion())
}

func TestTagsConfig(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.InitByFilePath("cli.test.yml")

	assert.Equal(t, []string{"login", "user"}, cliConfig.IncludeTags)
	assert.Equal(t, []string{"slow", "experimental"}, cliConfig.ExcludeTags)
}

func TestSetConcurrencyTo1WhenEnableBrowserIsEnabled(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.DisplayBrowser = true
	cliConfig.Concurrency = 10

	assert.Equal(t, 0, cliConfig.GetConcurrency())
}

func TestGetConcurrencyValueWhenEnableBrowserIsDisabled(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.DisplayBrowser = false
	cliConfig.Concurrency = 10

	assert.Equal(t, 10, cliConfig.GetConcurrency())
}

func TestTagsExpressionForOnlyIncludeTags(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.IncludeTags = []string{"LOGIN", "@USER-MANAGEMENT", "TEST_CHAP"}
	expression := "@LOGIN,@USER-MANAGEMENT,@TEST_CHAP"

	assert.Equal(t, cliConfig.GetTagsExpression(), expression)
}

func TestTagsExpressionForOnlyExcludeTags(t *testing.T) {
	cliConfig := config.ClI{}
	cliConfig.ExcludeTags = []string{"LOGIN", "@USER-MANAGEMENT", "TEST_CHAP"}
	expression := "~@LOGIN && ~@USER-MANAGEMENT && ~@TEST_CHAP"

	assert.Equal(t, cliConfig.GetTagsExpression(), expression)
}

func TestTagsExpressionReturnEmptyExpBecauseNoTagsProvided(t *testing.T) {
	cliConfig := config.ClI{}
	assert.Equal(t, "", cliConfig.GetTagsExpression())
}
