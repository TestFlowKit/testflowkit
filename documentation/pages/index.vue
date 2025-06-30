<template>
    <section class="text-center mb-12">
        <h1 class="text-4xl font-bold mb-4">TestFlowKit: Powerful Web Test Automation Framework</h1>
        <p class="text-xl mb-8">
            TestFlowKit is a powerful, open-source web test automation framework built in Go that simplifies the
            creation and execution of automated tests for web applications. It uses Gherkin syntax for test
            specification and provides comprehensive support for both frontend and backend testing.
        </p>

        <div class="flex flex-wrap justify-center gap-4">
            <nuxt-link :to="{ name: 'qa-guide' }"
                class="bg-green-600 text-white font-bold py-3 px-6 rounded text-xl hover:bg-green-700">
                🧪 For QA Teams
            </nuxt-link>
            <nuxt-link :to="{ name: 'get-started' }"
                class="primary-button-bg text-white font-bold py-3 px-6 rounded text-xl">
                Get Started Now
            </nuxt-link>
        </div>
    </section>

    <!-- QA Teams Section -->
    <section class="bg-gradient-to-r from-green-50 to-blue-50 p-8 rounded-lg mb-12">
        <div class="text-center mb-6">
            <h2 class="text-3xl font-bold mb-4 text-gray-800">New to Test Automation?</h2>
            <p class="text-lg text-gray-600 mb-6">
                No coding experience required! Our QA Guide is designed specifically for non-technical QA professionals.
            </p>
        </div>

        <div class="grid md:grid-cols-3 gap-6 mb-6">
            <div class="bg-white p-6 rounded-lg text-center">
                <div class="text-3xl mb-3">📝</div>
                <h3 class="font-semibold mb-2">Write in Plain English</h3>
                <p class="text-sm text-gray-600">Describe your tests using natural language - no programming required
                </p>
            </div>
            <div class="bg-white p-6 rounded-lg text-center">
                <div class="text-3xl mb-3">🎯</div>
                <h3 class="font-semibold mb-2">Visual Examples</h3>
                <p class="text-sm text-gray-600">See real examples of how to test common scenarios</p>
            </div>
            <div class="bg-white p-6 rounded-lg text-center">
                <div class="text-3xl mb-3">🚀</div>
                <h3 class="font-semibold mb-2">Quick Start</h3>
                <p class="text-sm text-gray-600">Get up and running with automated testing in minutes</p>
            </div>
        </div>

        <div class="text-center">
            <nuxt-link :to="{ name: 'qa-guide' }"
                class="bg-green-600 text-white font-bold py-3 px-8 rounded-lg text-lg hover:bg-green-700 inline-flex items-center">
                <span>Start with QA Guide</span>
                <svg class="ml-2 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
            </nuxt-link>
        </div>
    </section>

    <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
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

    <section class="mt-12 bg-gray-50 p-8 rounded-lg">
        <h2 class="text-2xl font-bold mb-6 text-center">Why Choose TestFlowKit?</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
                <h3 class="text-lg font-semibold mb-3">🚀 Performance</h3>
                <p>Built with Go for blazing fast execution, saving valuable time in your testing workflow.</p>
            </div>
            <div>
                <h3 class="text-lg font-semibold mb-3">🔧 Extensibility</h3>
                <p>Plugin-like step definition system with validation and documentation for easy customization.</p>
            </div>
            <div>
                <h3 class="text-lg font-semibold mb-3">📊 Comprehensive</h3>
                <p>Support for both frontend UI testing and backend API testing in a single framework.</p>
            </div>
            <div>
                <h3 class="text-lg font-semibold mb-3">🛠️ Developer-Friendly</h3>
                <p>Clear error messages, detailed logging, and excellent debugging capabilities.</p>
            </div>
        </div>
    </section>
</template>

<script setup lang="ts">
import FeatureCard from '../components/FeatureCard.vue';
</script>
