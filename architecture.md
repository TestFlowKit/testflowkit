# TestFlowKit Architecture

## Overview

TestFlowKit is a modular, extensible web test automation framework built in Go. The architecture follows clean architecture principles with clear separation of concerns, dependency inversion, and high cohesion. This document provides a detailed technical overview of the system's architecture, design patterns, and component interactions.

## 🏗️ High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Application Layer                        │
│  CLI Interface (Run, Init, Validate)                       │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Business Logic Layer                     │
│  Test Execution Engine                                      │
│  • Gherkin Parser  • Step Builder  • Scenario Context      │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Domain Layer                             │
│  Core Domain Models                                         │
│  • Browser Interface  • Config Management  • Reporter       │
│  • GraphQL Client     • HTTP Client       • Variables      │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Infrastructure Layer                     │
│  External Dependencies                                      │
│  • Rod Browser Engine  • HTTP Client  • File System        │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
testflowkit/
├── cmd/testflowkit/                    # Application Entry Point
├── internal/                          # Private application code
│   ├── actions/                       # Application actions
│   ├── browser/                       # Browser automation helpers
│   ├── config/                        # Configuration management
│   ├── step_definitions/              # Gherkin step implementations
│   │   ├── api/                       # API testing utilities
│   │   ├── backend/                   # Backend testing steps (REST & GraphQL)
│   │   ├── core/                      # Core step framework
│   │   ├── frontend/                  # Frontend testing steps
│   │   ├── variables/                 # Variable management steps
│   │   └── helpers/                   # Helper utilities
│   └── utils/                         # Internal utilities
├── pkg/                              # Public packages
│   ├── browser/                      # Browser interface and Rod implementation
│   ├── gherkinparser/                # Gherkin parsing and processing
│   ├── graphql/                      # GraphQL client and utilities
│   ├── logger/                       # Logging system
│   ├── reporters/                    # Test reporting system
│   └── utils/                        # Utility functions
├── e2e/                             # End-to-end test examples
│   ├── features/                    # Gherkin feature files
│   ├── test-files/                  # Test data files
│   └── server/                      # Test server for examples
├── documentation/                   # Project documentation website
├── scripts/                         # Build and utility scripts
├── readme.md                       # Project README
├── architecture.md                 # This architecture document
├── Makefile                        # Build automation
├── go.mod                          # Go module definition
└── go.sum                          # Go module checksums
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
settings:
  concurrency: 1
  think_time: 1000
  report_format: "html"
  gherkin_location: "./e2e/features"
  env_file: ".env.local.yml"  # Optional: default env file

# Inline environment variables (or use external file)
env:
  frontend_base_url: "http://localhost:3000"
  rest_api_base_url: "http://localhost:8080/api"
  graphql_endpoint: "http://localhost:8080/graphql"

frontend:
  default_timeout: 10000
  headless: false
  screenshot_on_failure: true
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

  graphql:
    default_headers:
      Content-Type: "application/json"
    operations:
      get_user_profile:
        type: "query"
        operation: |
          query GetUserProfile($userId: ID!) {
            user(id: $userId) {
              id
              name
              email
              profile {
                avatar
                bio
              }
            }
          }
        description: "Fetch user profile with nested data"

      search_users:
        type: "query"
        operation: |
          query SearchUsers($tags: [String!]!, $statuses: [UserStatus!]!) {
            users(tags: $tags, statuses: $statuses) {
              id
              name
              email
              tags
              status
            }
          }
        description: "Search users by tags and statuses with array parameters"
```

**Note:** The GraphQL endpoint URL is defined in environment variables as `graphql_endpoint`, not in the backend section.

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

### 3. Browser Automation Layer (`pkg/browser/` & `internal/browser/`)

Abstract browser interface with Rod implementation for Chrome-based automation.

**Key Components:**

- **Browser Interface** (`pkg/browser`): Abstract browser operations
- **Rod Implementation** (`pkg/browser/rod`): Concrete implementation using Rod
- **Browser Helpers** (`internal/browser`): High-level browser actions and factory
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

### 7. GraphQL Client Layer (`pkg/graphql/`)

Comprehensive GraphQL client implementation

**Key Components:**

- **GraphQL Client**: HTTP-based GraphQL request execution
- **Operation Validation**: Schema-based operation validation
- **Variable Parsing**: Intelligent parsing of complex variable types

**Design Patterns:**

- **Client Pattern**: Centralized GraphQL request handling
- **Strategy Pattern**: Multiple variable parsing strategies
- **Cache Pattern**: Schema caching for performance optimization
- **Validation Pattern**: Pre-execution operation validation

**GraphQL Client Architecture:**

```go
type Client struct {
    httpClient  *http.Client
    endpoint    string
    headers     map[string]string
}

type Request struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}

type Response struct {
    Data       json.RawMessage        `json:"data,omitempty"`
    Errors     []GraphQLError         `json:"errors,omitempty"`
    Extensions map[string]interface{} `json:"extensions,omitempty"`
}
```

**Variable Type Support:**

The GraphQL client supports comprehensive variable type parsing:

- **Primitive Types**: Strings, numbers, booleans, IDs
- **Array Types**: `["item1", "item2"]`, `[1, 2, 3]`, `[true, false]`
- **Object Types**: `{"key": "value", "nested": {"field": true}}`
- **Mixed Arrays**: `[{"id": 1, "name": "John"}, {"id": 2, "name": "Jane"}]`


### 4. Step Definition Framework (`internal/step_definitions/`)

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
- **GraphQL Context**: GraphQL request/response management with variable support
- **Page Management**: Current page tracking
- **Variable System**: Cross-step data storage and retrieval

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
            Categories: []stepbuilder.StepCategory{stepbuilder.Visual}
        },
    )
}
```

**GraphQL Step Definition Example:**

```go
func (steps) prepareGraphQLRequest() stepbuilder.Step {
    return stepbuilder.NewWithOneVariable(
        []string{`I prepare a request for the {string} operation`},
        func(ctx context.Context, operationName string) (context.Context, error) {
            scenarioCtx := scenario.MustFromContext(ctx)
            cfg := scenarioCtx.GetConfig()

            operation, err := cfg.GetGraphQLOperation(operationName)
            if err != nil {
                return ctx, fmt.Errorf("failed to resolve GraphQL operation '%s': %w", operationName, err)
            }

            scenarioCtx.SetGraphQLOperation(operationName, operation)

            logger.InfoFf("GraphQL request prepared for operation '%s'", operationName)
            return ctx, nil
        },
        nil,
        stepbuilder.DocParams{
            Description: "Prepares a GraphQL request for a configured operation.",
            Variables: []stepbuilder.DocVariable{
                {
                    Name:        "operationName",
                    Description: "The logical operation name as defined in configuration",
                    Type:        stepbuilder.VarTypeString,
                },
            },
            Example:  `Given I prepare a request to "get_user_profile"`,
            Categories: []stepbuilder.StepCategory{stepbuilder.GraphQL}
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
            Categories: []stepbuilder.StepCategory{stepbuilder.Custom}
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
- GraphQL client and variable parsing testing

### 2. Integration Testing

- Component interaction testing
- End-to-end workflow validation
- Configuration integration tests

### 3. End-to-End Testing

- Complete test execution flows
- Real browser automation testing
- Cross-platform compatibility testing
- GraphQL API integration testing

### 4. GraphQL Testing Patterns

- **Query Testing**: Validate GraphQL queries with various variable types
- **Mutation Testing**: Test GraphQL mutations with complex input objects
- **Array Variable Testing**: Comprehensive testing of array variable parsing and execution
- **Error Handling**: Test GraphQL error responses and validation
- **Schema Validation**: Test operation validation against GraphQL schemas
- **Data Flow Testing**: Test data extraction and variable storage from GraphQL responses

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
