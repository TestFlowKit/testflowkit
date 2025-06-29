<template>
    <div class="bg-white p-4 rounded-md shadow-md space-y-2">
        <div class="flex items-start justify-between">
            <h2 class="font-bold text-xl text-gray-800 flex-1" v-text="sentence"></h2>
            <span class="ml-2 px-2 py-1 text-xs font-medium bg-blue-100 text-blue-800 rounded-full capitalize">
                {{ category }}
            </span>
        </div>
        <div class="description">
            <h3 class="font-bold text-gray-700 inline">Description: </h3>
            <span class="text-gray-600">{{ description }}</span>
        </div>

        <div class="variables">
            <h3 class="font-bold text-gray-700 inline">Variables: </h3>
            <span v-if="!variables?.length" class="text-gray-500">None</span>
            <template v-else>
                <ul class="grid grid-cols-1 gap-4 mt-2" :class="{ 'grid-cols-2': variables.length > 1 }">
                    <li v-for="variable in variables" :key="variable.name" class="text-gray-600">
                        <span class="font-medium">{{ variable.name }}</span>
                        <span class="text-gray-400">({{ variable.type }})</span>
                        <p class="text-gray-500">{{ variable.description }}</p>
                    </li>
                </ul>
            </template>
        </div>

        <div>
            <h3 class="font-bold text-gray-700">Example:</h3>
            <code-block language='gherkin' :code="gherkinExample.trim()" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { Sentence } from '~/data/sentence';

defineProps<Sentence>();

</script>
