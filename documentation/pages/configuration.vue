<template>
  <div>
    <h1 class="text-3xl font-extrabold mb-4">TestFlowKit Configuration</h1>
    <p class="text-lg mb-8 text-gray-600">
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
