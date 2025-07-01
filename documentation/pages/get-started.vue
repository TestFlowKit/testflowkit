<template>
  <div class="space-y-8">
    <!-- Hero Section -->
    <section class="text-center hero-gradient p-8 rounded-xl relative overflow-hidden">
      <div class="relative z-10">
        <h1 class="heading-1 mb-4 text-gray-900 text-white">Getting Started with TestFlowKit</h1>
        <p class="body-large mb-6 text-gray-700 max-w-2xl mx-auto text-white">
          Set up TestFlowKit and write your first automated test in minutes
        </p>
      </div>
    </section>

    <!-- What You Need Section -->
    <section class="card p-8">
      <h2 class="heading-2 mb-6 text-white">What You Need to Get Started</h2>

      <div class="alert alert-success mb-6">
        <h3 class="heading-4 mb-4 text-green-800">For QA Teams & Non-Technical Users</h3>
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

      <div class="bg-orange-50 border border-orange-200 p-6 rounded-lg">
        <h3 class="heading-4 mb-4 text-orange-800">For Developers & Technical Users</h3>
        <p class="text-orange-700 mb-4">
          If you need technical setup instructions, configuration details, or advanced features, please visit our GitHub
          repository.
        </p>
        <a href="https://github.com/TestFlowKit/testflowkit" target="_blank"
          class="btn btn-primary inline-flex items-center">
          <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
            <path
              d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
          </svg>
          View on GitHub
        </a>
      </div>
    </section>

    <!-- Quick Setup Section -->
    <section class="card p-8">
      <h2 class="heading-2 mb-6 text-white">Quick Setup</h2>

      <div class="space-y-6">
        <div class="border-l-4 border-orange-400 pl-6 bg-orange-50 p-4 rounded-r-lg">
          <h3 class="heading-4 mb-3">Step 1: Download TestFlowKit</h3>
          <p class="body-normal text-gray-600 mb-4">
            Download TestFlowKit for your operating system and install it on your machine.
          </p>
          <div class="card-gradient p-6 rounded-lg">
            <h4 class="font-semibold mb-4 text-gray-800">Choose your platform and architecture:</h4>
            <DownloadSection />
          </div>
        </div>

        <div class="border-l-4 border-teal-400 pl-6 bg-teal-50 p-4 rounded-r-lg">
          <h3 class="heading-4 mb-3">Step 2: Ask Your Development Team</h3>
          <p class="body-normal text-gray-600 mb-4">
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
    <section class="bg-gradient-to-r from-orange-500 to-orange-600 text-white p-8 rounded-xl">
      <h2 class="heading-2 mb-6 text-center text-white">What's Next?</h2>

      <div class="text-center space-y-6">
        <div class="grid md:grid-cols-2 gap-6">
          <div class="bg-white/10 backdrop-blur-sm border border-white/20 p-6 rounded-lg">
            <h3 class="font-semibold mb-3 text-white">For QA Teams</h3>
            <ul class="space-y-2 text-sm text-orange-100 text-left">
              <li>• Learn to write tests in the <nuxt-link to="/quick-start" class="text-white hover:text-orange-200 underline">Quick
                  Start Tutorial</nuxt-link></li>
              <li>• Explore test patterns in the <nuxt-link to="/sentences"
                  class="text-white hover:text-orange-200 underline">Sentences Dictionary</nuxt-link></li>
              <li>• Check out our <nuxt-link to="/qa-guide" class="text-white hover:text-orange-200 underline">QA Guide</nuxt-link>
                for non-technical users</li>
            </ul>
          </div>

          <div class="bg-white/10 backdrop-blur-sm border border-white/20 p-6 rounded-lg">
            <h3 class="font-semibold mb-3 text-white">For Developers</h3>
            <p class="text-sm text-orange-100 mb-4">
              Need technical details, advanced configuration, or development information?
            </p>
            <a href="https://github.com/TestFlowKit/testflowkit" target="_blank"
              class="btn bg-white text-orange-600 hover:bg-gray-50 inline-flex items-center">
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
                <path
                  d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
              </svg>
              GitHub Repository
            </a>
          </div>
        </div>

        <div class="bg-white/10 backdrop-blur-sm border border-white/20 p-4 rounded-lg">
          <p class="text-sm text-orange-100">
            <strong>Need help?</strong> Check out our <nuxt-link to="/qa-guide"
              class="text-white hover:text-orange-200 underline font-medium">QA Guide</nuxt-link>
            for non-technical users or visit the
            <a href="https://testflowkit.dreamsfollowers.me/" target="_blank"
              class="text-white hover:text-orange-200 underline font-medium">official documentation</a>.
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// SEO and Meta
useHead({
  title: 'Getting Started - TestFlowKit Documentation',
  meta: [
    {
      name: 'description',
      content: 'Get started with TestFlowKit. Set up your test automation environment and write your first automated test in minutes.'
    }
  ]
})

import DownloadSection from '../components/DownloadSection.vue';
</script>
