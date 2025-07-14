<template>
    <div class="min-h-screen">
        <Navbar />
        <Sidebar>
            <div class="p-6 lg:p-8">
                <slot />
            </div>
        </Sidebar>
    </div>
</template>

<script setup lang="ts">
import Sidebar from '../components/Sidebar.vue';
import Navbar from '../components/Navbar.vue';
import useSidebar from '../composables/useSidebar';

useSidebar();
</script>
