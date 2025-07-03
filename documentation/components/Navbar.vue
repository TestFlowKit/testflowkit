<template>
    <header class="bg-blue-500 py-4 px-2 md:px-0">
        <div class="container mx-auto flex flex-col md:flex-row items-center justify-between text-white">
            <div class="w-full flex items-center justify-between md:w-auto md:inline">
                <nuxt-link tag="h1" to="/" class="text-3xl font-bold">Testflowkit</nuxt-link>
                <button class="md:hidden" @click="showMenu = !showMenu">
                    <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16m-7 6h7">
                        </path>
                    </svg>
                </button>
            </div>
            <nav class="md:block md:ml-4 mt-4 md:mt-0" :class="{ hidden: !showMenu }">
                <ul class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
                    <li><nuxt-link to="/" class="hover:underline" active-class="font-bold underline">Home</nuxt-link>
                    </li>

                    <li class="relative">
                        <button @click.stop="toggleGettingStartedMenu"
                            class="hover:underline flex items-center space-x-1"
                            :class="{ 'font-bold underline': menus.showGettingStartedMenu }">
                            <span>Getting Started</span>
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </button>
                        <div v-show="menus.showGettingStartedMenu"
                            class="dropdown-menu absolute top-full left-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 z-50">
                            <div class="p-2">
                                <nuxt-link :to="{ name: 'get-started' }"
                                    class="block px-3 py-2 text-gray-700 hover:bg-gray-100 rounded"
                                    @click="toggleSubMenu('showGettingStartedMenu')">
                                    Get Started
                                </nuxt-link>
                                <nuxt-link :to="{ name: 'quick-start' }"
                                    class="block px-3 py-2 text-gray-700 hover:bg-gray-100 rounded"
                                    @click="toggleSubMenu('showGettingStartedMenu')">
                                    Quick Start
                                </nuxt-link>
                                <nuxt-link :to="{ name: 'qa-guide' }"
                                    class="block px-3 py-2 text-gray-700 hover:bg-gray-100 rounded"
                                    @click="toggleSubMenu('showGettingStartedMenu')">
                                    QA Guide
                                </nuxt-link>
                            </div>
                        </div>
                    </li>

                    <li class="relative">
                        <button @click.stop="toggleDocsMenu" class="hover:underline flex items-center space-x-1"
                            :class="{ 'font-bold underline': menus.showDocsMenu }">
                            <span>Documentation</span>
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </button>
                        <div v-show="menus.showDocsMenu"
                            class="dropdown-menu absolute top-full left-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 z-50">
                            <div class="p-2">
                                <nuxt-link :to="{ name: 'configuration' }"
                                    class="block px-3 py-2 text-gray-700 hover:bg-gray-100 rounded"
                                    @click="toggleSubMenu('showDocsMenu')">
                                    Configuration
                                </nuxt-link>
                                <nuxt-link :to="{ name: 'sentences' }"
                                    class="block px-3 py-2 text-gray-700 hover:bg-gray-100 rounded"
                                    @click="toggleSubMenu('showDocsMenu')">
                                    Sentences
                                </nuxt-link>
                                <nuxt-link :to="{ name: 'troubleshooting' }"
                                    class="block px-3 py-2 text-gray-700 hover:bg-gray-100 rounded"
                                    @click="toggleSubMenu('showDocsMenu')">
                                    Troubleshooting
                                </nuxt-link>
                            </div>
                        </div>
                    </li>

                    <li class="relative">
                        <button @click.stop="toggleDownloadMenu" class="hover:underline flex items-center space-x-1"
                            :class="{ 'font-bold underline': menus.showDownloadMenu }">
                            <span>Download</span>
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </button>
                        <div v-show="menus.showDownloadMenu"
                            class="download-menu absolute top-full left-0 mt-2 w-64 bg-white rounded-lg shadow-lg border border-gray-200 z-50">
                            <div class="p-4">
                                <h3 class="font-semibold text-gray-800 mb-3">Choose Platform</h3>
                                <div class="space-y-2">
                                    <button @click.stop="selectPlatform('windows')"
                                        class="w-full text-left p-2 rounded hover:bg-gray-100 text-gray-700 flex items-center justify-between">
                                        <span>🪟 Windows (64-bit)</span>
                                        <svg v-if="selectedPlatform === 'windows'" class="w-4 h-4 text-blue-500"
                                            fill="currentColor" viewBox="0 0 20 20">
                                            <path fill-rule="evenodd"
                                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                                clip-rule="evenodd"></path>
                                        </svg>
                                    </button>
                                    <button @click.stop="selectPlatform('darwin')"
                                        class="w-full text-left p-2 rounded hover:bg-gray-100 text-gray-700 flex items-center justify-between">
                                        <span>🍎 macOS (Intel)</span>
                                        <svg v-if="selectedPlatform === 'darwin' && selectedArch === 'amd64'"
                                            class="w-4 h-4 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                                            <path fill-rule="evenodd"
                                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                                clip-rule="evenodd"></path>
                                        </svg>
                                    </button>
                                    <button @click.stop="selectPlatform('darwin-arm64')"
                                        class="w-full text-left p-2 rounded hover:bg-gray-100 text-gray-700 flex items-center justify-between">
                                        <span>🍎 macOS (Apple Silicon)</span>
                                        <svg v-if="selectedPlatform === 'darwin' && selectedArch === 'arm64'"
                                            class="w-4 h-4 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                                            <path fill-rule="evenodd"
                                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                                clip-rule="evenodd"></path>
                                        </svg>
                                    </button>
                                    <button @click.stop="selectPlatform('linux')"
                                        class="w-full text-left p-2 rounded hover:bg-gray-100 text-gray-700 flex items-center justify-between">
                                        <span>🐧 Linux (64-bit)</span>
                                        <svg v-if="selectedPlatform === 'linux' && selectedArch === 'amd64'"
                                            class="w-4 h-4 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                                            <path fill-rule="evenodd"
                                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                                clip-rule="evenodd"></path>
                                        </svg>
                                    </button>
                                    <button @click.stop="selectPlatform('linux-arm64')"
                                        class="w-full text-left p-2 rounded hover:bg-gray-100 text-gray-700 flex items-center justify-between">
                                        <span>🐧 Linux (ARM 64-bit)</span>
                                        <svg v-if="selectedPlatform === 'linux' && selectedArch === 'arm64'"
                                            class="w-4 h-4 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                                            <path fill-rule="evenodd"
                                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                                clip-rule="evenodd"></path>
                                        </svg>
                                    </button>
                                </div>
                                <div class="mt-4 pt-3 border-t border-gray-200">
                                    <button @click.stop="downloadSelected" :disabled="!selectedPlatform"
                                        class="w-full bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed">
                                        Download TestFlowKit
                                    </button>
                                </div>
                            </div>
                        </div>
                    </li>

                    <li><a target="_blank" :href="githubUrl" class="hover:underline">Github</a></li>
                </ul>
            </nav>
        </div>
    </header>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";

const router = useRouter();

const githubUrl = "https://github.com/TestFlowKit/testflowkit";
const menus = reactive({
    showDownloadMenu: false,
    showGettingStartedMenu: false,
    showDocsMenu: false,
});
const showMenu = ref(false);
const selectedPlatform = ref('');
const selectedArch = ref('');

const toggleDownloadMenu = () => {
    toggleSubMenu('showDownloadMenu');
};

const toggleGettingStartedMenu = () => {
    toggleSubMenu('showGettingStartedMenu');
};

const toggleDocsMenu = () => {
    toggleSubMenu('showDocsMenu');
};

const selectPlatform = (platform: string) => {
    if (platform === 'darwin-arm64') {
        selectedPlatform.value = 'darwin';
        selectedArch.value = 'arm64';
    } else if (platform === 'linux-arm64') {
        selectedPlatform.value = 'linux';
        selectedArch.value = 'arm64';
    } else {
        selectedPlatform.value = platform;
        selectedArch.value = 'amd64';
    }
};

const downloadSelected = () => {
    if (!selectedPlatform.value) return;

    const filename = `tkit-${selectedPlatform.value}-${selectedArch.value}.zip`;
    const downloadUrl = `https://github.com/TestFlowKit/testflowkit/releases/latest/download/${filename}`;

    // Create a temporary link element to trigger download
    const link = document.createElement('a');
    link.href = downloadUrl;
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    closeAllMenus();
};

const closeAllMenus = () => {
    menus.showDownloadMenu = false;
    menus.showGettingStartedMenu = false;
    menus.showDocsMenu = false;
};

const handleClickOutside = (event: Event) => {
    const target = event.target as HTMLElement;
    if (!target.closest('.download-menu') && !target.closest('.dropdown-menu') && !target.closest('button')) {
        closeAllMenus();
    }
};

const toggleSubMenu = (menu: keyof typeof menus) => {
    const currentValue = menus[menu];
    for (const key in menus) {
        menus[key as keyof typeof menus] = false;
    }
    menus[menu] = !currentValue;
};

watch(
    () => router.currentRoute.value.path,
    () => {
        showMenu.value = false;
        closeAllMenus();
    },
    { immediate: true }
);

onMounted(() => {
    document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside);
});
</script>