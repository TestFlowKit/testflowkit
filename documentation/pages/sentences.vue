<template>
    <div>
        <h1 class="text-3xl font-extrabold mb-4">Gherkin Step Definitions</h1>
        <p class="text-lg mb-8 text-gray-600">
            TestFlowKit provides a comprehensive set of pre-built Gherkin step definitions for frontend and backend
            testing. These steps support both CSS selectors and XPath expressions for flexible element selection.
        </p>

        <div class="bg-blue-100 p-6 rounded-lg mb-8">
            <h2 class="text-2xl font-semibold mb-4">Element Selector Support</h2>
            <p class="mb-4">TestFlowKit supports multiple selector types for robust element detection:</p>

            <AccordionItem title="CSS Selectors (Default)">
                <p>Standard CSS selectors are the default selector type in TestFlowKit:</p>
                <ul class="list-disc list-inside mb-4">
                    <li><strong>Element IDs:</strong> <code>#element-id</code></li>
                    <li><strong>CSS Classes:</strong> <code>.class-name</code></li>
                    <li><strong>Attribute Selectors:</strong> <code>[data-testid='value']</code></li>
                    <li><strong>Complex Selectors:</strong> <code>div.container > button[type='submit']</code></li>
                </ul>
                <CodeBlock :code="cssExamples" language="yaml" />
            </AccordionItem>

            <AccordionItem title="XPath Selectors">
                <p>TestFlowKit provides full XPath 1.0 support with the <code>xpath:</code> prefix for complex element
                    selection:</p>
                <ul class="list-disc list-inside mb-4">
                    <li><strong>Element Selection:</strong> <code>xpath://div[@class='container']</code></li>
                    <li><strong>Text Matching:</strong> <code>xpath://button[contains(text(), 'Submit')]</code></li>
                    <li><strong>Attribute Conditions:</strong> <code>xpath://input[@type='email' and @required]</code>
                    </li>
                    <li><strong>Complex Expressions:</strong>
                        <code>xpath://div[contains(@class, 'dynamic') and contains(text(), 'Loading')]</code>
                    </li>
                </ul>
                <CodeBlock :code="xpathExamples" language="yaml" />
            </AccordionItem>

            <AccordionItem title="Mixed Selectors with Fallback">
                <p>Combine CSS and XPath selectors for maximum flexibility and robustness:</p>
                <CodeBlock :code="mixedExamples" language="yaml" />
            </AccordionItem>
        </div>

        <div class="bg-purple-100 p-6 rounded-lg mb-8">
            <div class="flex justify-between items-start mb-4">
                <div>
                    <h2 class="text-2xl font-semibold">Macro System</h2>
                    <p class="mt-2">TestFlowKit provides a powerful macro system for creating reusable test scenarios and reducing code duplication:</p>
                </div>
                <nuxt-link to="/macros"
                    class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors">
                    Learn More About Macros
                </nuxt-link>
            </div>

            <AccordionItem title="Macro Basics">
                <p>Macros allow you to define reusable test scenarios with parameter support:</p>
                <ul class="list-disc list-inside mb-4">
                    <li><strong>Reusable Scenarios:</strong> Define common test patterns once</li>
                    <li><strong>Parameter Support:</strong> Pass different data to the same macro</li>
                    <li><strong>Better Organization:</strong> Separate macro definitions from test scenarios</li>
                    <li><strong>Maintainability:</strong> Update logic in one place</li>
                </ul>
                <CodeBlock :code="macroBasicsExample" language="gherkin" />
            </AccordionItem>

            <AccordionItem title="Macro Usage">
                <p>Use macros in your test scenarios by referencing the macro scenario name:</p>
                <CodeBlock :code="macroUsageExample" language="gherkin" />
            </AccordionItem>

            <AccordionItem title="Macro Organization">
                <p>Organize your macro files for better maintainability:</p>
                <CodeBlock :code="macroOrganizationExample" language="bash" />
            </AccordionItem>
        </div>

        <div class="bg-green-100 p-6 rounded-lg mb-8">
            <div class="flex justify-between items-start mb-4">
                <div>
                    <h2 class="text-2xl font-semibold">Variables Support</h2>
                    <p class="mt-2">TestFlowKit provides powerful variable support for dynamic test data and cross-step
                        data sharing:</p>
                </div>
                <nuxt-link to="/variables"
                    class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                    Learn More About Variables
                </nuxt-link>
            </div>

            <AccordionItem title="Variable Syntax">
                <p>Variables in TestFlowKit use the <code>&#123;&#123;variable_name&#125;&#125;</code> syntax and are
                    automatically substituted in all step parameters:</p>
                <ul class="list-disc list-inside mb-4">
                    <li><strong>Variable Declaration:</strong> <code>&#123;&#123;variable_name&#125;&#125;</code></li>
                    <li><strong>Automatic Substitution:</strong> Variables are replaced in strings, tables, and
                        parameters</li>
                    <li><strong>Scope:</strong> Variables persist throughout the entire scenario</li>
                    <li><strong>Type Support:</strong> Supports strings, numbers, booleans, and complex data structures
                    </li>
                </ul>
                <CodeBlock :code="variableSyntaxExamples" language="gherkin" />
            </AccordionItem>

            <AccordionItem title="Variable Types">
                <p>TestFlowKit supports three main types of variable storage:</p>
                <ul class="list-disc list-inside mb-4">
                    <li><strong>Custom Variables:</strong> Store any custom value for reuse</li>
                    <li><strong>JSON Path Variables:</strong> Extract data from API responses</li>
                    <li><strong>HTML Element Variables:</strong> Capture content from web page elements</li>
                </ul>
                <CodeBlock :code="variableTypesExamples" language="gherkin" />
            </AccordionItem>

            <AccordionItem title="Advanced Variable Usage">
                <p>Variables can be used in complex scenarios for data-driven testing:</p>
                <CodeBlock :code="advancedVariableExamples" language="gherkin" />
            </AccordionItem>
        </div>

        <div class="mb-6">
            <SentenceFilter :sentences="allSentences || []" v-model:search-query="searchQuery"
                v-model:selected-category="selectedCategory" @filtered="updateFilteredSentences" />
        </div>

        <div v-if="status === 'pending'" class="text-center mt-8">
            <p>Loading step definitions...</p>
        </div>

        <ClientOnly>
            <div v-if="status === 'success' && filteredSentences.length > 0">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <SentenceDefinitionCard v-for="(definition, index) in paginatedSentences" v-bind="definition"
                        :key="`${currentPage}-${index}-${definition.sentence}`" />
                </div>

                <div v-if="totalPages > 1" class="mt-8 flex justify-center">
                    <div class="flex items-center gap-2">
                        <button @click="currentPage = Math.max(1, currentPage - 1)" :disabled="currentPage === 1"
                            aria-label="Previous page"
                            class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
                            Previous
                        </button>

                        <div class="flex items-center gap-1">
                            <button v-for="page in visiblePages" :key="page" @click="currentPage = page" :class="[
                                'px-3 py-2 text-sm font-medium rounded-md',
                                page === currentPage
                                    ? 'bg-blue-600 text-white'
                                    : 'text-gray-500 bg-white border border-gray-300 hover:bg-gray-50'
                            ]">
                                {{ page }}
                            </button>
                        </div>

                        <button @click="currentPage = Math.min(totalPages, currentPage + 1)"
                            :disabled="currentPage === totalPages"
                            class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
                            Next
                        </button>
                    </div>
                </div>
            </div>

            <div v-else-if="status === 'success' && filteredSentences.length === 0" class="text-center mt-8">
                <p class="text-gray-600">No step definitions found matching your criteria.</p>
            </div>
        </ClientOnly>

        <div v-if="status === 'error'" class="text-center mt-8">
            <p>Error: {{ error }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import SentenceDefinitionCard from '../components/SentenceDefinitionCard.vue';
import SentenceFilter from '../components/SentenceFilter.vue';
import type { Sentence } from '~/data/sentence';
import AccordionItem from '../components/AccordionItem.vue';
import CodeBlock from '../components/global/CodeBlock.vue';

const { $rehighlight } = useNuxtApp();

const { data: allSentences, status, error } = await useAsyncData("get-sentences", () => queryCollection('sentence').all());

const searchQuery = ref('');
const selectedCategory = ref('');
const filteredSentences = ref<Sentence[]>([]);

const ITEMS_PER_PAGE = 12;
const currentPage = ref(1);
const itemsPerPage = ITEMS_PER_PAGE;

function updateFilteredSentences(sentences: Sentence[]) {
    filteredSentences.value = sentences;
    currentPage.value = 1;
}

const totalPages = computed(() => Math.ceil(filteredSentences.value.length / itemsPerPage));

const paginatedSentences = computed(() => {
    const startIndex = (currentPage.value - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    return filteredSentences.value.slice(startIndex, endIndex);
});

const visiblePages = computed(() => {
    const total = totalPages.value;
    const current = currentPage.value;
    const delta = 2;

    let start = Math.max(1, current - delta);
    let end = Math.min(total, current + delta);

    if (current - delta <= 1) {
        end = Math.min(total, 1 + 2 * delta);
    }
    if (current + delta >= total) {
        start = Math.max(1, total - 2 * delta);
    }

    const pages = [];
    for (let i = start; i <= end; i++) {
        pages.push(i);
    }
    return pages;
});

watch(allSentences, (newSentences) => {
    if (newSentences) {
        filteredSentences.value = newSentences;
    }
}, { immediate: true });

watch([currentPage, paginatedSentences], () => {
    nextTick(() => {
        $rehighlight();
    });
});

const cssExamples = `
frontend:
  elements:
    login_page:
      # CSS Selector Examples
      email_field:
        - "[data-testid='email-input']"
        - "input[name='email']"
        - "#email"
      password_field:
        - "input[type='password']"
        - "#password"
      submit_button:
        - "button[type='submit']"
        - ".btn-primary"
`.trim();

const xpathExamples = `
frontend:
  elements:
    login_page:
      # XPath Selector Examples
      complex_button:
        - "xpath://button[contains(@class, 'submit') and text()='Login']"
        - "xpath://div[@id='login-form']//button[@type='submit']"
      dynamic_text:
        - "xpath://div[contains(text(), 'Welcome') and @class='message']"
        - "xpath://span[text()='Hello, User!']"
      required_field:
        - "xpath://input[@type='email' and @required]"
        - "xpath://input[@name='email' and @data-required='true']"
`.trim();

const mixedExamples = `
frontend:
  elements:
    login_page:
      # Mixed CSS and XPath selectors
      flexible_element:
        - "xpath://div[contains(@class, 'dynamic') and contains(text(), 'Loading')]"
        - ".loading-indicator"
        - "[data-testid='loading']"
      complex_form:
        - "xpath://form[@id='login-form']//input[@type='email']"
        - "input[name='email']"
        - "#email"
`.trim();

const variableSyntaxExamples = `
# Basic variable usage
Given I store the "John Doe" into "user_name" variable
When the user enters "{{user_name}}" into the "name" field

# Variables in tables
When I set the following path params:
  | id | {{user_id}} |
  | name | {{user_name}} |

# Variables in assertions
Then the "welcome_message" should contain "{{user_name}}"
`.trim();

const variableTypesExamples = `
# Custom Variables
When I store the "test@example.com" into "email" variable

# JSON Path Variables
When I store the JSON path "data.user.id" from the response into "user_id" variable

# HTML Element Variables
When I store the content of "page_title" into "title" variable
`.trim();

const advancedVariableExamples = `
Scenario: End-to-end data flow with variables
  Given I prepare a request for the "get_user" endpoint
  And I set the following path params:
    | id | 123 |
  When I send the request
  And I store the JSON path "data.name" from the response into "api_user_name" variable
  And I store the JSON path "data.email" from the response into "api_user_email" variable
  Then the response status code should be 200
  
  When the user goes to the "profile" page
  And I store the content of "displayed_name" into "page_user_name" variable
  And the user enters "{{api_user_email}}" into the "email" field
  Then the "email" field should contain "{{api_user_email}}"
  And the "page_user_name" should equal "{{api_user_name}}"
`.trim();

const macroBasicsExample = `
@macro
Scenario: Login with credentials
  Given the user is on the homepage
  When the user goes to the "login" page
  And the user enters "test@example.com" into the "email" field
  And the user enters "password123" into the "password" field
  And the user clicks on the "login" button
`.trim();

const macroUsageExample = `
Scenario: Test with macro
  Given Login with credentials
  Then the user should be navigated to "dashboard" page
`.trim();

const macroOrganizationExample = `e2e/features/
├── macros/
│   ├── authentication.feature  # Contains @macro scenarios
│   ├── data-setup.feature      # Contains @macro scenarios
│   └── workflows.feature       # Contains @macro scenarios
└── tests/
    ├── login.feature           # Regular test scenarios
    └── user-management.feature # Regular test scenarios

# Note: The @macro tag identifies macro scenarios, not the file name`.trim();
</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

#sentences-menu {
    @apply mb-8 flex justify-center;
}
</style>