<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Introduction</h2>
    <p class="mb-4">
      This application allows you to write and execute automated test scenarios.
      It uses two YAML files for configuration: one for the CLI settings and one
      for test execution. You can also override CLI settings using command-line
      arguments for greater flexibility.
    </p>

    <div class="bg-gray-200 p-4 rounded-md mb-6">
      <h2 class="title">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
          />
        </svg>
        CLI Configuration - Settings for running the test application
      </h2>
      <p class="mb-4">
        This section describes the YAML file used to configure the CLI itself.
        This includes general settings like timeouts, how tests are run (e.g.,
        in parallel), and where your test files are located. It also includes
        information about the application being tested and how test results are
        reported.
      </p>

      <h3 class="text-xl font-semibold mb-4">Global Settings</h3>
      <p class="mb-4">
        These settings control the overall behavior of the test application.
      </p>

      <div class="overflow-x-auto">
        <table class="table-auto w-full mb-6">
          <thead>
            <tr class="bg-gray-300">
              <th class="px-4 py-2">Option</th>
              <th class="px-4 py-2">Description</th>
              <th class="px-4 py-2">Default Value</th>
              <th class="px-4 py-2">Example</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td class="border px-4 py-2"><code>timeout</code></td>
              <td class="border px-4 py-2">
                Maximum test execution time (in seconds)
              </td>
              <td class="border px-4 py-2"><code>15s</code></td>
              <td class="border px-4 py-2"><code>timeout: 30s</code></td>
            </tr>
            <tr>
              <td class="border px-4 py-2"><code>parallel</code></td>
              <td class="border px-4 py-2">
                Number of tests to run in parallel
              </td>
              <td class="border px-4 py-2"><code>5</code></td>
              <td class="border px-4 py-2"><code>parallel: 10</code></td>
            </tr>
            <tr>
              <td class="border px-4 py-2"><code>slowMotion</code></td>
              <td class="border px-4 py-2">
                Slow down test execution (in seconds) - useful for debugging
              </td>
              <td class="border px-4 py-2"><code>0s</code></td>
              <td class="border px-4 py-2"><code>slowMotion: 1s</code></td>
            </tr>
            <tr>
              <td class="border px-4 py-2"><code>gherkin_location</code></td>
              <td class="border px-4 py-2">
                Path to the directory containing the Gherkin feature files
              </td>
              <td class="border px-4 py-2"><code>e2e/features</code></td>
              <td class="border px-4 py-2">
                <code>gherkin_location: "tests/features"</code>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3 class="text-xl font-semibold mb-4">Application Information</h3>
      <p class="mb-4">
        You can configure information about the application under test:
      </p>
      <ul class="list-disc list-inside mb-6">
        <li><code>app_name</code>: Application name</li>
        <li><code>app_description</code>: Application description</li>
        <li><code>app_version</code>: Application version</li>
      </ul>

      <h3 class="text-xl font-semibold mb-4">Reporting Configuration</h3>
      <p class="mb-4">
        The <code>report_format</code> option allows you to choose the format of
        the test reports. Supported formats are:
      </p>

      <ul class="list-disc list-inside mb-6">
        <li><code>html</code>: HTML report</li>
        <li><code>json</code>: JSON report</li>
      </ul>

      <h3 class="text-xl font-semibold mb-4">CLI Configuration Example</h3>
      <code-block :code="cliConfigExample" language="yaml" />

      <h3 class="text-xl font-bold mb-2">Configuration via CLI Arguments</h3>

      <p class="mb-4">
        You can also configure testflowkit by using command-line arguments.
        These arguments provide flexibility and allow you to override settings
        from the YAML configuration file.
      </p>

      <ul class="list-disc list-inside mb-4">
        <li>
          <code>-l, --location &lt;path&gt;</code>: Specifies the path to the
          directory containing your Gherkin feature files. (e.g.,
          <code>--location ./features</code>)
        </li>
        <li>
          <code>-c, --config &lt;path&gt;</code>: Sets the path to the main
          testflowkit configuration YAML file. Defaults to "cli.yml". (e.g.,
          <code>--config config.yaml</code>)
        </li>
        <li>
          <code>-f, --front-config &lt;path&gt;</code>: Sets the path to a YAML
          file specifically for frontend testing configuration. Defaults to
          "frontend.yml". (e.g.,
          <code>--front-config frontend-tests.yaml</code>)
        </li>
        <li>
          <code>-t, --tags &lt;tags&gt;</code>: Filters tests to run based on
          specified tags. (e.g., <code>--tags "@smoke,@regression"</code>)
        </li>
        <li>
          <code>-p, --parallel &lt;number&gt;</code>: Defines the number of
          tests to execute concurrently. (e.g., <code>--parallel 4</code>)
        </li>
        <li>
          <code>--timeout &lt;duration&gt;</code>: Sets the maximum duration for
          the entire test suite to run before timing out. (e.g.,
          <code>--timeout 30s</code>)
        </li>
        <li>
          <code>--headless &lt;bool&gt;</code>: Controls whether to run the
          browser in headless mode (without a visible UI). Defaults to "true"
          (headless). (e.g., <code>--headless false</code> to show the browser)
        </li>
        <li>
          <code>-v, --version &lt;string&gt;</code>: Specifies the version of
          the application under test. Defaults to "1.0". (e.g.,
          <code>--version 2.1.0</code>)
        </li>
      </ul>
    </div>

    <div class="bg-blue-100 p-4 rounded-md mb-6">
      <h2 class="title">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"
          />
        </svg>
        Test Configuration - Settings for defining test variables and page
        objects
      </h2>
      <p class="mb-4">
        This section describes a separate YAML file used to define variables and
        settings for your tests. This file helps you organize your tests by
        creating reusable names for web page elements (like buttons and input
        fields) and the URLs of the pages you are testing.
      </p>

      <div>
        <h3 class="text-xl font-semibold mb-4">Global Section</h3>
        <p class="mb-4">
          The <code>global</code> section allows you to define global variables
          and settings for your tests. This section is further divided into
          <code>elements</code> and <code>pages</code>.
        </p>
      </div>

      <div>
        <h3 class="text-xl font-semibold mb-4">Elements Section</h3>
        <p class="mb-4">
          The <code>elements</code> section allows you to define reusable
          selectors for common UI elements used in your tests. You can use CSS
          selectors to identify elements.
        </p>
        <code-block :code="elementsSection" language="yaml" />
        <p class="mb-4">
          In this example, we define three different selectors for the "login"
          element and one selector for the "username_input" element.
        </p>
      </div>

      <div class="mb-6">
        <h3 class="text-xl font-semibold mb-4">Pages Section</h3>
        <p class="mb-4">
          The <code>pages</code> section allows you to define URLs for different
          pages used in your tests.
        </p>

        <p class="mb-4">
          In this example, we define the URL for the "home" page.
        </p>
        <code-block :code="pagesSection" language="yaml" />
      </div>

      <div>
        <h3 class="text-xl font-semibold mb-4">Base URL Section</h3>
        <p class="mb-4">
          The <code>base_url</code> section allows you to define base url of
          your app.<br />
          You can add some external urls by add her base url
        </p>
        <code-block :code="baseUrlSection" language="yaml" />
      </div>
      <div>
        <h2 class="text-2xl font-semibold mb-4">Test Configuration Example</h2>
        <code-block :code="testConfigExample" language="yaml" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const cliConfigExample = `
configuration:
  timeout: 30s
  parallel: 10
  slowMotion: 0s
  gherkin_location: "tests/features"

application:
  app_name: "MonApplication"
  app_description: "Une application géniale"
  app_version: "2.0.0"

reporting:
  report_format: "json"
`.trim();

const elementsSection = `
elements:
  login:
      - "#login"
      - ".login"
      - "button .login"
  username_input: 
      - "#username"
`.trim();

const pagesSection = `
pages:
  home: "https://google.com"
`.trim();


const baseUrlSection = `
pages:
  google: "https://google.com"
  configuration: "/configuration"
  quick_start: "quickstart"
base_url: "http://tkit.com"
`.trim();

const testConfigExample = `
global:
  elements:
    login:
      - "#login"
      - ".login"
      - "button .login"
    username_input: 
      - "#username"
  pages:
    home: "https://google.com"
`.trim();
</script>

<style scoped>
.sentences-grid {
  @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

.title {
  @apply text-xl md:text-2xl font-semibold mb-4 flex items-center;
}

.title > svg {
  @apply h-12 w-12 md:h-6 md:w-6 mr-2;
}
</style>