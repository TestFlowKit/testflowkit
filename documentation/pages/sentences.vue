<template>
    <div class="mb-8">
        <h1 class="text-3xl font-extrabold mb-4">Gherkin Step Definitions</h1>
        <p class="text-lg mb-6 text-gray-600">
            TestFlowKit provides a comprehensive set of pre-built Gherkin step definitions for frontend and backend
            testing.
            Search through the available steps to find the right ones for your test scenarios.
        </p>
    </div>

    <div class="max-w-6xl mx-auto p-4">
        <SentenceFilter :sentences="allSentences || []" v-model:search-query="searchQuery"
            v-model:selected-category="selectedCategory" @filtered="updateFilteredSentences" />
    </div>

    <div v-if="status === 'pending'" class="text-center mt-8">
        <p>Loading step definitions...</p>
    </div>

    <ClientOnly>
        <div v-if="status === 'success' && filteredSentences.length > 0" class="max-w-6xl mx-auto p-4">
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
</template>

<script setup lang="ts">
import SentenceDefinitionCard from '../components/SentenceDefinitionCard.vue';
import SentenceFilter from '../components/SentenceFilter.vue';
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

</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

#sentences-menu {
    @apply mb-8 flex justify-center;
}
</style>