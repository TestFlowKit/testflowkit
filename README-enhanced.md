# TestFlowKit Enhanced Configuration System

🚀 **Advanced Web Test Automation with Multi-Environment Support**

The Enhanced Configuration System for TestFlowKit provides enterprise-grade test automation capabilities with environment-aware configuration, robust element selection strategies, and comprehensive API testing support.

## ✨ New Features

### 🌍 **Multi-Environment Support**
```bash
# Switch environments easily
TEST_ENV=staging ./testflowkit -enhanced
TEST_ENV=production ./testflowkit -enhanced
```

### 🎯 **Smart Element Selection**
- **Fallback Strategies**: Automatically tries multiple selectors until one works
- **Page-Aware Organization**: Elements organized by page/component
- **Data-TestID Priority**: Follows testing best practices

### 🔧 **Comprehensive API Testing**
- **Named Endpoints**: Use logical names instead of raw URLs
- **Environment Variable Substitution**: Secure token management
- **Path Parameters**: Dynamic endpoint support
- **Request/Response Validation**: Built-in JSON validation

## 🚀 Quick Start

### 1. Create Enhanced Configuration

```bash
# Create your enhanced configuration file
cp config-enhanced.yml.example config-enhanced.yml
```

Edit `config-enhanced.yml`:

```yaml
active_environment: "local"

environments:
  local:
    frontend_base_url: "http://localhost:3000"
    backend_base_url: "http://localhost:8080/api"
  staging:
    frontend_base_url: "https://staging.yourapp.com"
    backend_base_url: "https://api.staging.yourapp.com"

frontend:
  elements:
    login_page:
      email_field:
        - "[data-testid='email-input']"
        - "input[name='email']"
        - "#email"
      login_button:
        - "[data-testid='login-btn']"
        - "button[type='submit']"

backend:
  default_headers:
    Authorization: "Bearer ${API_TOKEN}"
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Get user by ID"
```

### 2. Run with Enhanced Configuration

```bash
# Basic usage
./testflowkit -enhanced

# With custom configuration
./testflowkit -enhanced -config=my-config.yml

# Show configuration summary
./testflowkit -enhanced -show-config

# With environment override
TEST_ENV=staging ./testflowkit -enhanced
```

### 3. Set Environment Variables

```bash
# Set API token
export API_TOKEN="your-secret-token"

# Override settings
export TEST_HEADLESS=true
export TEST_CONCURRENCY=4
```

## 📝 Enhanced Gherkin Steps

### 🌐 **Frontend Testing**

```gherkin
# Navigate using page names
When the user goes to the "login" page
When the user goes to the "dashboard" page

# Element interaction with fallback selectors
When the user clicks on the "login_button" element on "login_page"
When the user fills the "email_field" on "login_page" with "user@example.com"

# Wait for elements with smart timeouts
Then the user waits for the "welcome_message" element on "dashboard_page"
```

### 🔌 **API Testing**

```gherkin
# Prepare requests using endpoint names
Given I prepare a request for the "get_user" endpoint

# Set path parameters for dynamic URLs
And I set the path parameter "id" to "123"

# Add custom headers
And I set the request header "X-Custom-Header" to "value"

# Add query parameters
And I set the query parameter "include" to "profile"

# Send request with automatic URL resolution
When I send the request

# Verify responses
Then the response status code should be 200
And the response body should contain "success"
And the response header "Content-Type" should be "application/json"

# Store values for later use
And I store the response body path "user.id" as "user_id"
```

### 🔄 **Integration Testing**

```gherkin
# Combine API and frontend testing
Given I prepare a request for the "create_user" endpoint
And I set the request body to:
  """
  {
    "name": "Test User",
    "email": "test@example.com"
  }
  """
When I send the request
Then the response status code should be 201
And I store the response body path "user.id" as "new_user_id"

# Use API data in frontend tests
When the user goes to the "users" page
Then the user should see on the page "Test User"
```

## 🏗️ Configuration Architecture

### **Environment Structure**

```yaml
environments:
  local:
    frontend_base_url: "http://localhost:3000"
    backend_base_url: "http://localhost:8080/api"
  
  integration:
    frontend_base_url: "https://integration.yourapp.com"
    backend_base_url: "https://api.integration.yourapp.com"
  
  staging:
    frontend_base_url: "https://staging.yourapp.com" 
    backend_base_url: "https://api.staging.yourapp.com"
  
  production:
    frontend_base_url: "https://yourapp.com"
    backend_base_url: "https://api.yourapp.com"
```

### **Element Organization**

```yaml
frontend:
  elements:
    # Shared elements across all pages
    common:
      loading_spinner:
        - "[data-testid='loading']"
        - ".spinner"
      error_message:
        - "[data-testid='error']"
        - ".alert-danger"
    
    # Page-specific elements
    login_page:
      email_field:
        - "[data-testid='email-input']"  # Preferred
        - "input[name='email']"          # Fallback
        - "#email"                       # Last resort
    
    dashboard_page:
      welcome_message:
        - "[data-testid='welcome']"
        - ".welcome-text"
        - "h1"
```

### **API Endpoint Configuration**

```yaml
backend:
  default_headers:
    Content-Type: "application/json"
    Accept: "application/json"
    Authorization: "Bearer ${API_TOKEN}"
  
  endpoints:
    # User management
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user by ID"
    
    create_user:
      method: "POST"
      path: "/users"
      description: "Create new user"
    
    # Product management
    get_products:
      method: "GET"
      path: "/products"
      description: "List all products"
    
    get_product:
      method: "GET"
      path: "/products/{id}"
      description: "Get product details"
```

## 🛠️ Advanced Configuration

### **Global Settings**

```yaml
settings:
  # Timeouts (milliseconds)
  default_timeout: 10000
  page_load_timeout: 30000
  
  # Execution settings
  concurrency: 1
  headless: false
  slow_motion: "100ms"
  
  # Reporting
  screenshot_on_failure: true
  video_recording: false
  report_format: "html"
  
  # Test selection
  gherkin_location: "./e2e/features"
  tags: "@smoke,@regression"
```

### **Environment Variable Substitution**

```yaml
backend:
  default_headers:
    Authorization: "Bearer ${API_TOKEN}"
    X-Client-ID: "${CLIENT_ID}"
    X-Environment: "${DEPLOYMENT_ENV:-local}"
```

Set environment variables:

```bash
export API_TOKEN="your-secret-api-token"
export CLIENT_ID="your-client-identifier"
export DEPLOYMENT_ENV="staging"
```

## 🚛 Migration Guide

### **From Legacy to Enhanced**

1. **Create Enhanced Configuration**
   ```bash
   cp config-enhanced.yml.example config-enhanced.yml
   ```

2. **Test Compatibility**
   ```bash
   # Test with existing scenarios
   ./testflowkit -enhanced -verbose
   ```

3. **Gradual Migration**
   - Start with environment configuration
   - Add element selectors gradually
   - Migrate API tests to use endpoint names

4. **Update CI/CD**
   ```yaml
   # GitHub Actions example
   - name: Run Enhanced Tests
     env:
       TEST_ENV: staging
       API_TOKEN: ${{ secrets.API_TOKEN }}
     run: ./testflowkit -enhanced
   ```

### **Compatibility Matrix**

| Feature | Legacy | Enhanced | Migration Effort |
|---------|--------|----------|------------------|
| Basic browser automation | ✅ | ✅ | None |
| Environment switching | ❌ | ✅ | Low |
| Element fallback strategies | ❌ | ✅ | Medium |
| Named API endpoints | ❌ | ✅ | Medium |
| Configuration validation | Basic | ✅ | Low |

## 🎯 Best Practices

### **1. Element Selector Strategy**

**✅ Recommended Hierarchy:**
```yaml
login_button:
  - "[data-testid='login-btn']"    # Most reliable
  - "button[type='submit']"        # Semantic
  - "#login-button"                # ID-based
  - ".btn-primary"                 # Class-based
```

**❌ Avoid:**
```yaml
login_button:
  - "div > span > button:nth-child(3)"  # Too fragile
  - ".css-generated-class-name"         # Changes frequently
```

### **2. Environment Management**

**✅ Good:**
```yaml
environments:
  local:
    frontend_base_url: "http://localhost:3000"
    backend_base_url: "http://localhost:8080/api"
  staging:
    frontend_base_url: "https://staging.example.com"
    backend_base_url: "https://api.staging.example.com"
```

### **3. Security**

**✅ Secure:**
```yaml
default_headers:
  Authorization: "Bearer ${API_TOKEN}"  # Environment variable
```

**❌ Insecure:**
```yaml
default_headers:
  Authorization: "Bearer hardcoded-token"  # Never do this
```

### **4. API Organization**

**✅ Well-organized:**
```yaml
endpoints:
  # Authentication
  auth_login:
    method: "POST"
    path: "/auth/login"
    description: "User authentication"
  
  # User management  
  user_get:
    method: "GET"
    path: "/users/{id}"
    description: "Get user by ID"
  
  user_create:
    method: "POST"
    path: "/users"
    description: "Create new user"
```

## 📊 Performance Tips

### **Selector Performance**

1. **data-testid** - Fastest, most reliable
2. **ID selectors** - Fast, unique
3. **Class selectors** - Moderate speed
4. **Complex CSS** - Slower, use sparingly

### **Concurrency Settings**

```yaml
# Local development
settings:
  concurrency: 1  # Single thread for debugging

# CI environment  
settings:
  concurrency: 4  # Parallel execution
```

### **Resource Management**

```yaml
# For CI/CD pipelines
settings:
  headless: true              # Faster execution
  video_recording: false      # Save disk space
  screenshot_on_failure: true # Keep for debugging
```

## 🐛 Troubleshooting

### **Common Issues**

#### **Configuration File Not Found**
```bash
# Check file exists
ls -la config-enhanced.yml

# Use absolute path
./testflowkit -enhanced -config=/full/path/to/config.yml
```

#### **Environment Variables Not Substituted**
```bash
# Check environment variables are set
echo $API_TOKEN

# Set if missing
export API_TOKEN="your-token-here"
```

#### **Element Not Found**
```gherkin
# Enable verbose logging
./testflowkit -enhanced -verbose

# Check selectors in browser dev tools
# F12 → Console → document.querySelector("[data-testid='element']")
```

#### **API Endpoint Not Found**
```yaml
# Verify endpoint is defined in configuration
endpoints:
  your_endpoint:
    method: "GET"
    path: "/your/path"
    description: "Description"
```

### **Debug Mode**

```bash
# Maximum verbosity
./testflowkit -enhanced -verbose

# Show configuration without running tests
./testflowkit -enhanced -show-config

# Test specific environment
TEST_ENV=staging ./testflowkit -enhanced -show-config
```

## 📚 Examples

### **Complete E2E Test**

```gherkin
@integration
Scenario: Complete User Journey with API and Frontend
  # Setup via API
  Given I prepare a request for the "create_user" endpoint
  And I set the request body to:
    """
    {
      "name": "Integration Test User",
      "email": "integration@test.com",
      "password": "TestPassword123"
    }
    """
  When I send the request
  Then the response status code should be 201
  And I store the response body path "user.id" as "test_user_id"

  # Frontend verification
  When the user goes to the "login" page
  And the user fills the "email_field" on "login_page" with "integration@test.com"
  And the user fills the "password_field" on "login_page" with "TestPassword123"
  And the user clicks on the "login_button" element on "login_page"
  
  Then the user should be navigated to "dashboard" page
  And the user should see on the page "Integration Test User"
  
  # Cleanup via API
  Given I prepare a request for the "delete_user" endpoint
  And I set the path parameter "id" to "${test_user_id}"
  When I send the request
  Then the response status code should be 204
```

### **Data-Driven Testing**

```gherkin
@api @data-driven
Scenario Outline: User Creation with Different Roles
  Given I prepare a request for the "create_user" endpoint
  And I set the request body to:
    """
    {
      "name": "<name>",
      "email": "<email>",
      "role": "<role>"
    }
    """
  When I send the request
  Then the response status code should be <status>

  Examples:
    | name        | email               | role  | status |
    | Admin User  | admin@example.com   | admin | 201    |
    | Regular User| user@example.com    | user  | 201    |
    | Guest User  | guest@example.com   | guest | 201    |
```

## 🤝 Contributing

The enhanced configuration system follows Go best practices:

- **Explicit Dependencies**: No hidden globals
- **Interface Segregation**: Small, focused interfaces  
- **Comprehensive Validation**: Fail fast with clear errors
- **Extensive Documentation**: Every public function documented

### **Development Setup**

```bash
# Clone repository
git clone https://github.com/yourorg/testflowkit.git
cd testflowkit

# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build -o testflowkit ./cmd/testflowkit
```

### **Adding New Features**

1. Update configuration types in `internal/config/enhanced_types.go`
2. Add validation in `internal/config/enhanced_loader.go`
3. Create step definitions in appropriate package
4. Add documentation and examples
5. Update configuration schema

## 🚀 Roadmap

- [ ] **Visual Testing Integration**: Screenshot comparison and visual regression testing
- [ ] **Performance Monitoring**: Built-in performance metrics collection  
- [ ] **Advanced Selectors**: XPath support and custom selector strategies
- [ ] **Test Data Management**: Built-in test data generation and cleanup
- [ ] **Parallel Execution**: Enhanced concurrency with smart load balancing
- [ ] **Cloud Integration**: Support for cloud testing platforms
- [ ] **AI-Powered Selectors**: Machine learning-based element detection

## 📞 Support

- **Documentation**: [Enhanced Configuration Guide](./docs/enhanced-configuration.md)
- **GitHub Issues**: [Report bugs and request features](https://github.com/yourorg/testflowkit/issues)
- **Discussions**: [Community support and questions](https://github.com/yourorg/testflowkit/discussions)

---

**TestFlowKit Enhanced Configuration** - Making web test automation more reliable, maintainable, and scalable. 🚀 