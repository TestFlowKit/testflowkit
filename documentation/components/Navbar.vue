<template>
    <nav class="fixed top-0 left-0 right-0 z-[60] bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 shadow-sm transition-colors duration-200">
        <div class="flex items-center justify-between h-16 px-4 lg:px-6">
            <div class="flex items-center space-x-4">
                <button 
                    v-if="showMenuButton" 
                    @click="handleToggle" 
                    class="lg:hidden text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white"
                    aria-label="Toggle navigation menu"
                >
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16"></path>
                    </svg>
                </button>

                <nuxt-link to="/" class="flex items-center space-x-2">
                    <span class="text-xl font-bold text-gray-900 dark:text-white">TestFlowKit</span>
                </nuxt-link>

                <!-- Navigation Links -->
                <div class="hidden md:flex items-center space-x-1 ml-6">
                    <nuxt-link 
                        to="/docs/getting-started/introduction" 
                        class="px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-gray-900 dark:hover:text-white transition-colors"
                    >
                        Docs
                    </nuxt-link>
                    <nuxt-link 
                        to="/sentences" 
                        class="px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-gray-900 dark:hover:text-white transition-colors"
                    >
                        Step Definitions
                    </nuxt-link>
                </div>
            </div>



            <!-- Right side - Actions -->
            <div class="flex items-center space-x-4">
                <!-- Dark mode toggle -->
                <button 
                    @click="handleDarkModeToggle" 
                    class="text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white transition-colors p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700"
                    :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
                >
                    <!-- Sun icon (shown in dark mode) -->
                    <svg v-if="isDark" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                    </svg>
                    <!-- Moon icon (shown in light mode) -->
                    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                    </svg>
                </button>

                <a :href="GITHUB_REPO_URL" target="_blank"
                    class="text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white transition-colors">
                    <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                        <path
                            d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
                    </svg>
                </a>
            </div>
        </div>
    </nav>
</template>

<script setup lang="ts">
import { useSidebarKey, type UseSidebar } from '~/composables/useSidebar';
import { useDarkModeKey, type UseDarkMode } from '~/composables/useDarkMode';

import { GITHUB_REPO_URL } from '~/constants/links';

const route = useRoute();
const showMenuButton = computed(() => route.path.startsWith('/docs'));

const sidebar = inject<UseSidebar>(useSidebarKey);
const darkMode = inject<UseDarkMode>(useDarkModeKey);

const isDark = computed(() => darkMode?.isDark.value ?? false);

function handleToggle() {
    if (sidebar?.toggleSidebar) {
        sidebar.toggleSidebar();
    }
}

function handleDarkModeToggle() {
    if (darkMode?.toggle) {
        darkMode.toggle();
    }
}
</script>