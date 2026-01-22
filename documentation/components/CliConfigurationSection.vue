<template>
    <div class="bg-gray-100 p-4 md:p-6 rounded-lg mb-8">
        <h2 class="text-xl md:text-2xl font-semibold mb-4">CLI Configuration</h2>
        <p class="text-sm md:text-base">Configure settings for running the test application. This YAML file controls
            global settings, environments,
            frontend elements, and backend endpoints.</p>

        <AccordionItem title="Global Settings">
            <div class="overflow-x-auto">
                <table class="table-auto w-full min-w-full text-sm">
                    <thead>
                        <tr>
                            <th class="px-2 md:px-4 py-2 text-left">Option</th>
                            <th class="px-2 md:px-4 py-2 text-left">Description</th>
                            <th class="px-2 md:px-4 py-2 text-left">Default</th>
                            <th class="px-2 md:px-4 py-2 text-left">Example</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">default_timeout</code>
                            </td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Maximum time to wait when searching
                                for elements by CSS
                                selectors or XPath (in milliseconds).</td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">10000</code></td>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">default_timeout: 30000</code></td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">concurrency</code>
                            </td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Number of tests to run in parallel
                                (1-20).</td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">1</code></td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">concurrency: 5</code>
                            </td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">headless</code></td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Run browser in headless mode.</td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">false</code></td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">headless: true</code>
                            </td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">think_time</code></td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Slow down test execution (in
                                milliseconds) - useful for
                                debugging.</td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">1000</code></td>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">think_time: 2000</code></td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">screenshot_on_failure</code></td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Take screenshots when tests fail.
                            </td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">true</code></td>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">screenshot_on_failure: false</code></td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">report_format</code>
                            </td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Format of test reports (html, json,
                                junit).</td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">html</code></td>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">report_format: json</code></td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">gherkin_location</code></td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Path to the directory containing the
                                Gherkin feature files.
                            </td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">./e2e/features</code>
                            </td>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">gherkin_location: "./tests/features"</code></td>
                        </tr>
                        <tr>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">tags</code></td>
                            <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Filter tests by tags.</td>
                            <td class="border px-2 md:px-4 py-2"><code class="text-xs md:text-sm">""</code></td>
                            <td class="border px-2 md:px-4 py-2"><code
                                    class="text-xs md:text-sm">tags: "@smoke,@regression"</code></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </AccordionItem>

        <AccordionItem title="Environment Variables">
            <p class="text-sm md:text-base">Define environment variables for your base URLs and configuration values. Can be inline or in external files.</p>
            <ul class="list-disc list-inside mb-4 text-sm md:text-base">
                <li><code>env</code>: Inline environment variables block.</li>
                <li><code>env_file</code>: Path to external environment file.</li>
                <li><code>frontend_base_url</code>: Base URL for frontend testing.</li>
                <li><code>rest_api_base_url</code>: Base URL for REST API testing.</li>
                <li><code>graphql_endpoint</code>: Endpoint for GraphQL operations.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="Frontend Configuration">
            <p class="text-sm md:text-base">Configure frontend elements and pages for UI testing. Supports both CSS
                selectors and XPath expressions
                for flexible element selection.</p>
            <ul class="list-disc list-inside mb-4 text-sm md:text-base">
                <li><code>elements</code>: Define reusable selectors for UI elements (CSS and XPath).</li>
                <li><code>pages</code>: Define page URLs and paths.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="Backend Configuration">
            <p class="text-sm md:text-base">Configure API endpoints and default headers for backend testing.</p>
            <ul class="list-disc list-inside mb-4 text-sm md:text-base">
                <li><code>default_headers</code>: Default HTTP headers for API requests.</li>
                <li><code>endpoints</code>: Define API endpoints with method, path, and description.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="CLI Configuration Example">
            <CodeBlock :code="cliConfigExample" language="yaml" />
        </AccordionItem>
    </div>
</template>

<script setup lang="ts">
import AccordionItem from '@/components/AccordionItem.vue';
import CodeBlock from '@/components/global/CodeBlock.vue';

const cliConfigExample = `
settings:
  # Element search timeout in milliseconds (1-300000ms)
  # Maximum time to wait when searching for elements by CSS selectors or XPath
  default_timeout: 30000
  concurrency: 5
  headless: false
  think_time: 1000
  screenshot_on_failure: true
  report_format: "html"
  gherkin_location: "./e2e/features"
  tags: "@smoke"
  env_file: ".env.local.yml"  # Optional: default env file

# Inline environment variables (or use external file)
env:
  frontend_base_url: "http://localhost:3000"
  rest_api_base_url: "http://localhost:8080"
  graphql_endpoint: "http://localhost:8080/graphql"

frontend:
  elements:
    common:
      loading_spinner:
        - "[data-testid='loading-spinner']"
        - ".spinner"
        - ".loading"
        - "xpath://div[contains(@class, 'loading')]"
    login:
      username_field: "#username"
      password_field: "#password"
      login_button: "#login-button"
      # XPath selectors for complex elements
      submit_button:
        - "xpath://button[contains(@class, 'submit') and text()='Login']"
        - "xpath://div[@id='login-form']//button[@type='submit']"
        - "[data-testid='login-button']"
  
  pages:
    home: "/"
    login: "/login"
    dashboard: "/dashboard"

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
