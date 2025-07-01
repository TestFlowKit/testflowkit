<template>
    <div class="mb-6">
        <div class="flex flex-wrap gap-4 items-center">
            <div class="flex-1 min-w-0">
                <input type="text" placeholder="Search step definitions..." :value="searchQuery" @input="handleSearch"
                    class="w-full px-4 py-3 text-base bg-white border-2 border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all duration-200 body-normal" />
            </div>

            <div class="flex items-center gap-2">
                <label for="category-filter" class="body-normal font-medium whitespace-nowrap">
                    Category:
                </label>
                <select id="category-filter" :value="selectedCategory" @change="handleCategoryChange"
                    class="px-3 py-3 text-base bg-white border-2 border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all duration-200 body-normal">
                    <option value="">All Categories</option>
                    <option v-for="category in availableCategories" :key="category" :value="category">
                        {{ formatCategoryName(category) }}
                    </option>
                </select>
            </div>
        </div>

        <div v-if="searchQuery || selectedCategory" class="mt-3 flex items-center gap-2 body-small">
            <span>Filtered results: {{ filteredCount }} of {{ totalCount }}</span>
            <button @click="clearFilters" class="text-orange-600 hover:text-orange-800 underline">
                Clear filters
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { Sentence } from '~/data/sentence';

interface Props {
    sentences: Sentence[];
    searchQuery: string;
    selectedCategory: string;
}

interface Emits {
    (e: 'update:searchQuery', value: string): void;
    (e: 'update:selectedCategory', value: string): void;
    (e: 'filtered', sentences: Sentence[]): void;
}

const { sentences, searchQuery, selectedCategory } = defineProps<Props>();
const emit = defineEmits<Emits>();

const availableCategories = computed(() => {
    const categories = new Set(sentences.map(s => s.category));
    return Array.from(categories).sort();
});

const filteredSentences = computed(() => {
    let filtered = sentences;

    if (selectedCategory) {
        filtered = filtered.filter(s => s.category === selectedCategory);
    }

    if (searchQuery.trim()) {
        const searchTerm = searchQuery.toUpperCase();
        filtered = filtered.filter(s => {
            const upperCaseTexts = [s.sentence, s.description].map(t => t.toUpperCase());
            return upperCaseTexts.some(upperText => upperText.includes(searchTerm));
        });
    }

    return filtered;
});

watch(filteredSentences, (newFiltered) => {
    emit('filtered', newFiltered);
}, { immediate: true });

const filteredCount = computed(() => filteredSentences.value.length);
const totalCount = computed(() => sentences.length);

function handleSearch(e: Event) {
    const value = (e.target as HTMLInputElement).value;
    emit('update:searchQuery', value);
}

function handleCategoryChange(e: Event) {
    const value = (e.target as HTMLSelectElement).value;
    emit('update:selectedCategory', value);
}

function clearFilters() {
    emit('update:searchQuery', '');
    emit('update:selectedCategory', '');
}

function formatCategoryName(category: string): string {
    return category.charAt(0).toUpperCase() + category.slice(1);
}
</script>