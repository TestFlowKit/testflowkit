package config

import (
	"strings"
)

// SelectorType is the query language for a browser element selector.
type SelectorType string

const (
	// SelectorTypeCSS uses CSS selector syntax (default).
	SelectorTypeCSS SelectorType = "css"
	// SelectorTypeXPath uses XPath; prefix the selector with "xpath:".
	SelectorTypeXPath SelectorType = "xpath"
)

// Selector pairs a query language with its expression string.
type Selector struct {
	// Type is the query language (css or xpath).
	Type SelectorType
	// Value is the selector expression without any language prefix.
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

// NewSelector parses a raw selector string into a typed Selector.
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

// GlobalSettings holds project-wide runner options; partially overridable via CLI flags.
type GlobalSettings struct {

	// PageLoadTimeout int `yaml:"page_load_timeout" validate:"min=1000,max=300000"`

	// VideoRecording bool `yaml:"video_recording"`

	// Concurrency is the number of parallel scenarios (1–20, default 1).
	Concurrency int `yaml:"concurrency" validate:"min=1,max=20"`

	// ReportFormat is the test report output format: html, json, or junit.
	ReportFormat string `yaml:"report_format" validate:"oneof=html json junit"`

	// GherkinLocation is the path to .feature files, e.g. "features". Required.
	GherkinLocation string `yaml:"gherkin_location" validate:"required"`

	// Tags is a Gherkin tag expression to filter scenarios, e.g. "@smoke".
	Tags string `yaml:"tags"`

	// EnvFile is the path to a YAML env file merged into the env block.
	EnvFile string `yaml:"env_file"`

	// Debug tunes debug output; debug mode itself is enabled via CLI --debug.
	Debug DebugConfig `yaml:"debug"`
}

// DebugConfig holds debug-mode settings; enabled only via the CLI --debug flag.
type DebugConfig struct {
	// PrettyPrint indents request/response bodies in debug output.
	PrettyPrint bool `yaml:"pretty_print"`

	// MaxBodySize caps debug body output in bytes (0 = unlimited, max 10 MiB).
	MaxBodySize int64 `yaml:"max_body_size" validate:"min=0,max=10485760"`

	// Enabled is set by CLI only and ignored during YAML unmarshalling.
	Enabled bool `yaml:"-"`
}

// FrontendElements maps group names to element names and their selector fallbacks.
type FrontendElements = map[string]map[string][]string

// FrontendPages maps page names to URL paths appended to frontend.base_url.
type FrontendPages = map[string]string

// FrontendConfig configures the browser driver for UI testing.
type FrontendConfig struct {
	// BaseURL is the app root URL; supports {{ env.* }} interpolation.
	BaseURL string `yaml:"base_url" validate:"required"`

	// DefaultTimeout is the element search timeout in ms (1000–300000).
	DefaultTimeout int `yaml:"default_timeout" validate:"min=1000,max=300000"`

	// ScreenshotOnFailure captures a PNG on step failure.
	ScreenshotOnFailure bool `yaml:"screenshot_on_failure"`

	// Headless runs the browser without a visible window.
	Headless bool `yaml:"headless"`

	// ThinkTime is a delay in ms between browser actions (0 in CI).
	ThinkTime int `yaml:"think_time" validate:"omitempty"`

	// Driver is the browser library: "rod" (default) or "playwright".
	Driver string `yaml:"driver" validate:"omitempty,oneof=rod playwright"`

	// UserAgent overrides the default browser user-agent string.
	UserAgent string `yaml:"user_agent"`

	// Locale sets the browser locale, e.g. "en-US".
	Locale string `yaml:"locale"`

	// TimezoneID sets the browser IANA timezone, e.g. "Europe/Paris".
	TimezoneID string `yaml:"timezone_id" validate:"omitempty,timezone_id"`

	// Elements maps groups to named CSS/XPath selectors referenced in Gherkin steps.
	Elements FrontendElements `yaml:"elements"`

	// Pages maps page names to paths, e.g. home: "/".
	Pages FrontendPages `yaml:"pages"`
}

// APIType is the protocol used by an API definition.
type APIType string

const (
	// APITypeREST is an HTTP REST API (requires base_url and endpoints).
	APITypeREST APIType = "rest"
	// APITypeGraphQL is a GraphQL API (requires endpoint and operations).
	APITypeGraphQL APIType = "graphql"
)

// APIsConfig holds the project-wide API timeout and named API definitions.
type APIsConfig struct {
	// DefaultTimeout is the fallback request timeout in ms (default 30000).
	DefaultTimeout int `yaml:"default_timeout" validate:"omitempty,min=1000,max=300000"`

	// Definitions maps API names to REST or GraphQL configs.
	Definitions map[string]APIDefinition `yaml:"definitions" validate:"required,min=1"`
}

// APIDefinition describes a single named REST or GraphQL API client.
type APIDefinition struct {
	// Type is the API protocol: "rest" or "graphql". Required.
	Type APIType `yaml:"type" validate:"required,oneof=rest graphql"`

	// BaseURL is the REST root URL; required when type is "rest".
	BaseURL string `yaml:"base_url" validate:"required_if=Type rest"`

	// Endpoint is the GraphQL URL; required when type is "graphql".
	Endpoint string `yaml:"endpoint" validate:"required_if=Type graphql"`

	// DefaultHeaders are HTTP headers sent with every request to this API.
	DefaultHeaders map[string]string `yaml:"default_headers"`

	// Timeout overrides apis.default_timeout for this API in ms.
	Timeout *int `yaml:"timeout" validate:"omitempty,min=1000,max=300000"`

	// Endpoints maps REST endpoint names; required when type is "rest".
	Endpoints map[string]Endpoint `yaml:"endpoints" validate:"required_if=Type rest,min=1"`

	// Operations maps GraphQL operation names; required when type is "graphql".
	Operations map[string]GraphQLOperation `yaml:"operations" validate:"required_if=Type graphql,min=1"`

	// SecurityRef overrides default_security for all requests to this API.
	SecurityRef *SecurityRef `yaml:"security_ref"`

	// SecurityOverrides patches fields of the resolved security scheme.
	SecurityOverrides *SecurityOverrides `yaml:"security_overrides"`
}

// Endpoint describes a named REST operation within an API definition.
type Endpoint struct {
	// Method is the HTTP verb (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS).
	Method string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`

	// Path is the URL path relative to base_url, e.g. "/users/{id}".
	Path string `yaml:"path" validate:"required"`

	// Description explains the endpoint purpose for AI scenario generation.
	Description string `yaml:"description" validate:"required"`

	// Timeout overrides the API-level timeout for this endpoint in ms.
	Timeout *int `yaml:"timeout" validate:"omitempty,min=1000,max=300000"`

	// SecurityRef overrides inherited auth for this endpoint.
	SecurityRef *SecurityRef `yaml:"security_ref"`
}

// FileConfig maps logical file names to paths for upload and body steps.
type FileConfig struct {
	// Definitions maps logical names to paths relative to base_directory.
	Definitions map[string]string `yaml:"definitions"`

	// BaseDirectory is the root path prepended to every file definition.
	BaseDirectory string `yaml:"base_directory"`
}

// GraphQLOperation describes a named GraphQL query or mutation.
type GraphQLOperation struct {
	// Type is the operation kind: "query" or "mutation".
	Type string `yaml:"type" validate:"required,oneof=query mutation"`

	// Operation is a .graphql file path or inline GraphQL document.
	Operation string `yaml:"operation" validate:"required"`

	// Description explains the operation purpose for AI scenario generation.
	Description string `yaml:"description"`

	// Timeout overrides the API-level timeout for this operation in ms.
	Timeout *int `yaml:"timeout" validate:"omitempty,min=1000,max=300000"`

	// SecurityRef overrides inherited auth for this operation.
	SecurityRef *SecurityRef `yaml:"security_ref"`
}
