<template>
  <div class="bg-blue-100 p-4 md:p-6 rounded-lg mb-8">
    <h2 class="text-xl md:text-2xl font-semibold mb-4">Test Configuration</h2>
    <p class="text-sm md:text-base">Define variables, page objects, and base URLs for your tests. This YAML file helps
      organize your tests and
      makes them more maintainable.</p>

    <AccordionItem title="Frontend Elements">
      <p class="text-sm md:text-base">The <code>frontend.elements</code> section allows you to define reusable selectors
        for UI elements. Elements
        can be organized by page or as common elements.</p>
      <ul class="list-disc list-inside mb-4 text-sm md:text-base">
        <li><code>common</code>: Elements that are used across multiple pages.</li>
        <li><code>page-specific</code>: Elements specific to a particular page.</li>
      </ul>
    </AccordionItem>

    <AccordionItem title="Elements Section">
      <p class="text-sm md:text-base">Define reusable selectors for common UI elements with multiple fallback options.
        TestFlowKit supports both CSS
        selectors and XPath expressions.</p>
      <CodeBlock :code="elementsSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="XPath Selector Support">
      <p class="text-sm md:text-base">TestFlowKit provides full XPath 1.0 support for complex element selection. Use the
        <code>xpath:</code> prefix
        to specify XPath expressions.</p>
      <CodeBlock :code="xpathSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Pages Section">
      <p class="text-sm md:text-base">Define URLs for different pages used in your tests. These can be relative paths or
        absolute URLs.</p>
      <CodeBlock :code="pagesSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Environment Variables">
      <p class="text-sm md:text-base">Define environment variables for your base URLs and configuration values. These can be inline or in external files.</p>
      <CodeBlock :code="environmentSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="Settings Configuration">
      <p class="text-sm md:text-base">Configure global test execution settings including timeouts, concurrency, and
        reporting options.</p>
      <CodeBlock :code="settingsSection" language="yaml" />
    </AccordionItem>

    <AccordionItem title="APIs Configuration">
      <p class="text-sm md:text-base">Configure multiple REST and GraphQL APIs with unified configuration. Each API can have its own base URL, headers, timeout, and endpoints/operations.</p>
      <CodeBlock :code="apisSection" language="yaml" />
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
settings:
  env_file: ".env.local.yml"  # Optional: default env file

# Inline environment variables
env:
  frontend_base_url: "http://localhost:3000"
  jsonplaceholder_base_url: "http://localhost:8080"
  my_graphql_endpoint: "http://localhost:8080/graphql"

# Or use external files:
# .env.local.yml
# frontend_base_url: "http://localhost:3000"
# jsonplaceholder_base_url: "http://localhost:8080"
#
# .env.staging.yml
# frontend_base_url: "https://staging.example.com"
# jsonplaceholder_base_url: "https://api-staging.example.com"
#
# Override at runtime: tkit run --env-file .env.staging.yml
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

const apisSection = `
apis:
  default_timeout: 30000  # Default timeout for all APIs (ms)
  
  definitions:
    # REST API example
    jsonplaceholder:
      type: rest
      base_url: "{{ env.jsonplaceholder_base_url }}"
      default_headers:
        Content-Type: "application/json"
        Accept: "application/json"
        User-Agent: "TestFlowKit/1.0"
      timeout: 5000  # Optional: override default timeout
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
    
    # GraphQL API example
    my_graphql:
      type: graphql
      endpoint: "{{ env.my_graphql_endpoint }}"
      default_headers:
        Content-Type: "application/json"
        Authorization: "Bearer {{ env.api_token }}"
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
        
        update_user:
          type: "mutation"
          operation: |
            mutation UpdateUser($id: ID!, $input: UserInput!) {
              updateUser(id: $id, input: $input) {
                id
                name
                email
                updatedAt
              }
            }
          description: "Update user information"
`.trim();

const testConfigExample = `
settings:
  default_timeout: 30000
  concurrency: 5
  headless: false
  think_time: 1000
  screenshot_on_failure: true
  report_format: "html"
  gherkin_location: "./e2e/features"
  env_file: ".env.local.yml"  # Optional

# Inline environment variables (or use external file)
env:
  frontend_base_url: "http://localhost:3000"
  jsonplaceholder_base_url: "http://localhost:8080"
  my_graphql_endpoint: "http://localhost:8080/graphql"
  api_token: "your-api-token"

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

apis:
  default_timeout: 30000
  
  definitions:
    jsonplaceholder:
      type: rest
      base_url: "{{ env.jsonplaceholder_base_url }}"
      default_headers:
        Content-Type: "application/json"
        Accept: "application/json"
      endpoints:
        get_user:
          method: "GET"
          path: "/users/{id}"
          description: "Retrieve user information by ID"
        create_user:
          method: "POST"
          path: "/users"
          description: "Create a new user"
    
    my_graphql:
      type: graphql
      endpoint: "{{ env.my_graphql_endpoint }}"
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
              }
            }
          description: "Fetch user profile"
`.trim();
</script>
