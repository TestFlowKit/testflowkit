<template>
    <div class="space-y-8">
        <section class="text-center bg-gradient-to-r from-blue-50 to-indigo-50 p-8 rounded-lg">
            <h1 class="text-4xl font-bold mb-4 text-gray-800">TestFlowKit for QA Teams</h1>
            <p class="text-xl text-gray-600 mb-6">
                Write automated tests in plain English - no coding required!
            </p>
            <div class="flex justify-center space-x-4">
                <button @click="scrollToSection('what-is')"
                    class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700">
                    What is TestFlowKit?
                </button>
                <button @click="scrollToSection('how-it-works')"
                    class="bg-green-600 text-white px-6 py-3 rounded-lg hover:bg-green-700">
                    How It Works
                </button>
            </div>
        </section>

        <section id="what-is" class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">What is TestFlowKit?</h2>

            <div class="grid md:grid-cols-2 gap-8 items-center">
                <div>
                    <p class="text-lg text-gray-600 mb-4">
                        TestFlowKit is like having a robot that can test your website automatically. Instead of manually
                        clicking buttons and checking pages, you write instructions in simple English, and the robot
                        does all the work for you.
                    </p>

                    <div class="bg-green-50 border-l-4 border-green-400 p-4 mb-4">
                        <h3 class="font-semibold text-green-800 mb-2">🎯 Perfect for QA Teams</h3>
                        <ul class="text-green-700 space-y-1">
                            <li>• No programming knowledge needed</li>
                            <li>• Write tests in natural language</li>
                            <li>• Visual feedback with screenshots</li>
                            <li>• Easy to understand test reports</li>
                        </ul>
                    </div>
                </div>

                <div class="bg-gray-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-3">Example: Testing a Login Form</h3>
                    <div class="space-y-2 text-sm">
                        <div class="bg-white p-3 rounded border">
                            <span class="text-blue-600 font-medium">Given</span> the user is on the login page
                        </div>
                        <div class="bg-white p-3 rounded border">
                            <span class="text-green-600 font-medium">When</span> the user enters "test@email.com" into
                            the email field
                        </div>
                        <div class="bg-white p-3 rounded border">
                            <span class="text-green-600 font-medium">And</span> the user enters "password123" into the
                            password field
                        </div>
                        <div class="bg-white p-3 rounded border">
                            <span class="text-green-600 font-medium">And</span> the user clicks on the login button
                        </div>
                        <div class="bg-white p-3 rounded border">
                            <span class="text-purple-600 font-medium">Then</span> the user should be navigated to
                            dashboard page
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section id="how-it-works" class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">How TestFlowKit Works</h2>

            <div class="grid md:grid-cols-3 gap-6">
                <div class="text-center p-6 bg-blue-50 rounded-lg">
                    <div class="text-4xl mb-4">📝</div>
                    <h3 class="text-xl font-semibold mb-3">1. Write Your Test</h3>
                    <p class="text-gray-600">
                        Describe what you want to test using simple English sentences. Think of it like writing a recipe
                        for testing.
                    </p>
                </div>

                <div class="text-center p-6 bg-green-50 rounded-lg">
                    <div class="text-4xl mb-4">🤖</div>
                    <h3 class="text-xl font-semibold mb-3">2. Run the Test</h3>
                    <p class="text-gray-600">
                        TestFlowKit reads your instructions and automatically performs the actions in a real browser.
                    </p>
                </div>

                <div class="text-center p-6 bg-purple-50 rounded-lg">
                    <div class="text-4xl mb-4">📊</div>
                    <h3 class="text-xl font-semibold mb-3">3. Get Results</h3>
                    <p class="text-gray-600">
                        Receive a detailed report showing what passed, what failed, and screenshots of any issues.
                    </p>
                </div>
            </div>
        </section>

        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Real Example: Testing an E-commerce Site</h2>

            <div class="bg-gray-50 p-6 rounded-lg mb-6">
                <h3 class="font-semibold mb-4">Scenario: Add Item to Shopping Cart</h3>
                <p class="text-gray-600 mb-4">Let's say you want to test that users can add items to their shopping
                    cart:</p>

                <div class="bg-white p-4 rounded border-l-4 border-blue-400">
                    <h4 class="font-medium mb-2">Your Test Instructions:</h4>
                    <ol class="space-y-2 text-sm">
                        <li>1. Go to the product catalog page</li>
                        <li>2. Click on the "Add to Cart" button for "Blue T-Shirt"</li>
                        <li>3. Check that the cart shows "1 item"</li>
                        <li>4. Verify the total price is correct</li>
                    </ol>
                </div>
            </div>

            <div class="bg-blue-50 p-6 rounded-lg">
                <h3 class="font-semibold mb-4">How You Write It in TestFlowKit:</h3>
                <div class="bg-white p-4 rounded font-mono text-sm">
                    <div class="text-blue-600">Feature: Shopping Cart</div>
                    <div class="text-blue-600"> As a customer</div>
                    <div class="text-blue-600"> I want to add items to my cart</div>
                    <div class="text-blue-600"> So that I can purchase them</div>
                    <br>
                    <div class="text-green-600"> Scenario: Add item to shopping cart</div>
                    <div class="text-purple-600"> Given the user is on the "catalog" page</div>
                    <div class="text-orange-600"> When the user clicks on the "Add to Cart" button</div>
                    <div class="text-purple-600"> Then the "cart_count" should contain "1"</div>
                    <div class="text-purple-600"> And the "total_price" should be visible</div>
                </div>
            </div>
        </section>

        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Getting Started (QA-Friendly Guide)</h2>

            <div class="space-y-6">
                <div class="border-l-4 border-green-400 pl-6">
                    <h3 class="text-xl font-semibold mb-3">Step 1: Ask Your Developer</h3>
                    <p class="text-gray-600 mb-2">
                        Ask your development team to set up TestFlowKit for your project. They'll need to:
                    </p>
                    <ul class="list-disc list-inside text-gray-600 ml-4 space-y-1">
                        <li>Install TestFlowKit on your testing environment</li>
                        <li>Configure it to work with your application</li>
                        <li>Set up the basic test structure</li>
                    </ul>
                </div>

                <div class="border-l-4 border-blue-400 pl-6">
                    <h3 class="text-xl font-semibold mb-3">Step 2: Learn the Basic Keywords</h3>
                    <p class="text-gray-600 mb-2">
                        You only need to know a few simple words to write tests:
                    </p>
                    <div class="grid md:grid-cols-2 gap-4 mt-4">
                        <div class="bg-gray-50 p-4 rounded">
                            <div class="font-semibold text-purple-600">Given</div>
                            <div class="text-sm text-gray-600">Sets up the starting point</div>
                        </div>
                        <div class="bg-gray-50 p-4 rounded">
                            <div class="font-semibold text-orange-600">When</div>
                            <div class="text-sm text-gray-600">Describes the action to take</div>
                        </div>
                        <div class="bg-gray-50 p-4 rounded">
                            <div class="font-semibold text-green-600">And</div>
                            <div class="text-sm text-gray-600">Adds more actions or checks</div>
                        </div>
                        <div class="bg-gray-50 p-4 rounded">
                            <div class="font-semibold text-purple-600">Then</div>
                            <div class="text-sm text-gray-600">Describes what should happen</div>
                        </div>
                    </div>
                </div>

                <div class="border-l-4 border-purple-400 pl-6">
                    <h3 class="text-xl font-semibold mb-3">Step 3: Write Your First Test</h3>
                    <p class="text-gray-600 mb-4">
                        Start with a simple test for a feature you know well. For example, testing that a page loads
                        correctly:
                    </p>
                    <div class="bg-gray-50 p-4 rounded font-mono text-sm">
                        <div class="text-blue-600">Feature: Homepage</div>
                        <div class="text-green-600"> Scenario: Homepage loads correctly</div>
                        <div class="text-purple-600"> Given the user is on the homepage</div>
                        <div class="text-purple-600"> Then the "welcome_message" should be visible</div>
                        <div class="text-purple-600"> And the "navigation_menu" should be visible</div>
                    </div>
                </div>
            </div>
        </section>

        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Common Test Patterns You'll Use</h2>

            <div class="grid md:grid-cols-2 gap-6">
                <div class="bg-yellow-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-3">🔍 Checking if Something is Visible</h3>
                    <div class="bg-white p-3 rounded font-mono text-sm">
                        <div class="text-purple-600">Then the "login_button" should be visible</div>
                    </div>
                </div>

                <div class="bg-green-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-3">📝 Filling Out Forms</h3>
                    <div class="bg-white p-3 rounded font-mono text-sm">
                        <div class="text-orange-600">When the user enters "john@email.com" into the "email" field</div>
                        <div class="text-orange-600">And the user enters "password123" into the "password" field</div>
                    </div>
                </div>

                <div class="bg-blue-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-3">🖱️ Clicking Buttons</h3>
                    <div class="bg-white p-3 rounded font-mono text-sm">
                        <div class="text-orange-600">When the user clicks on the "submit" button</div>
                    </div>
                </div>

                <div class="bg-purple-50 p-6 rounded-lg">
                    <h3 class="font-semibold mb-3">📄 Checking Text Content</h3>
                    <div class="bg-white p-3 rounded font-mono text-sm">
                        <div class="text-purple-600">Then the "success_message" should contain "Welcome back!"</div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Tips and Best Practices -->
        <section class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Tips for Writing Good Tests</h2>

            <div class="grid md:grid-cols-2 gap-6">
                <div class="space-y-4">
                    <div class="flex items-start space-x-3">
                        <div class="text-2xl">✅</div>
                        <div>
                            <h4 class="font-semibold">Keep Tests Simple</h4>
                            <p class="text-sm text-gray-600">One test should check one specific thing. Don't try to test
                                everything in one scenario.</p>
                        </div>
                    </div>

                    <div class="flex items-start space-x-3">
                        <div class="text-2xl">✅</div>
                        <div>
                            <h4 class="font-semibold">Use Clear Names</h4>
                            <p class="text-sm text-gray-600">Name your scenarios so anyone can understand what you're
                                testing.</p>
                        </div>
                    </div>

                    <div class="flex items-start space-x-3">
                        <div class="text-2xl">✅</div>
                        <div>
                            <h4 class="font-semibold">Test Real User Flows</h4>
                            <p class="text-sm text-gray-600">Focus on testing the paths that real users actually take.
                            </p>
                        </div>
                    </div>
                </div>

                <div class="space-y-4">
                    <div class="flex items-start space-x-3">
                        <div class="text-2xl">❌</div>
                        <div>
                            <h4 class="font-semibold">Don't Test Everything</h4>
                            <p class="text-sm text-gray-600">Focus on critical user journeys, not every possible
                                combination.</p>
                        </div>
                    </div>

                    <div class="flex items-start space-x-3">
                        <div class="text-2xl">❌</div>
                        <div>
                            <h4 class="font-semibold">Don't Make Tests Fragile</h4>
                            <p class="text-sm text-gray-600">Avoid testing things that change frequently, like specific
                                text or styling.</p>
                        </div>
                    </div>

                    <div class="flex items-start space-x-3">
                        <div class="text-2xl">❌</div>
                        <div>
                            <h4 class="font-semibold">Don't Skip the Basics</h4>
                            <p class="text-sm text-gray-600">Always test the happy path first before adding edge cases.
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section class="bg-gradient-to-r from-green-50 to-blue-50 p-8 rounded-lg">
            <h2 class="text-3xl font-bold mb-6 text-center text-gray-800">Ready to Get Started?</h2>

            <div class="text-center space-y-4">
                <p class="text-lg text-gray-600">
                    Now that you understand the basics, here's what you can do next:
                </p>

                <div class="flex flex-wrap justify-center gap-4">
                    <nuxt-link to="/sentences" class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700">
                        📚 Browse Available Keywords
                    </nuxt-link>
                    <nuxt-link to="/configuration"
                        class="bg-green-600 text-white px-6 py-3 rounded-lg hover:bg-green-700">
                        ⚙️ Learn About Configuration
                    </nuxt-link>
                    <a href="https://testflowkit.dreamsfollowers.me/" target="_blank"
                        class="bg-purple-600 text-white px-6 py-3 rounded-lg hover:bg-purple-700">
                        🌐 Full Documentation
                    </a>
                </div>

                <div class="bg-white p-4 rounded-lg mt-6">
                    <p class="text-sm text-gray-600">
                        <strong>Need help?</strong> Ask your development team or check out the examples in the
                        <nuxt-link to="/sentences" class="text-blue-600 hover:underline">Gherkin Sentences
                            Dictionary</nuxt-link>.
                    </p>
                </div>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
const scrollToSection = (sectionId: string) => {
    const element = document.getElementById(sectionId);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    }
};
</script>