<template>
  <footer class="bg-gray-900 text-gray-300 relative overflow-hidden">
    <!-- Decorative Elements -->
    <div class="absolute inset-0 bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900"></div>
    <div class="absolute top-0 left-0 w-full h-px bg-gradient-to-r from-transparent via-orange-500 to-transparent opacity-20"></div>
    
    <div class="relative z-10">
      <!-- Main Footer Content -->
      <div class="container mx-auto px-4 py-16">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          
          <!-- Brand Section -->
          <div class="lg:col-span-1">
            <div class="flex items-center space-x-3 mb-6">
              <div class="w-10 h-10 rounded-lg bg-gradient-to-br from-orange-500 to-orange-600 flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
                </svg>
              </div>
              <span class="text-xl font-bold text-white">TestFlowKit</span>
            </div>
            <p class="text-gray-400 text-sm leading-relaxed mb-6">
              Powerful web test automation framework built in Go. 
              Write tests in plain English and watch your applications come to life.
            </p>
            <div class="flex space-x-4">
              <a 
                href="https://github.com/TestFlowKit/testflowkit" 
                target="_blank" 
                rel="noopener noreferrer"
                class="w-10 h-10 bg-gray-800 hover:bg-orange-600 rounded-lg flex items-center justify-center transition-colors group"
                title="GitHub"
              >
                <svg class="w-5 h-5 text-gray-300 group-hover:text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                </svg>
              </a>
            </div>
          </div>

          <!-- Documentation Links -->
          <div>
            <h3 class="text-white font-semibold mb-4">Documentation</h3>
            <ul class="space-y-3 text-sm">
              <li>
                <nuxt-link to="/get-started" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Get Started
                </nuxt-link>
              </li>
              <li>
                <nuxt-link to="/quick-start" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Quick Start
                </nuxt-link>
              </li>
              <li>
                <nuxt-link to="/qa-guide" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  QA Guide
                </nuxt-link>
              </li>
              <li>
                <nuxt-link to="/configuration" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Configuration
                </nuxt-link>
              </li>
            </ul>
          </div>

          <!-- Reference Links -->
          <div>
            <h3 class="text-white font-semibold mb-4">Reference</h3>
            <ul class="space-y-3 text-sm">
              <li>
                <nuxt-link to="/sentences" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Gherkin Steps
                </nuxt-link>
              </li>
              <li>
                <a href="https://github.com/TestFlowKit/testflowkit/releases" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Releases
                </a>
              </li>
              <li>
                <a href="https://github.com/TestFlowKit/testflowkit/issues" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Issues & Support
                </a>
              </li>
              <li>
                <a href="https://github.com/TestFlowKit/testflowkit/blob/main/CONTRIBUTING.md" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Contributing
                </a>
              </li>
            </ul>
          </div>

          <!-- Community & Updates -->
          <div>
            <h3 class="text-white font-semibold mb-4">Community</h3>
            <ul class="space-y-3 text-sm">
              <li>
                <a href="https://github.com/TestFlowKit/testflowkit/discussions" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Discussions
                </a>
              </li>
              <li>
                <a href="https://github.com/TestFlowKit/testflowkit/blob/main/CHANGELOG.md" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  Changelog
                </a>
              </li>
              <li>
                <a href="https://github.com/TestFlowKit/testflowkit/blob/main/LICENSE" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors flex items-center group">
                  <svg class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                  MIT License
                </a>
              </li>
            </ul>
            
            <!-- Newsletter/Updates -->
            <div class="mt-6 p-4 bg-gray-800 rounded-lg border border-gray-700">
              <h4 class="text-white font-medium mb-2 text-sm">Stay Updated</h4>
              <p class="text-gray-400 text-xs mb-3">Get notified about new releases and features</p>
              <a 
                href="https://github.com/TestFlowKit/testflowkit/subscription" 
                target="_blank"
                class="inline-flex items-center text-xs bg-orange-600 hover:bg-orange-700 text-white px-3 py-2 rounded-md transition-colors"
              >
                <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM11 17H6l5 5v-5z"/>
                </svg>
                Watch on GitHub
              </a>
            </div>
          </div>
        </div>
      </div>

      <!-- Bottom Bar -->
      <div class="border-t border-gray-800">
        <div class="container mx-auto px-4 py-6">
          <div class="flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0">
            <div class="flex items-center space-x-6">
                             <p class="text-gray-400 text-sm">
                 © 2025 TestFlowKit. All rights reserved.
               </p>
              <span class="text-gray-600 text-sm">•</span>
              <p class="text-gray-400 text-sm">
                Built with ❤️ for the testing community
              </p>
            </div>
            
            <div class="flex items-center space-x-6 text-sm">
              <a href="https://github.com/TestFlowKit/testflowkit/blob/main/LICENSE" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors">
                MIT License
              </a>
              <span class="text-gray-600">•</span>
              <a href="https://github.com/TestFlowKit/testflowkit/blob/main/PRIVACY.md" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors">
                Privacy
              </a>
              <span class="text-gray-600">•</span>
              <a href="https://github.com/TestFlowKit/testflowkit/security" target="_blank" class="text-gray-400 hover:text-orange-400 transition-colors">
                Security
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
// Footer component for TestFlowKit documentation
</script>

<style scoped>
.container {
  max-width: var(--container-max-width);
}

/* Hover effects for links */
.group:hover svg {
  transform: translateX(2px);
  transition: transform 0.2s ease;
}

/* Custom scrollbar for any scrollable content */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
}

::-webkit-scrollbar-thumb {
  background: rgba(249, 115, 22, 0.3);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(249, 115, 22, 0.5);
}
</style>