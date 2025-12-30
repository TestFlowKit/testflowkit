<template>
    <nav v-if="totalPages > 1" class="mt-6 sm:mt-8 flex justify-center" aria-label="Pagination">
        <div class="flex flex-wrap items-center justify-center gap-1 sm:gap-2">
            <!-- Previous Button -->
            <button 
                @click="goToPage(currentPage - 1)" 
                :disabled="currentPage === 1"
                aria-label="Previous page"
                class="px-2 sm:px-3 py-1.5 sm:py-2 text-xs sm:text-sm font-medium text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded-md hover:bg-gray-50 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
                <span class="hidden sm:inline">Previous</span>
                <svg class="w-4 h-4 sm:hidden" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
                </svg>
            </button>

            <!-- Page Numbers -->
            <template v-for="(page, index) in displayedPages" :key="index">
                <!-- Ellipsis -->
                <span 
                    v-if="page === '...'" 
                    class="px-2 py-1.5 text-xs sm:text-sm text-gray-400 dark:text-gray-500"
                >
                    ...
                </span>
                
                <!-- Page Number Button -->
                <button 
                    v-else
                    @click="goToPage(page as number)"
                    :aria-label="`Go to page ${page}`"
                    :aria-current="page === currentPage ? 'page' : undefined"
                    :class="[
                        'px-2.5 sm:px-3.5 py-1.5 sm:py-2 text-xs sm:text-sm font-medium rounded-md transition-colors',
                        page === currentPage
                            ? 'bg-blue-600 text-white shadow-sm'
                            : 'text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700'
                    ]"
                >
                    {{ page }}
                </button>
            </template>

            <!-- Next Button -->
            <button 
                @click="goToPage(currentPage + 1)"
                :disabled="currentPage === totalPages"
                aria-label="Next page"
                class="px-2 sm:px-3 py-1.5 sm:py-2 text-xs sm:text-sm font-medium text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded-md hover:bg-gray-50 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
                <span class="hidden sm:inline">Next</span>
                <svg class="w-4 h-4 sm:hidden" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
            </button>

            <div class="w-full text-center mt-2 sm:hidden text-xs text-gray-500 dark:text-gray-400">
                Page {{ currentPage }} of {{ totalPages }}
            </div>
        </div>
    </nav>
</template>

<script setup lang="ts">
interface Props {
    currentPage: number;
    totalPages: number;
    siblingCount?: number; // How many pages to show on each side of current page
    boundaryCount?: number; // How many pages to show at start and end
}

interface Emits {
    (e: 'update:currentPage', value: number): void;
}

const props = withDefaults(defineProps<Props>(), {
    siblingCount: 1,
    boundaryCount: 1
});

const emit = defineEmits<Emits>();

// Generate the array of page numbers and ellipses to display
const displayedPages = computed(() => {
    const { currentPage, totalPages, siblingCount, boundaryCount } = props;
    
    // If total pages is small enough, show all pages
    const totalPageNumbers = boundaryCount * 2 + siblingCount * 2 + 3; // +3 for current page and two ellipses
    if (totalPages <= totalPageNumbers) {
        return Array.from({ length: totalPages }, (_, i) => i + 1);
    }

    const pages: (number | string)[] = [];
    
    const ranges = calculateRanges({ currentPage, siblingCount, boundaryCount, totalPages });
    const { showLeftEllipsis, leftSiblingIndex, rightSiblingIndex, showRightEllipsis } = ranges;
    // Add boundary pages at the start
    
    for (let i = 1; i <= boundaryCount; i++) {
        pages.push(i);
    }

    // Add left ellipsis or the page after boundary
    if (showLeftEllipsis) {
        pages.push('...');
    } else {
        // Add pages between boundary and sibling range
        for (let i = boundaryCount + 1; i < leftSiblingIndex; i++) {
            pages.push(i);
        }
    }

    // Add sibling pages and current page
    for (let i = leftSiblingIndex; i <= rightSiblingIndex; i++) {
        if (i > boundaryCount && i <= totalPages - boundaryCount) {
            pages.push(i);
        }
    }

    // Add right ellipsis or pages before end boundary
    if (showRightEllipsis) {
        pages.push('...');
    } else {
        // Add pages between sibling range and end boundary
        for (let i = rightSiblingIndex + 1; i <= totalPages - boundaryCount; i++) {
            pages.push(i);
        }
    }

    // Add boundary pages at the end
    for (let i = totalPages - boundaryCount + 1; i <= totalPages; i++) {
        if (i > boundaryCount) {
            pages.push(i);
        }
    }

    // Remove duplicates while preserving order
    const seen = new Set<number | string>();
    return pages.filter(page => {
        if (page === '...' || !seen.has(page)) {
            if (page !== '...') seen.add(page);
            return true;
        }
        return false;
    });
});

function calculateRanges(params: CalculateRangesParams): CalculateRangesReturn {
    const { currentPage, siblingCount, boundaryCount, totalPages } = params;
    const leftSiblingIndex = Math.max(currentPage - siblingCount, boundaryCount + 1);
    const rightSiblingIndex = Math.min(currentPage + siblingCount, totalPages - boundaryCount);

    const showLeftEllipsis = leftSiblingIndex > boundaryCount + 2;
    const showRightEllipsis = rightSiblingIndex < totalPages - boundaryCount - 1;
    return { showLeftEllipsis, leftSiblingIndex, rightSiblingIndex, showRightEllipsis };
}



function goToPage(page: number) {
    if (page >= 1 && page <= props.totalPages && page !== props.currentPage) {
        emit('update:currentPage', page);
    }
}
</script>

<script lang="ts">
    type CalculateRangesReturn = {
        showLeftEllipsis: boolean;
        leftSiblingIndex: number;
        rightSiblingIndex: number;
        showRightEllipsis: boolean;
    };

    type CalculateRangesParams = {
       currentPage: number; siblingCount: number; boundaryCount: number; totalPages: number;
}
</script>
