<template>
  <div>
    <!-- Hero Section -->
    <section class="hero-gradient relative py-24 lg:py-32 overflow-hidden">
      <!-- Hero Content -->
      <div class="relative z-10 text-center animate-fade-in-up">
        <div class="max-w-4xl mx-auto px-4">
          <!-- Main Heading -->
          <h1 class="heading-1 mb-6">
            Where Your Tests 
            <span class="gradient-text">Tell a Story</span>
          </h1>
          
          <!-- Subtitle -->
          <p class="body-large mb-8 max-w-2xl mx-auto text-gray-700">
            TestFlowKit is a powerful, open-source web test automation framework built in Go. 
            Write tests in plain English using Gherkin syntax and watch your applications come to life through automated testing.
          </p>
          
          <!-- CTA Buttons -->
          <div class="flex flex-col sm:flex-row items-center justify-center gap-4 mb-12">
            <nuxt-link to="/qa-guide" class="btn btn-primary text-lg px-8 py-4">
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              🧪 For QA Teams
            </nuxt-link>
            <nuxt-link to="/get-started" class="btn btn-secondary text-lg px-8 py-4">
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              Get Started Now
            </nuxt-link>
          </div>
          
          <!-- Quick Feature Highlights -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-3xl mx-auto">
            <div class="text-center p-4 bg-white/50 backdrop-blur-sm rounded-xl border border-white/20">
              <div class="text-2xl mb-2">📝</div>
              <h3 class="font-semibold text-gray-800 mb-1">Plain English</h3>
              <p class="text-sm text-gray-600">No coding required</p>
            </div>
            <div class="text-center p-4 bg-white/50 backdrop-blur-sm rounded-xl border border-white/20">
              <div class="text-2xl mb-2">🎯</div>
              <h3 class="font-semibold text-gray-800 mb-1">Visual Examples</h3>
              <p class="text-sm text-gray-600">Real-world scenarios</p>
            </div>
            <div class="text-center p-4 bg-white/50 backdrop-blur-sm rounded-xl border border-white/20">
              <div class="text-2xl mb-2">🚀</div>
              <h3 class="font-semibold text-gray-800 mb-1">Quick Start</h3>
              <p class="text-sm text-gray-600">Up & running in minutes</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- QA Teams Section -->
    <section class="py-20 bg-gradient-to-br from-teal-50 to-green-50 relative">
      <div class="container mx-auto px-4">
        <div class="text-center mb-12 animate-fade-in-up">
          <h2 class="heading-2 mb-6 text-gray-900">New to Test Automation?</h2>
          <p class="body-large text-gray-700 max-w-2xl mx-auto">
            Our QA Guide is designed specifically for non-technical QA professionals. 
            Start testing without any programming knowledge.
          </p>
        </div>

        <div class="grid md:grid-cols-3 gap-8 mb-12">
          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.1s">
            <div class="w-16 h-16 mx-auto mb-6 p-3 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-4">Write in Plain English</h3>
            <p class="body-normal text-gray-600">Describe your tests using natural language - no programming required</p>
          </div>
          
          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.2s">
            <div class="w-16 h-16 mx-auto mb-6 p-3 bg-gradient-to-br from-orange-500 to-red-500 rounded-xl flex items-center justify-center">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-4">Visual Examples</h3>
            <p class="body-normal text-gray-600">See real examples of how to test common scenarios</p>
          </div>
          
          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.3s">
            <div class="w-16 h-16 mx-auto mb-6 p-3 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-4">Quick Start</h3>
            <p class="body-normal text-gray-600">Get up and running with automated testing in minutes</p>
          </div>
        </div>

        <div class="text-center">
          <nuxt-link to="/qa-guide" class="btn btn-primary text-lg px-8 py-4 inline-flex items-center">
            <span>Start with QA Guide</span>
            <svg class="ml-2 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
            </svg>
          </nuxt-link>
        </div>
      </div>
    </section>

    <!-- Core Features Section -->
    <section class="py-20 bg-white">
      <div class="container mx-auto px-4">
        <div class="text-center mb-16 animate-fade-in-up">
          <h2 class="heading-2 mb-6">Powerful Features for Modern Testing</h2>
          <p class="body-large text-gray-700 max-w-3xl mx-auto">
            TestFlowKit combines the simplicity of natural language testing with the power of Go's performance, 
            delivering a comprehensive solution for web application testing.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.1s">
            <div class="w-12 h-12 mx-auto mb-4 p-2 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-3">Gherkin Syntax</h3>
            <p class="body-normal text-gray-600">Write tests using clear, readable BDD syntax with natural language that's easy to understand and maintain.</p>
          </div>

          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.2s">
            <div class="w-12 h-12 mx-auto mb-4 p-2 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-3">Multi-Environment Support</h3>
            <p class="body-normal text-gray-600">Configure and switch between different environments (local, staging, production) with ease.</p>
          </div>

          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.3s">
            <div class="w-12 h-12 mx-auto mb-4 p-2 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-3">Frontend & Backend Testing</h3>
            <p class="body-normal text-gray-600">Comprehensive web UI automation with smart element detection and full REST API testing capabilities.</p>
          </div>

          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.4s">
            <div class="w-12 h-12 mx-auto mb-4 p-2 bg-gradient-to-br from-amber-500 to-orange-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a1 1 0 01-1-1V9a1 1 0 011-1h1a2 2 0 100-4H4a1 1 0 01-1-1V4a1 1 0 011-1h3a1 1 0 001-1V1a2 2 0 114 0v1z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-3">Macro System</h3>
            <p class="body-normal text-gray-600">Reusable test scenarios to reduce code duplication and improve maintainability.</p>
          </div>

          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.5s">
            <div class="w-12 h-12 mx-auto mb-4 p-2 bg-gradient-to-br from-yellow-500 to-yellow-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-3">Parallel Execution</h3>
            <p class="body-normal text-gray-600">Run tests concurrently for faster execution, leveraging the power of Go for incredible performance.</p>
          </div>

          <div class="feature-card animate-fade-in-scale" style="animation-delay: 0.6s">
            <div class="w-12 h-12 mx-auto mb-4 p-2 bg-gradient-to-br from-pink-500 to-rose-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
            <h3 class="heading-4 mb-3">Rich Reporting</h3>
            <p class="body-normal text-gray-600">HTML and JSON report formats with detailed test results and automatic screenshot capture on failures.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Why Choose TestFlowKit Section -->
    <section class="py-20 bg-gray-50">
      <div class="container mx-auto px-4">
        <div class="text-center mb-16">
          <h2 class="heading-2 mb-6">Why Choose TestFlowKit?</h2>
          <p class="body-large text-gray-700 max-w-3xl mx-auto">
            Built with modern technologies and best practices, TestFlowKit delivers enterprise-grade 
            testing capabilities with developer-friendly simplicity.
          </p>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-12 max-w-6xl mx-auto">
          <div class="card p-8 animate-fade-in-scale" style="animation-delay: 0.1s">
            <div class="flex items-start space-x-4">
              <div class="w-12 h-12 p-2 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </div>
              <div>
                <h3 class="heading-4 mb-3">🚀 Performance</h3>
                <p class="body-normal text-gray-600">Built with Go for blazing fast execution, saving valuable time in your testing workflow.</p>
              </div>
            </div>
          </div>
          
          <div class="card p-8 animate-fade-in-scale" style="animation-delay: 0.2s">
            <div class="flex items-start space-x-4">
              <div class="w-12 h-12 p-2 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a1 1 0 01-1-1V9a1 1 0 011-1h1a2 2 0 100-4H4a1 1 0 01-1-1V4a1 1 0 011-1h3a1 1 0 001-1V1a2 2 0 114 0v1z" />
                </svg>
              </div>
              <div>
                <h3 class="heading-4 mb-3">🔧 Extensibility</h3>
                <p class="body-normal text-gray-600">Plugin-like step definition system with validation and documentation for easy customization.</p>
              </div>
            </div>
          </div>
          
          <div class="card p-8 animate-fade-in-scale" style="animation-delay: 0.3s">
            <div class="flex items-start space-x-4">
              <div class="w-12 h-12 p-2 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <div>
                <h3 class="heading-4 mb-3">📊 Comprehensive</h3>
                <p class="body-normal text-gray-600">Support for both frontend UI testing and backend API testing in a single framework.</p>
              </div>
            </div>
          </div>
          
          <div class="card p-8 animate-fade-in-scale" style="animation-delay: 0.4s">
            <div class="flex items-start space-x-4">
              <div class="w-12 h-12 p-2 bg-gradient-to-br from-orange-500 to-red-500 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <div>
                <h3 class="heading-4 mb-3">🛠️ Developer-Friendly</h3>
                <p class="body-normal text-gray-600">Clear error messages, detailed logging, and excellent debugging capabilities.</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action Section -->
    <section class="py-20 bg-gradient-to-r from-orange-500 to-orange-600 text-white relative overflow-hidden">
      <div class="absolute inset-0 bg-black/10"></div>
      <div class="container mx-auto px-4 text-center relative z-10">
        <div class="max-w-3xl mx-auto">
          <h2 class="heading-2 mb-6 text-white">Ready to Transform Your Testing?</h2>
          <p class="body-large mb-8 text-orange-100">
            Join thousands of teams who have made testing easier, faster, and more reliable with TestFlowKit. 
            Start your journey today.
          </p>
          <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
            <nuxt-link to="/get-started" class="btn bg-white text-orange-600 hover:bg-gray-50 text-lg px-8 py-4">
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              Get Started Free
            </nuxt-link>
            <a href="https://github.com/TestFlowKit/testflowkit" target="_blank" class="btn btn-ghost text-white border-white hover:bg-white/10 text-lg px-8 py-4">
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
              View on GitHub
            </a>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// SEO and Meta
useHead({
  title: 'TestFlowKit - Powerful Web Test Automation Framework',
  meta: [
    {
      name: 'description',
      content: 'TestFlowKit is a powerful, open-source web test automation framework built in Go. Write tests in plain English using Gherkin syntax for both frontend and backend testing.'
    },
    {
      property: 'og:title',
      content: 'TestFlowKit - Where Your Tests Tell a Story'
    },
    {
      property: 'og:description', 
      content: 'Write automated tests in plain English with TestFlowKit. No coding required - perfect for QA teams.'
    },
    {
      property: 'og:type',
      content: 'website'
    }
  ]
})
</script>

<style scoped>
/* Animation delays for staggered effects */
@keyframes gradient-shift {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

.hero-gradient {
  background-size: 200% 200%;
  animation: gradient-shift 8s ease infinite;
}

/* Ensure proper stacking context */
.relative {
  position: relative;
}

/* Custom animation timing */
.animate-fade-in-up {
  animation-fill-mode: both;
}

.animate-fade-in-scale {
  animation-fill-mode: both;
}
</style>
