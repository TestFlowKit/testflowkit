<template>
    <div class="space-y-8">
        <!-- Header Section -->
        <section class="text-center bg-gradient-to-r from-purple-50 to-indigo-50 p-8 rounded-lg">
            <h1 class="text-3xl md:text-4xl font-bold mb-4 text-gray-800">Macro System</h1>
            <p class="text-lg md:text-xl text-gray-600 mb-6">
                Create reusable test scenarios to reduce code duplication and improve maintainability
            </p>
            <div class="flex justify-center space-x-4">
                <div class="bg-white px-4 py-2 rounded-lg shadow-sm">
                    <span class="text-sm font-medium text-gray-600">Version:</span>
                    <span class="text-sm font-semibold text-purple-600 ml-1">1.0.0</span>
                </div>
                <div class="bg-white px-4 py-2 rounded-lg shadow-sm">
                    <span class="text-sm font-medium text-gray-600">Status:</span>
                    <span class="text-sm font-semibold text-green-600 ml-1">Stable</span>
                </div>
            </div>
        </section>

        <!-- Overview Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">What are Macros?</h2>
            <p class="text-gray-600 mb-4">
                Macros in TestFlowKit are reusable test scenarios that allow you to define common test patterns once and reuse them across multiple test scenarios. This helps reduce code duplication and makes your test suite more maintainable.
            </p>
            
            <div class="bg-blue-50 p-4 rounded-lg mb-6">
                <h3 class="font-semibold mb-2 text-blue-800">Key Benefits:</h3>
                <ul class="space-y-2 text-blue-700">
                    <li>• <strong>Reduce Duplication:</strong> Define common test patterns once</li>
                    <li>• <strong>Improve Maintainability:</strong> Update logic in one place</li>
                    <li>• <strong>Parameter Support:</strong> Pass different data to the same macro</li>
                    <li>• <strong>Better Organization:</strong> Separate macro definitions from test scenarios</li>
                </ul>
            </div>
        </section>

        <!-- How Macros Work Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">How Macros Work</h2>
            
            <div class="grid md:grid-cols-2 gap-6">
                <div>
                    <h3 class="text-lg font-semibold mb-3 text-gray-700">1. Define Macros</h3>
                    <p class="text-gray-600 mb-4">
                        Create macro scenarios in feature files with the <code>@macro</code> tag. These are groups of steps that can be reused.
                    </p>
                    <CodeBlock :code="macroDefinitionExample" language="gherkin" />
                </div>
                
                <div>
                    <h3 class="text-lg font-semibold mb-3 text-gray-700">2. Use Macros</h3>
                    <p class="text-gray-600 mb-4">
                        Reference macros in your test scenarios by using the macro scenario name as a step. The macro steps get substituted directly.
                    </p>
                    <CodeBlock :code="macroUsageExample" language="gherkin" />
                </div>
            </div>
        </section>

        <!-- Macro Syntax Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Macro Syntax</h2>
            
            <AccordionItem title="Macro Definition">
                <div class="space-y-4">
                    <p>Macros are defined using the <code>@macro</code> tag and follow standard Gherkin scenario syntax:</p>
                    <CodeBlock :code="macroSyntaxExample" language="gherkin" />
                    
                    <div class="bg-yellow-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-yellow-800">Important Notes:</h4>
                        <ul class="text-yellow-700 space-y-1 text-sm">
                            <li>• Use <code>@macro</code> tag to identify macro scenarios</li>
                            <li>• Macros are static groups of steps - no parameter passing</li>
                            <li>• Macro scenarios can be in any feature file - the <code>@macro</code> tag is what matters</li>
                            <li>• Macros can contain any valid Gherkin steps</li>
                        </ul>
                    </div>
                </div>
            </AccordionItem>

            <AccordionItem title="How Macros Work">
                <div class="space-y-4">
                    <p>Macros are simple groups of steps that get substituted directly into your test scenarios:</p>
                    <CodeBlock :code="macroSubstitutionExample" language="gherkin" />
                    
                    <div class="bg-green-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-green-800">Macro Behavior:</h4>
                        <ul class="text-green-700 space-y-1 text-sm">
                            <li>• Macros are groups of steps that get substituted directly</li>
                            <li>• No parameter passing - macros are static step groups</li>
                            <li>• The macro scenario name becomes a single step</li>
                            <li>• All steps from the macro are inserted in place of the macro step</li>
                        </ul>
                    </div>
                </div>
            </AccordionItem>

            <AccordionItem title="Macro Organization">
                <div class="space-y-4">
                    <p>Organize your macros for better maintainability. Macros are identified by the <code>@macro</code> tag, not by file name:</p>
                    <CodeBlock :code="fileOrganizationExample" language="bash" />
                    
                    <div class="bg-blue-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-blue-800">Best Practices:</h4>
                        <ul class="text-blue-700 space-y-1 text-sm">
                            <li>• Group related macros in the same file</li>
                            <li>• Use descriptive macro names</li>
                            <li>• Keep macros focused on single responsibilities</li>
                            <li>• Document parameter requirements</li>
                            <li>• The <code>@macro</code> tag is what defines a macro, not the file name</li>
                            <li>• Macros are static step groups - they don't accept parameters</li>
                        </ul>
                    </div>
                </div>
            </AccordionItem>
        </section>

        <!-- Real-World Examples Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Real-World Examples</h2>
            
            <AccordionItem title="User Authentication Macros">
                <div class="space-y-4">
                    <p>Common authentication patterns that can be reused across multiple test scenarios:</p>
                    <CodeBlock :code="authMacroExample" language="gherkin" />
                    
                    <p class="text-gray-600">Usage in test scenarios:</p>
                    <CodeBlock :code="authMacroUsageExample" language="gherkin" />
                </div>
            </AccordionItem>

            <AccordionItem title="Data Setup Macros">
                <div class="space-y-4">
                    <p>Macros for setting up test data and application state:</p>
                    <CodeBlock :code="dataSetupMacroExample" language="gherkin" />
                    
                    <p class="text-gray-600">Usage with different data sets:</p>
                    <CodeBlock :code="dataSetupMacroUsageExample" language="gherkin" />
                </div>
            </AccordionItem>

            <AccordionItem title="Complex Workflow Macros">
                <div class="space-y-4">
                    <p>Multi-step workflows that combine multiple actions:</p>
                    <CodeBlock :code="workflowMacroExample" language="gherkin" />
                    
                    <p class="text-gray-600">Using the workflow macro:</p>
                    <CodeBlock :code="workflowMacroUsageExample" language="gherkin" />
                </div>
            </AccordionItem>
        </section>

        <!-- Best Practices Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Best Practices</h2>
            
            <div class="grid md:grid-cols-2 gap-6">
                <div class="bg-green-50 p-4 rounded-lg">
                    <h3 class="font-semibold mb-3 text-green-800">✅ Do's</h3>
                    <ul class="space-y-2 text-green-700 text-sm">
                        <li>• Keep macros focused on single responsibilities</li>
                        <li>• Use descriptive parameter names</li>
                        <li>• Group related macros in the same file</li>
                        <li>• Document parameter requirements</li>
                        <li>• Test macros independently</li>
                        <li>• Use consistent naming conventions</li>
                    </ul>
                </div>
                
                <div class="bg-red-50 p-4 rounded-lg">
                    <h3 class="font-semibold mb-3 text-red-800">❌ Don'ts</h3>
                    <ul class="space-y-2 text-red-700 text-sm">
                        <li>• Don't create overly complex macros</li>
                        <li>• Don't use unclear parameter names</li>
                        <li>• Don't mix macro and test scenarios in the same file</li>
                        <li>• Don't forget to handle parameter validation</li>
                        <li>• Don't create macros that are too specific</li>
                        <li>• Don't ignore macro reusability</li>
                    </ul>
                </div>
            </div>
        </section>

        <!-- Troubleshooting Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Troubleshooting</h2>
            
            <AccordionItem title="Common Issues">
                <div class="space-y-4">
                    <div class="bg-red-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-red-800">Macro Not Found</h4>
                        <p class="text-red-700 text-sm mb-2">Error: "Macro 'Login with credentials' not found"</p>
                        <p class="text-gray-600 text-sm">Solution: Ensure the scenario has the <code>@macro</code> tag. The tag is what identifies a macro, not the file name.</p>
                    </div>
                    
                    <div class="bg-red-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-red-800">Macro Step Not Found</h4>
                        <p class="text-red-700 text-sm mb-2">Error: "Step 'Login with credentials' not found"</p>
                        <p class="text-gray-600 text-sm">Solution: Ensure the macro scenario name matches exactly with how you're calling it in your test.</p>
                    </div>
                    
                    <div class="bg-red-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-red-800">Circular Dependencies</h4>
                        <p class="text-red-700 text-sm mb-2">Error: "Circular dependency detected in macros"</p>
                        <p class="text-gray-600 text-sm">Solution: Avoid macros that reference each other in a circular manner.</p>
                    </div>
                </div>
            </AccordionItem>
        </section>

        <!-- Related Links Section -->
        <section class="bg-gray-50 p-6 rounded-lg">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Related Documentation</h2>
            <div class="grid md:grid-cols-2 gap-4">
                <nuxt-link to="/sentences" class="bg-white p-4 rounded-lg shadow-sm hover:shadow-md transition-shadow">
                    <h3 class="font-semibold text-blue-600 mb-2">Step Definitions</h3>
                    <p class="text-gray-600 text-sm">Learn about all available Gherkin step definitions</p>
                </nuxt-link>
                <nuxt-link to="/variables" class="bg-white p-4 rounded-lg shadow-sm hover:shadow-md transition-shadow">
                    <h3 class="font-semibold text-blue-600 mb-2">Variables System</h3>
                    <p class="text-gray-600 text-sm">Understand how variables work with macros</p>
                </nuxt-link>
                <nuxt-link to="/configuration" class="bg-white p-4 rounded-lg shadow-sm hover:shadow-md transition-shadow">
                    <h3 class="font-semibold text-blue-600 mb-2">Configuration</h3>
                    <p class="text-gray-600 text-sm">Configure your test environment and settings</p>
                </nuxt-link>
                <nuxt-link to="/quick-start" class="bg-white p-4 rounded-lg shadow-sm hover:shadow-md transition-shadow">
                    <h3 class="font-semibold text-blue-600 mb-2">Quick Start</h3>
                    <p class="text-gray-600 text-sm">Get started with TestFlowKit basics</p>
                </nuxt-link>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
import AccordionItem from '../components/AccordionItem.vue';
import CodeBlock from '../components/global/CodeBlock.vue';

// Example code blocks
const macroDefinitionExample = `@macro
Scenario: Login with credentials
  Given the user is on the homepage
  When the user goes to the "login" page
  And the user enters "test@example.com" into the "email" field
  And the user enters "password123" into the "password" field
  And the user clicks on the "login" button`;

const macroUsageExample = `Scenario: Test with macro
  Given Login with credentials
  Then the user should be navigated to "dashboard" page`;

const macroSyntaxExample = `@macro
Scenario: Macro Name
  Given some precondition
  When some action
  Then some verification`;

const macroSubstitutionExample = `# Original test scenario
Scenario: Test with macro
  Given Login with credentials
  Then the user should be navigated to "dashboard" page

# After macro substitution, it becomes:
Scenario: Test with macro
  Given the user is on the homepage
  When the user goes to the "login" page
  And the user enters "test@example.com" into the "email" field
  And the user enters "password123" into the "password" field
  And the user clicks on the "login" button
  Then the user should be navigated to "dashboard" page`;

const fileOrganizationExample = `e2e/features/
├── macros/
│   ├── authentication.feature  # Contains @macro scenarios
│   ├── data-setup.feature      # Contains @macro scenarios
│   └── workflows.feature       # Contains @macro scenarios
└── tests/
    ├── login.feature           # Regular test scenarios
    └── user-management.feature # Regular test scenarios

# Note: The @macro tag identifies macro scenarios, not the file name`;

const authMacroExample = `@macro
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
  Then the user should be navigated to "login" page`;

const authMacroUsageExample = `Scenario: Admin user login
  Given Login with credentials
  Then the user should be navigated to "dashboard" page

Scenario: Regular user login
  Given Login with credentials
  Then the user should be navigated to "dashboard" page`;

const dataSetupMacroExample = `@macro
Scenario: Create test user
  Given I prepare a request for the "create_user" endpoint
  When I set the following request body:
    | name  | email           | role    |
    | Admin | admin@test.com  | admin   |
  And I send the request
  Then the response status code should be 201`;

const dataSetupMacroUsageExample = `Scenario: Test with admin user
  Given Create test user
  When the user goes to the "admin" page
  Then the user should see "Admin Dashboard"`;

const workflowMacroExample = `@macro
Scenario: Complete user registration
  Given the user is on the "register" page
  When the user enters "John Doe" into the "name" field
  And the user enters "john@example.com" into the "email" field
  And the user enters "password123" into the "password" field
  And the user clicks on the "register" button
  And the user verifies their email with "123456"
  Then the user should be navigated to "dashboard" page`;

const workflowMacroUsageExample = `Scenario: New user registration flow
  Given Complete user registration
  When the user goes to the "profile" page
  Then the user should see "Welcome John"`;
</script> 