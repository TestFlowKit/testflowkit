<template>
    <main class="pb-12 px-4 sm:px-6 lg:px-8 max-w-7xl mx-auto pt-4">
        <h1 class="text-2xl sm:text-3xl font-extrabold mb-4 text-gray-900 dark:text-white">Gherkin Step Definitions</h1>
        <p class="text-base sm:text-lg mb-6 sm:mb-8 text-gray-600 dark:text-gray-300">
            TestFlowKit provides a comprehensive set of pre-built Gherkin step definitions for frontend and backend
            testing. These steps support both CSS selectors and XPath expressions for flexible element selection.
        </p>

        <div class="bg-blue-50 dark:bg-blue-900/30 p-3 sm:p-4 rounded-lg mb-4 sm:mb-6">
            <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-2">
                <div>
                    <h2 class="text-base sm:text-lg font-semibold text-blue-800 dark:text-blue-300">Element Selector Support</h2>
                    <p class="text-xs sm:text-sm text-blue-700 dark:text-blue-400">CSS selectors (default) and XPath expressions with
                        <code class="bg-blue-100 dark:bg-blue-800 px-1 rounded">xpath:</code> prefix
                    </p>
                </div>
                <nuxt-link to="/docs/concepts/selectors" class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 text-sm font-medium whitespace-nowrap">
                    View Examples →
                </nuxt-link>
            </div>
        </div>

        <div class="bg-purple-50 dark:bg-purple-900/30 p-3 sm:p-4 rounded-lg mb-4 sm:mb-6">
            <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-2">
                <div>
                    <h2 class="text-base sm:text-lg font-semibold text-purple-800 dark:text-purple-300">Macro System</h2>
                    <p class="text-xs sm:text-sm text-purple-700 dark:text-purple-400">Reusable test scenarios with parameter support to reduce code
                        duplication</p>
                </div>
                <nuxt-link to="/docs/features/macros" class="text-purple-600 dark:text-purple-400 hover:text-purple-800 dark:hover:text-purple-300 text-sm font-medium whitespace-nowrap">
                    Learn More →
                </nuxt-link>
            </div>
        </div>

        <div class="bg-green-50 dark:bg-green-900/30 p-3 sm:p-4 rounded-lg mb-4 sm:mb-6">
            <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-2">
                <div>
                    <h2 class="text-base sm:text-lg font-semibold text-green-800 dark:text-green-300">Variables Support</h2>
                    <p class="text-xs sm:text-sm text-green-700 dark:text-green-400">Dynamic test data with
                        <code class="bg-green-100 dark:bg-green-800 px-1 rounded">&#123;&#123;variable_name&#125;&#125;</code> syntax for cross-step data sharing
                    </p>
                </div>
                <nuxt-link to="/docs/features/variables" class="text-green-600 dark:text-green-400 hover:text-green-800 dark:hover:text-green-300 text-sm font-medium whitespace-nowrap">
                    Learn More →
                </nuxt-link>
            </div>
        </div>

        <div class="mb-4 sm:mb-6">
            <SentenceFilter :sentences="allSentences || []" v-model:search-query="searchQuery"
                v-model:selected-category="selectedCategory" @filtered="updateFilteredSentences" />
        </div>

        <div v-if="status === 'pending'" class="text-center mt-8">
            <p class="text-gray-600 dark:text-gray-400">Loading step definitions...</p>
        </div>

        <ClientOnly>
           <div id="available-sentences">
             <div v-show="status === 'success' && filteredSentences.length > 0">
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
                    <SentenceDefinitionCard v-for="(definition, index) in paginatedSentences" v-bind="definition"
                        :key="`${currentPage}-${index}-${definition.sentence}`" />
                </div>

                <Pagination 
                    v-model:current-page="currentPage" 
                    :total-pages="totalPages"
                    :sibling-count="1"
                    :boundary-count="1"
                />
            </div>

                <p v-show="status === 'success' && filteredSentences.length === 0" class="text-gray-600 dark:text-gray-400 text-center mt-8">No step definitions found matching your criteria.</p>
           </div>
        </ClientOnly>

        <div v-if="status === 'error'" class="text-center mt-8">
            <p class="text-red-600 dark:text-red-400">Error: {{ error }}</p>
        </div>
    </main>
</template>

<script setup lang="ts">
import SentenceDefinitionCard from '../components/SentenceDefinitionCard.vue';
import SentenceFilter from '../components/SentenceFilter.vue';
import Pagination from '../components/Pagination.vue';
import type { Sentence } from '~/data/sentence';



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




</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

#sentences-menu {
    @apply mb-8 flex justify-center;
}
</style>