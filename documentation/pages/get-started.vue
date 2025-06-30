<template>
  <div class="space-y-8">
    <!-- Hero Section -->
    <section class="text-center bg-gradient-to-r from-blue-50 to-indigo-50 p-8 rounded-lg">
      <h1 class="text-4xl font-bold mb-4 text-gray-800">Getting Started with TestFlowKit</h1>
      <p class="text-xl text-gray-600 mb-6">
        Set up TestFlowKit and write your first automated test in minutes
      </p>
    </section>

    <!-- What You Need Section -->
    <section class="bg-white p-8 rounded-lg shadow-md">
      <h2 class="text-3xl font-bold mb-6 text-gray-800">What You Need to Get Started</h2>

      <div class="bg-green-50 border-l-4 border-green-400 p-6 rounded-lg">
        <h3 class="text-xl font-semibold mb-4 text-green-800">For QA Teams & Non-Technical Users</h3>
        <div class="space-y-3">
          <div class="flex items-start space-x-3">
            <div class="text-2xl">✅</div>
            <div>
              <h4 class="font-semibold">TestFlowKit installed</h4>
              <p class="text-sm text-gray-600">Your development team will set this up for you</p>
            </div>
          </div>
          <div class="flex items-start space-x-3">
            <div class="text-2xl">✅</div>
            <div>
              <h4 class="font-semibold">Access to your application</h4>
              <p class="text-sm text-gray-600">The website or app you want to test</p>
            </div>
          </div>
          <div class="flex items-start space-x-3">
            <div class="text-2xl">✅</div>
            <div>
              <h4 class="font-semibold">Basic understanding of your app</h4>
              <p class="text-sm text-gray-600">Know what features you want to test</p>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-6 bg-blue-50 border-l-4 border-blue-400 p-6 rounded-lg">
        <h3 class="text-xl font-semibold mb-4 text-blue-800">For Developers & Technical Users</h3>
        <p class="text-blue-700 mb-4">
          If you need technical setup instructions, configuration details, or advanced features, please visit our GitHub
          repository.
        </p>
        <a href="https://github.com/TestFlowKit/testflowkit" target="_blank"
          class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 inline-flex items-center">
          <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
            <path
              d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
          </svg>
          View on GitHub
        </a>
      </div>
    </section>

    <!-- Quick Setup Section -->
    <section class="bg-white p-8 rounded-lg shadow-md">
      <h2 class="text-3xl font-bold mb-6 text-gray-800">Quick Setup</h2>

      <div class="space-y-6">
        <div class="border-l-4 border-green-400 pl-6">
          <h3 class="text-xl font-semibold mb-3">Step 1: Download TestFlowKit</h3>
          <p class="text-gray-600 mb-4">
            Download TestFlowKit for your operating system and install it on your machine.
          </p>
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="font-semibold mb-4 text-gray-800">Choose your platform and architecture:</h4>
            <DownloadSection />
          </div>
        </div>

        <div class="border-l-4 border-blue-400 pl-6">
          <h3 class="text-xl font-semibold mb-3">Step 2: Ask Your Development Team</h3>
          <p class="text-gray-600 mb-4">
            Have your development team set up TestFlowKit in your project. They'll need to:
          </p>
          <ul class="list-disc list-inside text-gray-600 ml-4 space-y-2">
            <li>Initialize TestFlowKit in your project directory</li>
            <li>Configure it to work with your application</li>
            <li>Set up the basic test structure</li>
            <li>Configure the application URLs and settings</li>
          </ul>
        </div>
      </div>
    </section>

    <!-- Next Steps Section -->
    <section class="bg-gradient-to-r from-green-50 to-blue-50 p-8 rounded-lg">
      <h2 class="text-3xl font-bold mb-6 text-center text-gray-800">What's Next?</h2>

      <div class="text-center space-y-6">
        <div class="grid md:grid-cols-2 gap-6">
          <div class="bg-white p-6 rounded-lg">
            <h3 class="font-semibold mb-3">For QA Teams</h3>
            <ul class="space-y-2 text-sm text-gray-600 text-left">
              <li>• Learn to write tests in the <nuxt-link to="/quick-start" class="text-blue-600 hover:underline">Quick
                  Start Tutorial</nuxt-link></li>
              <li>• Explore test patterns in the <nuxt-link to="/sentences"
                  class="text-blue-600 hover:underline">Sentences Dictionary</nuxt-link></li>
              <li>• Check out our <nuxt-link to="/qa-guide" class="text-blue-600 hover:underline">QA Guide</nuxt-link>
                for non-technical users</li>
            </ul>
          </div>

          <div class="bg-white p-6 rounded-lg">
            <h3 class="font-semibold mb-3">For Developers</h3>
            <p class="text-sm text-gray-600 mb-4">
              Need technical details, advanced configuration, or development information?
            </p>
            <a href="https://github.com/TestFlowKit/testflowkit" target="_blank"
              class="bg-gray-800 text-white px-6 py-3 rounded-lg hover:bg-gray-900 inline-flex items-center">
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
                <path
                  d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
              </svg>
              GitHub Repository
            </a>
          </div>
        </div>

        <div class="bg-white p-4 rounded-lg">
          <p class="text-sm text-gray-600">
            <strong>Need help?</strong> Check out our <nuxt-link to="/qa-guide"
              class="text-blue-600 hover:underline font-medium">QA Guide</nuxt-link>
            for non-technical users or visit the
            <a href="https://testflowkit.dreamsfollowers.me/" target="_blank"
              class="text-blue-600 hover:underline font-medium">official documentation</a>.
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import DownloadSection from '../components/DownloadSection.vue';
</script>
