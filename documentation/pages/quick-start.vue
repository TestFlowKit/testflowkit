<template>
    <div class="space-y-8">
        <!-- Hero Section -->
        <section class="text-center hero-gradient p-8 rounded-xl relative overflow-hidden">
            <div class="relative z-10">
                <h1 class="heading-1 mb-4 text-white">Quick Start Tutorial</h1>
                <p class="body-large mb-6 text-white max-w-2xl mx-auto">
                    Learn to write and run your first automated test
                </p>
            </div>
        </section>

        <!-- What We'll Test Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">What We'll Test</h2>

            <div class="card-gradient p-6 rounded-lg">
                <h3 class="font-semibold mb-4">Scenario: Search Functionality</h3>
                <p class="text-gray-600 mb-4">
                    We'll test that the search feature on our documentation website works correctly. This is a simple
                    example that demonstrates the basic concepts.
                </p>

                <div class="bg-white p-4 rounded border-l-4 border-orange-400">
                    <h4 class="font-medium mb-2">What we want to verify:</h4>
                    <ol class="space-y-2 text-sm">
                        <li>1. User can visit the search page</li>
                        <li>2. User can type "browser" in the search field</li>
                        <li>3. Search results show relevant content</li>
                        <li>4. Specific text appears in the results</li>
                    </ol>
                </div>
            </div>
        </section>

        <!-- Installation Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Installation</h2>

            <div class="space-y-6">
                <div class="card-gradient p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Download TestFlowKit</h3>
                    <p class="text-gray-600 mb-4">
                        TestFlowKit is a simple executable file that you can download and run. No complex installation
                        required!
                    </p>

                    <div class="bg-orange-50 border border-orange-200 p-6 rounded-lg">
                        <h4 class="font-medium mb-3">Step 1: Download</h4>
                        <p class="text-gray-600 mb-4">
                            Click the "Download" button in the navigation bar above to get the latest version for your
                            operating system:
                        </p>
                        <div class="bg-white p-4 rounded border-l-4 border-orange-400">
                            <p class="text-orange-600 font-medium">
                                📥 Use the "Download" button in the navigation bar
                            </p>
                        </div>
                    </div>

                    <div class="bg-teal-50 border border-teal-200 p-6 rounded-lg mt-4">
                        <h4 class="font-medium mb-3">Step 2: Extract & Run</h4>
                        <div class="space-y-3">
                            <div class="bg-white p-3 rounded border border-orange-200">
                                <div class="font-medium text-gray-800">Windows:</div>
                                <div class="text-sm text-gray-600">Extract the ZIP file and double-click
                                    <code class="bg-gray-100 px-1 rounded">testflowkit.exe</code>
                                </div>
                            </div>
                            <div class="bg-white p-3 rounded border border-orange-200">
                                <div class="font-medium text-gray-800">Mac:</div>
                                <div class="text-sm text-gray-600">Extract the file and run <code class="bg-gray-100 px-1 rounded">./testflowkit</code>
                                    in Terminal</div>
                            </div>
                            <div class="bg-white p-3 rounded border border-orange-200">
                                <div class="font-medium text-gray-800">Linux:</div>
                                <div class="text-sm text-gray-600">Extract the file and run <code class="bg-gray-100 px-1 rounded">./testflowkit</code>
                                    in Terminal</div>
                            </div>
                        </div>
                    </div>

                    <div class="mt-4 alert alert-info">
                        <h4 class="font-semibold text-blue-800 mb-2">💡 What you need:</h4>
                        <ul class="text-blue-700 space-y-1 text-sm">
                            <li>• A computer with Windows, Mac, or Linux</li>
                            <li>• Chrome or Edge browser installed</li>
                            <li>• That's it! No programming knowledge required</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>

        <!-- Configuration Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Setup Your Test Environment</h2>

            <div class="space-y-6">
                <div class="card-gradient p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Create a Simple Configuration File</h3>
                    <p class="text-gray-600 mb-4">
                        TestFlowKit needs to know about your website to run tests. Create a file called
                        <code class="bg-gray-100 px-1 rounded">config.yml</code> in your project folder:
                    </p>

                    <CodeBlock language="yaml" :code="yamlConfig" />

                    <div class="mt-4 bg-orange-50 border border-orange-200 p-4 rounded">
                        <h4 class="font-medium mb-2">📝 What this file does:</h4>
                        <ul class="text-sm text-gray-600 space-y-1">
                            <li>• Tells TestFlowKit where your website is running</li>
                            <li>• Defines the pages you want to test</li>
                            <li>• Lists the buttons and fields on your website</li>
                            <li>• Sets how fast or slow the tests should run</li>
                        </ul>
                    </div>
                </div>

                <div class="bg-teal-50 border border-teal-200 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Create Your Test Folder</h3>
                    <p class="text-gray-600 mb-4">
                        Create a simple folder structure for your tests:
                    </p>

                    <div class="bg-white p-4 rounded border border-orange-200">
                        <div class="text-sm text-gray-600 mb-2">Create these folders on your computer:</div>
                        <div class="code-block">
                            <pre class="code-content"><code>my-test-project/
├── config.yml
└── features/
    └── search.feature</code></pre>
                        </div>
                    </div>

                    <div class="mt-4 alert alert-warning">
                        <h4 class="font-semibold text-yellow-800 mb-2">💡 Quick Setup:</h4>
                        <ul class="text-yellow-700 space-y-1 text-sm">
                            <li>• Create a new folder called "my-test-project"</li>
                            <li>• Inside it, create a folder called "features"</li>
                            <li>• Put the config.yml file in the main folder</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>

        <!-- Writing the Test Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Writing Your Test</h2>

            <div class="space-y-6">
                <div class="card-gradient p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Understanding Gherkin Syntax</h3>
                    <p class="text-gray-600 mb-4">
                        Gherkin is a simple language that describes software behavior. It uses keywords that are easy to
                        understand:
                    </p>

                    <div class="grid md:grid-cols-2 gap-4">
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-orange-600 mb-2">Given</div>
                            <div class="text-sm text-gray-600">Sets up the starting point</div>
                        </div>
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-teal-600 mb-2">When</div>
                            <div class="text-sm text-gray-600">Describes the action to take</div>
                        </div>
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-teal-600 mb-2">And</div>
                            <div class="text-sm text-gray-600">Adds more actions or checks</div>
                        </div>
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-orange-600 mb-2">Then</div>
                            <div class="text-sm text-gray-600">Describes what should happen</div>
                        </div>
                    </div>
                </div>

                <div class="bg-orange-50 border border-orange-200 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Create Your Test File</h3>
                    <p class="text-gray-600 mb-4">
                        Create a new file called <code class="bg-gray-100 px-1 rounded">search.feature</code> in the <code class="bg-gray-100 px-1 rounded">features</code> directory and
                        add this content:
                    </p>

                    <CodeBlock language="gherkin" :code="gherkinTest" />
                </div>
            </div>
        </section>

        <!-- Running the Test Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Run Your Test</h2>

            <div class="space-y-4">
                <div class="card-gradient p-4 rounded-lg">
                    <h4 class="font-medium mb-2">Open your terminal/command prompt and run:</h4>
                    <CodeBlock language="bash" :code="runCommand" />
                </div>

                <div class="bg-orange-50 border border-orange-200 p-4 rounded-lg">
                    <h4 class="font-medium mb-2">Want to see the browser while it runs?</h4>
                    <CodeBlock language="bash" :code="visibleCommand" />
                </div>
            </div>
        </section>

        <!-- Understanding Results Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Understanding Your Test Results</h2>

            <div class="space-y-6">
                <div class="grid md:grid-cols-2 gap-6">
                    <div class="feature-card border border-green-200">
                        <div class="w-16 h-16 mx-auto mb-4 p-3 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center">
                            <span class="text-4xl">✅</span>
                        </div>
                        <h3 class="heading-4 mb-3">Test Passed</h3>
                        <p class="body-normal text-gray-600">
                            Great! Your test found everything it was looking for. The website behaved exactly as expected.
                        </p>
                    </div>

                    <div class="feature-card border border-red-200">
                        <div class="w-16 h-16 mx-auto mb-4 p-3 bg-gradient-to-br from-red-500 to-red-600 rounded-xl flex items-center justify-center">
                            <span class="text-4xl">❌</span>
                        </div>
                        <h3 class="heading-4 mb-3">Test Failed</h3>
                        <p class="body-normal text-gray-600">
                            Something didn't work as expected. Check the error message and screenshot to see what went wrong.
                        </p>
                    </div>
                </div>

                <div class="bg-teal-50 border border-teal-200 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">What You'll See</h3>
                    <ul class="space-y-2 text-gray-600">
                        <li>• <strong>Console Output:</strong> Step-by-step progress of your test</li>
                        <li>• <strong>Screenshots:</strong> Pictures of what went wrong (if anything failed)</li>
                        <li>• <strong>HTML Report:</strong> A detailed report you can open in your browser</li>
                        <li>• <strong>Timing Information:</strong> How long each step took to complete</li>
                    </ul>
                </div>
            </div>
        </section>

        <!-- Next Steps Section -->
        <section class="bg-gradient-to-r from-orange-500 to-orange-600 text-white p-8 rounded-xl text-center">
            <h2 class="heading-2 mb-4 text-white">What's Next?</h2>
            <p class="body-large mb-6 text-orange-100 max-w-2xl mx-auto">
                Congratulations! You've written and run your first automated test. Ready to learn more?
            </p>
            <div class="flex flex-col sm:flex-row justify-center gap-4">
                <nuxt-link to="/sentences" class="btn bg-white text-orange-600 hover:bg-gray-50">
                    Explore All Available Steps
                </nuxt-link>
                <nuxt-link to="/configuration" class="btn btn-ghost text-white border-white hover:bg-white/10 hover:text-white">
                    Advanced Configuration
                </nuxt-link>
                <nuxt-link to="/qa-guide" class="btn btn-ghost text-white border-white hover:bg-white/10 hover:text-white">
                    QA Team Guide
                </nuxt-link>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
// SEO and Meta
useHead({
  title: 'Quick Start Tutorial - TestFlowKit Documentation',
  meta: [
    {
      name: 'description',
      content: 'Learn to write and run your first automated test with TestFlowKit. Step-by-step tutorial for beginners.'
    }
  ]
})

// Import the CodeBlock component
import CodeBlock from '~/components/global/CodeBlock.vue'

// Computed properties for code blocks
const yamlConfig = `# config.yml - Copy this exactly
active_environment: local

environments:
  local:
    frontend_base_url: https://testflowkit.dreamsfollowers.me

settings:
  default_timeout: 15000
  headless: false
  think_time: 1000

frontend:
  pages:
    sentences: /sentences
    search: /search

  elements:
    sentences_page:
      search_field:
        - '[data-testid=search-input]'
        - input[type=search]
        - '#search'`

const gherkinTest = `Feature: Search Functionality
 As a user
 I want to search for content
 So that I can find relevant information

 Scenario: Search for browser content
   Given the user is on the "sentences" page
   When the user enters "browser" into the "search" field
   Then the user should see a element which contains "the user navigates back"`

const runCommand = './testflowkit run'

const visibleCommand = './testflowkit run --headless=false'
</script>
