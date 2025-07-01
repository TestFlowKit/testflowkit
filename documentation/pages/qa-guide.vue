<template>
    <div class="space-y-8">
        <!-- Hero Section -->
        <section class="text-center hero-gradient p-8 rounded-xl relative overflow-hidden">
            <div class="relative z-10">
                <h1 class="heading-1 mb-4 text-white">TestFlowKit for QA Teams</h1>
                <p class="body-large mb-6 text-gray-700 max-w-2xl mx-auto text-white">
                    Write automated tests in plain English - no coding required!
                </p>
                <div class="flex flex-col sm:flex-row justify-center gap-4">
                    <button @click="scrollToSection('what-is')"
                        class="btn btn-primary">
                        What is TestFlowKit?
                    </button>
                    <button @click="scrollToSection('how-it-works')"
                        class="btn btn-secondary">
                        How It Works
                    </button>
                </div>
            </div>
        </section>

        <!-- What is TestFlowKit Section -->
        <section id="what-is" class="card p-8">
            <h2 class="heading-2 mb-6 text-white">What is TestFlowKit?</h2>

            <div class="grid md:grid-cols-2 gap-8 items-center">
                <div>
                    <p class="body-normal mb-4 text-white">
                        TestFlowKit is like having a robot that can test your website automatically. Instead of manually
                        clicking buttons and checking pages, you write instructions in simple English, and the robot
                        does all the work for you.
                    </p>

                    <div class="alert alert-success">
                        <h3 class="font-semibold text-green-800 mb-2">🎯 Perfect for QA Teams</h3>
                        <ul class="text-green-700 space-y-1">
                            <li>• No programming knowledge needed</li>
                            <li>• Write tests in natural language</li>
                            <li>• Visual feedback with screenshots</li>
                            <li>• Easy to understand test reports</li>
                        </ul>
                    </div>
                </div>

                <div class="card-gradient p-6 rounded-lg">
                    <h3 class="font-semibold mb-3">Example: Testing a Login Form</h3>
                    <div class="space-y-2 text-sm">
                        <div class="bg-white p-3 rounded border border-orange-200">
                            <span class="text-orange-600 font-medium">Given</span> the user is on the login page
                        </div>
                        <div class="bg-white p-3 rounded border border-orange-200">
                            <span class="text-teal-600 font-medium">When</span> the user enters "test@email.com" into
                            the email field
                        </div>
                        <div class="bg-white p-3 rounded border border-orange-200">
                            <span class="text-teal-600 font-medium">And</span> the user enters "password123" into the
                            password field
                        </div>
                        <div class="bg-white p-3 rounded border border-orange-200">
                            <span class="text-teal-600 font-medium">And</span> the user clicks on the login button
                        </div>
                        <div class="bg-white p-3 rounded border border-orange-200">
                            <span class="text-orange-700 font-medium">Then</span> the user should be navigated to
                            dashboard page
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- How It Works Section -->
        <section id="how-it-works" class="card p-8">
            <h2 class="heading-2 mb-6 text-white">How TestFlowKit Works</h2>

            <div class="grid md:grid-cols-3 gap-6">
                <div class="feature-card">
                    <div class="w-16 h-16 mx-auto mb-4 p-3 bg-gradient-to-br from-orange-500 to-orange-600 rounded-xl flex items-center justify-center">
                        <span class="text-4xl">📝</span>
                    </div>
                    <h3 class="heading-4 mb-3">1. Write Your Test</h3>
                    <p class="body-normal text-gray-600">
                        Describe what you want to test using simple English sentences. Think of it like writing a recipe
                        for testing.
                    </p>
                </div>

                <div class="feature-card">
                    <div class="w-16 h-16 mx-auto mb-4 p-3 bg-gradient-to-br from-teal-500 to-teal-600 rounded-xl flex items-center justify-center">
                        <span class="text-4xl">🤖</span>
                    </div>
                    <h3 class="heading-4 mb-3">2. Run the Test</h3>
                    <p class="body-normal text-gray-600">
                        TestFlowKit reads your instructions and automatically performs the actions in a real browser.
                    </p>
                </div>

                <div class="feature-card">
                    <div class="w-16 h-16 mx-auto mb-4 p-3 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center">
                        <span class="text-4xl">📊</span>
                    </div>
                    <h3 class="heading-4 mb-3">3. Get Results</h3>
                    <p class="body-normal text-gray-600">
                        Receive a detailed report showing what passed, what failed, and screenshots of any issues.
                    </p>
                </div>
            </div>
        </section>

        <!-- Real Example Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Real Example: Testing an E-commerce Site</h2>

            <div class="card-gradient p-6 rounded-lg mb-6">
                <h3 class="font-semibold mb-4">Scenario: Add Item to Shopping Cart</h3>
                <p class="text-gray-600 mb-4">Let's say you want to test that users can add items to their shopping
                    cart:</p>

                <div class="bg-white p-4 rounded border-l-4 border-orange-400">
                    <h4 class="font-medium mb-2">Your Test Instructions:</h4>
                    <ol class="space-y-2 text-sm">
                        <li>1. Go to the product catalog page</li>
                        <li>2. Click on the "Add to Cart" button for "Blue T-Shirt"</li>
                        <li>3. Check that the cart shows "1 item"</li>
                        <li>4. Verify the total price is correct</li>
                    </ol>
                </div>
            </div>

            <div class="card-gradient p-6 rounded-lg">
                <h3 class="font-semibold mb-4">How You Write It in TestFlowKit:</h3>
                <div class="code-block">
                    <pre class="code-content"><code><span class="text-orange-400">Feature: Shopping Cart</span>
<span class="text-orange-400"> As a customer</span>
<span class="text-orange-400"> I want to add items to my cart</span>
<span class="text-orange-400"> So that I can purchase them</span>

<span class="text-teal-400"> Scenario: Add item to shopping cart</span>
<span class="text-orange-300"> Given the user is on the "catalog" page</span>
<span class="text-orange-300"> When the user clicks on the "Add to Cart" button</span>
<span class="text-orange-300"> Then the "cart_count" should contain "1"</span>
<span class="text-orange-300"> And the "total_price" should be visible</span></code></pre>
                </div>
            </div>
        </section>

        <!-- Getting Started Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Getting Started (QA-Friendly Guide)</h2>

            <div class="space-y-6">
                <div class="border-l-4 border-orange-400 pl-6 bg-orange-50 p-4 rounded-r-lg">
                    <h3 class="heading-4 mb-3">Step 1: Ask Your Developer</h3>
                    <p class="body-normal text-gray-600 mb-2">
                        Ask your development team to set up TestFlowKit for your project. They'll need to:
                    </p>
                    <ul class="list-disc list-inside text-gray-600 ml-4 space-y-1">
                        <li>Install TestFlowKit on your testing environment</li>
                        <li>Configure it to work with your application</li>
                        <li>Set up the basic test structure</li>
                    </ul>
                </div>

                <div class="border-l-4 border-teal-400 pl-6 bg-teal-50 p-4 rounded-r-lg">
                    <h3 class="heading-4 mb-3">Step 2: Learn the Basic Keywords</h3>
                    <p class="body-normal text-gray-600 mb-2">
                        You only need to know a few simple words to write tests:
                    </p>
                    <div class="grid md:grid-cols-2 gap-4 mt-4">
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-orange-600">Given</div>
                            <div class="text-sm text-gray-600">Sets up the starting point</div>
                        </div>
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-teal-600">When</div>
                            <div class="text-sm text-gray-600">Describes the action to take</div>
                        </div>
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-teal-600">And</div>
                            <div class="text-sm text-gray-600">Adds more actions or checks</div>
                        </div>
                        <div class="bg-white p-4 rounded border border-orange-200">
                            <div class="font-semibold text-orange-600">Then</div>
                            <div class="text-sm text-gray-600">Describes what should happen</div>
                        </div>
                    </div>
                </div>

                <div class="border-l-4 border-green-400 pl-6 bg-green-50 p-4 rounded-r-lg">
                    <h3 class="heading-4 mb-3">Step 3: Write Your First Test</h3>
                    <p class="body-normal text-gray-600 mb-4">
                        Start with a simple test for a feature you know well. For example, testing that a page loads
                        correctly:
                    </p>
                    <div class="code-block">
                        <pre class="code-content"><code><span class="text-orange-400">Feature: Homepage</span>
<span class="text-teal-400"> Scenario: Homepage loads correctly</span>
<span class="text-orange-300"> Given the user is on the homepage</span>
<span class="text-orange-300"> Then the "welcome_message" should be visible</span>
<span class="text-orange-300"> And the "navigation_menu" should be visible</span></code></pre>
                    </div>
                </div>
            </div>
        </section>

        <!-- Common Test Patterns Section -->
        <section class="card p-8">
            <h2 class="heading-2 mb-6 text-white">Common Test Patterns You'll Use</h2>

            <div class="grid md:grid-cols-2 gap-6">
                <div class="card-gradient p-6 rounded-lg border border-orange-200">
                    <h3 class="font-semibold mb-3">🔍 Checking if Something is Visible</h3>
                    <div class="code-block">
                        <pre class="code-content"><code><span class="text-orange-300">Then the "login_button" should be visible</span></code></pre>
                    </div>
                </div>

                <div class="card-gradient p-6 rounded-lg border border-orange-200">
                    <h3 class="font-semibold mb-3">✏️ Filling Out Forms</h3>
                    <div class="code-block">
                        <pre class="code-content"><code><span class="text-teal-300">When the user enters "john@email.com" into the "email" field</span></code></pre>
                    </div>
                </div>

                <div class="card-gradient p-6 rounded-lg border border-orange-200">
                    <h3 class="font-semibold mb-3">🖱️ Clicking Buttons</h3>
                    <div class="code-block">
                        <pre class="code-content"><code><span class="text-teal-300">When the user clicks on the "submit" button</span></code></pre>
                    </div>
                </div>

                <div class="card-gradient p-6 rounded-lg border border-orange-200">
                    <h3 class="font-semibold mb-3">📋 Verifying Text Content</h3>
                    <div class="code-block">
                        <pre class="code-content"><code><span class="text-orange-300">Then the "success_message" should contain "Thank you"</span></code></pre>
                    </div>
                </div>
            </div>
        </section>

        <!-- Call to Action -->
        <section class="bg-gradient-to-r from-orange-500 to-orange-600 text-white p-8 rounded-xl text-center">
            <h2 class="heading-2 mb-4 text-white">Ready to Start Testing?</h2>
            <p class="body-large mb-6 text-orange-100 max-w-2xl mx-auto">
                Get TestFlowKit set up for your team and start writing your first automated tests today.
            </p>
            <div class="flex flex-col sm:flex-row justify-center gap-4">
                <nuxt-link to="/get-started" class="btn bg-white text-orange-600 hover:bg-gray-50">
                    Get Started Guide
                </nuxt-link>
                <nuxt-link to="/sentences" class="btn btn-ghost text-white border-white hover:bg-white/10">
                    View All Gherkin Steps
                </nuxt-link>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
// SEO and Meta
useHead({
  title: 'QA Guide - TestFlowKit Documentation',
  meta: [
    {
      name: 'description',
      content: 'Complete guide for QA teams to get started with TestFlowKit. Write automated tests in plain English with no coding required.'
    }
  ]
})

const scrollToSection = (sectionId: string) => {
    const element = document.getElementById(sectionId);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    }
};
</script>