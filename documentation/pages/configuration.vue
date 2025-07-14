<template>
  <div class="p-4 md:p-0">
    <h1 class="text-2xl md:text-3xl font-extrabold mb-4">TestFlowKit Configuration</h1>
    <p class="text-base md:text-lg mb-8 text-gray-600">
      TestFlowKit uses YAML configuration files to define test environments, element selectors, and execution settings.
      The configuration system supports multiple environments and provides comprehensive validation.
    </p>

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
