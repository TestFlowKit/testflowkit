<template>
    <div class="bg-gray-200 p-6 rounded-lg mb-8">
        <h2 class="text-2xl font-semibold mb-4">CLI Arguments</h2>
        <p>Override settings from the YAML configuration files using command-line arguments.</p>
        <table class="table-auto w-full">
            <thead>
                <tr>
                    <th class="px-4 py-2">Argument</th>
                    <th class="px-4 py-2">Description</th>
                    <th class="px-4 py-2">Example</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td class="border px-4 py-2"><code>-l, --location &lt;path&gt;</code></td>
                    <td class="border px-4 py-2">Specifies the path to the Gherkin feature files.</td>
                    <td class="border px-4 py-2"><code>--location ./features</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>-c, --config &lt;path&gt;</code></td>
                    <td class="border px-4 py-2">Sets the path to the main configuration YAML file.</td>
                    <td class="border px-4 py-2"><code>--config config.yaml</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>-f, --front-config &lt;path&gt;</code></td>
                    <td class="border px-4 py-2">Sets the path to the frontend configuration YAML file.</td>
                    <td class="border px-4 py-2"><code>--front-config frontend.yaml</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>-t, --tags &lt;tags&gt;</code></td>
                    <td class="border px-4 py-2">Filters tests based on tags.</td>
                    <td class="border px-4 py-2"><code>--tags "@smoke,@regression"</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>-p, --parallel &lt;number&gt;</code></td>
                    <td class="border px-4 py-2">Defines the number of parallel tests.</td>
                    <td class="border px-4 py-2"><code>--parallel 4</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>--timeout &lt;duration&gt;</code></td>
                    <td class="border px-4 py-2">Sets the maximum test suite duration.</td>
                    <td class="border px-4 py-2"><code>--timeout 30s</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>--headless &lt;bool&gt;</code></td>
                    <td class="border px-4 py-2">Run browser in headless mode (true/false).</td>
                    <td class="border px-4 py-2"><code>--headless false</code></td>
                </tr>
                <tr>
                    <td class="border px-4 py-2"><code>-v, --version &lt;string&gt;</code></td>
                    <td class="border px-4 py-2">Specifies the application version.</td>
                    <td class="border px-4 py-2"><code>--version 2.1.0</code></td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
</script>
