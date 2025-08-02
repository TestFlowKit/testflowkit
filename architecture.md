# TestFlowKit Architecture

## Overview

TestFlowKit is a modular, extensible web test automation framework built in Go. The architecture follows clean architecture principles with clear separation of concerns, dependency inversion, and high cohesion. This document provides a detailed technical overview of the system's architecture, design patterns, and component interactions.

## 🏗️ High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Application Layer                        │
├─────────────────────────────────────────────────────────────┤
│  Command Line Interface (CLI)                               │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │   Run Mode  │  │  Init Mode  │  │ Validate    │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Business Logic Layer                     │
├─────────────────────────────────────────────────────────────┤
│  Test Execution Engine                                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │ Gherkin     │  │ Step        │  │ Scenario    │         │
│  │ Parser      │  │ Builder     │  │ Context     │         │
│  │             │  │             │  │             │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Domain Layer                             │
├─────────────────────────────────────────────────────────────┤
│  Core Domain Models                                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │ Browser     │  │ Config      │  │ Reporter    │         │
│  │ Interface   │  │ Management  │  │ System      │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Infrastructure Layer                     │
├─────────────────────────────────────────────────────────────┤
│  External Dependencies                                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │ Rod Browser │  │ HTTP Client │  │ File System │         │
│  │ Engine      │  │             │  │             │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
testflowkit/
├── cmd/testflowkit/                    # Application Entry Point
│   ├── main.go                        # Main application entry point
│   └── args.config.go                 # CLI argument parsing
├── internal/                          # Private application code
│   ├── actions/                       # Application actions
│   │   ├── main.go                    # Action dispatcher
│   │   ├── run.go                     # Test execution logic
│   │   ├── init.go                    # Project initialization
│   │   ├── validate.go                # Configuration validation
│   │   └── boilerplate/               # Template generation
│   ├── browser/                       # Browser automation layer
│   │   ├── main.go                    # Browser factory
│   │   ├── common/                    # Common interfaces
│   │   │   └── types.go               # Browser interface definitions
│   │   └── rod/                       # Rod browser implementation
│   │       ├── browser.go             # Browser instance management
│   │       ├── page.go                # Page operations
│   │       ├── element.go             # Element interactions
│   │       └── keyboard.go            # Keyboard input handling
│   ├── config/                        # Configuration management
│   │   ├── loader.go                  # Configuration loading
│   │   ├── types.go                   # Configuration structures
│   │   ├── mode.go                    # Execution modes
│   │   └── utils.go                   # Configuration utilities
│   └── steps_definitions/             # Gherkin step implementations
│       ├── core/                      # Core step framework
│       │   ├── scenario/              # Scenario context management
│       │   │   ├── main.go            # Context factory
│       │   │   ├── frontend.go        # Frontend context operations
│       │   │   ├── rest_api_steps_context.go # API context
│       │   │   └── helpers.go         # Context utilities
│       │   └── stepbuilder/           # Step definition builder
│       │       ├── types.go           # Step definition structures
│       │       ├── step_no_var.go     # No-variable steps
│       │       ├── step_one_var.go    # Single-variable steps
│       │       ├── step_two_vars.go   # Two-variable steps
│       │       ├── documentation.go   # Step documentation
│       │       └── validation_errors.go # Validation framework
│       ├── frontend/                  # Frontend testing steps
│       │   ├── main.go                # Frontend step registry
│       │   ├── assertions/            # Assertion steps
│       │   ├── form/                  # Form interaction steps
│       │   ├── navigation/            # Navigation steps
│       │   ├── mouse/                 # Mouse interaction steps
│       │   ├── keyboard/              # Keyboard input steps
│       │   └── visual/                # Visual verification steps
│       └── restapi/                   # API testing steps
│           ├── main.go                # API step registry
│           ├── prepare_request.go     # Request preparation
│           ├── send_request.go        # Request execution
│           ├── check_response_status_code.go # Status validation
│           └── response_body_should_contain.go # Response validation
├── pkg/                              # Public packages
│   ├── gherkinparser/                # Gherkin parsing and processing
│   │   ├── main.go                   # Main parser entry point
│   │   ├── types.go                  # Gherkin data structures
│   │   ├── macro.go                  # Macro processing logic
│   │   └── apply_macros.go           # Macro application
│   ├── logger/                       # Logging system
│   │   ├── common.go                 # Common logging interface
│   │   ├── info.go                   # Info level logging
│   │   ├── error.go                  # Error level logging
│   │   ├── warn.go                   # Warning level logging
│   │   └── success.go                # Success level logging
│   ├── reporters/                    # Test reporting system
│   │   ├── main.go                   # Report factory
│   │   ├── report_formatter.go       # Report formatting interface
│   │   ├── html_report.formatter.go  # HTML report generation
│   │   ├── json_report_formatter.go  # JSON report generation
│   │   ├── html_report.template.html # HTML report template
│   │   └── scenario.go               # Scenario result tracking
│   └── utils/                        # Utility functions
│       └── text_writer.go            # Text output utilities
├── e2e/                             # End-to-end test examples
│   ├── features/                    # Gherkin feature files
│   │   ├── frontend/                # Frontend test examples
│   │   └── restapi/                 # API test examples
│   ├── compose.yml                  # Docker compose for test environment
│   └── server/                      # Test server for examples
├── documentation/                   # Project documentation
├── scripts/                         # Build and utility scripts
└── build/                          # Build artifacts
```

## 🔧 Core Components

### 1. Application Entry Point (`cmd/testflowkit/`)

The application entry point handles CLI argument parsing, configuration loading, and action dispatching.

**Key Responsibilities:**

- Parse command-line arguments
- Load and validate configuration
- Dispatch to appropriate action handler
- Display configuration summary

**Design Patterns:**

- **Command Pattern**: Different execution modes (run, init, validate)
- **Factory Pattern**: Configuration and action creation

### 2. Configuration Management (`internal/config/`)

Centralized configuration management with environment-specific settings and validation.

**Key Features:**

- YAML-based configuration
- Environment-specific settings
- Configuration validation
- Default value management

**Configuration Structure:**

```yaml
active_environment: "local"
settings:
  default_timeout: 10000
  concurrency: 1
  headless: false
  screenshot_on_failure: true
  report_format: "html"
  gherkin_location: "./e2e/features"

environments:
  local:
    frontend_base_url: "http://localhost:3000"
    api_base_url: "http://localhost:8080/api"

frontend:
  elements:
    page_name:
      element_name:
        - "selector1"
        - "selector2"
        - "xpath://complex/selector[@attribute='value']"
  pages:
    page_name: "/path"

backend:
  endpoints:
    endpoint_name:
      method: "GET"
      path: "/api/endpoint"
      description: "Endpoint description"
```

**Selector Configuration Examples:**

```yaml
frontend:
  elements:
    login_page:
      # CSS Selectors (default)
      email_field:
        - "[data-testid='email-input']"
        - "input[name='email']"
        - "#email"

      # XPath Selectors (with xpath: prefix)
      complex_button:
        - "xpath://button[contains(@class, 'submit') and text()='Login']"
        - "xpath://div[@id='login-form']//button[@type='submit']"
        - "[data-testid='login-button']" # CSS fallback

      # Mixed selectors for robust detection
      dynamic_element:
        - "xpath://div[contains(@class, 'dynamic') and contains(text(), 'Loading')]"
        - ".loading-indicator"
        - "[data-testid='loading']"
```

### 3. Browser Automation Layer (`internal/browser/`)

Abstract browser interface with Rod implementation for Chrome-based automation.

**Key Components:**

- **Browser Interface**: Abstract browser operations
- **Page Interface**: Page-level operations
- **Element Interface**: Element interaction methods
- **Keyboard Interface**: Keyboard input handling

**Design Patterns:**

- **Interface Segregation**: Separate interfaces for different browser operations
- **Strategy Pattern**: Multiple selector strategies with fallback
- **Factory Pattern**: Browser instance creation

**Smart Element Detection:**

```go
// Multiple selector strategies with parallel execution
func getElementBySelectors(page common.Page, selectors []string) common.Element {
    ctx, cancel := context.WithCancel(context.Background())
    ch := make(chan common.Element, 1)

    for _, selector := range selectors {
        go searchForSelector(ctx, page, selector, ch)
    }

    <-ctx.Done()
    cancel()
    return <-ch
}
```

**XPath Support:**

TestFlowKit provides comprehensive XPath 1.0 support alongside CSS selectors for flexible element selection:

```go
// XPath selector detection and execution
func searchForSelector(ctx contextWrapper, mu *sync.RWMutex, p page, selector config.Selector, ch chan<- element) {
    var elt element
    var err error

    value := selector.String()
    if selector.IsXPath() {
        elt, err = p.GetOneByXPath(value)
    } else {
        elt, err = p.GetOneBySelector(value)
    }
    // ... error handling and result processing
}
```

**Selector Types:**

- **CSS Selectors**: Standard CSS selectors (default behavior)

  - Element IDs: `#element-id`
  - CSS classes: `.class-name`
  - Attribute selectors: `[data-testid='value']`
  - Complex selectors: `div.container > button[type='submit']`

- **XPath Selectors**: Full XPath 1.0 support with `xpath:` prefix
  - Element selection: `xpath://div[@class='container']`
  - Text matching: `xpath://button[contains(text(), 'Submit')]`
  - Attribute conditions: `xpath://input[@type='email' and @required]`
  - Complex expressions: `xpath://div[contains(@class, 'dynamic') and contains(text(), 'Loading')]`

**Parallel Selector Execution:**

The framework executes multiple selectors in parallel, automatically selecting the first successful match, providing robust element detection even when page structures change.

### 4. Step Definition Framework (`internal/steps_definitions/`)

Modular step definition system with validation and documentation support.

**Core Components:**

#### Step Builder (`core/stepbuilder/`)

- **Type Safety**: Strongly typed step definitions
- **Validation**: Pre-execution validation
- **Documentation**: Automatic documentation generation
- **Variable Support**: 0, 1, or 2 variable steps

#### Scenario Context (`core/scenario/`)

- **State Management**: Test execution state
- **Browser Management**: Browser instance lifecycle
- **API Context**: HTTP request/response management
- **Page Management**: Current page tracking

**Step Definition Example:**

```go
func (steps) elementShouldBeVisible() stepbuilder.Step {
    return stepbuilder.NewWithOneVariable(
        []string{`^the {string} should be visible$`},
        func(ctx context.Context, elementName string) (context.Context, error) {
            scenarioCtx := scenario.MustFromContext(ctx)
            currentPage, pageName := scenarioCtx.GetCurrentPage()
            element, err := browser.GetElementByLabel(currentPage, pageName, elementName)
            if err != nil {
                return ctx, err
            }

            if !element.IsVisible() {
                return ctx, fmt.Errorf("%s is not visible", elementName)
            }

            return ctx, nil
        },
        func(name string) stepbuilder.ValidationErrors {
            vc := stepbuilder.ValidationErrors{}
            if !config.IsElementDefined(name) {
                vc.AddMissingElement(name)
            }
            return vc
        },
        stepbuilder.DocParams{
            Description: "This assertion checks if the element is present in the DOM and displayed on the page.",
            Variables: []stepbuilder.DocVariable{
                {Name: "name", Description: "The logical name of the element.", Type: stepbuilder.VarTypeString},
            },
            Example:  "Then the submit button should be visible",
            Category: stepbuilder.Visual,
        },
    )
}
```

### 5. Gherkin Parser (`pkg/gherkinparser/`)

Advanced Gherkin parsing with macro support and feature processing.

**Key Features:**

- **Macro Processing**: Reusable scenario definitions
- **Feature Separation**: Macro vs test feature distinction
- **Parallel Processing**: Concurrent feature parsing
- **Error Handling**: Graceful parsing error recovery

**Macro System:**

Advanced Gherkin parsing with macro support for reusable test scenarios.

**Key Features:**

- **Macro Processing**: Reusable scenario definitions with direct step substitution
- **Feature Separation**: Macro vs test feature distinction using `@macro` tags
- **Parallel Processing**: Concurrent feature parsing and macro application
- **Error Handling**: Graceful parsing error recovery and circular dependency detection

**Macro Definition:**

```gherkin
# login.feature (or any feature file)
@macro
Scenario: Login with credentials
  Given the user is on the homepage
  When the user goes to the "login" page
  And the user enters "test@example.com" into the "email" field
  And the user enters "password123" into the "password" field
  And the user clicks on the "login" button

@macro
Scenario: Logout user
  Given the user is logged in
  When the user clicks on the "logout" button
  Then the user should be navigated to "login" page
```

**Note:** The `@macro` tag identifies macro scenarios, not the file name. Macros are static groups of steps that get substituted directly.

**Macro Usage:**

```gherkin
# test.feature
Scenario: Admin user login
  Given Login with credentials
  Then the user should be navigated to "dashboard" page

Scenario: Regular user login
  Given Login with credentials
  Then the user should be navigated to "dashboard" page
```

**Implementation Details:**

```go
// Macro processing workflow
func applyMacros(macros []*scenario, featuresContainingMacros []*Feature) {
    macroTitles := getMacroTitles(macros)
    mustContainsMacro := regexp.MustCompile(strings.Join(macroTitles, "|"))

    for _, f := range featuresContainingMacros {
        // Process each feature file for macro references
        if f.background != nil {
            applyMacro(f.background.Steps, macroTitles, macros, featureContent)
        }

        for _, sc := range f.scenarios {
            applyMacro(sc.Steps, macroTitles, macros, featureContent)
        }
    }
}
```

### 6. Reporting System (`pkg/reporters/`)

Multi-format test reporting with detailed execution information.

**Supported Formats:**

- **HTML**: Rich interactive reports with screenshots
- **JSON**: Machine-readable structured data
- **JUnit**: CI/CD integration format

**Report Structure:**

```go
type Report struct {
    scenarios         []Scenario
    startDate         time.Time
    formatter         formatter
    AreAllTestsPassed bool
}

type Scenario struct {
    Title    string
    Result   scenarioResult
    Duration time.Duration
    Error    error
}
```

## 🔄 Execution Flow

### 1. Application Startup

```
main() → parseArgs() → loadConfig() → validateConfig() → executeAction()
```

### 2. Test Execution Flow

```
run() → parseGherkin() → processMacros() → executeScenarios() → generateReport()
```

### 3. Scenario Execution

```
scenario → setupContext() → executeSteps() → teardownContext() → recordResult()
```

### 4. Step Execution

```
step → validateStep() → executeStep() → handleError() → updateContext()
```

## 🎯 Design Patterns

### 1. Dependency Injection

Configuration and dependencies are injected through interfaces, enabling easy testing and modularity.

### 2. Strategy Pattern

Multiple selector strategies for element detection with fallback mechanisms.

### 3. Factory Pattern

Browser instances, step definitions, and report formatters are created through factory methods.

### 4. Command Pattern

Different execution modes (run, init, validate) are implemented as separate commands.

### 5. Observer Pattern

Logging and reporting systems observe test execution events.

### 6. Template Method Pattern

Step definitions follow a common template with customizable behavior.

## 🔒 Error Handling

### 1. Graceful Degradation

- Element not found: Try alternative selectors
- Configuration errors: Use defaults where possible
- Network issues: Retry with exponential backoff

### 2. Comprehensive Logging

- Structured logging with context
- Different log levels (info, warn, error, success)
- Error categorization and suggestions

### 3. Screenshot Capture

- Automatic screenshots on test failures
- Configurable screenshot behavior
- Screenshot integration in reports

## 🚀 Performance Optimizations

### 1. Parallel Execution

- Concurrent test scenario execution
- Parallel element selector evaluation
- Multi-threaded report generation

### 2. Resource Management

- Browser instance pooling
- Memory-efficient element handling
- Automatic cleanup of resources

### 3. Caching

- Configuration caching
- Element selector caching
- Page state caching

## 🔧 Extensibility

### 1. Custom Step Definitions

```go
func (steps) customStep() stepbuilder.Step {
    return stepbuilder.NewWithOneVariable(
        []string{`^custom step with {string}$`},
        func(ctx context.Context, param string) (context.Context, error) {
            // Custom implementation
            return ctx, nil
        },
        nil,
        stepbuilder.DocParams{
            Description: "Custom step description",
            Variables: []stepbuilder.DocVariable{
                {Name: "param", Description: "Parameter description", Type: stepbuilder.VarTypeString},
            },
            Example:  "When custom step with \"value\"",
            Category: stepbuilder.Custom,
        },
    )
}
```

### 2. Custom Report Formats

```go
type customFormatter struct{}

func (f customFormatter) WriteReport(details TestSuiteDetails) error {
    // Custom report generation
    return nil
}
```

### 3. Custom Browser Implementation

```go
type CustomBrowser struct{}

func (b *CustomBrowser) NewPage() (common.Page, error) {
    // Custom page implementation
    return &CustomPage{}, nil
}
```

## 🧪 Testing Strategy

### 1. Unit Testing

- Individual component testing
- Mock interfaces for isolation
- High test coverage requirements

### 2. Integration Testing

- Component interaction testing
- End-to-end workflow validation
- Configuration integration tests

### 3. End-to-End Testing

- Complete test execution flows
- Real browser automation testing
- Cross-platform compatibility testing

## 🔄 CI/CD Integration

### 1. Build Pipeline

- Multi-platform builds
- Automated testing
- Code quality checks

### 2. Release Process

- Semantic versioning
- Automated release notes
- Binary distribution

### 3. Documentation

- Automated documentation generation
- API documentation updates
- Example generation

## 📊 Monitoring and Observability

### 1. Metrics Collection

- Test execution metrics
- Performance monitoring
- Error rate tracking

### 2. Logging Strategy

- Structured logging
- Log level management
- Log aggregation

### 3. Health Checks

- Configuration validation
- Browser connectivity
- Resource availability

## 🔮 Future Architecture Considerations

### 1. Microservices Architecture

- Separate services for different testing domains
- API-first design for integration

### 2. AI/ML Integration

- Intelligent test generation
- Anomaly detection
- Predictive maintenance

---

This architecture document provides a comprehensive overview of TestFlowKit's technical design. For implementation details, refer to the source code and API documentation.
