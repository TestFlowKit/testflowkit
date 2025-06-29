<template>
  <section id="prerequisites" class="mb-12">
    <h2 class="heading-2">Prerequisites</h2>
    <p class="mb-4">Before you begin, ensure you have the following:</p>
    <PrerequisitesList />
  </section>

  <section id="installation" class="mb-12">
    <h2 class="heading-2">Installation</h2>
    
    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 class="text-xl font-semibold mb-4">From Source</h3>
      <p class="mb-4">Clone the repository and build from source:</p>
      <code-block language="bash" :code="sourceInstallCommand" />
    </div>

    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 class="text-xl font-semibold mb-4">From Pre-built Binaries</h3>
      <p class="mb-4">Download the latest release for your platform:</p>
      <DownloadSection />
    </div>
  </section>

  <section id="quick-start" class="mb-12">
    <h2 class="heading-2">Quick Start</h2>
    
    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 class="text-xl font-semibold mb-4">1. Initialize Project</h3>
      <p class="mb-4">Initialize a new TestFlowKit project:</p>
      <code-block language="bash" :code="initCommand" />
    </div>

    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 class="text-xl font-semibold mb-4">2. Configure Your Application</h3>
      <p class="mb-4">Edit the generated <code>config.yml</code> file with your application details:</p>
      <code-block language="yaml" :code="configExample" />
    </div>

    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 class="text-xl font-semibold mb-4">3. Write Your First Test</h3>
      <p class="mb-4">Create a feature file at <code>e2e/features/login.feature</code>:</p>
      <code-block language="gherkin" :code="featureExample" />
    </div>

    <div class="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 class="text-xl font-semibold mb-4">4. Run Tests</h3>
      <p class="mb-4">Execute your tests:</p>
      <code-block language="bash" :code="runCommands" />
    </div>
  </section>

  <section class="mb-12">
    <h2 class="heading-2">Next Steps</h2>
    <p>After getting started, explore these resources:</p>
    <ol class="list-decimal list-inside ml-6 space-y-2">
      <li><strong>Check out the <router-link to="/quick-start" class="text-blue-600 hover:underline">Quick Start Guide</router-link> for a hands-on tutorial.</strong></li>
      <li><strong>Refer to the <router-link to="/configuration" class="text-blue-600 hover:underline">Configuration Guide</router-link> for detailed setup instructions.</strong></li>
      <li><strong>Explore the <router-link to="/sentences" class="text-blue-600 hover:underline">Gherkin Sentences Dictionary</router-link> to understand available keywords.</strong></li>
      <li><strong>Visit the <a href="https://testflowkit.dreamsfollowers.me/" target="_blank" class="text-blue-600 hover:underline">official documentation</a> for comprehensive guides.</strong></li>
    </ol>
  </section>
</template>

<script setup lang="ts">
import PrerequisitesList from '../components/PrerequisitesList.vue';
import DownloadSection from '../components/DownloadSection.vue';

const sourceInstallCommand = `# Clone the repository
git clone https://github.com/TestFlowKit/testflowkit.git
cd testflowkit

# Install dependencies
go mod tidy

# Build the application
make build GOOS=linux GOARCH=amd64  # or your target platform`;

const initCommand = `# Initialize a new TestFlowKit project
./testflowkit init`;

const configExample = `active_environment: "local"

settings:
  default_timeout: 10000
  concurrency: 1
  headless: false
  screenshot_on_failure: true
  report_format: "html"
  gherkin_location: "./e2e/features"

environments:
  local:
    frontend_base_url: "http://localhost:3000"
    api_base_url: "http://localhost:8080/api"

frontend:
  elements:
    login_page:
      email_field:
        - "[data-testid='email-input']"
        - "input[name='email']"
      password_field:
        - "[data-testid='password-input']"
        - "input[name='password']"
      login_button:
        - "[data-testid='login-btn']"
        - "button[type='submit']"

  pages:
    login: "/login"
    dashboard: "/dashboard"

backend:
  endpoints:
    get_user:
      method: "GET"
      path: "/users/{id}"
      description: "Retrieve user by ID"`;

const featureExample = `Feature: User Authentication
  As a user
  I want to log into the application
  So that I can access my account

  Scenario: Successful login with valid credentials
    Given the user is on the homepage
    When the user goes to the "login" page
    And the user enters "test@example.com" into the "email" field
    And the user enters "password123" into the "password" field
    And the user clicks on the "login" button
    Then the user should be navigated to "dashboard" page
    And the "welcome_message" should be visible`;

const runCommands = `# Run all tests
./testflowkit run

# Run specific tags
./testflowkit run --tags "@smoke"

# Run with specific configuration
./testflowkit run --config ./custom-config.yml`;
</script>
