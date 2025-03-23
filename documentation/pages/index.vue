<template>
    <section class="text-center mb-12">
        <h1 class="text-4xl font-bold mb-4">testflowkit: Automated Front-End Testing Made Easy</h1>
        <p class="text-xl mb-8">
            testflowkit is an open-source tool built with Go that simplifies automated testing for your web
            applications. Generate tests from Gherkin files quickly and easily, perfect for developers and QA testers
            alike.
        </p>

        <nuxt-link :to="{ name: 'get-started' }"
            class="primary-button-bg text-white font-bold py-3 px-6 rounded text-xl">
            Get Started Now
        </nuxt-link>
    </section>

    <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <FeatureCard title="Effortless Setup" icon="cog">
            <p>Quick and easy installation and configuration. Get started testing in minutes.</p>
        </FeatureCard>
        <FeatureCard title="Blazing Fast Execution" icon="bolt">
            <p>Leverage the power of Go for incredibly fast test execution, saving you valuable time.</p>
        </FeatureCard>
        <FeatureCard title="Seamless CI/CD Integration" icon="pipeline">
            <p>Integrate testflowkit into your existing CI/CD pipelines for continuous testing and automated feedback.
            </p>
        </FeatureCard>
        <FeatureCard title="Open Source & Extensible" icon="code">
            <p>Contribute to the project, extend its functionality, and tailor it to your specific needs.</p>
        </FeatureCard>
        <FeatureCard title="Human-Readable Tests" icon="chat">
            <p>Write your tests in Gherkin, a natural language format, making them easy to understand and maintain.</p>
        </FeatureCard>
        <FeatureCard title="Comprehensive Reporting" icon="chart-bar">
            <p>Receive detailed reports on your test results, making it easy to identify and fix issues.</p>
        </FeatureCard>
    </section>
</template>

<script setup lang="ts">
import FeatureCard from '../components/FeatureCard.vue';
</script>
