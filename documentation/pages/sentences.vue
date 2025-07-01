<template>
    <div>
        <h1 class="heading-1 mb-4">Gherkin Step Definitions</h1>
        <p class="body-large mb-8">
            TestFlowKit provides a comprehensive set of pre-built Gherkin step definitions for frontend and backend
            testing. These steps support both CSS selectors and XPath expressions for flexible element selection.
        </p>

        <div class="card-gradient p-6 rounded-lg mb-8">
            <h2 class="heading-2 mb-4">Element Selector Support</h2>
            <p class="body-normal mb-4">TestFlowKit supports multiple selector types for robust element detection:</p>

            <AccordionItem title="CSS Selectors (Default)">
                <p class="body-normal">Standard CSS selectors are the default selector type in TestFlowKit:</p>
                <ul class="list-disc list-inside mb-4 body-normal">
                    <li><strong>Element IDs:</strong> <code>#element-id</code></li>
                    <li><strong>CSS Classes:</strong> <code>.class-name</code></li>
                    <li><strong>Attribute Selectors:</strong> <code>[data-testid='value']</code></li>
                    <li><strong>Complex Selectors:</strong> <code>div.container > button[type='submit']</code></li>
                </ul>
                <CodeBlock :code="cssExamples" language="yaml" />
            </AccordionItem>

            <AccordionItem title="XPath Selectors">
                <p class="body-normal">TestFlowKit provides full XPath 1.0 support with the <code>xpath:</code> prefix for complex element
                    selection:</p>
                <ul class="list-disc list-inside mb-4 body-normal">
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
                <p class="body-normal">Combine CSS and XPath selectors for maximum flexibility and robustness:</p>
                <CodeBlock :code="mixedExamples" language="yaml" />
            </AccordionItem>
        </div>

        <div class="max-w-6xl mx-auto p-4">
            <SentenceFilter :sentences="allSentences || []" v-model:search-query="searchQuery"
                v-model:selected-category="selectedCategory" @filtered="updateFilteredSentences" />
        </div>

        <div v-if="status === 'pending'" class="text-center mt-8">
            <p class="body-normal">Loading step definitions...</p>
        </div>

        <ClientOnly>
            <div v-if="status === 'success' && filteredSentences.length > 0" class="max-w-6xl mx-auto p-4 text-white">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <SentenceDefinitionCard v-for="(definition, index) in paginatedSentences" v-bind="definition"
                        :key="`${currentPage}-${index}-${definition.sentence}`" />
                </div>

                <div v-if="totalPages > 1" class="mt-8 flex justify-center">
                    <div class="flex items-center gap-2">
                        <button @click="currentPage = Math.max(1, currentPage - 1)" :disabled="currentPage === 1"
                            aria-label="Previous page"
                            class="btn btn-secondary disabled:opacity-50 disabled:cursor-not-allowed">
                            Previous
                        </button>

                        <div class="flex items-center gap-1">
                            <button v-for="page in visiblePages" :key="page" @click="currentPage = page" :class="[
                                'px-3 py-2 text-sm font-medium rounded-md',
                                page === currentPage
                                    ? 'btn-primary'
                                    : 'btn btn-secondary'
                            ]">
                                {{ page }}
                            </button>
                        </div>

                        <button @click="currentPage = Math.min(totalPages, currentPage + 1)"
                            :disabled="currentPage === totalPages"
                            class="btn btn-secondary disabled:opacity-50 disabled:cursor-not-allowed">
                            Next
                        </button>
                    </div>
                </div>
            </div>

            <div v-else-if="status === 'success' && filteredSentences.length === 0" class="text-center mt-8 text-white">
                <p class="body-normal">No step definitions found matching your criteria.</p>
            </div>
        </ClientOnly>

        <div v-if="status === 'error'" class="text-center mt-8">
            <p class="body-normal">Error: {{ error }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
// SEO and Meta
useHead({
  title: 'Gherkin Step Definitions - TestFlowKit Documentation',
  meta: [
    {
      name: 'description',
      content: 'Complete reference of all available Gherkin step definitions in TestFlowKit for frontend and backend testing automation.'
    }
  ]
})

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
</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

#sentences-menu {
    @apply mb-8 flex justify-center;
}
</style>