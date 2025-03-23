<template>
    <div class="border border-gray-300 rounded-lg shadow-sm mb-4">
        <button type="button"
            class="w-full flex justify-between items-center px-4 py-3 bg-gray-100 hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-t-lg"
            @click="isOpen = !isOpen" aria-expanded="false" :aria-expanded="isOpen">
            <span class="text-lg font-medium">{{ title }}</span>
            <svg class="h-5 w-5 transform transition-transform duration-200" :class="{ 'rotate-90': isOpen }"
                fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
        </button>
        <transition name="accordion">
            <div v-show="isOpen" class="px-4 py-3 bg-white rounded-b-lg">
                <slot />
            </div>
        </transition>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

defineProps({
    title: {
        type: String,
        required: true,
    },
});

const isOpen = ref(false);
</script>

<style scoped>
.accordion-enter-active,
.accordion-leave-active {
    transition: height 0.3s ease-in-out;
}

.accordion-enter-from,
.accordion-leave-to {
    height: 0;
    overflow: hidden;
}
</style>
