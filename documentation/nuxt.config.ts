export default defineNuxtConfig({
  nitro: {
    preset: 'github_pages',
    routeRules: {
      '/docs/concepts/gherkin-basics': { redirect: '/docs/guides/writing-tests' },
      '/docs/getting-started/qa-guide': { redirect: '/docs/guides/writing-tests' },
      '/docs/concepts/configuration': { redirect: '/docs/config/overview' },
      '/docs/concepts/selectors': { redirect: '/docs/config/selectors' },
      '/docs/features/frontend-testing': { redirect: '/docs/guides/frontend-testing' },
      '/docs/features/api-testing': { redirect: '/docs/guides/api-testing' },
      '/docs/getting-started/ide-agent': { redirect: '/docs/guides/ide-agent' },
      '/docs/features/variables': { redirect: '/docs/patterns/variables' },
      '/docs/features/random-data-generation': { redirect: '/docs/patterns/random-data' },
      '/docs/features/macros': { redirect: '/docs/patterns/macros' },
      '/docs/features/global-hooks': { redirect: '/docs/patterns/global-hooks' },
      '/docs/features/skip-tag': { redirect: '/docs/patterns/skip-tag' },
      '/docs/reference/step-definitions': { redirect: '/sentences' },
    },
  },
  app: {
    baseURL: process.env.NUXT_APP_BASE_URL || '/',
    head: {
      title: "TestFlowKit - Behavior-Driven Testing Framework"
    }
  },
  compatibilityDate: "2024-11-01",
  devtools: { enabled: false },
  css: ["~/assets/css/main.css"],
  modules: ["@nuxtjs/tailwindcss", "@nuxt/content"],
  tailwindcss: {
    viewer: false,
  },
  plugins: ["~/plugins/highlightjs.ts"],
  components: [{ path: "components/global", global: true }],
});
