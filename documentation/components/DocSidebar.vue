<template>
  <div class="flex min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-200">
    <div 
      v-if="isOpen" 
      class="mobile-overlay"
      @click="toggleSidebar"
    ></div>

    <aside 
      class="sidebar"
      :class="{ '-translate-x-full': !isOpen, 'translate-x-0': isOpen }"
    >
      <nav class="px-4 py-6">
        <template v-for="group in navigation" :key="group.title">
          <div class="mb-6">
            <h3 class="text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase tracking-wider mb-3">
              {{ group.title }}
            </h3>
            <ul class="space-y-1">
              <li v-for="item in group.children" :key="item.path">
                <NuxtLink
                  :to="item.path"
                  class="flex items-center px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-gray-900 dark:hover:text-white transition-colors"
                  :class="{ 'bg-blue-50 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400 border-r-2 border-blue-700 dark:border-blue-400': isActive(item.path) }"
                  @click="closeSidebarOnMobile"
                >
                  <component :is="getIcon(group.title)" class="w-4 h-4 mr-3 flex-shrink-0" />
                  {{ item.title }}
                </NuxtLink>
              </li>
            </ul>
          </div>
        </template>

        <!-- Resources Section (External Links) -->
        <div class="mb-6">
          <h3 class="text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase tracking-wider mb-3">Resources</h3>
          <ul class="space-y-1">
            <li>
              <NuxtLink
                to="/sentences"
                class="flex items-center px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-gray-900 dark:hover:text-white transition-colors"
                :class="{ 'bg-blue-50 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400 border-r-2 border-blue-700 dark:border-blue-400': $route.path === '/sentences' }"
                @click="closeSidebarOnMobile"
              >
                <svg class="w-4 h-4 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
                Step Definitions
              </NuxtLink>
            </li>
            <li>
              <a
                :href="GITHUB_REPO_URL"
                target="_blank"
                class="flex items-center px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-gray-900 dark:hover:text-white transition-colors"
              >
                <svg class="w-4 h-4 mr-3" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
                </svg>
                GitHub
              </a>
            </li>
            <li>
              <a
                :href="GITHUB_RELEASES_URL"
                target="_blank"
                class="flex items-center px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-gray-900 dark:hover:text-white transition-colors"
              >
                <svg class="w-4 h-4 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                </svg>
                Downloads
              </a>
            </li>
          </ul>
        </div>
      </nav>
    </aside>

    <main class="content-area">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import { h } from 'vue';
import { GITHUB_REPO_URL, GITHUB_RELEASES_URL } from '~/constants/links';
import { useSidebarKey, type UseSidebar } from '~/composables/useSidebar';

const sidebar = inject<UseSidebar>(useSidebarKey);

const isOpen = computed(() => sidebar?.isOpen.value ?? false);


function toggleSidebar() {
  sidebar?.toggleSidebar();
}

const route = useRoute();

function closeSidebarOnMobile() {
  if (typeof window !== 'undefined' && window.innerWidth < 1024 && isOpen.value) {
    toggleSidebar();
  }
}

const navigation = [
  {
    title: 'Getting Started',
    children: [
      { path: '/docs/getting-started/introduction', title: 'Introduction' },
      { path: '/docs/getting-started/installation', title: 'Installation' },
      { path: '/docs/getting-started/quick-start', title: 'Quick Start' },
      { path: '/docs/getting-started/qa-guide', title: 'QA Guide' },
    ],
  },
  {
    title: 'Core Concepts',
    children: [
      { path: '/docs/concepts/gherkin-basics', title: 'Gherkin Basics' },
      { path: '/docs/concepts/configuration', title: 'Configuration' },
      { path: '/docs/concepts/selectors', title: 'Selectors' },
    ],
  },
  {
    title: 'Features',
    children: [
      { path: '/docs/features/variables', title: 'Variables' },
      { path: '/docs/features/macros', title: 'Macros' },
      { path: '/docs/features/frontend-testing', title: 'Frontend Testing' },
      { path: '/docs/features/api-testing', title: 'API Testing' },
      { path: '/docs/features/global-hooks', title: 'Global Hooks' },
    ],
  },
  {
    title: 'Reference',
    children: [
      { path: '/docs/reference/cli', title: 'CLI Reference' },
      { path: '/docs/reference/step-definitions', title: 'Step Definitions' },
    ],
  },
  {
    title: 'Troubleshooting',
    children: [
      { path: '/docs/troubleshooting/common-issues', title: 'Common Issues' },
      { path: '/docs/troubleshooting/platform-issues', title: 'Platform Issues' },
    ],
  },
];

function isActive(path: string) {
  return route.path === path;
}

// Icon components based on section
function getIcon(sectionTitle: string) {
  const icons: Record<string, any> = {
    'Getting Started': {
      render() {
        return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M13 10V3L4 14h7v7l9-11h-7z' })
        ]);
      }
    },
    'Core Concepts': {
      render() {
        return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z' })
        ]);
      }
    },
    'Features': {
      render() {
        return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 110-4h1a1 1 0 001-1V7a1 1 0 011-1h3a1 1 0 001-1V4z' })
        ]);
      }
    },
    'Reference': {
      render() {
        return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253' })
        ]);
      }
    },
    'Troubleshooting': {
      render() {
        return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z' }),
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M15 12a3 3 0 11-6 0 3 3 0 016 0z' })
        ]);
      }
    },
  };
  
  return icons[sectionTitle] || icons['Getting Started'];
}
</script>

<style scoped> 
  .mobile-overlay {
    @apply fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden;
  }

  .sidebar {
    @apply fixed inset-y-0 left-0 z-50 w-64 bg-white dark:bg-gray-800 shadow-lg transform transition-transform duration-300 ease-in-out pt-16 overflow-y-auto lg:translate-x-0;
  }

  .content-area {
    @apply flex-1 w-full lg:ml-64 transition-all duration-300;
  }
</style>
