<template>
    <div class="bg-gray-200 p-4 md:p-6 rounded-lg mb-8">
        <h2 class="text-xl md:text-2xl font-semibold mb-4">CLI Arguments</h2>
        <p class="mb-4 text-sm md:text-base">
            Override settings from the YAML configuration files using command-line
            arguments.
        </p>
        <div class="overflow-x-auto">
            <table class="table-auto w-full min-w-full text-sm">
                <thead>
                    <tr>
                        <th class="px-2 md:px-4 py-2 text-left">Argument</th>
                        <th class="px-2 md:px-4 py-2 text-left">Description</th>
                        <th class="px-2 md:px-4 py-2 text-left">Example</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">-l, --location &lt;path&gt;</code>
                        </td>
                        <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">
                            Specifies the path to the Gherkin feature files.
                        </td>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--location ./features</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">-c, --config &lt;path&gt;</code>
                        </td>
                        <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">
                            Sets the path to the main configuration YAML file.
                        </td>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--config config.yml</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">-t, --tags &lt;tags&gt;</code>
                        </td>
                        <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">Filters tests based on tags.</td>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--tags "@smoke,@regression"</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--env-file &lt;path&gt;</code>
                        </td>
                        <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">
                            Specifies the environment variables file to use (overrides settings.env_file).
                        </td>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--env-file .env.staging.yml</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--timeout &lt;duration&gt;</code>
                        </td>
                        <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">
                            Sets the maximum test execution time (e.g., 10s, 1m, 2h).
                        </td>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--timeout 30s</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--headless &lt;bool&gt;</code>
                        </td>
                        <td class="border px-2 md:px-4 py-2 text-xs md:text-sm">
                            Run browser in headless mode (true/false).
                        </td>
                        <td class="border px-2 md:px-4 py-2 whitespace-nowrap">
                            <code class="text-xs md:text-sm">--headless false</code>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div class="mt-6">
            <h3 class="text-lg font-semibold mb-3">Available Commands</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="bg-white p-3 md:p-4 rounded border">
                    <h4 class="font-semibold text-blue-600 text-sm md:text-base">run</h4>
                    <p class="text-xs md:text-sm text-gray-600">Execute test scenarios</p>
                    <code class="text-xs">tkit run [options]</code>
                </div>
                <div class="bg-white p-3 md:p-4 rounded border">
                    <h4 class="font-semibold text-green-600 text-sm md:text-base">init</h4>
                    <p class="text-xs md:text-sm text-gray-600">Initialize project configuration</p>
                    <code class="text-xs">tkit init</code>
                </div>
                <div class="bg-white p-3 md:p-4 rounded border">
                    <h4 class="font-semibold text-orange-600 text-sm md:text-base">validate</h4>
                    <p class="text-xs md:text-sm text-gray-600">Validate Gherkin files</p>
                    <code class="text-xs">tkit validate [options]</code>
                </div>
            </div>
        </div>

        <div class="mt-6">
            <h3 class="text-lg font-semibold mb-3">Usage Examples</h3>
            <div class="space-y-2">
                <div class="bg-gray-800 text-white p-3 rounded text-xs md:text-sm">
                    <code># Run all tests with default configuration</code><br>
                    <code>tkit run</code>
                </div>
                <div class="bg-gray-800 text-white p-3 rounded text-xs md:text-sm">
                    <code># Run tests with custom configuration</code><br>
                    <code>tkit run --config my-config.yml --env staging</code>
                </div>
                <div class="bg-gray-800 text-white p-3 rounded text-xs md:text-sm">
                    <code># Run specific tagged tests</code><br>
                    <code>tkit run --tags "@smoke" --headless true</code>
                </div>
                <div class="bg-gray-800 text-white p-3 rounded text-xs md:text-sm">
                    <code># Validate Gherkin files</code><br>
                    <code>tkit validate --location ./features</code>
                </div>
            </div>
        </div>
    </div>
</template>