<template>
    <ul class="list-disc list-inside mb-4 space-y-3">
        <li>
            <b>Go 1.23+:</b> Required for building from source.
            <ul class="list-inside ml-4 mt-1">
                <li>
                    <a href="https://golang.org/dl/" target="_blank" rel="noopener noreferrer"
                        class="text-blue-500 hover:underline">
                        Download and install Go
                    </a>
                </li>
            </ul>
        </li>
        <li>
            <b>Git:</b> For cloning the repository and version control.
            <ul class="list-inside ml-4 mt-1">
                <li>
                    <a href="https://git-scm.com/downloads" target="_blank" rel="noopener noreferrer"
                        class="text-blue-500 hover:underline">
                        Download Git
                    </a>
                </li>
            </ul>
        </li>
        <li>
            <b>Make:</b> For build automation (optional but recommended).
            <ul class="list-inside ml-4 mt-1">
                <li>Usually pre-installed on Linux/macOS</li>
                <li>Windows users can use WSL or install via package managers</li>
            </ul>
        </li>
        <li>
            <b>Chromium Browser:</b> TestFlowKit uses Chromium for web testing.
            <ul class="list-inside ml-4 mt-1">
                <li>TestFlowKit automatically installs Chromium via Rod</li>
                <li>
                    <a href="https://www.chromium.org/getting-involved/download-chromium/" target="_blank"
                        rel="noopener noreferrer" class="text-blue-500 hover:underline">
                        Manual installation instructions
                    </a>
                </li>
            </ul>
        </li>
        <li>
            <b>Understanding of Gherkin:</b> Gherkin is the language used to write your test scenarios.
            <ul class="list-inside ml-4 mt-1">
                <li>
                    <a href="https://cucumber.io/docs/gherkin/reference/" target="_blank" rel="noopener noreferrer"
                        class="text-blue-500 hover:underline">
                        Gherkin reference guide
                    </a>
                </li>
            </ul>
        </li>
        <li>
            <b>Understanding of YAML:</b> TestFlowKit uses YAML for configuration.
            <ul class="list-inside ml-4 mt-1">
                <li>
                    <a href="https://learnxinyminutes.com/docs/yaml/" target="_blank" rel="noopener noreferrer"
                        class="text-blue-500 hover:underline">
                        YAML tutorial
                    </a>
                </li>
            </ul>
        </li>
        <li>
            <b>Command Line Interface (CLI):</b> You'll need to be comfortable using a CLI to run TestFlowKit commands.
        </li>
    </ul>
</template>