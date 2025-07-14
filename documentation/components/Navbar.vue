<template>
    <nav class="fixed top-0 left-0 right-0 z-[60] bg-white border-b border-gray-200 shadow-sm">
        <div class="flex items-center justify-between h-16 px-4 lg:px-6">
            <!-- Left side - Logo and mobile menu button -->
            <div class="flex items-center space-x-4">
                <!-- Mobile menu button -->
                <button @click="toggleSidebar" class="lg:hidden text-gray-600 hover:text-gray-900">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16"></path>
                    </svg>
                </button>

                <nuxt-link to="/" class="flex items-center space-x-2">
                    <span class="text-xl font-bold text-gray-900">TestFlowKit</span>
                </nuxt-link>
            </div>



            <!-- Right side - Actions -->
            <div class="flex items-center space-x-4">
                <!-- Download button -->
                <a href="https://github.com/TestFlowKit/testflowkit/releases" target="_blank"
                    class="hidden sm:inline-flex items-center px-3 py-2 text-sm font-medium text-blue-600 bg-blue-50 rounded-lg hover:bg-blue-100 transition-colors">
                    <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
                    </svg>
                    Download
                </a>

                <!-- GitHub link -->
                <a href="https://github.com/TestFlowKit/testflowkit" target="_blank"
                    class="text-gray-600 hover:text-gray-900 transition-colors">
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

const { toggleSidebar } = inject<UseSidebar>(useSidebarKey, {
    isOpen: ref(true),
    toggleSidebar: () => {
        throw new Error('toggleSidebar not found');
    },
});
</script>