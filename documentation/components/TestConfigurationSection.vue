<template>
    <div class="bg-blue-100 p-6 rounded-lg mb-8">
        <h2 class="text-2xl font-semibold mb-4">Test Configuration</h2>
        <p>Define variables, page objects, and base URLs for your tests. This YAML file helps organize your tests and
            makes them more maintainable.</p>

        <AccordionItem title="Global Section">
            <p>The <code>global</code> section allows you to define global variables and settings for your tests. This
                section is further divided into <code>elements</code> and <code>pages</code>.</p>
        </AccordionItem>

        <AccordionItem title="Elements Section">
            <p>Define reusable selectors for common UI elements.</p>
            <CodeBlock :code="elementsSection" language="yaml" />
        </AccordionItem>

        <AccordionItem title="Pages Section">
            <p>Define URLs for different pages used in your tests.</p>
            <CodeBlock :code="pagesSection" language="yaml" />
        </AccordionItem>

        <AccordionItem title="Base URL Section">
            <p>Define the base URL for your application. Relative paths in the `pages` section will be relative to this
                base URL.</p>
            <CodeBlock :code="baseUrlSection" language="yaml" />
        </AccordionItem>

        <AccordionItem title="Test Configuration Example">
            <CodeBlock :code="testConfigExample" language="yaml" />
        </AccordionItem>
    </div>
</template>

<script setup lang="ts">
import AccordionItem from './AccordionItem.vue';

const elementsSection = `
elements:
  login_button:
    - "#login-button"
    - ".login-button"
  username_field: "#username"
  password_field: "#password"
`.trim();

const pagesSection = `
pages:
  home: "/home"
  login: "/login"
`.trim();

const baseUrlSection = `
base_url: "http://localhost:3000"
`.trim();

const testConfigExample = `
global:
  elements:
    login_button:
      - "#login-button"
      - ".login-button"
    username_field: "#username"
    password_field: "#password"
  pages:
    home: "/home"
    login: "/login"
base_url: "http://localhost:3000"
`.trim();
</script>
