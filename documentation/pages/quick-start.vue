<template>
    <div class="space-y-8">
        <section class="text-center bg-gradient-to-r from-blue-50 to-indigo-50 p-8 rounded-lg">
            <h1 class="text-4xl font-bold mb-4 text-gray-800">Quick Start Tutorial</h1>
            <p class="text-xl text-gray-600 mb-6">
                Learn to write and run your first automated test
            </p>
        </section>

        <!-- What We'll Test Section -->
        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">What We'll Test</h2>

            <div class="bg-gray-50 p-6 rounded-lg">
                <h3 class="font-semibold mb-4">Scenario: Search Functionality</h3>
                <p class="text-gray-600 mb-4">
                    We'll test that the search feature on our documentation website works correctly. This is a simple
                    example that demonstrates the basic concepts.
                </p>

                <div class="bg-white p-4 rounded border-l-4 border-blue-400">
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
        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Installation</h2>

            <div class="space-y-6">
                <div class="bg-gray-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Download TestFlowKit</h3>
                    <p class="text-gray-600 mb-4">
                        TestFlowKit is a simple executable file that you can download and run. No complex installation
                        required!
                    </p>

                    <div class="bg-blue-50 p-6 rounded-lg">
                        <h4 class="font-medium mb-3">Step 1: Download</h4>
                        <p class="text-gray-600 mb-4">
                            Click the "Download" button in the navigation bar above to get the latest version for your
                            operating system:
                        </p>
                        <div class="bg-white p-4 rounded border-l-4 border-blue-400">
                            <p class="text-blue-600 font-medium">
                                📥 Use the "Download" button in the navigation bar
                            </p>
                        </div>
                    </div>

                    <div class="bg-green-50 p-6 rounded-lg mt-4">
                        <h4 class="font-medium mb-3">Step 2: Extract & Run</h4>
                        <div class="space-y-3">
                            <div class="bg-white p-3 rounded">
                                <div class="font-medium text-gray-800">Windows:</div>
                                <div class="text-sm text-gray-600">Extract the ZIP file and double-click
                                    <code>tkit.exe</code>
                                </div>
                            </div>
                            <div class="bg-white p-3 rounded">
                                <div class="font-medium text-gray-800">Mac:</div>
                                <div class="text-sm text-gray-600">Extract the file and run <code>./tkit</code>
                                    in Terminal</div>
                            </div>
                            <div class="bg-white p-3 rounded">
                                <div class="font-medium text-gray-800">Linux:</div>
                                <div class="text-sm text-gray-600">Extract the file and run <code>./tkit</code>
                                    in Terminal</div>
                            </div>
                        </div>
                    </div>

                    <div class="mt-4 bg-yellow-50 border-l-4 border-yellow-400 p-4">
                        <h4 class="font-semibold text-yellow-800 mb-2">💡 What you need:</h4>
                        <ul class="text-yellow-700 space-y-1 text-sm">
                            <li>• A computer with Windows, Mac, or Linux</li>
                            <li>• Chrome or Edge browser installed</li>
                            <li>• That's it! No programming knowledge required</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>

        <!-- Configuration Section -->
        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Setup Your Test Environment</h2>

            <div class="space-y-6">
                <div class="bg-gray-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Create a Simple Configuration File</h3>
                    <p class="text-gray-600 mb-4">
                        TestFlowKit needs to know about your website to run tests. Create a file called
                        <code>config.yml</code> in your project folder:
                    </p>

                    <CodeBlock language="yaml" :code="yamlConfig" />

                    <div class="mt-4 bg-blue-50 p-4 rounded">
                        <h4 class="font-medium mb-2">📝 What this file does:</h4>
                        <ul class="text-sm text-gray-600 space-y-1">
                            <li>• Tells TestFlowKit where your website is running</li>
                            <li>• Defines the pages you want to test</li>
                            <li>• Lists the buttons and fields on your website</li>
                            <li>• Sets how fast or slow the tests should run</li>
                        </ul>
                    </div>
                </div>

                <div class="bg-green-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Create Your Test Folder</h3>
                    <p class="text-gray-600 mb-4">
                        Create a simple folder structure for your tests:
                    </p>

                    <div class="bg-white p-4 rounded">
                        <div class="text-sm text-gray-600 mb-2">Create these folders on your computer:</div>
                        <div class="font-mono text-sm bg-gray-100 p-3 rounded">
                            <div class="text-gray-600">my-test-project/</div>
                            <div class="text-gray-600">├── config.yml</div>
                            <div class="text-gray-600">└── features/</div>
                            <div class="text-gray-600"> └── search.feature</div>
                        </div>
                    </div>

                    <div class="mt-4 bg-yellow-50 border-l-4 border-yellow-400 p-4">
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
        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Writing Your Test</h2>

            <div class="space-y-6">
                <div class="bg-gray-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Understanding Gherkin Syntax</h3>
                    <p class="text-gray-600 mb-4">
                        Gherkin is a simple language that describes software behavior. It uses keywords that are easy to
                        understand:
                    </p>

                    <div class="grid md:grid-cols-2 gap-4">
                        <div class="bg-white p-4 rounded">
                            <div class="font-semibold text-purple-600 mb-2">Given</div>
                            <div class="text-sm text-gray-600">Sets up the starting point</div>
                        </div>
                        <div class="bg-white p-4 rounded">
                            <div class="font-semibold text-orange-600 mb-2">When</div>
                            <div class="text-sm text-gray-600">Describes the action to take</div>
                        </div>
                        <div class="bg-white p-4 rounded">
                            <div class="font-semibold text-green-600 mb-2">And</div>
                            <div class="text-sm text-gray-600">Adds more actions or checks</div>
                        </div>
                        <div class="bg-white p-4 rounded">
                            <div class="font-semibold text-purple-600 mb-2">Then</div>
                            <div class="text-sm text-gray-600">Describes what should happen</div>
                        </div>
                    </div>
                </div>

                <div class="bg-blue-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-4">Create Your Test File</h3>
                    <p class="text-gray-600 mb-4">
                        Create a new file called <code>search.feature</code> in the <code>features</code> directory and
                        add this content:
                    </p>

                    <CodeBlock language="gherkin" :code="gherkinTest" />
                </div>
            </div>
        </section>

        <!-- Running the Test Section -->
        <section class=" bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Run Your Test</h2>

            <div class="space-y-4">
                <div class="bg-gray-50 p-4 rounded">
                    <h4 class="font-medium mb-2">Open your terminal/command prompt and run:</h4>
                    <CodeBlock language="bash" :code="runCommand" />
                </div>

                <div class="bg-blue-50 p-4 rounded">
                    <h4 class="font-medium mb-2">Want to see the browser while it runs?</h4>
                    <CodeBlock language="bash" :code="runCommandVisible" />
                </div>
            </div>

            <div class="mt-6 bg-green-50 border-l-4 border-green-400 p-4">
                <h4 class="font-semibold text-green-800 mb-2">What You Should See:</h4>
                <ul class="text-green-700 space-y-1 text-sm">
                    <li>• A browser window opens automatically (no need to explicitly open it first!)</li>
                    <li>• It navigates to the sentences page</li>
                    <li>• It types "browser" in the search field</li>
                    <li>• It checks that "browser automation" appears on the page</li>
                    <li>• A test report is generated showing the results</li>
                </ul>
                <div class="mt-3 p-3 bg-blue-50 rounded">
                    <p class="text-blue-800 text-sm">
                        <strong>💡 Pro Tip:</strong> TestFlowKit automatically opens the browser when
                        needed, so you
                        don't have to remember to add browser setup steps to your tests!
                    </p>
                </div>
            </div>
        </section>

        <!-- Next Steps Section -->
        <section class="bg-gradient-to-r from-green-50 to-blue-50 p-8 rounded-lg">
            <h2 class="text-3xl font-bold mb-6 text-center text-gray-800">Congratulations! 🎉</h2>

            <div class="text-center space-y-4">
                <p class="text-lg text-gray-600">
                    You've successfully created and run your first automated test! Here's what you can do next:
                </p>

                <div class="flex flex-wrap justify-center gap-4">
                    <nuxt-link to="/sentences" class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700">
                        📚 Explore More Test Patterns
                    </nuxt-link>
                    <nuxt-link to="/qa-guide" class="bg-purple-600 text-white px-6 py-3 rounded-lg hover:bg-purple-700">
                        🧪 QA Guide for Non-Technical Users
                    </nuxt-link>
                </div>

                <div class="bg-white p-4 rounded-lg mt-6">
                    <p class="text-sm text-gray-600">
                        <strong>Need technical help?</strong> Check out the
                        <a href="https://github.com/TestFlowKit/testflowkit" target="_blank"
                            class="text-blue-600 hover:underline">GitHub README</a>
                        for technical details or visit the
                        <a href="https://testflowkit.dreamsfollowers.me/" target="_blank"
                            class="text-blue-600 hover:underline">official documentation</a>.
                    </p>
                </div>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
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
  concurrency: 1
  gherkin_location: "features"
  report_format: html

frontend:
  pages:
    sentences: /sentences

  elements:
    sentences:
      search_field:
        - "#search-input"
`

const gherkinTest = `Feature: Search Functionality
 As a user
 I want to search for content
 So that I can find relevant information

 Scenario: Search for browser navigation phrase
   Given the user is on the "sentences" page
   When the user enters "browser" into the "search" field
   Then the user should see a element which contains "the user navigates back"`

const runCommand = './tkit run'

const runCommandVisible = './tkit run --headless=false'
</script>
