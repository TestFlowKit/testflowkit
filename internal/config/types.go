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
	// Type indicates whether the selector uses CSS or XPath syntax.
	Type SelectorType
	// Value is the raw selector expression consumed by the browser driver.
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

	// Concurrency is the number of scenarios executed in parallel.
	Concurrency int `yaml:"concurrency" validate:"min=1,max=20"`

	// ReportFormat selects the output report format.
	ReportFormat string `yaml:"report_format" validate:"oneof=html json junit"`

	// GherkinLocation points to the directory containing .feature files.
	GherkinLocation string `yaml:"gherkin_location" validate:"required"`

	// Tags is an optional Cucumber expression used to filter executed scenarios.
	Tags string `yaml:"tags"`

	// EnvFile points to a YAML or .env file merged into the env section.
	EnvFile string `yaml:"env_file"`

	// Debug controls debug logging behavior for HTTP payloads.
	Debug DebugConfig `yaml:"debug"`
}

// DebugConfig holds debug-mode settings. All fields default to their zero
// values so that existing config files without a "debug" block continue to
// work unchanged.
type DebugConfig struct {
	// PrettyPrint formats JSON bodies in logs for readability.
	PrettyPrint bool `yaml:"pretty_print"`

	// MaxBodySize caps response bytes included in debug output.
	MaxBodySize int64 `yaml:"max_body_size" validate:"min=0,max=10485760"`
}

type FrontendElements = map[string]map[string][]string

type FrontendPages = map[string]string
type FrontendConfig struct {
	// BaseURL is the base URL for the frontend application
	// Supports variable interpolation (e.g., "{{ env.HOST }}:3000")
	BaseURL string `yaml:"base_url" validate:"required"`

	// DefaultTimeout is the maximum time (in milliseconds)
	// to wait for an element to be found during element search operations.
	DefaultTimeout int `yaml:"default_timeout" validate:"min=1000,max=300000"`

	// ScreenshotOnFailure captures a screenshot when a frontend step fails.
	ScreenshotOnFailure bool `yaml:"screenshot_on_failure"`

	// Headless runs the browser without a visible UI window.
	Headless bool `yaml:"headless"`

	// ThinkTime is the delay (in milliseconds) between browser actions to simulate human behavior
	ThinkTime int `yaml:"think_time" validate:"omitempty"`

	// Driver specifies the browser driver to use ("rod" or "playwright")
	Driver string `yaml:"driver" validate:"omitempty,oneof=rod playwright"`

	// UserAgent sets a custom user agent string for the browser
	UserAgent string `yaml:"user_agent"`

	// Locale sets the browser locale (e.g., "en-US", "fr-FR")
	Locale string `yaml:"locale"`

	// TimezoneID sets the browser timezone using IANA format (e.g., "America/New_York", "Europe/Paris")
	TimezoneID string `yaml:"timezone_id" validate:"omitempty,timezone_id"`

	// Elements maps feature/page groups to named selectors used in Gherkin steps.
	Elements FrontendElements `yaml:"elements"`

	// Pages maps logical page names to URL paths relative to base_url.
	Pages FrontendPages `yaml:"pages"`
}

type APIType string

const (
	APITypeREST    APIType = "rest"
	APITypeGraphQL APIType = "graphql"
)

type APIsConfig struct {
	// DefaultTimeout is the fallback timeout for API requests in milliseconds.
	DefaultTimeout int `yaml:"default_timeout" validate:"omitempty,min=1000,max=300000"`
	// Definitions contains named API definitions referenced by step expressions.
	Definitions map[string]APIDefinition `yaml:"definitions" validate:"required,min=1"`
}

type APIDefinition struct {
	// Type selects the API protocol: rest or graphql.
	Type APIType `yaml:"type" validate:"required,oneof=rest graphql"`
	// BaseURL is required for REST APIs and prefixes endpoint paths.
	BaseURL string `yaml:"base_url" validate:"required_if=Type rest"`
	// Endpoint is required for GraphQL APIs and points to the GraphQL HTTP endpoint.
	Endpoint string `yaml:"endpoint" validate:"required_if=Type graphql"`
	// DefaultHeaders are attached to every request for this API.
	DefaultHeaders map[string]string `yaml:"default_headers"`
	// Timeout overrides the global APIs default timeout for this API.
	Timeout *int `yaml:"timeout" validate:"omitempty,min=1000,max=300000"`
	// Endpoints contains REST endpoint definitions keyed by operation name.
	Endpoints map[string]Endpoint `yaml:"endpoints" validate:"required_if=Type rest,min=1"`
	// Operations contains GraphQL operations keyed by operation name.
	Operations map[string]GraphQLOperation `yaml:"operations" validate:"required_if=Type graphql,min=1"`

	// Security – API-level auth (overrides project default_security).
	// Use SecurityRef to reference a named scheme or embed an inline one.
	// Use SecurityOverrides to patch specific fields (e.g. scopes) without
	// duplicating the whole scheme.
	SecurityRef       *SecurityRef       `yaml:"security_ref"`
	SecurityOverrides *SecurityOverrides `yaml:"security_overrides"`
}

type Endpoint struct {
	// Method is the HTTP verb used for this REST endpoint.
	Method string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`

	// Path is appended to API base_url to build the request URL.
	Path string `yaml:"path" validate:"required"`

	// Description documents the endpoint purpose for humans and agents.
	Description string `yaml:"description" validate:"required"`

	// Timeout overrides the API-level timeout for this specific REST endpoint.
	Timeout *int `yaml:"timeout" validate:"omitempty,min=1000,max=300000"`

	// SecurityRef overrides the API-level and project-level security for this
	// specific endpoint. Set Name to "none" to disable all inherited auth.
	SecurityRef *SecurityRef `yaml:"security_ref"`
}

type FileConfig struct {
	// Definitions maps logical file names to relative file paths.
	Definitions map[string]string `yaml:"definitions"`
	// BaseDirectory prefixes all file definition paths at runtime.
	BaseDirectory string `yaml:"base_directory"`
}

type GraphQLOperation struct {
	// Type declares whether the operation is a query or mutation.
	Type string `yaml:"type" validate:"required,oneof=query mutation"`
	// Operation is the GraphQL document string sent to the endpoint.
	Operation string `yaml:"operation" validate:"required"`
	// Description documents the operation purpose for humans and agents.
	Description string `yaml:"description"`

	// Timeout overrides the API-level timeout for this specific GraphQL operation.
	Timeout *int `yaml:"timeout" validate:"omitempty,min=1000,max=300000"`

	// SecurityRef overrides the API-level and project-level security for this
	// specific GraphQL operation. Set Name to "none" to disable all inherited auth.
	SecurityRef *SecurityRef `yaml:"security_ref"`
}
