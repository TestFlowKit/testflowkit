export default defineNuxtConfig({
  nitro: {
    preset: 'github_pages'
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
