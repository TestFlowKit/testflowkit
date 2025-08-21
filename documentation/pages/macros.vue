<template>
    <div class="space-y-8">
        <!-- Header Section -->
        <section class="text-center bg-gradient-to-r from-purple-50 to-indigo-50 p-8 rounded-lg">
            <h1 class="text-3xl md:text-4xl font-bold mb-4 text-gray-800">Macro System</h1>
            <p class="text-lg md:text-xl text-gray-600 mb-6">
                Create reusable, parameterized test scenarios with variable substitution to reduce code duplication and
                improve maintainability
            </p>
            <div class="flex justify-center space-x-4">
                <div class="bg-white px-4 py-2 rounded-lg shadow-sm">
                    <span class="text-sm font-medium text-gray-600">Version:</span>
                    <span class="text-sm font-semibold text-purple-600 ml-1">2.0.0</span>
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
                Macros in TestFlowKit are reusable, parameterized test scenarios that allow you to define common test
                patterns once and reuse them across multiple test scenarios with different data sets. This helps reduce
                code duplication and makes your test suite more maintainable and flexible.
            </p>

            <div class="bg-blue-50 p-4 rounded-lg mb-6">
                <h3 class="font-semibold mb-2 text-blue-800">Key Benefits:</h3>
                <ul class="space-y-2 text-blue-700">
                    <li>• <strong>Parameterized Reusability:</strong> Use the same macro with different data sets</li>
                    <li>• <strong>Table-Driven Testing:</strong> Easy to test multiple scenarios with different
                        parameters</li>
                    <li>• <strong>Improve Maintainability:</strong> Update logic in one place, affects all usages</li>
                    <li>• <strong>Better Organization:</strong> Separate macro definitions from test scenarios</li>
                    <li>• <strong>Dynamic Behavior:</strong> Adapt macros to different test requirements</li>
                </ul>
            </div>
        </section>

        <!-- How Macros Work Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">How Macros Work</h2>

            <div class="grid md:grid-cols-2 gap-6">
                <div>
                    <h3 class="text-lg font-semibold mb-3 text-gray-700">1. Define Parameterized Macros</h3>
                    <p class="text-gray-600 mb-4">
                        Create macro scenarios with variable placeholders using <code>|variable_name|</code> syntax.
                        These placeholders get replaced with actual values during execution.
                    </p>
                    <CodeBlock :code="macroDefinitionExample" language="gherkin" />
                </div>

                <div>
                    <h3 class="text-lg font-semibold mb-3 text-gray-700">2. Use Macros with Parameters</h3>
                    <p class="text-gray-600 mb-4">
                        Invoke macros with data tables containing variable values. The system automatically substitutes
                        placeholders with the provided values.
                    </p>
                    <CodeBlock :code="macroUsageExample" language="gherkin" />
                </div>
            </div>
        </section>

        <!-- Macro Syntax Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Macro Syntax</h2>

            <AccordionItem title="Parameterized Macro Definition">
                <div class="space-y-4">
                    <p>Macros use <code>|variable_name|</code> syntax to define variable placeholders:</p>
                    <CodeBlock :code="macroSyntaxExample" language="gherkin" />

                    <div class="bg-yellow-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-yellow-800">Important Notes:</h4>
                        <ul class="text-yellow-700 space-y-1 text-sm">
                            <li>• Use <code>@macro</code> tag to identify macro scenarios</li>
                            <li>• Use <code>|variable_name|</code> syntax for variable placeholders</li>
                            <li>• Variables are automatically substituted during macro expansion</li>
                            <li>• Support for multiple variables in a single macro</li>
                        </ul>
                    </div>
                </div>
            </AccordionItem>

            <AccordionItem title="Variable Substitution">
                <div class="space-y-4">
                    <p>Variables are automatically replaced with actual values from the data table:</p>
                    <CodeBlock :code="variableSubstitutionExample" language="gherkin" />

                    <div class="bg-green-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-green-800">Substitution Behavior:</h4>
                        <ul class="text-green-700 space-y-1 text-sm">
                            <li>• Variables are replaced in real-time during macro expansion</li>
                            <li>• Table headers define variable names, data rows provide values</li>
                            <li>• The macro scenario name becomes a single step</li>
                            <li>• All steps from the macro are inserted with substituted variables</li>
                        </ul>
                    </div>
                </div>
            </AccordionItem>

            <AccordionItem title="Macro Organization">
                <div class="space-y-4">
                    <p>Organize your macros for better maintainability. Macros are identified by the <code>@macro</code>
                        tag, not by file name:</p>
                    <CodeBlock :code="fileOrganizationExample" language="bash" />

                    <div class="bg-blue-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-blue-800">Best Practices:</h4>
                        <ul class="text-blue-700 space-y-1 text-sm">
                            <li>• Group related macros in the same file</li>
                            <li>• Use descriptive macro names that indicate their purpose</li>
                            <li>• Keep macros focused on single responsibilities</li>
                            <li>• Document variable requirements and expected values</li>
                            <li>• The <code>@macro</code> tag is what defines a macro, not the file name</li>
                            <li>• Use consistent variable naming patterns across related macros</li>
                        </ul>
                    </div>
                </div>
            </AccordionItem>
        </section>

        <!-- Real-World Examples Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Real-World Examples</h2>

            <AccordionItem title="Parameterized Login Macros">
                <div class="space-y-4">
                    <p>Parameterized authentication patterns that can be reused with different credentials:</p>
                    <CodeBlock :code="authMacroExample" language="gherkin" />

                    <p class="text-gray-600">Usage with different data sets:</p>
                    <CodeBlock :code="authMacroUsageExample" language="gherkin" />
                </div>
            </AccordionItem>

            <AccordionItem title="Dynamic Data Setup Macros">
                <div class="space-y-4">
                    <p>Macros for setting up test data with configurable parameters:</p>
                    <CodeBlock :code="dataSetupMacroExample" language="gherkin" />

                    <p class="text-gray-600">Usage with different data sets:</p>
                    <CodeBlock :code="dataSetupMacroUsageExample" language="gherkin" />
                </div>
            </AccordionItem>

            <AccordionItem title="Complex Workflow Macros">
                <div class="space-y-4">
                    <p>Multi-step workflows with multiple configurable parameters:</p>
                    <CodeBlock :code="workflowMacroExample" language="gherkin" />

                    <p class="text-gray-600">Using the workflow macro with different parameters:</p>
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
                        <li>• Use descriptive variable names that are self-explanatory</li>
                        <li>• Group related macros in the same file</li>
                        <li>• Document variable requirements and expected values</li>
                        <li>• Test macros with different variable combinations</li>
                        <li>• Use consistent variable naming patterns across related macros</li>
                        <li>• Keep variables focused and specific to the macro's purpose</li>
                    </ul>
                </div>

                <div class="bg-red-50 p-4 rounded-lg">
                    <h3 class="font-semibold mb-3 text-red-800">❌ Don'ts</h3>
                    <ul class="space-y-2 text-red-700 text-sm">
                        <li>• Don't create overly complex macros with too many variables</li>
                        <li>• Don't use unclear or ambiguous variable names</li>
                        <li>• Don't mix macro and test scenarios in the same file</li>
                        <li>• Don't forget to validate variable values in your tests</li>
                        <li>• Don't create macros that are too specific to one use case</li>
                        <li>• Don't ignore macro reusability across different test scenarios</li>
                        <li>• Don't use too many variables in a single macro (aim for 2-5)</li>
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
                        <p class="text-gray-600 text-sm">Solution: Ensure the scenario has the <code>@macro</code> tag.
                            The tag is what identifies a macro, not the file name.</p>
                    </div>

                    <div class="bg-red-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-red-800">Macro Step Not Found</h4>
                        <p class="text-red-700 text-sm mb-2">Error: "Step 'Login with credentials' not found"</p>
                        <p class="text-gray-600 text-sm">Solution: Ensure the macro scenario name matches exactly with
                            how you're calling it in your test.</p>
                    </div>

                    <div class="bg-red-50 p-4 rounded-lg">
                        <h4 class="font-semibold mb-2 text-red-800">Circular Dependencies</h4>
                        <p class="text-red-700 text-sm mb-2">Error: "Circular dependency detected in macros"</p>
                        <p class="text-gray-600 text-sm">Solution: Avoid macros that reference each other in a circular
                            manner.</p>
                    </div>
                </div>
            </AccordionItem>
        </section>

        <!-- Pro Tips Section -->
        <section class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">Pro Tips</h2>

            <div class="bg-yellow-50 p-4 rounded-lg">
                <h3 class="font-semibold mb-2 text-yellow-800">💡 Best Practices for Success</h3>
                <ul class="text-yellow-700 space-y-1 text-sm">
                    <li>• Use descriptive variable names that clearly indicate their purpose</li>
                    <li>• Keep macros focused on single responsibilities for better reusability</li>
                    <li>• Test macros with different variable combinations to ensure reliability</li>
                    <li>• Organize related macros in dedicated feature files for better maintainability</li>
                    <li>• Document variable requirements and expected values for team collaboration</li>
                    <li>• Use consistent variable naming patterns across related macros</li>
                    <li>• Keep variables focused and specific to the macro's purpose</li>
                    <li>• Aim for 2-5 variables per macro for optimal maintainability</li>
                </ul>
            </div>

            <div class="mt-4 bg-blue-50 p-4 rounded-lg">
                <h3 class="font-semibold mb-2 text-blue-800">🚀 Advanced Usage Patterns</h3>
                <ul class="text-blue-700 space-y-1 text-sm">
                    <li>• Combine multiple variables for complex workflows</li>
                    <li>• Use macros for data-driven testing scenarios</li>
                    <li>• Create reusable authentication and setup patterns</li>
                    <li>• Build consistent test workflows across your test suite</li>
                    <li>• Leverage table-driven testing for multiple data sets</li>
                </ul>
            </div>
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
                <nuxt-link to="/configuration"
                    class="bg-white p-4 rounded-lg shadow-sm hover:shadow-md transition-shadow">
                    <h3 class="font-semibold text-blue-600 mb-2">Configuration</h3>
                    <p class="text-gray-600 text-sm">Configure your test environment and settings</p>
                </nuxt-link>
                <nuxt-link to="/quick-start"
                    class="bg-white p-4 rounded-lg shadow-sm hover:shadow-md transition-shadow">
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

// Example code blocks for macros
const macroDefinitionExample = `@macro
Scenario: user login with credentials
  Given the user is on the login page
  When the user fills the "username" field with |username|
  And the user fills the "password" field with |password|
  And the user clicks the "login" button
  Then the user should be logged in successfully`;

const macroUsageExample = `Scenario: Test login with valid credentials
  When user login with credentials
    | username | password |
    | oki     | ler123   |
  Then the user should be logged in successfully`;

const macroSyntaxExample = `@macro
Scenario: Macro Name with Variables
  Given some precondition with |variable1|
  When some action with |variable2|
  Then some verification for |variable3|`;

const variableSubstitutionExample = `# Macro definition with variables
When the user fills the "username" field with |username|

# Macro invocation with table
| username | password |
| oki     | ler123   |

# Result after substitution
When the user fills the "username" field with oki`;

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
Scenario: user login with credentials
  Given the user is on the login page
  When the user fills the "username" field with |username|
  And the user fills the "password" field with |password|
  And the user clicks the "login" button
  Then the user should be logged in successfully

@macro
Scenario: user logout
  Given the user is logged in
  When the user clicks the "logout" button
  Then the user should be logged out successfully`;

const authMacroUsageExample = `Scenario: Test login with valid credentials
  When user login with credentials
    | username | password |
    | oki     | ler123   |
  Then the user should be logged in successfully

Scenario: Test login with different credentials
  When user login with credentials
    | username | password |
    | admin   | secret   |
  Then the user should be logged in successfully`;

const dataSetupMacroExample = `@macro
Scenario: create test user with role
  Given I prepare a request for the "create_user" endpoint
  When I set the following request body:
    | name  | email           | role    |
    | |user_name| | |user_email| | |user_role| |
  And I send the request
  Then the response status code should be 201`;

const dataSetupMacroUsageExample = `Scenario: Test with admin user
  When create test user with role
    | user_name | user_email        | user_role |
    | Admin     | admin@test.com    | admin     |
  Then the user should see "Admin Dashboard"

Scenario: Test with regular user
  When create test user with role
    | user_name | user_email        | user_role |
    | John      | john@test.com     | user      |
  Then the user should see "User Dashboard"`;

const workflowMacroExample = `@macro
Scenario: complete user workflow
  Given the user is on the login page
  When the user fills the "username" field with |username|
  And the user fills the "password" field with |password|
  And the user clicks the "login" button
  Then the user should be logged in successfully
  When the user navigates to the |page| page
  Then the user should see the |page| page`;

const workflowMacroUsageExample = `Scenario: Test complete workflow
  When complete user workflow
    | username | password | page      |
    | oki     | ler123   | dashboard |
  Then the user should see the dashboard page

Scenario: Test workflow with different page
  When complete user workflow
    | username | password | page      |
    | admin    | secret   | profile   |
  Then the user should see the profile page`;


</script>