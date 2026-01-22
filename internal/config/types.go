package config

import (
	"strings"
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

type GlobalSettings struct {

	// PageLoadTimeout int `yaml:"page_load_timeout" validate:"min=1000,max=300000"`

	// VideoRecording bool `yaml:"video_recording"`

	Concurrency int `yaml:"concurrency" validate:"min=1,max=20"`

	ThinkTime int `yaml:"think_time" validate:"omitempty" `

	ReportFormat string `yaml:"report_format" validate:"oneof=html json junit"`

	GherkinLocation string `yaml:"gherkin_location" validate:"required"`

	Tags string `yaml:"tags"`

	EnvFile string `yaml:"env_file"`
}

type FrontendElements = map[string]map[string][]string

type FrontendPages = map[string]string
type FrontendConfig struct {
	// DefaultTimeout is the maximum time (in milliseconds)
	// to wait for an element to be found during element search operations.
	DefaultTimeout int `yaml:"default_timeout" validate:"min=1000,max=300000"`

	ScreenshotOnFailure bool `yaml:"screenshot_on_failure"`

	Headless bool `yaml:"headless"`

	Elements FrontendElements `yaml:"elements"`

	Pages FrontendPages `yaml:"pages"`
}

type BackendConfig struct {
	DefaultHeaders map[string]string   `yaml:"default_headers"`
	Endpoints      map[string]Endpoint `yaml:"endpoints"`
	GraphQL        *GraphQLConfig      `yaml:"graphql"`
}

type Endpoint struct {
	Method string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`

	Path string `yaml:"path" validate:"required"`

	Description string `yaml:"description" validate:"required"`
}

type FileConfig struct {
	Definitions   map[string]string `yaml:"definitions"`
	BaseDirectory string            `yaml:"base_directory"`
}

type GraphQLConfig struct {
	DefaultHeaders map[string]string           `yaml:"default_headers"`
	Operations     map[string]GraphQLOperation `yaml:"operations" validate:"required,min=1"`
}

type GraphQLOperation struct {
	Type        string `yaml:"type" validate:"required,oneof=query mutation"`
	Operation   string `yaml:"operation" validate:"required"`
	Description string `yaml:"description"`
}
