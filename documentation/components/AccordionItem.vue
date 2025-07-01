<template>
    <div class="card mb-4 overflow-hidden">
        <button type="button"
            class="w-full flex justify-between items-center px-6 py-4 bg-orange-50 hover:bg-orange-100 focus:outline-none focus:ring-2 focus:ring-orange-500 border-b border-orange-200 transition-colors duration-200"
            @click="isOpen = !isOpen" aria-expanded="false" :aria-expanded="isOpen">
            <span class="heading-4">{{ title }}</span>
            <svg class="h-5 w-5 transform transition-transform duration-200 text-orange-600" :class="{ 'rotate-90': isOpen }"
                fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
        </button>
        <transition name="accordion">
            <div v-show="isOpen" class="px-6 py-4 bg-white">
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
