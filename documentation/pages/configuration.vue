<template>
  <div>
    <h1 class="heading-1 mb-4">TestFlowKit Configuration</h1>
    <p class="body-large mb-8">
      TestFlowKit uses YAML configuration files to define test environments, element selectors, and execution settings. 
      The configuration system supports multiple environments and provides comprehensive validation.
    </p>

    <CliConfigurationSection />
    <TestConfigurationSection />
    <CliArgumentsSection />
  </div>
</template>

<script setup lang="ts">
// SEO and Meta
useHead({
  title: 'Configuration - TestFlowKit Documentation',
  meta: [
    {
      name: 'description',
      content: 'Complete guide to TestFlowKit configuration including YAML configuration files, test environments, element selectors, and execution settings.'
    }
  ]
})

import CliConfigurationSection from '../components/CliConfigurationSection.vue';
import TestConfigurationSection from '../components/TestConfigurationSection.vue';
import CliArgumentsSection from '../components/CliArgumentsSection.vue';
</script>
