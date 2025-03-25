<template>
    <div class="bg-gray-100 p-6 rounded-lg mb-8">
        <h2 class="text-2xl font-semibold mb-4">CLI Configuration</h2>
        <p>Configure settings for running the test application. This YAML file controls global settings, application
            information, and reporting options.</p>

        <AccordionItem title="Global Settings">
            <div class="overflow-x-auto">
                <table class="table-auto w-full">
                    <thead>
                        <tr>
                            <th class="px-4 py-2">Option</th>
                            <th class="px-4 py-2">Description</th>
                            <th class="px-4 py-2">Default Value</th>
                            <th class="px-4 py-2">Example</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td class="border px-4 py-2"><code>timeout</code></td>
                            <td class="border px-4 py-2">Maximum test execution time (in seconds).</td>
                            <td class="border px-4 py-2"><code>15s</code></td>
                            <td class="border px-4 py-2"><code>timeout: 30s</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>parallel</code></td>
                            <td class="border px-4 py-2">Number of tests to run in parallel.</td>
                            <td class="border px-4 py-2"><code>5</code></td>
                            <td class="border px-4 py-2"><code>parallel: 10</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>slowMotion</code></td>
                            <td class="border px-4 py-2">Slow down test execution (in seconds) - useful for debugging.
                            </td>
                            <td class="border px-4 py-2"><code>0s</code></td>
                            <td class="border px-4 py-2"><code>slowMotion: 1s</code></td>
                        </tr>
                        <tr>
                            <td class="border px-4 py-2"><code>gherkin_location</code></td>
                            <td class="border px-4 py-2">Path to the directory containing the Gherkin feature files.
                            </td>
                            <td class="border px-4 py-2"><code>e2e/features</code></td>
                            <td class="border px-4 py-2"><code>gherkin_location: "tests/features"</code></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </AccordionItem>

        <AccordionItem title="Application Information">
            <ul class="list-disc list-inside">
                <li><code>app_name</code>: Name of the application under test.</li>
                <li><code>app_description</code>: Description of the application.</li>
                <li><code>app_version</code>: Version of the application.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="Reporting Configuration">
            <p>Configure the format of the generated test reports.</p>
            <ul class="list-disc list-inside">
                <li><code>report_format: html</code>: Generates an HTML report.</li>
                <li><code>report_format: json</code>: Generates a JSON report.</li>
            </ul>
        </AccordionItem>

        <AccordionItem title="CLI Configuration Example">
            <CodeBlock :code="cliConfigExample" language="yaml" />
        </AccordionItem>
    </div>
</template>

<script setup lang="ts">
import AccordionItem from '@/components/AccordionItem.vue';

const cliConfigExample = `
  configuration:
    timeout: 30s
    parallel: 10
    slowMotion: 0s
    gherkin_location: "tests/features"
  
  application:
    app_name: "My Application"
    app_description: "A sample application for testing"
    app_version: "1.0.0"
  
  reporting:
    report_format: "json"
  `.trim();
</script>
