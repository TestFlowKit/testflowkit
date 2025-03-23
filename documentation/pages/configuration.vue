<template>
  <div>
    <h1 class="text-3xl font-extrabold mb-8">testflowkit Configuration</h1>

    <CliConfigurationSection />
    <TestConfigurationSection />
    <CliArgumentsSection />
  </div>
</template>

<script setup lang="ts">
import CliConfigurationSection from '../components/CliConfigurationSection.vue';
import TestConfigurationSection from '../components/TestConfigurationSection.vue';
import CliArgumentsSection from '../components/CliArgumentsSection.vue';
</script>
