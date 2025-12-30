<template>
  <DocSidebar>
    <div class="p-4 sm:p-6 lg:p-8 max-w-4xl mx-auto">
      <template v-if="page">
        <!-- Page Header -->
        <header class="mb-6 sm:mb-8">
          <h1 class="text-2xl sm:text-3xl md:text-4xl font-bold text-gray-900 dark:text-white mb-2">
            {{ page.title }}
          </h1>
          <p v-if="page.description" class="text-base sm:text-lg text-gray-600 dark:text-gray-400">
            {{ page.description }}
          </p>
        </header>

        <!-- Content -->
        <article class="prose prose-blue dark:prose-invert max-w-none prose-sm sm:prose-base">
          <ContentRenderer :value="page" />
        </article>

        <!-- Navigation -->
        <DocNavigation :prev="prevPage" :next="nextPage" />
      </template>

      <template v-else>
        <div class="text-center py-12">
          <h1 class="text-xl sm:text-2xl font-bold text-gray-900 dark:text-white mb-4">Page Not Found</h1>
          <p class="text-gray-600 dark:text-gray-400 mb-6">The documentation page you're looking for doesn't exist.</p>
          <NuxtLink to="/docs/getting-started/introduction" class="text-blue-600 dark:text-blue-400 hover:underline">
            Go to Introduction
          </NuxtLink>
        </div>
      </template>
    </div>
  </DocSidebar>
</template>

<script setup lang="ts">
import DocSidebar from '~/components/DocSidebar.vue';
import DocNavigation from '~/components/DocNavigation.vue';


const { $rehighlight } = useNuxtApp();

const route = useRoute();
const path = computed(() => route.path);

// Fetch the current page
const { data: page } = await useAsyncData(`docs-${path.value}`, () => {
  return queryCollection('docs').path(path.value).first();
});

// Define all docs in order for navigation
const allDocs = [
  // Getting Started
  { path: '/docs/getting-started/introduction', title: 'Introduction' },
  { path: '/docs/getting-started/installation', title: 'Installation' },
  { path: '/docs/getting-started/quick-start', title: 'Quick Start' },
  { path: '/docs/getting-started/qa-guide', title: 'QA Guide' },
  // Core Concepts
  { path: '/docs/concepts/gherkin-basics', title: 'Gherkin Basics' },
  { path: '/docs/concepts/configuration', title: 'Configuration' },
  { path: '/docs/concepts/selectors', title: 'Selectors' },
  // Features
  { path: '/docs/features/variables', title: 'Variables' },
  { path: '/docs/features/macros', title: 'Macros' },
  { path: '/docs/features/frontend-testing', title: 'Frontend Testing' },
  { path: '/docs/features/api-testing', title: 'API Testing' },
  { path: '/docs/features/global-hooks', title: 'Global Hooks' },
  // Reference
  { path: '/docs/reference/cli', title: 'CLI Reference' },
  { path: '/docs/reference/step-definitions', title: 'Step Definitions' },
  // Troubleshooting
  { path: '/docs/troubleshooting/common-issues', title: 'Common Issues' },
  { path: '/docs/troubleshooting/platform-issues', title: 'Platform Issues' },
];

// Find current index and get prev/next
const currentIndex = computed(() => {
  return allDocs.findIndex(doc => doc.path === path.value);
});

const prevPage = computed(() => {
  if (currentIndex.value > 0) {
    return allDocs[currentIndex.value - 1];
  }
  return null;
});

const nextPage = computed(() => {
  if (currentIndex.value < allDocs.length - 1 && currentIndex.value >= 0) {
    return allDocs[currentIndex.value + 1];
  }
  return null;
});

// Apply syntax highlighting after content renders
onMounted(() => {
  nextTick(() => {
    $rehighlight();
  });
});

watch(page, () => {
  nextTick(() => {
    $rehighlight();
  });
});

// Set page meta
useHead({
  title: computed(() => page.value?.title ? `${page.value.title} | TestFlowKit` : 'Documentation | TestFlowKit'),
});
</script>

<style>
/* Prose overrides for better content styling */
.prose h1 {
  @apply text-3xl font-bold text-gray-900 dark:text-white mt-8 mb-4;
}

.prose h2 {
  @apply text-2xl font-bold text-gray-800 dark:text-gray-100 mt-8 mb-4 pb-2 border-b border-gray-200 dark:border-gray-700;
}

.prose h3 {
  @apply text-xl font-semibold text-gray-800 dark:text-gray-100 mt-6 mb-3;
}

.prose h4 {
  @apply text-lg font-semibold text-gray-700 dark:text-gray-200 mt-4 mb-2;
}

.prose p {
  @apply text-gray-600 dark:text-gray-300 leading-relaxed mb-4;
}

.prose code {
  @apply bg-gray-100 dark:bg-gray-800 text-pink-600 dark:text-pink-400 px-1.5 py-0.5 rounded text-sm font-mono;
}

.prose pre {
  @apply bg-[#282c34] rounded-lg p-4 overflow-x-auto my-4;
}

.prose pre code {
  @apply bg-transparent p-0 font-mono text-sm;
  color: #abb2bf !important;
}

.prose pre code * {
  color: inherit;
}

/* Ensure hljs styles are applied with proper colors */
.prose pre code.hljs {
  @apply p-0;
  background: transparent !important;
  color: #abb2bf !important;
}

/* Highlight.js syntax colors for atom-one-dark theme */
.prose .hljs-keyword,
.prose .hljs-selector-tag,
.prose .hljs-literal,
.prose .hljs-section,
.prose .hljs-link {
  color: #c678dd !important;
}

.prose .hljs-string,
.prose .hljs-title,
.prose .hljs-name,
.prose .hljs-type,
.prose .hljs-attribute,
.prose .hljs-symbol,
.prose .hljs-bullet,
.prose .hljs-addition,
.prose .hljs-variable,
.prose .hljs-template-tag,
.prose .hljs-template-variable {
  color: #98c379 !important;
}

.prose .hljs-comment,
.prose .hljs-quote,
.prose .hljs-deletion {
  color: #5c6370 !important;
}

.prose .hljs-meta,
.prose .hljs-tag {
  color: #e5c07b !important;
}

.prose .hljs-number {
  color: #d19a66 !important;
}

.prose .hljs-attr {
  color: #d19a66 !important;
}

.prose .hljs-built_in,
.prose .hljs-builtin-name {
  color: #e6c07b !important;
}

.prose ul, .prose ol {
  @apply my-4 pl-6;
}

.prose li {
  @apply text-gray-600 dark:text-gray-300 mb-2;
}

.prose a {
  @apply text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 hover:underline;
}

.prose table {
  @apply w-full my-4 border-collapse;
}

.prose th {
  @apply bg-gray-50 dark:bg-gray-800 text-left p-3 font-semibold text-gray-700 dark:text-gray-200 border border-gray-200 dark:border-gray-700;
}

.prose td {
  @apply p-3 border border-gray-200 dark:border-gray-700 text-gray-600 dark:text-gray-300;
}

.prose blockquote {
  @apply border-l-4 border-blue-500 dark:border-blue-400 pl-4 italic text-gray-600 dark:text-gray-300 my-4;
}

.prose hr {
  @apply my-8 border-gray-200 dark:border-gray-700;
}
</style>
