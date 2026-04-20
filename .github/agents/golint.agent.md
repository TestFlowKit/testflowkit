---
name: golint
description: Run go lint and fix lint errors
model: Auto (copilot)
tools: ['vscode', 'execute', 'read', 'agent', 'edit'] # specify the tools this agent can use. If not set, all enabled tools are allowed.
---

<!-- Tip: Use /create-agent in chat to generate content with agent assistance -->

Run the command `golangci-lint run` to identify linting issues in the Go codebase. Review the output for any lint errors and warnings. For each identified issue, determine the appropriate fix based on the linting rules and best practices. Implement the necessary code changes to resolve the lint errors, ensuring that the code adheres to the specified style and quality guidelines. After making the fixes, run `golangci-lint run` again to verify that all linting issues have been addressed successfully.