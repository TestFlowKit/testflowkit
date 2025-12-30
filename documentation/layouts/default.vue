<template>
    <div class="min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-200">
        <Navbar />
        <div class="pt-16">
            <slot />
        </div>
    </div>
</template>

<script setup lang="ts">
import Navbar from '~/components/Navbar.vue';
import useSidebar from '~/composables/useSidebar';
import useDarkMode from '~/composables/useDarkMode';

useSidebar();
useDarkMode();
</script>


