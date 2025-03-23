<template>
  <section id="prerequisites" class="mb-12">
    <h2 class="heading-2">Prerequisites</h2>
    <p class="mb-4">Before you begin, ensure you have the following:</p>
    <PrerequisitesList />
  </section>

  <section id="download" class="mb-12">
    <h2 class="heading-2">Download testflowkit</h2>
    <p class="mb-6">Select the binary matching your operating system and architecture:</p>

    <DownloadSection />

  </section>

  <section class="mb-12">
    <h2 class="heading-2">Next Steps</h2>
    <p>After downloading, follow these steps:</p>
    <ol class="list-decimal list-inside ml-6">
      <li><strong>Unzip the downloaded archive.</strong></li>
      <li><strong>Navigate to the extracted directory in your terminal.</strong></li>
      <li><strong>Run the executable:</strong> <code>./tkit --help</code> (or `tkit.exe --help` on Windows) to see
        available commands.</li>
      <li><strong>Refer to the <router-link to="/configuration">Configuration Guide</router-link> for detailed setup
          instructions.</strong></li>
      <li><strong>Check out the <router-link to="/quick-start">Quick Start Guide</router-link> for a hands-on
          tutorial.</strong></li>
      <li><strong>Explore the <router-link to="/sentences">Gherkin Sentences Dictionary</router-link> to understand
          available keywords.</strong></li>
    </ol>
  </section>


</template>

<script setup lang="ts">
import PrerequisitesList from '../components/PrerequisitesList.vue';
import DownloadSection from '../components/DownloadSection.vue';
</script>
