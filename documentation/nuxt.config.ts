export default defineNuxtConfig({
  app: {
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
