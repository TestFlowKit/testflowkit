<template>
  <nav class="mt-8 sm:mt-12 pt-4 sm:pt-6 border-t border-gray-200 dark:border-gray-700" aria-label="Page navigation">
    <div class="flex flex-col sm:flex-row justify-between items-stretch sm:items-center gap-3 sm:gap-4">
      <!-- Previous Link -->
      <div class="flex-1">
        <NuxtLink
          v-if="prev"
          :to="prev.path"
          class="group flex flex-col items-start p-3 sm:p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors border border-transparent hover:border-gray-200 dark:hover:border-gray-700"
        >
          <span class="text-xs sm:text-sm text-gray-500 dark:text-gray-400 mb-1 flex items-center">
            <svg class="w-3 h-3 sm:w-4 sm:h-4 mr-1 transition-transform group-hover:-translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            Previous
          </span>
          <span class="text-sm sm:text-base text-blue-600 dark:text-blue-400 font-medium group-hover:text-blue-700 dark:group-hover:text-blue-300">
            {{ prev.title }}
          </span>
        </NuxtLink>
      </div>

      <!-- Next Link -->
      <div class="flex-1 flex justify-end">
        <NuxtLink
          v-if="next"
          :to="next.path"
          class="group flex flex-col items-end p-3 sm:p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors border border-transparent hover:border-gray-200 dark:hover:border-gray-700 w-full sm:w-auto"
        >
          <span class="text-xs sm:text-sm text-gray-500 dark:text-gray-400 mb-1 flex items-center">
            Next
            <svg class="w-3 h-3 sm:w-4 sm:h-4 ml-1 transition-transform group-hover:translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </span>
          <span class="text-sm sm:text-base text-blue-600 dark:text-blue-400 font-medium group-hover:text-blue-700 dark:group-hover:text-blue-300">
            {{ next.title }}
          </span>
        </NuxtLink>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
interface NavItem {
  path: string;
  title: string;
}

defineProps<{
  prev?: NavItem | null;
  next?: NavItem | null;
}>();
</script>
