package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

type SelectorType string

const (
	SelectorTypeCSS   SelectorType = "css"
	SelectorTypeXPath SelectorType = "xpath"
)

type Selector struct {
	Type  SelectorType
	Value string
}

func (s *Selector) String() string {
	return s.Value
}

func (s *Selector) IsXPath() bool {
	return s.Type == SelectorTypeXPath
}

func (s *Selector) IsCSS() bool {
	return s.Type == SelectorTypeCSS
}

func NewSelector(selectorStr string) Selector {
	selectorStr = strings.TrimSpace(selectorStr)
	if strings.HasPrefix(selectorStr, "xpath:") {
		return Selector{
			Type:  SelectorTypeXPath,
			Value: strings.TrimPrefix(selectorStr, "xpath:"),
		}
	}
	return Selector{
		Type:  SelectorTypeCSS,
		Value: selectorStr,
	}
}

type Config struct {
	ActiveEnvironment string `yaml:"active_environment" validate:"required"`

	Settings GlobalSettings `yaml:"settings" validate:"required"`

	Environments map[string]Environment `yaml:"environments" validate:"required,min=1"`

	Frontend FrontendConfig `yaml:"frontend" validate:"required"`

	Backend BackendConfig `yaml:"backend"`
}

type GlobalSettings struct {
	DefaultTimeout int `yaml:"default_timeout" validate:"min=1000,max=300000"`

	// PageLoadTimeout int `yaml:"page_load_timeout" validate:"min=1000,max=300000"`

	ScreenshotOnFailure bool `yaml:"screenshot_on_failure"`

	// VideoRecording bool `yaml:"video_recording"`

	Concurrency int `yaml:"concurrency" validate:"min=1,max=20"`

	Headless bool `yaml:"headless"`

	SlowMotion int `yaml:"slow_motion" validate:"omitempty" `

	ReportFormat string `yaml:"report_format" validate:"oneof=html json junit"`

	GherkinLocation string `yaml:"gherkin_location" validate:"required"`

	Tags string `yaml:"tags"`
}

type Environment struct {
	FrontendBaseURL string `yaml:"frontend_base_url" validate:"required,url"`

	APIBaseURL string `yaml:"api_base_url" validate:"required,url"`
}

type FrontendConfig struct {
	Elements map[string]map[string][]string `yaml:"elements" validate:"required"`

	Pages map[string]string `yaml:"pages" validate:"required"`
}

type BackendConfig struct {
	DefaultHeaders map[string]string   `yaml:"default_headers"`
	Endpoints      map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Method string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`

	Path string `yaml:"path" validate:"required"`

	Description string `yaml:"description" validate:"required"`
}

func (c *Config) GetCurrentEnvironment() (Environment, error) {
	env, exists := c.Environments[c.ActiveEnvironment]
	if !exists {
		return Environment{}, fmt.Errorf("active environment '%s' not found in configuration", c.ActiveEnvironment)
	}
	return env, nil
}

func (c *Config) GetAPIEndpoint(endpointName string) (string, Endpoint, error) {
	endpoint, exists := c.Backend.Endpoints[endpointName]
	if !exists {
		return "", Endpoint{}, fmt.Errorf("endpoint '%s' not found in configuration", endpointName)
	}

	parsedURL, err := url.Parse(endpoint.Path)
	if err != nil {
		return "", Endpoint{}, fmt.Errorf("failed to parse endpoint path: %w", err)
	}

	if parsedURL.Scheme != "" {
		return parsedURL.String(), endpoint, nil
	}

	fullURL, err := url.JoinPath(c.GetBackendBaseURL(), parsedURL.Path)
	if err != nil {
		return "", Endpoint{}, fmt.Errorf("failed to join base URL and endpoint path: %w", err)
	}

	return fullURL, endpoint, nil
}

func (c *Config) GetFrontendURL(page string) (string, error) {
	env, err := c.GetCurrentEnvironment()
	if err != nil {
		return "", err
	}

	if pagePath, ok := c.Frontend.Pages[GetLabelKey(page)]; ok {
		if strings.HasPrefix(pagePath, "http://") || strings.HasPrefix(pagePath, "https://") {
			return pagePath, nil
		}

		fullURL, errJoin := url.JoinPath(env.FrontendBaseURL, pagePath)
		if errJoin != nil {
			return "", err
		}

		return fullURL, nil
	}

	return env.FrontendBaseURL, nil
}

func (c *Config) GetElementSelectors(page, elementName string) []Selector {
	elementName = GetLabelKey(elementName)
	var selectors []Selector

	chainOfResponsibility := []func(string) []Selector{
		c.getElementSelectors(page),
		c.getElementSelectors("common"),
	}

	for _, chain := range chainOfResponsibility {
		selectors = append(selectors, chain(elementName)...)
		if len(selectors) > 0 {
			return selectors
		}
	}

	return []Selector{}
}

func (c *Config) getElementSelectors(key string) func(string) []Selector {
	return func(elementName string) []Selector {
		var selectors []Selector
		if pageElements, isFound := c.Frontend.Elements[key]; isFound {
			if selectorStrs, isSelectorsFound := pageElements[elementName]; isSelectorsFound {
				for _, selectorStr := range selectorStrs {
					selectors = append(selectors, NewSelector(selectorStr))
				}
				return selectors
			}
		}
		return selectors
	}
}

func (c *Config) GetSlowMotion() time.Duration {
	if c.Settings.SlowMotion == 0 || c.Settings.Headless {
		return 0
	}

	duration, err := time.ParseDuration(fmt.Sprintf("%dms", c.Settings.SlowMotion))
	if err != nil {
		log.Printf("Invalid slow motion duration: %d, using 0", c.Settings.SlowMotion)
		return 0
	}
	return duration
}

func (c *Config) IsHeadlessModeEnabled() bool {
	return c.Settings.Headless
}

func (c *Config) GetConcurrency() int {
	if c.Settings.Concurrency <= 0 {
		return 1
	}
	return c.Settings.Concurrency
}

func (c *Config) GetTimeout() time.Duration {
	return time.Duration(c.Settings.DefaultTimeout) * time.Millisecond
}

func (c *Config) IsScreenshotOnFailureEnabled() bool {
	return c.Settings.ScreenshotOnFailure
}

func (c *Config) GetFrontendBaseURL() string {
	return c.Environments[c.ActiveEnvironment].FrontendBaseURL
}

func (c *Config) GetBackendBaseURL() string {
	return c.Environments[c.ActiveEnvironment].APIBaseURL
}

// func (c *Config) GetPageLoadTimeout() time.Duration {
// 	return time.Duration(c.Settings.PageLoadTimeout) * time.Millisecond
// }

func (c *Config) ValidateConfiguration() error {
	if err := c.validateGlobalSettings(); err != nil {
		return err
	}

	if err := c.validateFrontend(); err != nil {
		return err
	}

	return nil
}

func (c *Config) validateFrontend() error {
	if len(c.Frontend.Elements) == 0 {
		return errors.New("frontend elements configuration is required")
	}

	if len(c.Frontend.Pages) == 0 {
		return errors.New("frontend pages configuration is required")
	}

	if c.Settings.DefaultTimeout < 1000 || c.Settings.DefaultTimeout > 300000 {
		return errors.New("default_timeout must be between 1000 and 300000 milliseconds")
	}

	// if c.Settings.PageLoadTimeout < 5000 || c.Settings.PageLoadTimeout > 300000 {
	// 	return errors.New("page_load_timeout must be between 5000 and 300000 milliseconds")
	// }
	return nil
}

func (c *Config) validateGlobalSettings() error {
	if c.ActiveEnvironment == "" {
		return errors.New("active_environment is required")
	}

	if _, exists := c.Environments[c.ActiveEnvironment]; !exists {
		return fmt.Errorf("active environment '%s' not found in environments", c.ActiveEnvironment)
	}

	if len(c.Environments) == 0 {
		return errors.New("at least one environment must be defined")
	}

	if c.Settings.Concurrency < 1 || c.Settings.Concurrency > 20 {
		return errors.New("concurrency must be between 1 and 20")
	}

	if c.Settings.GherkinLocation == "" {
		return errors.New("gherkin_location is required")
	}

	validReportFormats := []string{"html", "json", "junit"}
	reportFormatValid := false
	for _, validFormat := range validReportFormats {
		if c.Settings.ReportFormat == validFormat {
			reportFormatValid = true
			break
		}
	}
	if !reportFormatValid {
		return fmt.Errorf("report_format must be one of: %s", strings.Join(validReportFormats, ", "))
	}
	return nil
}
