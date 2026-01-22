# TestFlowKit

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)](https://github.com/TestFlowKit/testflowkit/actions)
[![GitHub release](https://img.shields.io/github/v/release/TestFlowKit/testflowkit)](https://github.com/TestFlowKit/testflowkit/releases)

**TestFlowKit** is a powerful, open-source web test automation framework built in Go that simplifies the creation and execution of automated tests for web applications. It uses Gherkin syntax for test specification and provides comprehensive support for both frontend and backend testing.

## 🚀 Features

### Core Capabilities

- **Gherkin Syntax**: Write tests using clear, readable BDD syntax
- **Environment Variables**: Flexible environment configuration using env variables and external .env files
- **Frontend Testing**: Comprehensive web UI automation with smart element detection
- **Backend API Testing**: Full REST API testing capabilities with request/response validation
- **GraphQL Support**: Complete GraphQL testing with queries, mutations, schema validation
- **Macro System**: Reusable test scenarios to reduce code duplication
- **Global Hooks**: Setup and teardown logic with `@BeforeAll` and `@AfterAll` tags
- **Parallel Execution**: Run tests concurrently for faster execution
- **Rich Reporting**: HTML and JSON report formats with detailed test results
- **XPath Support**: Full XPath 1.0 support alongside CSS selectors for flexible element selection

### Advanced Features

- **Smart Element Detection**: Multiple selector strategies with fallback mechanisms
- **Screenshot on Failure**: Automatic screenshot capture for failed tests
- **Headless Mode**: Run tests without browser UI for CI/CD environments
- **Slow Motion Mode**: Debug-friendly execution with configurable delays
- **Cross-Browser Support**: Chrome-based automation with Rod browser engine
- **Configuration Management**: YAML-based configuration with environment-specific settings
- **Variable System**: Dynamic data storage and substitution throughout test scenarios

## 📋 Prerequisites

- **Go 1.23+**: [Download and install Go](https://golang.org/dl/)
- **Git**: For cloning the repository
- **Make**: For build automation (optional but recommended)

## 🛠️ Installation

### Using npm (Recommended)

The easiest way to install TestFlowKit is via npm:

```bash
# Global installation
npm install -g @testflowkit/cli

# Verify installation
tkit --version
```

Or use it directly with npx:

```bash
npx @testflowkit/cli --version
```

### From Pre-built Binaries

Download the latest release from [GitHub Releases](https://github.com/TestFlowKit/testflowkit/releases) for your platform.

### From Source

```bash
# Clone the repository
git clone https://github.com/TestFlowKit/testflowkit.git
cd testflowkit

# Install dependencies
go mod tidy

# Build the application
make build GOOS=linux GOARCH=amd64  # or your target platform
```

## 🚀 Quick Start

### 1. Initialize Project

```bash
# Initialize a new TestFlowKit project
./tkit init
```

### 2. Configure Your Application

Edit the generated `config.yml` file with your application details:

```yaml
settings:
  concurrency: 1
  think_time: 1000
  report_format: "html"
  gherkin_location: "./e2e/features"
  env_file: ".env.yml"

env:
  frontend_base_url: "http://localhost:3000"
  api_base_url: "http://localhost:8080/api"

frontend:
  # Element search timeout in milliseconds (1-300000ms)
  # Maximum time to wait when searching for elements by CSS selectors or XPath
  default_timeout: 10000
  headless: false
  screenshot_on_failure: true

  elements:
    login_page:
      email_field:
        - "[data-testid='email-input']"
        - "input[name='email']"
      password_field:
        - "[data-testid='password-input']"
        - "input[name='password']"
      login_button:
        - "[data-testid='login-btn']"
        - "button[type='submit']"

  pages:
    login: "/login"
    dashboard: "/dashboard"

backend:
  # REST API endpoints
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user by ID"
  
  # GraphQL configuration
  graphql:
    endpoint: "/graphql"
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

files:
  base_directory: "./"
  definitions:
    profile_image: "images/profile.jpg"
    test_document: "documents/test.pdf"
    sample_csv: "data/sample.csv"
```

### 3. Write Your First Test

Create a feature file at `e2e/features/login.feature`:

```gherkin
Feature: User Authentication
  As a user
  I want to log into the application
  So that I can access my account

  Scenario: Successful login with valid credentials
    Given the user is on the homepage
    When the user goes to the "login" page
    And the user enters "test@example.com" into the "email" field
    And the user enters "password123" into the "password" field
    And the user clicks on the "login" button
    Then the user should be navigated to "dashboard" page
    And the "welcome_message" should be visible
```

### 4. Using Variables for Dynamic Testing

TestFlowKit provides powerful variable support for dynamic test data:

```gherkin
Feature: Dynamic User Testing
  As a user
  I want to test with dynamic data
  So that my tests are more flexible

  Scenario: Test with API data and variables
    Given I prepare a request for the "get_user" endpoint
    And I set the following path parameters:
      | id | 123 |
    When I send the request
    And I store the JSON path "data.name" from the response into "user_name" variable
    And I store the JSON path "data.email" from the response into "user_email" variable
    Then the response status code should be 200

    When the user goes to the "profile" page
    And the user enters "{{user_name}}" into the "name" field
    And the user enters "{{user_email}}" into the "email" field
    And the user uploads the "profile_image" file into the "avatar" field
    Then the "name" field should contain "{{user_name}}"
    And the "avatar" field should contain the uploaded file
```

### 5. Run Tests

```bash
# Run all tests
./tkit run

# Run specific tags
./tkit run --tags "@smoke"

# Run with specific configuration
./tkit run --config ./custom-config.yml
```

## 📚 Documentation

For comprehensive documentation, visit the [official TestFlowKit documentation](https://testflowkit.github.io/testflowkit/).

### Key Documentation Sections

- [Getting Started](https://testflowkit.github.io/testflowkit/get-started)
- [Configuration Guide](https://testflowkit.github.io/testflowkit/configuration)
- [Step Definitions](https://testflowkit.github.io/testflowkit/sentences)
- [Variables System](https://testflowkit.github.io/testflowkit/variables)
- [FAQ & Troubleshooting](https://testflowkit.github.io/testflowkit/troubleshooting)
- [Test Execution Design (TED)](https://testflowkit.github.io/testflowkit/docs/category/test-execution-design-ted)

## 🏗️ Project Structure

> 📖 **For detailed technical architecture and design patterns, see [ARCHITECTURE.md](architecture.md)**

TestFlowKit follows a clean, modular architecture with clear separation of concerns:

```
testflowkit/
├── cmd/testflowkit/          # Application entry point
├── internal/                 # Core application logic
│   ├── actions/             # Test execution, initialization, validation
│   ├── browser/             # Browser automation layer
│   ├── config/              # Configuration management
│   └── steps_definitions/   # Gherkin step implementations
├── pkg/                     # Reusable packages
│   ├── gherkinparser/       # Gherkin parsing and macro processing
│   ├── logger/              # Logging utilities
│   ├── reporters/           # Test report generation
│   └── utils/               # General utilities
├── e2e/                     # End-to-end test examples
├── documentation/           # Project documentation website
└── scripts/                 # Build and utility scripts
```

**Key Architectural Highlights:**

- **Modular Design**: Clear separation between application logic, domain models, and infrastructure
- **Interface-Based**: Browser automation abstracted through interfaces for testability
- **Configuration-Driven**: YAML-based configuration with environment support
- **Extensible**: Plugin-like step definition system with validation and documentation

## 🔧 Configuration

TestFlowKit uses YAML configuration files to define test environments, element selectors, and execution settings. Key configuration sections include:

### Global Settings

- **Element Search Timeout**: Maximum time to wait when searching for elements by CSS selectors or XPath expressions
- **Execution**: Concurrency, headless mode, screenshot settings
- **Reporting**: Output formats and locations

### Environment Variables

- **env block**: Define environment variables directly in config.yml
- **env_file**: Load variables from external .env.yml file
- **Variable Access**: Use `{{ env.VARIABLE_NAME }}` syntax in tests and configuration
- **CLI Override**: Override variables using `--env-file` option

### Frontend Configuration

- **Element Selectors**: CSS selectors, XPath expressions, and data attributes
- **Page Definitions**: Logical page names and their URLs
- **Fallback Strategies**: Multiple selector options for robust element detection
- **XPath Support**: Full XPath 1.0 support with `xpath:` prefix for complex element selection

### Backend Configuration

- **REST API Endpoints**: REST API definitions with methods and paths
- **GraphQL Operations**: GraphQL queries and mutations with comprehensive variable support
- **Array Variables**: Full support for GraphQL array variables including strings, numbers, and complex objects
- **Default Headers**: Common HTTP headers for API requests
- **Authentication**: API authentication configuration
- **Variable Parsing**: Intelligent parsing of JSON objects, arrays, and primitive types

### File Configuration

- **File Directory**: Base directory where test files are stored
- **File Definitions**: Logical file names mapped to actual file paths
- **Support for Subdirectories**: Organized file structure with nested folders

## 🧪 Writing Tests

### Frontend Testing

TestFlowKit provides comprehensive frontend testing capabilities with support for both CSS selectors and XPath expressions:

```gherkin
# Navigation
Given the user is on the homepage
When the user goes to the "login" page

# Note: Browser auto-initialization
# TestFlowKit automatically opens a browser if you forget to initialize it
# before navigating to a page. This makes your tests more robust and
# reduces the need for explicit browser setup steps.

# Form Interactions
And the user enters "test@example.com" into the "email" field
And the user selects "Option 1" from the "dropdown" dropdown
And the user checks the "remember_me" checkbox
And the user uploads the "profile_image" file into the "avatar" field
And the user uploads the "image1, image2, image3" files into the "gallery" field

# Assertions
Then the "welcome_message" should be visible
And the "email" field should contain "test@example.com"
```

### Element Selector Configuration

TestFlowKit supports multiple selector types for robust element detection:

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
        - "[data-testid='login-button']"

      # Mixed selectors with fallback
      dynamic_element:
        - "xpath://div[contains(@class, 'dynamic') and contains(text(), 'Loading')]"
        - ".loading-indicator"
        - "[data-testid='loading']"
```

**Selector Types:**

- **CSS Selectors**: Standard CSS selectors (default)
  - `#id`, `.class`, `[attribute=value]`, etc.
- **XPath Selectors**: Full XPath 1.0 support with `xpath:` prefix
  - `xpath://div[@class='container']`
  - `xpath://button[contains(text(), 'Submit')]`
  - `xpath://input[@type='email' and @required]`

**Smart Element Detection:**

TestFlowKit automatically tries multiple selectors in parallel and uses the first one that finds an element, providing robust element detection even when page structures change.

### File Upload Testing

TestFlowKit provides comprehensive file upload testing for both frontend forms and REST APIs:

#### Frontend File Upload

Upload files to web forms using logical file names defined in configuration:

```gherkin
# Single file upload
When the user uploads the "profile_image" file into the "avatar" field

# Multiple file upload
When the user uploads the "image1, image2, image3" files into the "gallery" field
```

#### File Configuration

Configure files in your `config.yml`:

```yaml
files:
  base_directory: "test"
  definitions:
    profile_image: "images/profile.jpg"
    avatar_image: "images/avatar.png"
    gallery_image1: "images/gallery/image1.jpg"
    gallery_image2: "images/gallery/image2.jpg"
    gallery_image3: "images/gallery/image3.jpg"
    test_document: "documents/test.pdf"
    sample_csv: "data/sample.csv"
```

**File Configuration Features:**

- **Logical Names**: Use descriptive names instead of file paths in tests
- **Subdirectory Support**: Organize files in nested folders
- **Multiple File Types**: Support for images, documents, data files, etc.
- **Path Resolution**: Automatic path resolution relative to file directory

### API Testing

#### REST API Testing

Full REST API testing with request/response validation:

```gherkin
# REST API Requests
Given I prepare a request for the "get_user" endpoint
When I set the following path parameters:
  | id | 1 |
And I send the request

# Response Validation
Then the response status code should be 200
And the response should have field "userId"
And the response should have field "id"
```

#### GraphQL Testing

Comprehensive GraphQL testing with queries, mutations, array variables, and schema validation:

```gherkin
# GraphQL Query with Variables
Given I prepare a GraphQL request to "get_user_profile"
And I set the following GraphQL variables:
  | userId | 123 |
When I send the request
Then the GraphQL response should not contain errors
And the response should have field "user.name"
And I store the GraphQL data at path "user.id" into "userId" variable

# GraphQL with Array Variables
Given I prepare a GraphQL request to "search_users"
And I set the following GraphQL variables:
  | tags     | ["frontend", "testing", "automation"] |
  | statuses | ["active", "verified"]               |
  | limit    | 10                                   |
When I send the request
Then the GraphQL response should not contain errors
And I store the GraphQL array at path "users" into "foundUsers" variable

# GraphQL with Individual Array Variables
Given I prepare a GraphQL request to "update_user_tags"
And I set the GraphQL variable "userId" to "123"
And I set the GraphQL variable "tags" to array ["frontend", "backend", "testing"]
When I send the request
Then the GraphQL response should not contain errors
And the response should have field "updateUser.tags"

# GraphQL Mutations with Complex Input
Given I prepare a GraphQL request to "create_post"
And I set the following GraphQL variables:
  | input | {"title": "New Post", "content": "Post content", "tags": ["tech", "tutorial"]} |
When I send the request
Then the GraphQL response should not contain errors
And the response should have field "createPost.post.id"

# GraphQL Error Handling
Given I prepare a GraphQL request to "get_user_profile"
And I set the following GraphQL variables:
  | userId | invalid_id |
When I send the request
Then the GraphQL response should contain errors

# GraphQL with Custom Headers
Given I prepare a GraphQL request to "get_user_profile"
And I set the following GraphQL headers:
  | Authorization | Bearer token123 |
  | X-Client-ID   | test-client     |
And I set the following GraphQL variables:
  | userId | 123 |
When I send the request
Then the GraphQL response should not contain errors
```

#### GraphQL Testing

Comprehensive GraphQL testing with queries, mutations, array variables, and schema validation:

```gherkin
# GraphQL Query with Variables
Given I prepare a GraphQL request for the "get_user_profile" operation
And I set the following GraphQL variables:
  | userId | 123 |
When I send the GraphQL request
Then the GraphQL response should not contain errors
And the GraphQL response should contain data at path "user.name"
And I store the GraphQL data at path "user.id" into "userId" variable

# GraphQL with Array Variables
Given I prepare a GraphQL request for the "search_users" operation
And I set the following GraphQL variables:
  | tags     | ["frontend", "testing", "automation"] |
  | statuses | ["active", "verified"]               |
  | limit    | 10                                   |
When I send the GraphQL request
Then the GraphQL response should not contain errors
And I store the GraphQL array at path "users" into "foundUsers" variable

# GraphQL with Individual Array Variables
Given I prepare a GraphQL request for the "update_user_tags" operation
And I set the GraphQL variable "userId" to "123"
And I set the GraphQL variable "tags" to array ["frontend", "backend", "testing"]
When I send the GraphQL request
Then the GraphQL response should not contain errors
And the GraphQL response should contain data at path "updateUser.tags"

# GraphQL Mutations with Complex Input
Given I prepare a GraphQL request for the "create_post" operation
And I set the following GraphQL variables:
  | input | {"title": "New Post", "content": "Post content", "tags": ["tech", "tutorial"]} |
When I send the GraphQL request
Then the GraphQL response should not contain errors
And the GraphQL response should contain data at path "createPost.post.id"

# GraphQL Error Handling
Given I prepare a GraphQL request for the "get_user_profile" operation
And I set the following GraphQL variables:
  | userId | invalid_id |
When I send the GraphQL request
Then the GraphQL response should contain errors

# GraphQL with Custom Headers
Given I prepare a GraphQL request for the "get_user_profile" operation
And I set the following GraphQL headers:
  | Authorization | Bearer token123 |
  | X-Client-ID   | test-client     |
And I set the following GraphQL variables:
  | userId | 123 |
When I send the GraphQL request
Then the GraphQL response should not contain errors
```

### Variables System

TestFlowKit provides a powerful variable system for dynamic data management:

#### Environment Variables

Access environment variables defined in config.yml or external .env files using `{{ env.VARIABLE_NAME }}` syntax:

```gherkin
# Access environment variables in tests
Given I prepare a request for the "get_user" endpoint with base URL "{{ env.API_BASE_URL }}"
When the user goes to "{{ env.FRONTEND_BASE_URL }}/login"
And the user enters "{{ env.TEST_USERNAME }}" into the "email" field
```

**Configuration:**

```yaml
# In config.yml
env:
  API_BASE_URL: "http://localhost:8080/api"
  FRONTEND_BASE_URL: "http://localhost:3000"
  TEST_USERNAME: "test@example.com"

# Or reference external file
settings:
  env_file: ".env.yml"
```

#### Variable Syntax

Variables use the `{{variable_name}}` syntax and are automatically substituted in all step parameters:

```gherkin
# Store custom values
When I store the "John Doe" into "user_name" variable
And I store the "test@example.com" into "user_email" variable

# Use variables in other steps
When the user enters "{{user_name}}" into the "name" field
And the user enters "{{user_email}}" into the "email" field
```

#### Variable Types

**Environment Variables**: Access configuration values using `{{ env.VARIABLE_NAME }}`

```gherkin
When the user enters "{{ env.TEST_EMAIL }}" into the "email" field
```

**Custom Variables**: Store any custom value for reuse

```gherkin
When I store the "Active" into "status" variable
```

**JSON Path Variables**: Extract data from API responses

```gherkin
When I store the JSON path "data.user.id" from the response into "user_id" variable
And I store the JSON path "items[0].name" from the response into "first_item" variable
```

**HTML Element Variables**: Capture content from web page elements

```gherkin
When I store the content of "page_title" into "title" variable
And I store the content of "user_name_label" into "displayed_name" variable
```

#### Advanced Variable Usage

Variables can be used in complex scenarios for data-driven testing:

```gherkin
Scenario: End-to-end data flow with variables
  Given I prepare a request for the "get_user" endpoint
  And I set the following path parameters:
    | id | 123 |
  When I send the request
  And I store the JSON path "data.name" from the response into "api_user_name" variable
  And I store the JSON path "data.email" from the response into "api_user_email" variable
  Then the response status code should be 200

  When the user goes to the "profile" page
  And I store the content of "displayed_name" into "page_user_name" variable
  And the user enters "{{api_user_email}}" into the "email" field
  Then the "email" field should contain "{{api_user_email}}"
  And the "page_user_name" should equal "{{api_user_name}}"
```

#### Variable Features

- **Automatic Substitution**: Variables are replaced in strings, tables, and parameters
- **Scope**: Variables persist throughout the entire scenario
- **Type Support**: Supports strings, numbers, booleans, and complex data structures
- **Cross-Step Usage**: Use variables across different step types (API, frontend, assertions)

### Macros with Parameterized Variables

Create reusable, parameterized test scenarios to reduce code duplication and improve maintainability. The macro system supports variable substitution through table definitions.

#### Macro Definition with Variables

```gherkin
# In login.feature
@macro
Scenario: user login with credentials
  Given the user is on the login page
  When the user fills the "username" field with ${username}
  And the user fills the "password" field with ${password}
  And the user clicks the "login" button
  Then the user should be logged in successfully

@macro
Scenario: user logout
  Given the user is logged in
  When the user clicks the "logout" button
  Then the user should be logged out successfully
```

**Key Features:**

- Use `${variable_name}` syntax to define variable placeholders
- Variables are automatically substituted during macro expansion
- Support for multiple variables in a single macro

#### Using Parameterized Macros

```gherkin
# In test.feature
Scenario: Test login with valid credentials
  When user login with credentials
    | username | password |
    | oki     | ler123   |
  Then the user should be logged in successfully

Scenario: Test login with different credentials
  When user login with credentials
    | username | password |
    | admin   | secret   |
  Then the user should be logged in successfully
```

### Global Hooks and Variables

TestFlowKit supports global setup and teardown hooks using `@BeforeAll` and `@AfterAll` tags. These hooks run sequentially before and after the main test suite, allowing you to perform environment setup, data seeding, and cleanup.

#### Global Setup and Teardown 

```gherkin
@BeforeAll
Scenario: Global Setup
  Given I call the API "POST" "/auth/login"
  Then I save the response path "token" as global variable "AUTH_TOKEN"

@AfterAll
Scenario: Global Teardown
  Given I call the API "DELETE" "/cleanup"
  
```

#### Global Variables

Global variables are shared across all scenarios and can be set in the `config.yml` or dynamically during the `@BeforeAll` phase.

**In Scenarios:**

You can access global variables just like local variables using the `{{VARIABLE_NAME}}` syntax.

#### Complex Workflow Macros

```gherkin
@macro
Scenario: complete user workflow
  Given the user is on the login page
  When the user fills the "username" field with ${username}
  And the user fills the "password" field with ${password}
  And the user clicks the "login" button
  Then the user should be logged in successfully
  When the user navigates to the ${page} page
  Then the user should see the ${page} page

# Usage
Scenario: Test complete workflow
  When complete user workflow
    | username | password | page      |
    | oki     | ler123   | dashboard |
  Then the user should see the dashboard page
```

#### Macro Benefits

- **Parameterized Reusability**: Use the same macro with different data sets
- **Table-Driven Testing**: Easy to test multiple scenarios with different parameters
- **Maintainability**: Update logic in one place, affects all usages
- **Better Organization**: Separate macro definitions from test scenarios
- **Dynamic Behavior**: Adapt macros to different test requirements

#### File Organization

```
e2e/features/
├── macros/
│   ├── authentication.feature  # Contains @macro scenarios
│   ├── data-setup.feature      # Contains @macro scenarios
│   └── workflows.feature       # Contains @macro scenarios
└── tests/
    ├── login.feature           # Regular test scenarios
    └── user-management.feature # Regular test scenarios
```

**Note:** The `@macro` tag identifies macro scenarios, not the file name. You can organize macros in any feature file structure you prefer.

#### Variable Substitution

The system automatically replaces `${variable_name}` placeholders with actual values from the data table:

**Macro definition:**

```gherkin
When the user fills the "username" field with ${username}
```

**Macro invocation:**

```gherkin
| username | password |
| oki     | ler123   |
```

**Result:**

```gherkin
When the user fills the "username" field with oki
```

#### Best Practices

- **Naming**: Use descriptive macro names that clearly indicate their purpose
- **Variables**: Keep variables focused and specific to the macro's purpose
- **Documentation**: Use clear, descriptive step text that explains the intent
- **Testing**: Test macros with different variable combinations to ensure reliability

## 🚀 Advanced Usage

### Parallel Execution

```bash
# Run tests with 4 parallel workers
./tkit run --concurrency 4
```

### Environment-Specific Execution

```bash
# Run tests with specific environment file
./tkit run --env-file .env.staging.yml

# Override environment variables from CLI
./tkit run --env-file .env.yml
```

### Custom Configuration

```bash
# Use custom configuration file
./tkit run --config ./custom-config.yml
```

### Tag-Based Execution

```bash
# Run only smoke tests
./tkit run --tags "@smoke"

# Exclude slow tests
./tkit run --tags "~@slow"
```

## 🛠️ Development

### Building from Source

```bash
# Install dependencies
go mod tidy

# Run tests
make test

# Build for all platforms
make releases

# Build for specific platform
make build GOOS=linux GOARCH=amd64
```

### Running Tests

```bash
# Run all tests
make test

# Run with coverage
go test -v -race -coverprofile=coverage.out ./...

# Run end-to-end tests
make run_e2e
```

### Code Quality

```bash
# Run linter
make lint

# Format code
go fmt ./...

# Generate documentation
make generate_doc
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feat/amazing-feature`
3. **Follow the coding standards**:
   - Use conventional commit messages
   - Add tests for new features
   - Update documentation as needed
4. **Commit your changes**: `git commit -m 'feat: add amazing feature'`
5. **Push to the branch**: `git push origin feat/amazing-feature`
6. **Open a Pull Request**

### Branch Naming Convention

- `feat/`: New features
- `fix/`: Bug fixes
- `docs/`: Documentation updates
- `style/`: Code style changes
- `refactor/`: Code refactoring
- `test/`: Test additions or updates
- `chore/`: Maintenance tasks

### Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
type(scope): description

[optional body]

[optional footer]
```

## 🐛 Bug Reports & Troubleshooting

If you encounter issues with TestFlowKit, please check our [Troubleshooting Guide](https://testflowkit.github.io/testflowkit/troubleshooting) first for common solutions.

For bugs not covered in the troubleshooting guide, please [create a GitHub issue](https://github.com/TestFlowKit/testflowkit/issues) with:

- **Description**: Clear description of the problem
- **Steps to Reproduce**: Detailed steps to reproduce the issue
- **Expected vs Actual Behavior**: What you expected vs what happened
- **Environment**: OS, Go version, TestFlowKit version
- **Configuration**: Relevant parts of your config.yml (without sensitive data)
- **Logs**: Error messages and stack traces

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Marc-Henry Nanguy** - _Initial work_ - [marckent04](https://github.com/marckent04)

## 🙏 Acknowledgments

### Contributors

- **Stéphane Salou** - [stephsalou](https://github.com/stephsalou)

### Dependencies

- **[alexflint/go-arg](https://github.com/alexflint/go-arg)** - Command-line argument parsing
- **[cucumber/godog](https://github.com/cucumber/godog)** - Gherkin parsing and BDD test execution framework
- **[fatih/color](https://github.com/fatih/color)** - Colorized terminal output
- **[go-rod/rod](https://github.com/go-rod/rod)** - Chrome DevTools Protocol automation library
- **[goccy/go-yaml](https://github.com/goccy/go-yaml)** - High-performance YAML parser and emitter
- **[gofrs/uuid/v5](https://github.com/gofrs/uuid/v5)** - UUID v5 implementation
- **[stretchr/testify](https://github.com/stretchr/testify)** - Testing utilities and assertions
- **[tdewolff/parse](https://github.com/tdewolff/parse)** - HTML/CSS parsing utilities
- **[tidwall/gjson](https://github.com/tidwall/gjson)** - Fast JSON parser and getter

## 📊 Project Status

- **Status**: Active Development
- **Go Version**: 1.23+
- **License**: MIT

## 🔗 Links

- **Documentation**: [https://testflowkit.github.io/testflowkit/](https://testflowkit.github.io/testflowkit/)
- **GitHub**: [https://github.com/TestFlowKit/testflowkit](https://github.com/TestFlowKit/testflowkit)
- **Issues**: [https://github.com/TestFlowKit/testflowkit/issues](https://github.com/TestFlowKit/testflowkit/issues)
- **Releases**: [https://github.com/TestFlowKit/testflowkit/releases](https://github.com/TestFlowKit/testflowkit/releases)

---

**TestFlowKit** - Simplifying web test automation with the power of Go and Gherkin syntax.
