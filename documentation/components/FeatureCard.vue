<template>
    <div class="bg-white p-6 rounded-lg shadow-md">
        <div class="flex items-center mb-4">
            <i :class="['text-blue-500 text-2xl mr-4', icon]"></i> {{ title }}
        </div>
        <slot />
    </div>
</template>

<script setup lang="ts">
defineProps({
    title: { type: String, required: true },
    icon: { type: String, required: true }
});
</script>
