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
    <p class="mb-6">
      Choose the binary corresponding to your operating system and architecture:
    </p>

    <div class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-8">
      <div class="text-center space-y-2">
        <h2 class="text-lg font-semibold">Windows</h2>
        <a
          href="testflowkit-windows-amd64"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded block"
          >Windows (amd64)</a
        >
        <a
          href="testflowkit-windows-386"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded block"
          >Windows (386)</a
        >
      </div>

      <div class="text-center space-y-2">
        <h2 class="text-lg font-semibold">macOS</h2>
        <a
          href="testflowkit-macos-arm64"
          class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded block"
          >macOS (arm64)</a
        >
        <a
          href="testflowkit-macos-amd64"
          class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded block"
          >macOS (amd64)</a
        >
      </div>

      <div class="text-center space-y-2">
        <h2 class="text-lg font-semibold">Linux</h2>
        <a
          href="testflowkit-linux-amd64"
          class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded block"
          >Linux (amd64)</a
        >
        <a
          href="testflowkit-linux-386"
          class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded block"
          >Linux (386)</a
        >
        <a
          href="testflowkit-linux-arm64"
          class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded block"
          >Linux (arm64)</a
        >
      </div>
    </div>


    
  </section>

  <section>
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
  <!-- <section id="download" class="mb-12">
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
    </section> -->
</template>

<script setup lang="ts">
import PrerequisitesList from "../components/PrerequisitesList.vue";
</script> 
