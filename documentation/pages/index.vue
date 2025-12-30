<template>
    <main class="pb-12 px-4 sm:px-6 lg:px-8 max-w-7xl mx-auto pt-4">
        <!-- Hero Section -->
        <section class="text-center mb-8 md:mb-12">
            <h1 class="text-2xl sm:text-3xl md:text-4xl font-bold mb-4 text-gray-900 dark:text-white">TestFlowKit: Powerful Web Test Automation Framework</h1>
            <p class="text-base sm:text-lg md:text-xl mb-6 md:mb-8 text-gray-600 dark:text-gray-300 max-w-3xl mx-auto">
                TestFlowKit is a powerful, open-source web test automation framework built in Go that simplifies the
                creation and execution of automated tests for web applications. It uses Gherkin syntax for test
                specification and provides comprehensive support for both frontend and backend testing.
            </p>

            <div class="flex flex-col sm:flex-row flex-wrap justify-center gap-3 sm:gap-4">
                <nuxt-link to="/docs/getting-started/qa-guide"
                    class="bg-green-600 text-white font-bold py-3 px-6 rounded-lg text-base sm:text-lg md:text-xl hover:bg-green-700 transition-colors">
                    🧪 For QA Teams
                </nuxt-link>
                <nuxt-link to="/docs/getting-started/introduction"
                    class="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-lg text-base sm:text-lg md:text-xl transition-colors">
                    Get Started Now
                </nuxt-link>
            </div>
        </section>

        <!-- QA Teams Section -->
        <section class="bg-gradient-to-r from-green-50 to-blue-50 dark:from-green-900/20 dark:to-blue-900/20 p-4 sm:p-6 md:p-8 rounded-lg mb-8 md:mb-12 border border-green-100 dark:border-green-800">
            <div class="text-center mb-6">
                <h2 class="text-xl sm:text-2xl md:text-3xl font-bold mb-4 text-gray-800 dark:text-white">New to Test Automation?</h2>
                <p class="text-base sm:text-lg text-gray-600 dark:text-gray-300 mb-6">
                    No coding experience required! Our QA Guide is designed specifically for non-technical QA professionals.
                </p>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 md:gap-6 mb-6">
                <div class="bg-white dark:bg-gray-800 p-4 sm:p-6 rounded-lg text-center shadow-sm border border-gray-100 dark:border-gray-700">
                    <div class="text-2xl sm:text-3xl mb-3">📝</div>
                    <h3 class="font-semibold mb-2 text-gray-800 dark:text-white">Write in Plain English</h3>
                    <p class="text-sm text-gray-600 dark:text-gray-400">Describe your tests using natural language - no programming required</p>
                </div>
                <div class="bg-white dark:bg-gray-800 p-4 sm:p-6 rounded-lg text-center shadow-sm border border-gray-100 dark:border-gray-700">
                    <div class="text-2xl sm:text-3xl mb-3">🎯</div>
                    <h3 class="font-semibold mb-2 text-gray-800 dark:text-white">Visual Examples</h3>
                    <p class="text-sm text-gray-600 dark:text-gray-400">See real examples of how to test common scenarios</p>
                </div>
                <div class="bg-white dark:bg-gray-800 p-4 sm:p-6 rounded-lg text-center shadow-sm border border-gray-100 dark:border-gray-700 sm:col-span-2 md:col-span-1">
                    <div class="text-2xl sm:text-3xl mb-3">🚀</div>
                    <h3 class="font-semibold mb-2 text-gray-800 dark:text-white">Quick Start</h3>
                    <p class="text-sm text-gray-600 dark:text-gray-400">Get up and running with automated testing in minutes</p>
                </div>
            </div>

            <div class="text-center">
                <nuxt-link to="/docs/getting-started/qa-guide"
                    class="bg-green-600 text-white font-bold py-3 px-6 sm:px-8 rounded-lg text-base sm:text-lg hover:bg-green-700 inline-flex items-center transition-colors">
                    <span>Start with QA Guide</span>
                    <svg class="ml-2 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                    </svg>
                </nuxt-link>
            </div>
        </section>

        <!-- Features Grid -->
        <section class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6 mb-8 md:mb-12">
            <FeatureCard title="Gherkin Syntax" icon="chat">
                <p>Write tests using clear, readable BDD syntax with natural language that's easy to understand and
                    maintain.</p>
            </FeatureCard>
            <FeatureCard title="Multi-Environment Support" icon="cog">
                <p>Configure and switch between different environments (local, staging, production) with ease.</p>
            </FeatureCard>
            <FeatureCard title="Frontend & Backend Testing" icon="code">
                <p>Comprehensive web UI automation with smart element detection and full REST API testing capabilities.</p>
            </FeatureCard>
            <FeatureCard title="Macro System" icon="puzzle">
                <p>Reusable test scenarios to reduce code duplication and improve maintainability.</p>
            </FeatureCard>
            <FeatureCard title="Parallel Execution" icon="bolt">
                <p>Run tests concurrently for faster execution, leveraging the power of Go for incredible performance.</p>
            </FeatureCard>
            <FeatureCard title="Rich Reporting" icon="chart-bar">
                <p>HTML and JSON report formats with detailed test results and automatic screenshot capture on failures.</p>
            </FeatureCard>
            <FeatureCard title="Smart Element Detection" icon="search">
                <p>Multiple selector strategies with fallback mechanisms for robust element detection.</p>
            </FeatureCard>
            <FeatureCard title="Headless Mode" icon="monitor">
                <p>Run tests without browser UI for CI/CD environments with configurable headless execution.</p>
            </FeatureCard>
            <FeatureCard title="Configuration Management" icon="settings">
                <p>YAML-based configuration with environment-specific settings and comprehensive validation.</p>
            </FeatureCard>
        </section>

        <!-- Why Choose Section -->
        <section class="bg-white dark:bg-gray-800 p-4 sm:p-6 md:p-8 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700">
            <h2 class="text-xl sm:text-2xl font-bold mb-6 text-center text-gray-900 dark:text-white">Why Choose TestFlowKit?</h2>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 md:gap-6">
                <div class="p-3 sm:p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
                    <h3 class="text-base sm:text-lg font-semibold mb-2 sm:mb-3 text-gray-800 dark:text-white">🚀 Performance</h3>
                    <p class="text-sm sm:text-base text-gray-600 dark:text-gray-300">Built with Go for blazing fast execution, saving valuable time in your testing workflow.</p>
                </div>
                <div class="p-3 sm:p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
                    <h3 class="text-base sm:text-lg font-semibold mb-2 sm:mb-3 text-gray-800 dark:text-white">🔧 Extensibility</h3>
                    <p class="text-sm sm:text-base text-gray-600 dark:text-gray-300">Plugin-like step definition system with validation and documentation for easy customization.</p>
                </div>
                <div class="p-3 sm:p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
                    <h3 class="text-base sm:text-lg font-semibold mb-2 sm:mb-3 text-gray-800 dark:text-white">📊 Comprehensive</h3>
                    <p class="text-sm sm:text-base text-gray-600 dark:text-gray-300">Support for both frontend UI testing and backend API testing in a single framework.</p>
                </div>
                <div class="p-3 sm:p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
                    <h3 class="text-base sm:text-lg font-semibold mb-2 sm:mb-3 text-gray-800 dark:text-white">🛠️ Developer-Friendly</h3>
                    <p class="text-sm sm:text-base text-gray-600 dark:text-gray-300">Clear error messages, detailed logging, and excellent debugging capabilities.</p>
                </div>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import FeatureCard from '~/components/FeatureCard.vue';
</script>
