<template>
    <div class="bg-gray-200 p-6 rounded-lg mb-8">
        <h2 class="text-2xl font-semibold mb-4">CLI Arguments</h2>
        <p class="mb-4">
            Override settings from the YAML configuration files using command-line
            arguments.
        </p>
        <div class="overflow-x-auto">
            <table class="table-auto w-full min-w-full">
                <thead>
                    <tr>
                        <th class="px-4 py-2 text-left">Argument</th>
                        <th class="px-4 py-2 text-left">Description</th>
                        <th class="px-4 py-2 text-left">Example</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>-l, --location &lt;path&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Specifies the path to the Gherkin feature files.
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--location ./features</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>-c, --config &lt;path&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Sets the path to the main configuration YAML file.
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--config config.yaml</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>-f, --front-config &lt;path&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Sets the path to the frontend configuration YAML file.
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--front-config frontend.yaml</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>-t, --tags &lt;tags&gt;</code>
                        </td>
                        <td class="border px-4 py-2">Filters tests based on tags.</td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--tags "@smoke,@regression"</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>-p, --parallel &lt;number&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Defines the number of parallel tests.
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--parallel 4</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--timeout &lt;duration&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Sets the maximum test suite duration.
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--timeout 30s</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--headless &lt;bool&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Run browser in headless mode (true/false).
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--headless false</code>
                        </td>
                    </tr>
                    <tr>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>-v, --version &lt;string&gt;</code>
                        </td>
                        <td class="border px-4 py-2">
                            Specifies the application version.
                        </td>
                        <td class="border px-4 py-2 whitespace-nowrap">
                            <code>--version 2.1.0</code>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>