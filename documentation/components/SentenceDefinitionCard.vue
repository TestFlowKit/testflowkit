<template>
    <div class="bg-white dark:bg-gray-800 p-4 rounded-md shadow-md space-y-2 border border-gray-100 dark:border-gray-700">
        <div class="flex items-start justify-between">
            <h2 class="font-bold text-xl text-gray-800 dark:text-white flex-1" v-text="sentence"></h2>
            <span class="ml-2 px-2 py-1 text-xs font-medium bg-blue-100 dark:bg-blue-900/50 text-blue-800 dark:text-blue-300 rounded-full capitalize">
                {{ category }}
            </span>
        </div>
        <div class="description">
            <h3 class="font-bold text-gray-700 dark:text-gray-300 inline">Description: </h3>
            <span class="text-gray-600 dark:text-gray-400">{{ description }}</span>
        </div>

        <div class="variables">
            <h3 class="font-bold text-gray-700 dark:text-gray-300 inline">Variables: </h3>
            <span v-if="!variables?.length" class="text-gray-500 dark:text-gray-500">None</span>
            <template v-else>
                <ul class="grid grid-cols-1 gap-4 mt-2" :class="{ 'grid-cols-2': variables.length > 1 }">
                    <li v-for="variable in variables" :key="variable.name" class="text-gray-600 dark:text-gray-400">
                        <span class="font-medium text-gray-700 dark:text-gray-300">{{ variable.name }}</span>
                        <span class="text-gray-400 dark:text-gray-500">({{ variable.type }})</span>
                        <p class="text-gray-500 dark:text-gray-500">{{ variable.description }}</p>
                    </li>
                </ul>
            </template>
        </div>

        <div>
            <h3 class="font-bold text-gray-700 dark:text-gray-300">Example:</h3>
            <code-block language='gherkin' :code="gherkinExample.trim()" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { Sentence } from '~/data/sentence';

defineProps<Sentence>();

</script>
