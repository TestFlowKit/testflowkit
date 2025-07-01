<template>
    <div class="bg-gray-100 p-6 rounded-lg mb-8">
        <h2 class="text-2xl font-semibold mb-4">CLI Configuration</h2>
        <p>Configure settings for running the test application. This YAML file controls global settings, environments,
            frontend elements, and backend endpoints.</p>

        <AccordionItem title="Global Settings">
            <div class="overflow-x-auto">
                <table class="table-auto w-full">
                    <thead>
                        <tr>
                            <th class="px-4 py-2">Option</th>
                            <th class="px-4 py-2">Description</th>
                            <th class="px-4 py-2">Default Value</th>
                            <th class="px-4 py-2">Example</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td class="border px-4 py-2"><code>default_timeout</code></td>
                            <td class="border px-4 py-2">Maximum test execution time (in milliseconds).</td>
                            <td class="border px-4 py-2"><code>10000</code></td>
                            <td class="border px-4 py-2"><code>default_timeout: 30000</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>concurrency</code></td>
                            <td class="border px-4 py-2">Number of tests to run in parallel (1-20).</td>
                            <td class="border px-4 py-2"><code>1</code></td>
                            <td class="border px-4 py-2"><code>concurrency: 5</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>headless</code></td>
                            <td class="border px-4 py-2">Run browser in headless mode.</td>
                            <td class="border px-4 py-2"><code>false</code></td>
                            <td class="border px-4 py-2"><code>headless: true</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>think_time</code></td>
                            <td class="border px-4 py-2">Slow down test execution (in milliseconds) - useful for
                                debugging.</td>
                            <td class="border px-4 py-2"><code>1000</code></td>
                            <td class="border px-4 py-2"><code>think_time: 2000</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>screenshot_on_failure</code></td>
                            <td class="border px-4 py-2">Take screenshots when tests fail.</td>
                            <td class="border px-4 py-2"><code>true</code></td>
                            <td class="border px-4 py-2"><code>screenshot_on_failure: false</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>report_format</code></td>
                            <td class="border px-4 py-2">Format of test reports (html, json, junit).</td>
                            <td class="border px-4 py-2"><code>html</code></td>
                            <td class="border px-4 py-2"><code>report_format: json</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>gherkin_location</code></td>
                            <td class="border px-4 py-2">Path to the directory containing the Gherkin feature files.
                            </td>
                            <td class="border px-4 py-2"><code>./e2e/features</code></td>
                            <td class="border px-4 py-2"><code>gherkin_location: "./tests/features"</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>tags</code></td>
                            <td class="border px-4 py-2">Filter tests by tags.</td>
                            <td class="border px-4 py-2"><code>""</code></td>
                            <td class="border px-4 py-2"><code>tags: "@smoke,@regression"</code></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </AccordionItem>

        <AccordionItem title="Environments">
            <p>Define different environments with their base URLs for frontend and API testing.</p>
            <ul class="list-disc list-inside mb-4">
                <li><code>active_environment</code>: The currently active environment.</li>
                <li><code>frontend_base_url</code>: Base URL for frontend testing.</li>
                <li><code>api_base_url</code>: Base URL for API testing.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="Frontend Configuration">
            <p>Configure frontend elements and pages for UI testing. Supports both CSS selectors and XPath expressions
                for flexible element selection.</p>
            <ul class="list-disc list-inside mb-4">
                <li><code>elements</code>: Define reusable selectors for UI elements (CSS and XPath).</li>
                <li><code>pages</code>: Define page URLs and paths.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="Backend Configuration">
            <p>Configure API endpoints and default headers for backend testing.</p>
            <ul class="list-disc list-inside mb-4">
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
active_environment: "local"

settings:
  default_timeout: 30000
  concurrency: 5
  headless: false
  think_time: 1000
  screenshot_on_failure: true
  report_format: "html"
  gherkin_location: "./e2e/features"
  tags: "@smoke"

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
