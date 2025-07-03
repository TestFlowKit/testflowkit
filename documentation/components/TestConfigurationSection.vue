<template>
  <div class="bg-blue-100 p-6 rounded-lg mb-8">
    <h2 class="text-2xl font-semibold mb-4">Test Configuration</h2>
    <p>Define variables, page objects, and base URLs for your tests. This YAML file helps organize your tests and
      makes them more maintainable.</p>

    <AccordionItem title="Frontend Elements">
      <p>The <code>frontend.elements</code> section allows you to define reusable selectors for UI elements. Elements
        can be organized by page or as common elements.</p>
      <ul class="list-disc list-inside mb-4">
        <li><code>common</code>: Elements that are used across multiple pages.</li>
        <li><code>page-specific</code>: Elements specific to a particular page.</li>
      </ul>
    </AccordionItem>

    <AccordionItem title="Elements Section">
      <p>Define reusable selectors for common UI elements with multiple fallback options. TestFlowKit supports both CSS
        selectors and XPath expressions.</p>
      <CodeBlock :code="elementsSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="XPath Selector Support">
      <p>TestFlowKit provides full XPath 1.0 support for complex element selection. Use the <code>xpath:</code> prefix
        to specify XPath expressions.</p>
      <CodeBlock :code="xpathSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Pages Section">
      <p>Define URLs for different pages used in your tests. These can be relative paths or absolute URLs.</p>
      <CodeBlock :code="pagesSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Environment Configuration">
      <p>Define different environments with their base URLs. The <code>active_environment</code> determines which
        environment to use.</p>
      <CodeBlock :code="environmentSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Settings Configuration">
      <p>Configure global test execution settings including timeouts, concurrency, and reporting options.</p>
      <CodeBlock :code="settingsSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Backend Configuration">
      <p>Configure API endpoints and default headers for backend testing.</p>
      <CodeBlock :code="backendSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Complete Test Configuration Example">
      <CodeBlock :code="testConfigExample" language="yaml" />
    </AccordionItem>
  </div>
</template>

<script setup lang="ts">
import AccordionItem from './AccordionItem.vue';
import CodeBlock from './global/CodeBlock.vue';

const elementsSection = `
frontend:
  elements:
    common:
      loading_spinner:
        - "[data-testid='loading-spinner']"
        - ".spinner"
        - ".loading"
      error_message:
        - ".error-message"
        - "[data-testid='error']"
    login:
      username_field: "#username"
      password_field: "#password"
      login_button: "#login-button"
      forgot_password_link: ".forgot-password"
`.trim();

const xpathSection = `
frontend:
  elements:
    login_page:
      # XPath selectors for complex element selection
      complex_button:
        - "xpath://button[contains(@class, 'submit') and text()='Login']"
        - "xpath://div[@id='login-form']//button[@type='submit']"
        - "[data-testid='login-button']"  # CSS fallback
      
      # XPath with text content matching
      dynamic_text:
        - "xpath://div[contains(text(), 'Welcome') and @class='message']"
        - "xpath://span[text()='Hello, User!']"
      
      # XPath with attribute conditions
      required_field:
        - "xpath://input[@type='email' and @required]"
        - "xpath://input[@name='email' and @data-required='true']"
      
      # Mixed selectors with XPath and CSS fallbacks
      flexible_element:
        - "xpath://div[contains(@class, 'dynamic') and contains(text(), 'Loading')]"
        - ".loading-indicator"
        - "[data-testid='loading']"
`.trim();

const pagesSection = `
frontend:
  pages:
    home: "/"
    login: "/login"
    dashboard: "/dashboard"
    profile: "/profile"
    settings: "/settings"
`.trim();

const environmentSection = `
active_environment: "local"

environments:
  local:
    frontend_base_url: "http://localhost:3000"
    api_base_url: "http://localhost:8080"
  
  staging:
    frontend_base_url: "https://staging.example.com"
    api_base_url: "https://api-staging.example.com"
  
  production:
    frontend_base_url: "https://example.com"
    api_base_url: "https://api.example.com"
`.trim();

const settingsSection = `
settings:
  # Element search timeout in milliseconds (1-300000ms)
  # Maximum time to wait when searching for elements by CSS selectors or XPath
  default_timeout: 30000
  
  # Number of parallel test executions
  concurrency: 5
  
  # Run browser in headless mode (no UI)
  headless: false
  
  # Delay between actions in milliseconds (for debugging)
  think_time: 1000
  
  # Take screenshot on test failure
  screenshot_on_failure: true
  
  # Test report format
  report_format: "html"
  
  # Location of Gherkin feature files
  gherkin_location: "./e2e/features"
  
  # Filter tests by tags
  tags: "@smoke"
`.trim();

const backendSection = `
backend:
  default_headers:
    Content-Type: "application/json"
    Accept: "application/json"
    User-Agent: "TestFlowKit/1.0"
    Authorization: "Bearer {token}"
  
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user information by ID"
    create_user:
      method: "POST"
      path: "/users"
      description: "Create a new user"
    update_user:
      method: "PUT"
      path: "/users/{id}"
      description: "Update user information"
    delete_user:
      method: "DELETE"
      path: "/users/{id}"
      description: "Delete a user"
`.trim();

const testConfigExample = `
active_environment: "local"

settings:
  default_timeout: 30000
  concurrency: 5
  headless: false
  think_time: 1000
  screenshot_on_failure: true
  report_format: "html"
  gherkin_location: "./e2e/features"

environments:
  local:
    frontend_base_url: "http://localhost:3000"
    api_base_url: "http://localhost:8080"
  
  staging:
    frontend_base_url: "https://staging.example.com"
    api_base_url: "https://api-staging.example.com"

frontend:
  elements:
    common:
      loading_spinner:
        - "[data-testid='loading-spinner']"
        - ".spinner"
        - ".loading"
      error_message:
        - ".error-message"
        - "[data-testid='error']"
    login:
      username_field: "#username"
      password_field: "#password"
      login_button: "#login-button"
      forgot_password_link: ".forgot-password"
      # XPath selectors for complex elements
      submit_button:
        - "xpath://button[contains(@class, 'submit') and text()='Login']"
        - "xpath://div[@id='login-form']//button[@type='submit']"
        - "[data-testid='login-button']"
    dashboard:
      welcome_message: ".welcome-message"
      logout_button: "#logout"
      profile_link: ".profile-link"
      # XPath for dynamic content
      user_greeting:
        - "xpath://div[contains(text(), 'Welcome') and @class='greeting']"
        - ".user-greeting"
  
  pages:
    home: "/"
    login: "/login"
    dashboard: "/dashboard"
    profile: "/profile"

backend:
  default_headers:
    Content-Type: "application/json"
    Accept: "application/json"
    User-Agent: "TestFlowKit/1.0"
  
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user information by ID"
    create_user:
      method: "POST"
      path: "/users"
      description: "Create a new user"
`.trim();
</script>
