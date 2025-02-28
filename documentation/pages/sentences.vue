<template>

    <div class="max-w-4xl mx-auto p-4">
        <div class="relative">
            <input type="text" placeholder="Search..." ref="searchInput" @input="search"
                class="w-full px-6 py-4 text-lg bg-white border-2 border-gray-300 rounded-full shadow-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200" />

            <div class="absolute inset-y-0 right-0 flex items-center pr-5">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-500" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
            </div>
        </div>
    </div>

    <div v-if="status === 'pending'" class="text-center mt-8">
        <p>Loading...</p>
    </div>
    <ClientOnly>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4" v-if="status === 'success'">
            <SentenceDefinitionCard v-for="definition in data" v-bind="definition" :key="definition" />
        </div>
    </ClientOnly>


    <div v-if="status === 'error'" class="text-center mt-8">
        <p>Error: {{ error }}</p>
    </div>

</template>



<script setup lang="ts">
import SentenceDefinitionCard from '../components/SentenceDefinitionCard.vue';
const { data, status, error } = await useAsyncData("get-sentences", () => queryCollection('sentence').all())
const searchInput = ref<HTMLInputElement>();


async function search(e: Event) {
    const value = ((e as InputEvent).target as HTMLInputElement).value;
    const queryBuilder = queryCollection('sentence')

    if (value.trim() === '') {
        data.value = await queryBuilder.all();
        return;
    }

    const searchValue = `%${value}%`
    const result = await queryBuilder
        .orWhere(query => {
            return query.where('description', 'LIKE', searchValue)
                .where('sentence', 'LIKE', searchValue)
        })
        .all();

    data.value = result;
}
</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

#sentences-menu {
    @apply mb-8 flex justify-center;
}
</style>