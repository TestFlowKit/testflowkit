<template>
    <div class="mb-6">
        <div class="flex flex-wrap gap-4 items-center">
            <div class="flex-1 min-w-0">
                <input type="text" placeholder="Search step definitions..." :value="searchQuery" @input="handleSearch"
                    id="search-input"
                    class="w-full px-4 py-2 text-base bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-600 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400" />
            </div>

            <div class="flex items-center gap-2">
                <label for="category-filter" class="text-sm font-medium text-gray-700 dark:text-gray-300 whitespace-nowrap">
                    Category:
                </label>
                <select id="category-filter" :value="selectedCategory" @change="handleCategoryChange"
                    class="px-3 py-2 text-base bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-600 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 text-gray-900 dark:text-white">
                    <option value="">All Categories</option>
                    <option v-for="category in availableCategories" :key="category" :value="category">
                        {{ formatCategoryName(category) }}
                    </option>
                </select>
            </div>
        </div>

        <div v-if="searchQuery || selectedCategory" class="mt-3 flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
            <span>Filtered results: {{ filteredCount }} of {{ totalCount }}</span>
            <button @click="clearFilters" class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 underline">
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
    (e: 'filtered', sentences: Sentence[]): void;
}

const [selectedCategory] = defineModel<string>("selectedCategory", {
    default: '',
});
const [searchQuery] = defineModel<string>("searchQuery", {
    default: '',
});

const { sentences } = defineProps<Props>();
const emit = defineEmits<Emits>();

const availableCategories = computed(() => {
    const allCategories = new Set<string>();
        for (const {categories} of sentences) {
            categories?.forEach(c => allCategories.add(c));
        }
        
    return Array.from(allCategories).sort();
});

const filteredSentences = computed(() => {
    let filtered = sentences;

    if (selectedCategory.value) {
        filtered = filtered.filter(s => s.categories.includes(selectedCategory.value));
    }

    if (searchQuery.value.trim()) {
        const searchTerm = searchQuery.value.toUpperCase();
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
    searchQuery.value = value;
}

function handleCategoryChange(e: Event) {
    const value = (e.target as HTMLSelectElement).value;
    selectedCategory.value = value;
}

function clearFilters() {
    searchQuery.value = '';
    selectedCategory.value = '';
}

function formatCategoryName(category: string): string {
    return category.charAt(0).toUpperCase() + category.slice(1);
}
</script>