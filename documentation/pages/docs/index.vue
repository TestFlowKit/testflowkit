<template>
  <DocSidebar>
    <div class="p-4 sm:p-6 lg:p-8 max-w-4xl mx-auto">
      <header class="mb-8 sm:mb-10">
        <h1 class="text-2xl sm:text-3xl md:text-4xl font-bold text-gray-900 dark:text-white mb-2">
          Documentation
        </h1>
        <p class="text-base sm:text-lg text-gray-600 dark:text-gray-400">
          Pick a path based on your role, or browse the sidebar.
        </p>
      </header>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 sm:gap-6 mb-10">
        <section
          v-for="path in hubPaths"
          :key="path.id"
          class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 p-5 sm:p-6"
        >
          <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">{{ path.title }}</h2>
          <p class="text-sm text-gray-600 dark:text-gray-400 mb-4">{{ path.description }}</p>
          <ul class="space-y-2">
            <li v-for="link in path.links" :key="link.path">
              <NuxtLink
                :to="link.path"
                class="text-sm text-blue-600 dark:text-blue-400 hover:underline"
              >
                {{ link.title }} →
              </NuxtLink>
            </li>
          </ul>
        </section>
      </div>

      <section class="border-t border-gray-200 dark:border-gray-700 pt-8">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">New here?</h2>
        <div class="flex flex-wrap gap-3">
          <NuxtLink
            to="/docs/getting-started/installation"
            class="inline-flex items-center px-4 py-2 rounded-lg bg-blue-600 text-white text-sm font-medium hover:bg-blue-700 transition-colors"
          >
            Installation
          </NuxtLink>
          <NuxtLink
            to="/docs/getting-started/quick-start"
            class="inline-flex items-center px-4 py-2 rounded-lg bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-200 text-sm font-medium hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors"
          >
            Quick Start
          </NuxtLink>
          <NuxtLink
            to="/sentences"
            class="inline-flex items-center px-4 py-2 rounded-lg bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-200 text-sm font-medium hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors"
          >
            Step Catalog
          </NuxtLink>
        </div>
      </section>
    </div>
  </DocSidebar>
</template>

<script setup lang="ts">
import DocSidebar from '~/components/DocSidebar.vue';
import { docHubPaths } from '~/navigation';

const hubPaths = [
  {
    id: 'qa',
    title: 'QA / Test author',
    description: 'Write tests in plain English without coding.',
    links: docHubPaths.qa,
  },
  {
    id: 'developer',
    title: 'Developer',
    description: 'Set up config, APIs, and CI integration.',
    links: docHubPaths.developer,
  },
  {
    id: 'ai',
    title: 'AI-assisted (Cursor / VS Code)',
    description: 'Generate tests with the MCP server.',
    links: docHubPaths.ai,
  },
];

useHead({ title: 'Documentation | TestFlowKit' });
</script>
