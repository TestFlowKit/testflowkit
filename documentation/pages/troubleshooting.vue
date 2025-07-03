<template>
    <div class="space-y-8">
        <section class="text-center bg-gradient-to-r from-orange-50 to-red-50 p-8 rounded-lg">
            <h1 class="text-4xl font-bold mb-4 text-gray-800">FAQ & Troubleshooting</h1>
            <p class="text-xl text-gray-600 mb-6">
                Common issues and solutions for TestFlowKit installation and usage
            </p>
            <div class="flex justify-center space-x-4">
                <button @click="scrollToSection('installation')"
                    class="bg-orange-600 text-white px-6 py-3 rounded-lg hover:bg-orange-700">
                    Installation Issues
                </button>
                <button @click="scrollToSection('usage')"
                    class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700">
                    Usage Problems
                </button>
                <button @click="scrollToSection('faq')"
                    class="bg-green-600 text-white px-6 py-3 rounded-lg hover:bg-green-700">
                    FAQ
                </button>
            </div>
        </section>

        <!-- Installation Issues Section -->
        <section id="installation" class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Installation Issues</h2>

            <!-- macOS Issues -->
            <div class="mb-8">
                <h3 class="text-2xl font-semibold mb-4 text-gray-800 flex items-center">
                    <span class="text-2xl mr-2">🍎</span>
                    macOS Installation Issues
                </h3>

                <div class="bg-red-50 border-l-4 border-red-400 p-6 rounded-lg mb-6">
                    <h4 class="text-lg font-semibold text-red-800 mb-3">"tkit cannot be opened because it is from an
                        unidentified developer"</h4>
                    <p class="text-red-700 mb-4">
                        This is a common macOS security feature that blocks applications from unidentified developers.
                        Here's how to fix it:
                    </p>

                    <div class="bg-white p-4 rounded-lg">
                        <h5 class="font-semibold mb-2">Solution 1: Remove Quarantine Attribute (Recommended)</h5>
                        <div class="bg-gray-100 p-3 rounded font-mono text-sm mb-3">
                            xattr -d com.apple.quarantine tkit
                        </div>
                        <p class="text-sm text-gray-600 mb-3">
                            Run this command in Terminal from the directory where you downloaded TestFlowKit. This
                            removes the quarantine attribute that macOS adds to downloaded files.
                        </p>

                        <h5 class="font-semibold mb-2">Solution 2: Allow in System Preferences</h5>
                        <ol class="text-sm text-gray-600 list-decimal list-inside space-y-1">
                            <li>Go to System Preferences → Security & Privacy</li>
                            <li>Click the "General" tab</li>
                            <li>Look for a message about "tkit was blocked from opening"</li>
                            <li>Click "Allow Anyway"</li>
                            <li>Try running tkit again</li>
                        </ol>
                    </div>
                </div>

                <div class="bg-yellow-50 border-l-4 border-yellow-400 p-6 rounded-lg">
                    <h4 class="text-lg font-semibold text-yellow-800 mb-3">Permission Denied Error</h4>
                    <p class="text-yellow-700 mb-3">
                        If you get a "permission denied" error when trying to run TestFlowKit:
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <div class="bg-gray-100 p-3 rounded font-mono text-sm mb-3">
                            chmod +x tkit
                        </div>
                        <p class="text-sm text-gray-600">
                            This command makes the TestFlowKit executable file executable.
                        </p>
                    </div>
                </div>
            </div>

            <!-- Windows Issues -->
            <div class="mb-8">
                <h3 class="text-2xl font-semibold mb-4 text-gray-800 flex items-center">
                    <span class="text-2xl mr-2">🪟</span>
                    Windows Installation Issues
                </h3>

                <div class="bg-red-50 border-l-4 border-red-400 p-6 rounded-lg mb-6">
                    <h4 class="text-lg font-semibold text-red-800 mb-3">Antivirus Software Blocks TestFlowKit</h4>
                    <p class="text-red-700 mb-4">
                        Some antivirus software may flag TestFlowKit as suspicious because it's a command-line tool that
                        can interact with browsers. This is a false positive.
                    </p>

                    <div class="bg-white p-4 rounded-lg">
                        <h5 class="font-semibold mb-2">How to Fix:</h5>
                        <ol class="text-sm text-gray-600 list-decimal list-inside space-y-2">
                            <li><strong>Add to Exclusions:</strong> Add the TestFlowKit folder to your antivirus
                                exclusions/whitelist</li>
                            <li><strong>Windows Defender:</strong> Go to Windows Security → Virus & threat protection →
                                Manage settings → Add or remove exclusions</li>
                            <li><strong>Other Antivirus:</strong> Look for "Exclusions", "Whitelist", or "Trusted Files"
                                in your antivirus settings</li>
                            <li><strong>Download from Official Source:</strong> Always download TestFlowKit from the
                                official GitHub releases</li>
                        </ol>
                    </div>
                </div>

                <div class="bg-yellow-50 border-l-4 border-yellow-400 p-6 rounded-lg">
                    <h4 class="text-lg font-semibold text-yellow-800 mb-3">"tkit is not recognized as an internal or
                        external command"</h4>
                    <p class="text-yellow-700 mb-3">
                        This error occurs when Windows can't find the TestFlowKit executable in the system PATH.
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <h5 class="font-semibold mb-2">Solutions:</h5>
                        <ol class="text-sm text-gray-600 list-decimal list-inside space-y-2">
                            <li><strong>Run from Directory:</strong> Navigate to the folder containing tkit.exe and run
                                it from there</li>
                            <li><strong>Add to PATH:</strong> Add the TestFlowKit directory to your system PATH
                                environment variable</li>
                            <li><strong>Use Full Path:</strong> Run with the full path to the executable: <code
                                    class="bg-gray-100 px-1 rounded">C:\path\to\tkit.exe</code></li>
                        </ol>
                    </div>
                </div>
            </div>

            <!-- Linux Issues -->
            <div>
                <h3 class="text-2xl font-semibold mb-4 text-gray-800 flex items-center">
                    <span class="text-2xl mr-2">🐧</span>
                    Linux Installation Issues
                </h3>

                <div class="bg-yellow-50 border-l-4 border-yellow-400 p-6 rounded-lg">
                    <h4 class="text-lg font-semibold text-yellow-800 mb-3">Permission Denied or Executable Issues</h4>
                    <p class="text-yellow-700 mb-3">
                        Common permission issues on Linux systems:
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <div class="bg-gray-100 p-3 rounded font-mono text-sm mb-3">
                            chmod +x tkit
                        </div>
                        <p class="text-sm text-gray-600 mb-3">
                            Make the file executable if you get permission denied errors.
                        </p>

                        <h5 class="font-semibold mb-2">Missing Dependencies:</h5>
                        <p class="text-sm text-gray-600 mb-2">
                            TestFlowKit normally installs Chromium automatically. If a version already exists on your
                            system, TestFlowKit will use it. If you need to install Chromium manually:
                        </p>
                        <div class="bg-gray-100 p-3 rounded font-mono text-sm">
                            # Ubuntu/Debian<br>
                            sudo apt update && sudo apt install chromium-browser<br><br>
                            # CentOS/RHEL/Fedora<br>
                            sudo yum install chromium # or sudo dnf install chromium
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Usage Problems Section -->
        <section id="usage" class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Usage Problems</h2>

            <div class="grid md:grid-cols-2 gap-6">
                <div class="bg-blue-50 p-6 rounded-lg">
                    <h3 class="text-xl font-semibold mb-3 text-blue-800">Browser Won't Start</h3>
                    <p class="text-blue-700 mb-3">
                        TestFlowKit can't launch the browser for automation.
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <h4 class="font-semibold mb-2">Possible Causes:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Chromium installation failed</li>
                            <li>Browser already running in headless mode</li>
                            <li>Insufficient system resources</li>
                            <li>Firewall blocking browser launch</li>
                        </ul>
                        <h4 class="font-semibold mb-2 mt-3">Solutions:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Check if Chromium is properly installed (TestFlowKit installs it automatically)</li>
                            <li>Close other browser instances</li>
                            <li>Check system memory and CPU usage</li>
                            <li>Try running with <code class="bg-gray-100 px-1 rounded">--headless=false</code></li>
                        </ul>
                    </div>
                </div>

                <div class="bg-green-50 p-6 rounded-lg">
                    <h3 class="text-xl font-semibold mb-3 text-green-800">Tests Fail with Element Not Found</h3>
                    <p class="text-green-700 mb-3">
                        TestFlowKit can't find elements on the page.
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <h4 class="font-semibold mb-2">Common Issues:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Page not fully loaded</li>
                            <li>Element selectors are incorrect</li>
                            <li>Dynamic content not rendered</li>
                            <li>Element is in an iframe</li>
                        </ul>
                        <h4 class="font-semibold mb-2 mt-3">Solutions:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Increase timeout in configuration</li>
                            <li>Verify element selectors in browser dev tools</li>
                            <li>Add wait conditions for dynamic content</li>
                            <li>Use XPath selectors for complex elements</li>
                        </ul>
                    </div>
                </div>

                <div class="bg-purple-50 p-6 rounded-lg">
                    <h3 class="text-xl font-semibold mb-3 text-purple-800">Configuration File Issues</h3>
                    <p class="text-purple-700 mb-3">
                        Problems with config.yml or configuration loading.
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <h4 class="font-semibold mb-2">Common Problems:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>YAML syntax errors</li>
                            <li>Missing required fields</li>
                            <li>Invalid URLs or selectors</li>
                            <li>File not found</li>
                        </ul>
                        <h4 class="font-semibold mb-2 mt-3">Solutions:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Use a YAML validator to check syntax</li>
                            <li>Run <code class="bg-gray-100 px-1 rounded">./tkit validate</code> to check config</li>
                            <li>Verify all URLs are accessible</li>
                            <li>Check file permissions and location</li>
                        </ul>
                    </div>
                </div>

                <div class="bg-orange-50 p-6 rounded-lg">
                    <h3 class="text-xl font-semibold mb-3 text-orange-800">Performance Issues</h3>
                    <p class="text-orange-700 mb-3">
                        Tests are running slowly or timing out.
                    </p>
                    <div class="bg-white p-4 rounded-lg">
                        <h4 class="font-semibold mb-2">Performance Factors:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Network latency</li>
                            <li>Page load times</li>
                            <li>System resources</li>
                            <li>Browser performance</li>
                        </ul>
                        <h4 class="font-semibold mb-2 mt-3">Optimizations:</h4>
                        <ul class="text-sm text-gray-600 list-disc list-inside space-y-1">
                            <li>Use headless mode for faster execution</li>
                            <li>Increase concurrency for parallel tests</li>
                            <li>Optimize element selectors</li>
                            <li>Close unnecessary browser tabs</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>

        <!-- FAQ Section -->
        <section id="faq" class="bg-white p-8 rounded-lg shadow-md">
            <h2 class="text-3xl font-bold mb-6 text-gray-800">Frequently Asked Questions</h2>

            <div class="space-y-6">
                <div class="border border-gray-200 rounded-lg">
                    <button @click="toggleFAQ(1)"
                        class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-lg">
                        <div class="flex justify-between items-center">
                            <span class="font-semibold text-gray-800">What browsers does TestFlowKit support?</span>
                            <svg :class="['w-5 h-5 transform transition-transform', { 'rotate-180': openFAQ === 1 }]"
                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </div>
                    </button>
                    <div v-show="openFAQ === 1" class="px-6 py-4 bg-white border-t border-gray-200">
                        <p class="text-gray-600">
                            TestFlowKit supports Chrome and Chromium-based browsers. The framework automatically
                            downloads and
                            manages the appropriate browser version for your system.
                        </p>
                    </div>
                </div>

                <div class="border border-gray-200 rounded-lg">
                    <button @click="toggleFAQ(2)"
                        class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-lg">
                        <div class="flex justify-between items-center">
                            <span class="font-semibold text-gray-800">Can I run tests in parallel?</span>
                            <svg :class="['w-5 h-5 transform transition-transform', { 'rotate-180': openFAQ === 2 }]"
                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </div>
                    </button>
                    <div v-show="openFAQ === 2" class="px-6 py-4 bg-white border-t border-gray-200">
                        <p class="text-gray-600 mb-3">
                            Yes! TestFlowKit supports parallel test execution. You can configure the concurrency level
                            in your config.yml file or use the --concurrency flag when running tests.
                        </p>
                        <div class="bg-gray-100 p-3 rounded font-mono text-sm">
                            ./tkit run --concurrency 4
                        </div>
                    </div>
                </div>

                <div class="border border-gray-200 rounded-lg">
                    <button @click="toggleFAQ(3)"
                        class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-lg">
                        <div class="flex justify-between items-center">
                            <span class="font-semibold text-gray-800">How do I handle dynamic content and AJAX?</span>
                            <svg :class="['w-5 h-5 transform transition-transform', { 'rotate-180': openFAQ === 3 }]"
                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </div>
                    </button>
                    <div v-show="openFAQ === 3" class="px-6 py-4 bg-white border-t border-gray-200">
                        <p class="text-gray-600 mb-3">
                            TestFlowKit automatically waits for elements to be present and visible. For dynamic content,
                            you can:
                        </p>
                        <ul class="text-gray-600 list-disc list-inside space-y-1">
                            <li>Increase the default timeout for slow-loading content</li>
                            <li>Use XPath selectors that can wait for specific text or attributes</li>
                            <li>Add custom wait steps in your feature files</li>
                        </ul>
                    </div>
                </div>

                <div class="border border-gray-200 rounded-lg">
                    <button @click="toggleFAQ(4)"
                        class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-lg">
                        <div class="flex justify-between items-center">
                            <span class="font-semibold text-gray-800">Can I integrate TestFlowKit with CI/CD
                                pipelines?</span>
                            <svg :class="['w-5 h-5 transform transition-transform', { 'rotate-180': openFAQ === 4 }]"
                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </div>
                    </button>
                    <div v-show="openFAQ === 4" class="px-6 py-4 bg-white border-t border-gray-200">
                        <p class="text-gray-600 mb-3">
                            Absolutely! TestFlowKit is designed for CI/CD integration. It supports:
                        </p>
                        <ul class="text-gray-600 list-disc list-inside space-y-1">
                            <li>Headless mode for server environments</li>
                            <li>JSON and HTML report generation</li>
                            <li>Exit codes for pass/fail detection</li>
                            <li>Environment-specific configurations</li>
                            <li>Docker containerization support</li>
                        </ul>
                    </div>
                </div>

                <div class="border border-gray-200 rounded-lg">
                    <button @click="toggleFAQ(5)"
                        class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-lg">
                        <div class="flex justify-between items-center">
                            <span class="font-semibold text-gray-800">How do I debug failing tests?</span>
                            <svg :class="['w-5 h-5 transform transition-transform', { 'rotate-180': openFAQ === 5 }]"
                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </div>
                    </button>
                    <div v-show="openFAQ === 5" class="px-6 py-4 bg-white border-t border-gray-200">
                        <p class="text-gray-600 mb-3">
                            TestFlowKit provides several debugging features:
                        </p>
                        <ul class="text-gray-600 list-disc list-inside space-y-1">
                            <li>Automatic screenshots on test failures</li>
                            <li>Detailed HTML reports with step-by-step execution</li>
                            <li>Slow motion mode for visual debugging</li>
                            <li>Browser dev tools integration</li>
                        </ul>
                    </div>
                </div>

                <div class="border border-gray-200 rounded-lg">
                    <button @click="toggleFAQ(6)"
                        class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-lg">
                        <div class="flex justify-between items-center">
                            <span class="font-semibold text-gray-800">Is TestFlowKit free to use?</span>
                            <svg :class="['w-5 h-5 transform transition-transform', { 'rotate-180': openFAQ === 6 }]"
                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </div>
                    </button>
                    <div v-show="openFAQ === 6" class="px-6 py-4 bg-white border-t border-gray-200">
                        <p class="text-gray-600">
                            Yes! TestFlowKit is completely free and open-source under the MIT license. You can use it
                            for personal projects, commercial applications, and contribute to its development. The
                            source code is available on GitHub for transparency and community contributions.
                        </p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Getting Help Section -->
        <section class="bg-gradient-to-r from-blue-50 to-indigo-50 p-8 rounded-lg">
            <h2 class="text-3xl font-bold mb-6 text-gray-800 text-center">Still Need Help?</h2>
            <div class="grid md:grid-cols-3 gap-6">
                <div class="text-center p-6 bg-white rounded-lg shadow-md">
                    <div class="text-4xl mb-4">📚</div>
                    <h3 class="text-xl font-semibold mb-3">Documentation</h3>
                    <p class="text-gray-600 mb-4">
                        Check our comprehensive documentation for detailed guides and examples.
                    </p>
                    <nuxt-link to="/get-started" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
                        Get Started
                    </nuxt-link>
                </div>

                <div class="text-center p-6 bg-white rounded-lg shadow-md">
                    <div class="text-4xl mb-4">🐛</div>
                    <h3 class="text-xl font-semibold mb-3">GitHub Issues</h3>
                    <p class="text-gray-600 mb-4">
                        Found a bug? Report it on GitHub with detailed information.
                    </p>
                    <a href="https://github.com/TestFlowKit/testflowkit/issues" target="_blank"
                        class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700">
                        Report Issue
                    </a>
                </div>

                <div class="text-center p-6 bg-white rounded-lg shadow-md">
                    <div class="text-4xl mb-4">💬</div>
                    <h3 class="text-xl font-semibold mb-3">Community</h3>
                    <p class="text-gray-600 mb-4">
                        Join our community discussions and get help from other users.
                    </p>
                    <a href="https://github.com/TestFlowKit/testflowkit/discussions" target="_blank"
                        class="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700">
                        Join Discussion
                    </a>
                </div>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const openFAQ = ref<number | null>(null);

const scrollToSection = (sectionId: string) => {
    const element = document.getElementById(sectionId);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    }
};

const toggleFAQ = (faqNumber: number) => {
    openFAQ.value = openFAQ.value === faqNumber ? null : faqNumber;
};
</script>
