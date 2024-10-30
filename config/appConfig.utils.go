package config

import (
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
)

type appConfig struct {
	AppName        string `yaml:"app_name"`
	AppDescription string `yaml:"app_description,omitempty"`
	AppVersion     string `yaml:"app_version"`
}

type reportingConfig struct {
	ReportEnabled bool   `yaml:"enable"`
	ReportFormat  string `yaml:"report_format"`
}

type testingConfig struct {
	DisplayBrowser  bool   `yaml:"display_browser"`
	Timeout         string `yaml:"timeout"`
	Concurrency     int    `yaml:"concurrency"`
	GherkinLocation string `yaml:"gherkin_location"`
}

type tagsConfig struct {
	IncludeTags []string `yaml:"include"`
	ExcludeTags []string `yaml:"exclude"`
}

func (c tagsConfig) GetTagsExpression() string {
	var expression string
	if len(c.IncludeTags) == 0 && len(c.ExcludeTags) == 0 {
		return expression
	}

	if len(c.IncludeTags) > 0 {
		expression = strings.Join(c.formatTags(c.IncludeTags), ",")
	}

	if len(c.ExcludeTags) == 0 {
		return expression
	}

	excludeExp := fmt.Sprintf("~%s", strings.Join(c.formatTags(c.ExcludeTags), " && ~"))

	if len(expression) > 0 {
		return fmt.Sprintf("%s && %s", expression, excludeExp)
	}
	return excludeExp
}

func (c tagsConfig) formatTags(tags []string) []string {
	var formattedTags []string
	for _, tag := range tags {
		if !strings.HasPrefix(tag, "@") {
			tag = fmt.Sprintf("@%s", tag)
		}
		formattedTags = append(formattedTags, tag)
	}
	return formattedTags
}

type configType interface {
	tagsConfig | testingConfig | reportingConfig | appConfig
}

func getConfig[T configType](file, path string, config *T) error {
	configPath, err := yaml.PathString(path)
	if err != nil {
		return err
	}

	if err = configPath.Read(strings.NewReader(file), config); err != nil {
		return err
	}

	return nil
}
