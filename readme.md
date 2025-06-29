# TestFlowKit

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)](https://github.com/TestFlowKit/testflowkit/actions)

**TestFlowKit** is a powerful, open-source web test automation framework built in Go that simplifies the creation and execution of automated tests for web applications. It uses Gherkin syntax for test specification and provides comprehensive support for both frontend and backend testing.

## 🚀 Features

### Core Capabilities

- **Gherkin Syntax**: Write tests using clear, readable BDD syntax
- **Multi-Environment Support**: Configure and switch between different environments (local, staging, production)
- **Frontend Testing**: Comprehensive web UI automation with smart element detection
- **Backend API Testing**: Full REST API testing capabilities with request/response validation
- **Macro System**: Reusable test scenarios to reduce code duplication
- **Parallel Execution**: Run tests concurrently for faster execution
- **Rich Reporting**: HTML and JSON report formats with detailed test results

### Advanced Features

- **Smart Element Detection**: Multiple selector strategies with fallback mechanisms
- **Screenshot on Failure**: Automatic screenshot capture for failed tests
- **Headless Mode**: Run tests without browser UI for CI/CD environments
- **Slow Motion Mode**: Debug-friendly execution with configurable delays
- **Cross-Browser Support**: Chrome-based automation with Rod browser engine
- **Configuration Management**: YAML-based configuration with environment-specific settings

## 📋 Prerequisites

- **Go 1.23+**: [Download and install Go](https://golang.org/dl/)
- **Git**: For cloning the repository
- **Make**: For build automation (optional but recommended)

## 🛠️ Installation

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

### From Pre-built Binaries

Download the latest release from [GitHub Releases](https://github.com/TestFlowKit/testflowkit/releases) for your platform.

## 🚀 Quick Start

### 1. Initialize Project

```bash
# Initialize a new TestFlowKit project
./testflowkit init
```

### 2. Configure Your Application

Edit the generated `config.yml` file with your application details:

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
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user by ID"
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

### 4. Run Tests

```bash
# Run all tests
./testflowkit run

# Run specific tags
./testflowkit run --tags "@smoke"

# Run with specific configuration
./testflowkit run --config ./custom-config.yml
```

## 📚 Documentation

For comprehensive documentation, visit the [official TestFlowKit documentation](https://testflowkit.dreamsfollowers.me/).

### Key Documentation Sections

- [Getting Started](https://testflowkit.dreamsfollowers.me/get-started)
- [Configuration Guide](https://testflowkit.dreamsfollowers.me/configuration)
- [Step Definitions](https://testflowkit.dreamsfollowers.me/sentences)
- [Test Execution Design (TED)](https://testflowkit.dreamsfollowers.me/docs/category/test-execution-design-ted)

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

- **Timeouts**: Element wait times and page load timeouts
- **Execution**: Concurrency, headless mode, screenshot settings
- **Reporting**: Output formats and locations

### Environments

- **Frontend URLs**: Base URLs for different environments
- **API URLs**: Backend service endpoints
- **Environment-specific settings**

### Frontend Configuration

- **Element Selectors**: CSS selectors, XPath, and data attributes
- **Page Definitions**: Logical page names and their URLs
- **Fallback Strategies**: Multiple selector options for robust element detection

### Backend Configuration

- **API Endpoints**: REST API definitions with methods and paths
- **Default Headers**: Common HTTP headers for API requests
- **Authentication**: API authentication configuration

## 🧪 Writing Tests

### Frontend Testing

TestFlowKit provides comprehensive frontend testing capabilities:

```gherkin
# Navigation
Given the user is on the homepage
When the user goes to the "login" page

# Form Interactions
And the user enters "test@example.com" into the "email" field
And the user selects "Option 1" from the "dropdown" dropdown
And the user checks the "remember_me" checkbox

# Assertions
Then the "welcome_message" should be visible
And the "email" field should contain "test@example.com"
```

### API Testing

Full REST API testing with request/response validation:

```gherkin
# API Requests
Given I prepare a request for the "get_user" endpoint
When I set the following path params:
  | id | 1 |
And I send the request

# Response Validation
Then the response status code should be 200
And the response body should contain "userId"
And the response body path "id" should exist
```

### Macros

Create reusable test scenarios:

```gherkin
# In login.macro.feature
Scenario: Login with credentials
  Given the user is on the homepage
  When the user goes to the "login" page
  And the user enters "{email}" into the "email" field
  And the user enters "{password}" into the "password" field
  And the user clicks on the "login" button

# In test.feature
Scenario: Test with macro
  Given Login with credentials
    | email    | password   |
    | user@test| pass123    |
```

## 🚀 Advanced Usage

### Parallel Execution

```bash
# Run tests with 4 parallel workers
./testflowkit run --concurrency 4
```

### Environment-Specific Execution

```bash
# Run tests against staging environment
TEST_ENV=staging ./testflowkit run
```

### Custom Configuration

```bash
# Use custom configuration file
./testflowkit run --config ./custom-config.yml
```

### Tag-Based Execution

```bash
# Run only smoke tests
./testflowkit run --tags "@smoke"

# Exclude slow tests
./testflowkit run --tags "~@slow"
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

## 🐛 Bug Reports

If you encounter a bug, please [create a GitHub issue](https://github.com/TestFlowKit/testflowkit/issues) with:

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

- **Stéphane Salou** - [stephsalou](https://github.com/stephsalou)
- [godog](https://github.com/cucumber/godog) - Gherkin parsing and execution
- [go-rod](https://github.com/go-rod/rod) - Browser automation
- [go-yaml](https://github.com/goccy/go-yaml) - YAML configuration handling

## 📊 Project Status

- **Version**: 1.0.0
- **Status**: Active Development
- **Go Version**: 1.23+
- **License**: MIT

## 🔗 Links

- **Documentation**: [https://testflowkit.dreamsfollowers.me/](https://testflowkit.dreamsfollowers.me/)
- **GitHub**: [https://github.com/TestFlowKit/testflowkit](https://github.com/TestFlowKit/testflowkit)
- **Issues**: [https://github.com/TestFlowKit/testflowkit/issues](https://github.com/TestFlowKit/testflowkit/issues)
- **Releases**: [https://github.com/TestFlowKit/testflowkit/releases](https://github.com/TestFlowKit/testflowkit/releases)

---

**TestFlowKit** - Simplifying web test automation with the power of Go and Gherkin syntax.
