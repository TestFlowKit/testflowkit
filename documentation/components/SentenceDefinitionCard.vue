<template>
    <div class="card p-6 space-y-4">
        <div class="flex items-start justify-between">
            <h2 class="heading-4 flex-1" v-text="sentence"></h2>
            <span class="ml-2 px-3 py-1 text-xs font-medium bg-orange-100 text-orange-700 rounded-full capitalize border border-orange-200">
                {{ category }}
            </span>
        </div>
        <div class="description">
            <h3 class="font-bold body-normal inline">Description: </h3>
            <span class="body-normal">{{ description }}</span>
        </div>

        <div class="variables">
            <h3 class="font-bold body-normal inline">Variables: </h3>
            <span v-if="!variables?.length" class="body-small">None</span>
            <template v-else>
                <ul class="grid grid-cols-1 gap-4 mt-2" :class="{ 'grid-cols-2': variables.length > 1 }">
                    <li v-for="variable in variables" :key="variable.name" class="body-normal">
                        <span class="font-medium">{{ variable.name }}</span>
                        <span class="body-small">({{ variable.type }})</span>
                        <p class="body-small">{{ variable.description }}</p>
                    </li>
                </ul>
            </template>
        </div>

        <div>
            <h3 class="font-bold body-normal">Example:</h3>
            <code-block language='gherkin' :code="gherkinExample.trim()" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { Sentence } from '~/data/sentence';

defineProps<Sentence>();

</script>
