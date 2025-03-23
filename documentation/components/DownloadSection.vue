<template>
    <div class="relative inline-block text-left">
        <button type="button"
            class="inline-flex w-full justify-center gap-x-1.5 rounded-md primary-button-bg px-3 py-2 text-sm font-semibold text-white shadow-sm"
            id="download-menu-button" @click="isOpen = !isOpen" aria-expanded="true" aria-haspopup="true">
            Download
            <svg class="-mr-1 h-5 w-5 text-white" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd"
                    d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
                    clip-rule="evenodd" />
            </svg>
        </button>

        <div v-show="isOpen"
            class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
            role="menu" aria-orientation="vertical" aria-labelledby="download-menu-button" tabindex="-1"
            @mouseleave="isOpen = false">
            <DownloadItem os="windows" arch="amd64" />
            <DownloadItem os="windows" arch="386" />
            <DownloadItem os="darwin" arch="arm64" />
            <DownloadItem os="darwin" arch="amd64" />
            <DownloadItem os="linux" arch="amd64" />
            <DownloadItem os="linux" arch="386" />
            <DownloadItem os="linux" arch="arm64" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import DownloadItem from './DownloadItem.vue';

const isOpen = ref(false);
</script>
