<template>
    <section id="prerequisites" class="mb-12">
        <h2 class="text-2xl font-bold mb-4">Prerequisites</h2>

        <p class="mb-4">
            Before you get started with testflowkit, make sure you have the following:
        </p>

        <prerequisites-list></prerequisites-list>
    </section>

    <section id="download" class="mb-12">
        <h2 class="text-2xl font-bold mb-4">Download testflowkit</h2>

        <p class="mb-4">
            testflowkit is available for Windows, macOS, and Linux. Choose the download link that matches your operating
            system:
        </p>

        <div class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-8">
            <a href="#" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex items-center">
                <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/windows8/windows8-original.svg"
                    alt="Windows logo" class="h-6 mr-2"> Download for Windows
            </a>
            <a href="#" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex items-center">
                <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/apple/apple-original.svg" alt="macOS logo"
                    class="h-6 mr-2"> Download for macOS
            </a>
            <a href="#" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex items-center">
                <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/linux/linux-original.svg" alt="Linux logo"
                    class="h-6 mr-2"> Download for Linux
            </a>
        </div>

        <p class="mt-4 text-sm text-gray-600">
            For configuration instructions, please refer to the
            <router-link :to="{ name: 'configuration' }"
                class="text-blue-500 hover:underline">Configuration</router-link> guide.
        </p>

        <p class="mt-4 text-sm text-gray-600">
            For quick start tutorial, please refer to the
            <router-link :to="{ name: 'quick-start' }" class="text-blue-500 hover:underline">Quick
                start</router-link>.
        </p>
        <p class="mt-4 text-sm text-gray-600">
            For available gherkin phrases, please refer to the
            <router-link :to="{ name: 'sentences' }" class="text-blue-500 hover:underline">Sentences</router-link>
            dictionary.
        </p>
    </section>
</template>

<script setup lang="ts">
import PrerequisitesList from '../components/PrerequisitesList.vue';

</script>
