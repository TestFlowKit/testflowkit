<template>
    <a :href="downloadLink" class="text-gray-700 block px-4 py-2 text-sm hover:bg-gray-100" role="menuitem"
        tabindex="-1">
        {{ osLabel }} ({{ archLabel }})
    </a>
</template>

<script setup lang="ts">
import { computed } from 'vue';
const props = defineProps({
    os: { type: String, required: true },
    arch: { type: String, required: true },
});

const osLabel = computed(() => {
    switch (props.os) {
        case 'windows': return 'Windows';
        case 'darwin': return 'macOS';
        case 'linux': return 'Linux';
        default: return props.os;
    }
});

const archLabel = computed(() => {
    switch (props.arch) {
        case 'amd64': return 'amd64';
        case '386': return '386';
        case 'arm64': return 'arm64';
        default: return props.arch;
    }
});

const downloadLink = computed(() => {
    const baseUrl = "https://github.com/TestFlowKit/testflowkit/releases/latest/download/tkit";
    return `${baseUrl}-${props.os}-${props.arch}`;
});
</script>
