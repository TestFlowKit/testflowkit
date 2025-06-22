# TestFlowKit Enhanced Configuration System

## Overview

The Enhanced Configuration System for TestFlowKit provides a comprehensive, environment-aware configuration management solution that supports multi-environment deployments, organized element selectors with fallback strategies, and comprehensive API testing capabilities.

## Key Features

### 🌍 **Multi-Environment Support**
- Configure multiple environments (local, staging, production)
- Environment-specific URLs for frontend and backend
- Runtime environment switching via `TEST_ENV` environment variable

### 🎯 **Organized Element Selectors**
- Page-aware element organization
- Fallback selector strategies (data-testid → id → class → generic)
- Shared common elements across pages

### 🔧 **Comprehensive API Testing**
- Detailed endpoint definitions with method, path, and description
- Environment variable substitution for sensitive data
- Default headers applied to all requests
- Path parameter support for dynamic endpoints

### ⚙️ **Advanced Settings Management**
- Global timeout configurations
- Screenshot and video recording options
- Concurrency control
- Environment variable overrides

## Configuration File Structure

The enhanced configuration uses YAML format with the following structure:

```yaml
# ========= GENERAL SETTINGS =========
active_environment: "local"

settings:
  default_timeout: 10000
  page_load_timeout: 30000
  screenshot_on_failure: true
  video_recording: false
  concurrency: 1
  headless: false
  slow_motion: "100ms"
  report_format: "html"
  gherkin_location: "./e2e/features"
  tags: ""

# ========= ENVIRONMENTS =========
environments:
  local:
    frontend_base_url: "http://localhost:3000"
    backend_base_url: "http://localhost:8080/api"
  staging:
    frontend_base_url: "https://staging.example.com"
    backend_base_url: "https://api.staging.example.com"

# ========= FRONTEND CONFIGURATION =========
frontend:
  elements:
    common:
      loading_spinner:
        - "[data-testid='loading-spinner']"
        - ".spinner"
        - ".loading"
    login_page:
      email_field:
        - "[data-testid='email-input']"
        - "input[name='email']"
        - "#email"
  pages:
    home: "/"
    login: "/login"

# ========= BACKEND CONFIGURATION =========
backend:
  default_headers:
    Content-Type: "application/json"
    Authorization: "Bearer ${API_TOKEN}"
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user information by ID"
```

## Getting Started

### 1. Create Enhanced Configuration

Create a `config-enhanced.yml` file in your project root:

```bash
# Copy the example configuration
cp config-enhanced.yml.example config-enhanced.yml

# Edit with your specific settings
nano config-enhanced.yml
```

### 2. Run with Enhanced Configuration

```bash
# Use enhanced configuration system
./testflowkit -enhanced

# Specify custom config file
./testflowkit -enhanced -config=path/to/config.yml

# Show configuration summary
./testflowkit -enhanced -show-config

# Run with specific environment
TEST_ENV=staging ./testflowkit -enhanced
```

### 3. Environment Variable Overrides

Override configuration at runtime:

```bash
# Override environment
export TEST_ENV=staging

# Override browser mode
export TEST_HEADLESS=true

# Override concurrency
export TEST_CONCURRENCY=4

# Override test tags
export TEST_TAGS="@smoke"

# Set API token for authentication
export API_TOKEN="your-api-token-here"
```

## Configuration Sections

### Global Settings

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `default_timeout` | int | 10000 | Default timeout for element waits (ms) |
| `page_load_timeout` | int | 30000 | Maximum page load timeout (ms) |
| `screenshot_on_failure` | bool | true | Take screenshots on test failures |
| `video_recording` | bool | false | Enable video recording |
| `concurrency` | int | 1 | Number of parallel test executions |
| `headless` | bool | false | Run browser in headless mode |
| `slow_motion` | string | "100ms" | Delay between actions for debugging |
| `report_format` | string | "html" | Report format: html, json, junit |
| `gherkin_location` | string | "./e2e/features" | Feature files directory |
| `tags` | string | "" | Filter scenarios by tags |

### Environment Configuration

Each environment must define:

- `frontend_base_url`: Base URL for the web application
- `backend_base_url`: Base URL for API endpoints

### Frontend Element Selectors

Element selectors follow a hierarchical fallback strategy:

1. **Page-specific selectors** (most specific)
2. **Common selectors** (shared across pages)
3. **Generic selectors** (fallback options)

```yaml
frontend:
  elements:
    # Common elements available on all pages
    common:
      logout_button:
        - "[data-testid='logout-btn']"
        - "button[aria-label='Logout']"
        - ".logout-button"
    
    # Page-specific elements
    login_page:
      email_field:
        - "[data-testid='email-input']"  # Preferred: data-testid
        - "input[name='email']"          # Fallback: name attribute
        - "#email"                       # Fallback: ID
        - "input[type='email']"          # Generic: type
```

### Backend API Configuration

API endpoints support:

- **Method specification** (GET, POST, PUT, DELETE, etc.)
- **Path parameters** using `{parameter}` syntax
- **Self-documenting descriptions**
- **Default headers** applied to all requests

```yaml
backend:
  default_headers:
    Content-Type: "application/json"
    Authorization: "Bearer ${API_TOKEN}"
  
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user information by ID"
    
    create_user:
      method: "POST"
      path: "/users"
      description: "Create a new user account"
```

## Step Definitions Usage

### Enhanced Navigation Steps

```gherkin
# Navigate to configured pages
When the user goes to the "login" page
When the user goes to the "dashboard" page

# Direct URL navigation
When the user navigates to URL "https://example.com/login"

# Page verification
Then the user should be navigated to "dashboard" page
```

### Enhanced Element Interaction

```gherkin
# The framework automatically uses fallback selectors
When the user clicks on the "login_button" element on "login_page"
When the user fills the "email_field" on "login_page" with "user@example.com"

# Wait for elements with configuration-based timeouts
Then the user waits for the "welcome_message" element on "dashboard_page"
```

### Enhanced API Testing Steps

```gherkin
# Prepare requests using configured endpoints
Given I prepare a request for the "get_user" endpoint

# Set path parameters for dynamic endpoints
And I set the path parameter "id" to "123"

# Set custom headers (overrides defaults)
And I set the request header "Authorization" to "Bearer custom-token"

# Set query parameters
And I set the query parameter "include" to "profile"

# Send request with automatic URL resolution
When I send the request

# Verify responses
Then the response status code should be 200
And the response body should contain "success"

# Store response values for later use
And I store the response body path "user.id" as "user_id"
```

## Environment Variable Substitution

The enhanced configuration supports dynamic environment variable substitution using the `${VAR_NAME}` syntax:

```yaml
backend:
  default_headers:
    Authorization: "Bearer ${API_TOKEN}"
    X-Client-ID: "${CLIENT_ID}"
    X-Environment: "${DEPLOYMENT_ENV:-local}"
```

Set environment variables:

```bash
export API_TOKEN="your-secret-token"
export CLIENT_ID="your-client-id"
export DEPLOYMENT_ENV="production"
```

## Migration from Legacy Configuration

### Automatic Conversion

The enhanced system provides automatic conversion from enhanced to legacy configuration for backward compatibility:

```go
// Enhanced configuration is automatically converted
legacyConfig := convertToLegacyConfig(enhancedConfig)
```

### Migration Steps

1. **Create enhanced configuration file**
2. **Test with existing scenarios** using `-enhanced` flag
3. **Gradually migrate step definitions** to use enhanced features
4. **Update CI/CD pipelines** to use enhanced configuration

### Compatibility Matrix

| Feature | Legacy | Enhanced | Notes |
|---------|--------|----------|-------|
| Basic browser automation | ✅ | ✅ | Full compatibility |
| Environment switching | ❌ | ✅ | New feature |
| Fallback selectors | ❌ | ✅ | Enhanced robustness |
| API testing | Basic | ✅ | Comprehensive support |
| Configuration validation | Basic | ✅ | Extensive validation |

## Best Practices

### 1. Element Selector Strategy

**Preferred selector hierarchy:**

1. `[data-testid='element-name']` - Most reliable
2. `[name='element-name']` - Semantic meaning
3. `#element-id` - Unique identifier
4. `.element-class` - Style-based
5. `tagname[attribute='value']` - Generic fallback

```yaml
login_button:
  - "[data-testid='login-btn']"    # Best practice
  - "button[type='submit']"        # Semantic
  - "#login-button"                # ID-based
  - ".btn-primary"                 # Class-based
```

### 2. Environment Management

**Organize by deployment stage:**

```yaml
environments:
  local:        # Development environment
  integration:  # Integration testing
  staging:      # Pre-production
  production:   # Live environment
```

### 3. API Endpoint Organization

**Group by functional area:**

```yaml
endpoints:
  # Authentication
  login: { method: "POST", path: "/auth/login", description: "User login" }
  logout: { method: "POST", path: "/auth/logout", description: "User logout" }
  
  # User management
  get_user: { method: "GET", path: "/users/{id}", description: "Get user by ID" }
  update_user: { method: "PUT", path: "/users/{id}", description: "Update user" }
```

### 4. Security Considerations

**Sensitive data handling:**

```yaml
# ✅ Good: Use environment variables
Authorization: "Bearer ${API_TOKEN}"

# ❌ Bad: Hardcode secrets
Authorization: "Bearer hardcoded-secret-token"
```

### 5. Configuration Validation

**Always validate before execution:**

```bash
# Check configuration validity
./testflowkit -enhanced -show-config

# Validate specific environment
TEST_ENV=staging ./testflowkit -enhanced -show-config
```

## Troubleshooting

### Common Issues

#### 1. Configuration File Not Found

```
Failed to load enhanced configuration from 'config-enhanced.yml'
```

**Solution:** Ensure the configuration file exists and path is correct:

```bash
# Check file exists
ls -la config-enhanced.yml

# Use custom path
./testflowkit -enhanced -config=/path/to/config.yml
```

#### 2. Environment Variable Not Substituted

```
Required environment variables not set: API_TOKEN
```

**Solution:** Set the required environment variable:

```bash
export API_TOKEN="your-token"
./testflowkit -enhanced
```

#### 3. Element Not Found with Fallback Selectors

```
Element 'login_button' not found on page 'login_page' with any selector
```

**Solution:** Verify selectors in browser dev tools:

1. Open browser dev tools (F12)
2. Test each selector in console: `document.querySelector("[data-testid='login-btn']")`
3. Update configuration with working selectors

#### 4. Invalid Environment Configuration

```
Active environment 'production' not found in configuration
```

**Solution:** Ensure environment is defined:

```yaml
environments:
  production:  # Add missing environment
    frontend_base_url: "https://prod.example.com"
    backend_base_url: "https://api.prod.example.com"
```

### Debug Mode

Enable verbose logging for troubleshooting:

```bash
./testflowkit -enhanced -verbose
```

### Configuration Validation

The enhanced configuration system provides comprehensive validation:

```bash
# Validate configuration structure
./testflowkit -enhanced -show-config

# Test specific environment
TEST_ENV=staging ./testflowkit -enhanced -show-config
```

## Advanced Usage

### Custom Step Definitions

Create custom step definitions that leverage the enhanced configuration:

```go
func (st steps) myCustomStep() stepbuilder.Step {
    return stepbuilder.NewWithOneVariable(
        []string{`^I perform a custom action on {string}$`},
        func(ctx *scenario.Context) func(string) error {
            return func(elementName string) error {
                // Get enhanced configuration
                cfg, err := config.GetEnhancedConfig()
                if err != nil {
                    return err
                }
                
                // Use configuration-aware element finding
                selectors := cfg.GetElementSelectors("current_page", elementName)
                // ... implement custom logic
                
                return nil
            }
        },
        nil,
        stepbuilder.DocParams{
            Description: "Performs a custom action using enhanced configuration",
        },
    )
}
```

### Dynamic Configuration Updates

Update configuration at runtime:

```go
// Reset configuration for testing
config.ResetEnhancedConfig()

// Load different configuration
newConfig, err := config.LoadEnhancedConfig("test-config.yml")
```

### Integration with CI/CD

Example GitHub Actions workflow:

```yaml
name: TestFlowKit Enhanced Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        environment: [staging, production]
    
    steps:
      - uses: actions/checkout@v2
      
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Run Enhanced Tests
        env:
          TEST_ENV: ${{ matrix.environment }}
          API_TOKEN: ${{ secrets.API_TOKEN }}
          TEST_HEADLESS: true
          TEST_CONCURRENCY: 2
        run: |
          ./testflowkit -enhanced -verbose
```

## Performance Considerations

### Selector Performance

**Optimize selectors for speed:**

1. **data-testid** - Fastest, most reliable
2. **ID selectors** - Fast, unique
3. **Class selectors** - Moderate speed
4. **Complex CSS** - Slower, avoid if possible

### Concurrent Execution

**Configure concurrency based on resources:**

```yaml
settings:
  concurrency: 4  # Adjust based on available CPU/memory
```

**Environment-specific concurrency:**

```bash
# Local development
TEST_CONCURRENCY=1 ./testflowkit -enhanced

# CI environment
TEST_CONCURRENCY=4 ./testflowkit -enhanced
```

### Memory Management

**Large test suites:**

```yaml
settings:
  video_recording: false      # Disable for CI
  screenshot_on_failure: true # Keep for debugging
```

## Support and Community

### Documentation

- [Configuration Reference](./configuration-reference.md)
- [Step Definitions Guide](./step-definitions.md)
- [API Testing Guide](./api-testing.md)

### Contributing

The enhanced configuration system follows Go best practices:

1. **Explicit dependencies** - No hidden globals
2. **Interface segregation** - Small, focused interfaces
3. **Comprehensive validation** - Fail fast with clear errors
4. **Extensive documentation** - Every public function documented

### Feedback

Report issues or suggest improvements:

- GitHub Issues: [testflowkit/issues](https://github.com/testflowkit/testflowkit/issues)
- Discussions: [testflowkit/discussions](https://github.com/testflowkit/testflowkit/discussions) 