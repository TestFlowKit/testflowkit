<template>
    <div>
        <h1 class="text-3xl font-extrabold mb-4">Variables System</h1>
        <p class="text-lg mb-8 text-gray-600">
            TestFlowKit provides a powerful variable system that allows you to store and reuse data throughout your test
            scenarios,
            enabling dynamic and data-driven testing.
        </p>

        <div class="bg-blue-100 p-6 rounded-lg mb-8">
            <h2 class="text-2xl font-semibold mb-4">Variable Syntax</h2>
            <p class="mb-4">Variables in TestFlowKit use the <code>&#123;&#123;variable_name&#125;&#125;</code> syntax
                and are automatically substituted in all step parameters.</p>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <h3 class="text-lg font-semibold mb-2">Key Features</h3>
                    <ul class="list-disc list-inside space-y-1">
                        <li><strong>Automatic Substitution:</strong> Variables are replaced in strings, tables, and
                            parameters</li>
                        <li><strong>Scenario Scope:</strong> Variables persist throughout the entire scenario</li>
                        <li><strong>Type Support:</strong> Supports strings, numbers, booleans, and complex data
                            structures</li>
                        <li><strong>Cross-Step Usage:</strong> Use variables across different step types</li>
                    </ul>
                </div>
                <div>
                    <h3 class="text-lg font-semibold mb-2">Basic Example</h3>
                    <CodeBlock :code="basicExample" language="gherkin" />
                </div>
            </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mb-8">
            <div class="bg-white p-6 rounded-lg shadow-md">
                <h3 class="text-xl font-semibold mb-4 text-blue-600">Custom Variables</h3>
                <p class="mb-4">Store any custom value for reuse throughout your scenario.</p>
                <CodeBlock :code="customVariableExample" language="gherkin" />
                <div class="mt-4">
                    <h4 class="font-semibold mb-2">Use Cases:</h4>
                    <ul class="list-disc list-inside text-sm space-y-1">
                        <li>Store hardcoded values for reuse</li>
                        <li>Create variables for dynamic test data</li>
                        <li>Set up test data used in multiple steps</li>
                    </ul>
                </div>
            </div>

            <div class="bg-white p-6 rounded-lg shadow-md">
                <h3 class="text-xl font-semibold mb-4 text-green-600">JSON Path Variables</h3>
                <p class="mb-4">Extract data from JSON API responses using JSON path expressions.</p>
                <CodeBlock :code="jsonPathExample" language="gherkin" />
                <div class="mt-4">
                    <h4 class="font-semibold mb-2">Use Cases:</h4>
                    <ul class="list-disc list-inside text-sm space-y-1">
                        <li>Extract specific data from API responses</li>
                        <li>Store response data for use in subsequent steps</li>
                        <li>Create dynamic test data based on API responses</li>
                    </ul>
                </div>
            </div>

            <div class="bg-white p-6 rounded-lg shadow-md">
                <h3 class="text-xl font-semibold mb-4 text-purple-600">HTML Element Variables</h3>
                <p class="mb-4">Capture text content from HTML elements on web pages.</p>
                <CodeBlock :code="htmlElementExample" language="gherkin" />
                <div class="mt-4">
                    <h4 class="font-semibold mb-2">Use Cases:</h4>
                    <ul class="list-disc list-inside text-sm space-y-1">
                        <li>Capture dynamic content from web pages</li>
                        <li>Store element text for comparison or reuse</li>
                        <li>Extract data displayed on the page</li>
                    </ul>
                </div>
            </div>
        </div>

        <div class="bg-green-100 p-6 rounded-lg mb-8">
            <h2 class="text-2xl font-semibold mb-4">Advanced Variable Usage</h2>
            <p class="mb-4">Variables can be used in complex scenarios for data-driven testing and end-to-end workflows.
            </p>

            <AccordionItem title="End-to-End Data Flow">
                <p class="mb-4">Combine API testing with frontend testing using variables:</p>
                <CodeBlock :code="endToEndExample" language="gherkin" />
            </AccordionItem>

            <AccordionItem title="Data-Driven Testing">
                <p class="mb-4">Use variables to create dynamic test scenarios:</p>
                <CodeBlock :code="dataDrivenExample" language="gherkin" />
            </AccordionItem>

            <AccordionItem title="Cross-Page Data Transfer">
                <p class="mb-4">Transfer data between different pages in your application:</p>
                <CodeBlock :code="crossPageExample" language="gherkin" />
            </AccordionItem>
        </div>

        <div class="bg-yellow-100 p-6 rounded-lg mb-8">
            <h2 class="text-2xl font-semibold mb-4">Best Practices</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <h3 class="text-lg font-semibold mb-3">Variable Naming</h3>
                    <ul class="list-disc list-inside space-y-2">
                        <li>Use descriptive names that clearly indicate the data being stored</li>
                        <li>Use snake_case for variable names (e.g., <code>user_name</code>, <code>api_response</code>)
                        </li>
                        <li>Include the data source in the name when relevant (e.g., <code>api_user_id</code>,
                            <code>page_title</code>)
                        </li>
                        <li>Avoid generic names like <code>data</code> or <code>value</code></li>
                    </ul>
                </div>
                <div>
                    <h3 class="text-lg font-semibold mb-3">Variable Usage</h3>
                    <ul class="list-disc list-inside space-y-2">
                        <li>Store variables early in your scenario for better readability</li>
                        <li>Use variables to make your tests more maintainable and reusable</li>
                        <li>Consider using variables for test data that might change between environments</li>
                        <li>Verify variables exist before using them in critical steps</li>
                    </ul>
                </div>
            </div>
        </div>

        <div class="bg-red-100 p-6 rounded-lg mb-8">
            <h2 class="text-2xl font-semibold mb-4">Common Patterns</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <h3 class="text-lg font-semibold mb-3">API to Frontend Flow</h3>
                    <CodeBlock :code="apiToFrontendPattern" language="gherkin" />
                </div>
                <div>
                    <h3 class="text-lg font-semibold mb-3">Page Content Capture</h3>
                    <CodeBlock :code="pageContentPattern" language="gherkin" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import AccordionItem from '../components/AccordionItem.vue';
import CodeBlock from '../components/global/CodeBlock.vue';

const basicExample = `
# Store a custom value
When I store the "John Doe" into "user_name" variable

# Use the variable in another step
When the user enters "{{user_name}}" into the "name" field

# Use in assertions
Then the "name" field should contain "{{user_name}}"
`.trim();

const customVariableExample = `
# Store different types of data
When I store the "test@example.com" into "email" variable
And I store the "12345" into "user_id" variable
And I store the "Active" into "status" variable

# Use in form filling
When the user enters "{{email}}" into the "email" field
And the user enters "{{user_id}}" into the "id" field
`.trim();

const jsonPathExample = `
# Extract data from API response
Given I prepare a request for the "get_user" endpoint
And I set the following path params:
  | id | 123 |
When I send the request
And I store the JSON path "data.name" from the response into "user_name" variable
And I store the JSON path "data.email" from the response into "user_email" variable
Then the response status code should be 200
`.trim();

const htmlElementExample = `
# Capture content from web elements
Given the user goes to the "profile" page
When I store the content of "page_title" into "title" variable
And I store the content of "user_name_label" into "displayed_name" variable
And I store the content of "welcome_message" into "welcome_text" variable
Then the "title" should contain "Profile"
`.trim();

const endToEndExample = `
Scenario: Complete user workflow with variables
  Given I prepare a request for the "create_user" endpoint
  And I set the request body to:
    """
    {
      "name": "Test User",
      "email": "test@example.com"
    }
    """
  When I send the request
  And I store the JSON path "data.id" from the response into "created_user_id" variable
  Then the response status code should be 201
  
  When the user goes to the "user_list" page
  And I store the content of "first_user_name" into "list_user_name" variable
  And the user clicks on the "first_user" link
  Then the user should be on the "user_details" page
  And the "user_name" should contain "{{list_user_name}}"
`.trim();

const dataDrivenExample = `
Scenario Outline: Test with dynamic data
  Given I store the "<email>" into "test_email" variable
  When the user goes to the "login" page
  And the user enters "{{test_email}}" into the "email" field
  And the user enters "<password>" into the "password" field
  And the user clicks on the "login" button
  Then the user should be navigated to "dashboard" page

  Examples:
    | email              | password   |
    | user1@test.com     | pass123    |
    | user2@test.com     | pass456    |
    | admin@test.com     | admin123   |
`.trim();

const crossPageExample = `
Scenario: Transfer data between pages
  Given the user goes to the "product_list" page
  When I store the content of "first_product_name" into "product_name" variable
  And I store the content of "first_product_price" into "product_price" variable
  And the user clicks on the "first_product" link
  Then the user should be on the "product_details" page
  And the "product_title" should contain "{{product_name}}"
  And the "product_price_display" should contain "{{product_price}}"
`.trim();

const apiToFrontendPattern = `
# Pattern: Extract API data and use in frontend
Given I prepare a request for the "get_user" endpoint
And I set the following path params:
  | id | 123 |
When I send the request
And I store the JSON path "data.name" from the response into "api_user_name" variable
And I store the JSON path "data.email" from the response into "api_user_email" variable
Then the response status code should be 200

When the user goes to the "profile" page
And the user enters "{{api_user_name}}" into the "name" field
And the user enters "{{api_user_email}}" into the "email" field
Then the "name" field should contain "{{api_user_name}}"
`.trim();

const pageContentPattern = `
# Pattern: Capture page content for verification
Given the user goes to the "dashboard" page
When I store the content of "page_title" into "dashboard_title" variable
And I store the content of "user_welcome" into "welcome_message" variable
Then the variable "dashboard_title" should contain "Dashboard"
And the variable "welcome_message" should contain "Welcome"
`.trim();
</script>