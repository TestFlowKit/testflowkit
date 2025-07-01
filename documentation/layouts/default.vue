<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation Header -->
    <Navbar />
    
    <!-- Main Content Layout -->
    <div class="relative">
      <!-- Decorative Background Elements -->
      <div class="decoration-blob decoration-blob-1"></div>
      <div class="decoration-blob decoration-blob-2"></div>
      
      <!-- Content Container -->
      <div class="relative z-10">
        <main class="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div class="max-w-7xl mx-auto">
            <slot />
          </div>
        </main>
      </div>
    </div>
    
    <!-- Footer -->
    <Footer />
  </div>
</template>

<script setup lang="ts">
// Import components
import Navbar from '~/components/Navbar.vue'
import Footer from '~/components/Footer.vue'
</script>

<style scoped>
.container {
  max-width: var(--container-max-width);
}
</style>
