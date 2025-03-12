<template>

    <div class="space-y-6">
        <h1 class="text-3xl font-bold">Quick Start with testflowkit</h1>

        <p>This guide is for developers or QA testers who want to learn how to use testflowkit to launch e2e tests on
            their
            web applications. It assumes that you are comfortable with the command line, yaml and gherkin.</p>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-bold mb-4">Prerequisites</h2>
            <prerequisites-list></prerequisites-list>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-bold mb-4">Objective</h2>
            <p>In this exercise, we will test our own documentation. We will make sure that the search for phrases
                works
                correctly.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Installation and configuration</h2>
            <ol class="list-decimal list-inside mb-4">
                <li>Download testflowkit for your operating system: <a href="[download link]" target="_blank"
                        class="underline">Download link</a>. Place the executable in the folder of your choice.</li>
                        <li>Extract tkit for the downloaded zip in the folder of your choice. its contains a binary named tkit</li>
                <li>Open your terminal and type the following command:
                    <code>tkit init --appname=test_app</code>
                    <br />
                    <em>This command creates the configuration files `frontend.yml` and `cli.yml` at the
                        root of the folder.</em><br />
                    <em>The `frontend.yml` file contains the settings specific to your web application, while the
                        `cli.yml`
                        file contains the general settings for testflowkit. and a folder containing an gherkin example</em>
                </li>
            </ol>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Writing the test</h2>
            <p>Let's start by writing our tests in natural language. You can go to the reference page to perform the
                actions
                and note them. Once finished you can unroll our suggestion.</p>
            <div>
                <a class="suggestion-toggle" @click="toggleSuggestion">{{ suggestionShowed ? 'Hide' : 'Show' }}
                    suggestion</a>
                <div v-show="suggestionShowed">
                    <p class="italic">To verify that the search is working correctly, we will perform the following
                        actions:</p>
                    <ol class="list-decimal ml-10">
                        <li>Open the browser.</li>
                        <li>Go to the sentence referencing page.</li>
                        <li>Write "browser" in the search field.</li>
                        <li>Check that the following sentence appears on the screen:
                            <code>I open a new browser tab</code>.
                        </li>
                    </ol>
                </div>
            </div>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Interactive Tutorial: Writing a Gherkin Scenario</h2>

            <p>Let's walk through creating a Gherkin scenario step-by-step. Our goal is to write a test that
                verifies
                the search functionality on the sentence referencing page. </p>

            <p>Here are the actions we want to test:</p>
            <ol class="list-decimal list-inside mb-4">
                <li>Open the browser.</li>
                <li>Go to the sentence referencing page.</li>
                <li>Type "browser" into the search field.</li>
                <li>Verify that the following sentence appears on the screen: <code>I open a new browser tab.</code>
                </li>
            </ol>

            <p>Now, let's translate these actions into Gherkin code.</p>

            <h3 class="text-lg font-medium mt-4">Step 1: Create a Feature File</h3>

            <p>First, create a new file named <code>search.feature</code> in your <code>features</code> directory.
                This
                file will contain our Gherkin scenarios.</p>

            <h3 class="text-lg font-medium mt-4">Step 2: Define the Feature</h3>

            <p>Start by defining the feature we are testing. In this case, it's the search functionality. Add the
                following line to your <code>search.feature</code> file:</p>

            <code-block language="gherkin" :code="`Feature: Search Functionality`" />

            <h3 class="text-lg font-medium mt-4">Step 3: Write the Scenario</h3>

            <p>Next, define a scenario within our feature. A scenario outlines a specific test case. We'll do this
                in
                stages:</p>

            <h4 class="text-md font-medium mt-3">Stage 1: Setting the Scene</h4>
            <p>Add the following line to your file:</p>
            <code-block language="gherkin" :code="`Scenario: Search for a sentence`" />
            <p>This gives our scenario a descriptive name.</p>

            <h4 class="text-md font-medium mt-3">Stage 2: Starting Point</h4>
            <p>Now, let's define the starting point of our test. Add this line:</p>
            <code-block language="gherkin"
                :code="`Given I open a new browser tab\nAnd I navigate to sentences referencing page`" />
            <p>This "Given" step sets the initial context for the scenario. It says we begin on the sentence
                referencing
                page.</p>

            <h4 class="text-md font-medium mt-3">Stage 3: User Action</h4>
            <p>Next, describe the action the user performs. Add this line:</p>
            <code-block language="gherkin" :code='`When I type "browser" into the search field`' />
            <p>This "When" step specifies the user action – typing "browser" into the search field.</p>

            <h4 class="text-md font-medium mt-3">Stage 4: Expected Outcome</h4>
            <p>Finally, define the expected result of the user's action. Add this line:</p>
            <code-block language="gherkin" :code='`Then I should see on page "I open a new browser tab."`' />
            <p>This "Then" step states the expected outcome – seeing a specific sentence on the page.</p>

            <p>Save the <code>search.feature</code> file.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Defining Variables</h2>
            <p>You can define variables in the <code>frontend.yml</code> file to make your scenarios work smoothly.
                These variables represent elements on your web page and make your Gherkin code more readable and
                maintainable.</p>

            <sm><em><strong>Tip</strong>: you can run the tkit validate command in order to validate your work: <code>./tkit validate</code></em></sm>

            <p>Here's how it works:</p>
            <ul class="list-disc list-inside ml-2">
                <li><b>Page URLs:</b> Store the URLs of frequently used pages, like the sentence referencing page.
                    This
                    way, you can refer to them by name in your scenarios. For example:
                    <code>sentences_reference: "http://&lt;doc-base-url&gt;/sentences"</code>
                </li>
                <li><b>UI Elements:</b> Define variables for interactive elements like buttons, input fields, etc.
                    This
                    makes your Gherkin steps clearer and easier to understand. For example, instead of writing a
                    complex
                    CSS selector in your Gherkin step, you can use a variable like <code>search_field</code>.</li>
            </ul>

            <p><b>Important Tip:</b> You can even specify multiple CSS selectors for a single element! This is
                helpful
                if the element might be identified differently in various parts of your application. Etools will try
                each selector in order until it finds a match.</p>

            <p><b>Naming Convention:</b> To keep your code organized, use underscores to separate words in variable
                names (e.g., <code>login_button</code> instead of <code>loginButton</code>). This is a common
                practice
                in YAML files.</p>

            <p>Here's an example of how your <code>frontend.yml</code> file might look:</p>
            <code-block language="yml" :code='frontendYml' />

            <p><b>Remember:</b> Replace <code>&lt;doc-base-url&gt;</code> with the actual base URL of your
                documentation.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Running the tests</h2>
            <p>Run the following command in your terminal to launch the e2e tests:</p>
            <code-block language="bash" :code="launchCommand" />

            <p>This command will execute your Gherkin scenario and provide you with a report of the test results.
            </p>

            <p>Congratulations! You've successfully written a Gherkin scenario to test the search functionality. You
                can
                now expand on this by adding more scenarios to cover other test cases.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Troubleshooting</h2>
            <p>If you encounter errors while running your tests, here are a few things to check:</p>
            <ul class="list-disc list-inside ml-2">
                <li>Make sure your antivirus is not blocking the installation of Chromium.</li>
                <li>Verify that you have correctly defined the variables in your <code>frontend.yml</code> file,
                    especially the base URL of your documentation and the CSS selectors for the UI elements.</li>
                <li>Consult the testflowkit documentation for more detailed information on troubleshooting and resolving
                    errors.</li>
            </ul>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Useful Resources</h2>
            <ul class="list-disc list-inside ml-2">
                <li><router-link :to="{ name: 'configuration' }" target="_blank" class="underline">testflowkit
                        Configuration</router-link></li>
                <li><router-link :to="{ name: 'sentences' }" target="_blank" class="underline">testflowkit
                        Sentence Repository</router-link></li>
            </ul>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">Next Steps</h2>
            <p>Now that you've learned the basics of testflowkit, here are some suggestions to go further:</p>
            <ul class="list-disc list-inside ml-2">
                <li>Write more complex Gherkin scenarios to test different parts of your web application.</li>
                <li>Explore the possibilities of integrating testflowkit with other testing and automation tools.</li>
                <li>Contribute to the testflowkit project by reporting bugs, suggesting improvements, or developing new
                    features.</li>
            </ul>
        </div>
    </div>

</template>

<script setup lang="ts">

import PrerequisitesList from '../components/PrerequisitesList.vue';
import { ref } from 'vue';

const suggestionShowed = ref(false);
function toggleSuggestion() {
    suggestionShowed.value = !suggestionShowed.value;
}


const frontendYml = `
# frontend.yml
global:
  elements:
    search_field:
      - input[type='search'] 
      - .search-input 
    category_button: 
      - .menu-link 
        
   pages:
    sentences_referencing: "/sentences" 
   
   base_url: "doc-base-url"
    ` 



const launchCommand = "./tkit run";

</script>

<style scoped>
.suggestion-toggle {
    @apply cursor-pointer text-blue-500 hover:underline;
}
</style>