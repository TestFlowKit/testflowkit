<template>
    <section class="text-center">
        <p class="text-xl mb-8">
            testflowkit is an open source tool developed in Go that allows you to generate automated tests from Gherkin
            files. Simple, fast and extensible, it is aimed at developers and QA testers.
        </p>

        <nuxt-link :to="{ name: 'get-started' }"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-4 px-6 rounded text-xl">
            Get started now
        </nuxt-link>
    </section>

    <section class="mt-12">
        <h2 class="text-2xl font-bold mb-4">Advantages</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
            <div class="bg-white p-4 rounded shadow">
                <h3 class="text-lg font-bold mb-2">Ease of Use</h3>
                <p>Write your tests in natural language, testflowkit takes care of the rest.</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
                <h3 class="text-lg font-bold mb-2">Fast Execution</h3>
                <p>Go, a compiled language, ensures optimal performance for your tests.</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
                <h3 class="text-lg font-bold mb-2">Continuous Integration</h3>
                <p>Easily integrate testflowkit into your CI/CD pipelines.</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
                <h3 class="text-lg font-bold mb-2">Open Source and Extensible</h3>
                <p>Contribute to the project and adapt it to your specific needs.</p>
            </div>
        </div>
    </section>
</template>

<script setup lang="ts">


</script>

<style scoped></style>
