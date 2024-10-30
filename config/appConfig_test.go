package config_test

import (
	"cucumber/config"
	"testing"

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
	assert.Equal(t, "./features", cliConfig.GherkinLocation)
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
