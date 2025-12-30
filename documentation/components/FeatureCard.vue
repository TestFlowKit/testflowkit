<template>
    <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md border border-gray-100 dark:border-gray-700">
        <div class="flex items-center mb-4 text-gray-900 dark:text-white font-semibold">
            <i :class="['text-blue-500 dark:text-blue-400 text-2xl mr-4', icon]"></i> {{ title }}
        </div>
        <div class="text-gray-600 dark:text-gray-300">
            <slot />
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps({
    title: { type: String, required: true },
    icon: { type: String, required: true }
});
</script>
