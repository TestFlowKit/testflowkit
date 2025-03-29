# Web Test Automation Tool: TestFlowKit

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Key Concepts of TestFlowKit](#key-concepts-of-testflowkit)
- [Contribution](#contribution)
- [Bug Reports](#bug-reports)
- [License](#license)
- [Author](#author)
- [Acknowledgments](#acknowledgments)

## Description

TestFlowKit is an open-source command-line tool designed to simplify the creation of automated tests for web applications. It is developed in Go and uses predefined Gherkin phrases. This README is primarily intended for developers and contributors. For detailed usage instructions, please refer to the [official documentation](https://testflowkit.dreamsfollowers.me/).

## Features

- **Gherkin Syntax:** Write tests using a clear and concise syntax.
- **Command-Line Interface:** Manage and execute your tests directly from the terminal.
- **Open Source:** Benefit from the flexibility and transparency of an open-source project.
- **Extensible:** Ability to add new features and adapt to your specific needs.
- **Developed in Go:** Benefit from the performance and speed of Go.
- **Based on TED:** Organize your tests efficiently using the principles of Test Execution Design.
- **Macros:** Define reusable scenarios to simplify test writing.

## Installation

### Prerequisites

Ensure you have Go installed and configured.

### Installation

1.  **Installation:**

    ```
    go mod tidy
    ```

2.  **Binary Installation:**

    You can download the pre-built binary from the [GitHub Releases](https://github.com/TestFlowKit/testflowkit/releases) page or from the [official documentation](https://testflowkit.dreamsfollowers.me/).

## Usage

**For detailed usage instructions, please refer to the [official documentation](https://testflowkit.dreamsfollowers.me/).** This README is primarily intended for developers and contributors.

### Basic Syntax

Tests are written using Gherkin syntax. Here is an example:

```gherkin
Feature: User Login
Scenario: Successful login
Given the user is on the login page
When they enter a valid username and a valid password
Then they should be logged in
```

### Running Tests

To run your tests, use the following command:

```bash
make test
```

## Key Concepts of TestFlowKit

TestFlowKit is built around the concepts of Test Execution Design (TED). Here are the main ones:

- **Test Plan:** Organizes the entire set of tests.
- **Test Suite:** Groups associated test cases.
- **Test Case:** Defines an individual test scenario.
- **Step:** Individual actions performed in a test case.
- **Macros:** Reusable scenarios that can be called from other scenarios.

For more information, see the [Test Execution Design documentation](https://testflowkit.dreamsfollowers.me/docs/category/test-execution-design-ted).

## Contribution

Contributions are welcome! To contribute to the project, please follow these steps:

1.  Clone the repository.
2.  Create a branch for your feature, following the format \`type/feature-name\` (e.g., \`feat/new-feature\` or \`chore/update-dependencies\`). Allowed branch types are:

    - \`feat\`: Adding a new feature.
    - \`fix\`: Fixing a bug.
    - \`docs\`: Modifying documentation.
    - \`style\`: Changes that do not affect the meaning of the code (spaces, formatting, semicolons, etc.).
    - \`refactor\`: Refactoring code without adding a feature or fixing a bug.
    - \`perf\`: Improving performance.
    - \`test\`: Adding or modifying tests.
    - \`build\`: Changes that affect the build system or external dependencies.
    - \`ci\`: Changes to our CI configuration files and scripts.
    - \`chore\`: Other changes that do not modify the source code or tests.

3.  Commit your changes using the [Conventional Commits](https://www.conventionalcommits.org/en/) format.
4.  Push to the branch (\`git push origin type/feature-name\`).
5.  Create a pull request to the main repository: [TestFlowKit/testflowkit](https://github.com/TestFlowKit/testflowkit).

## Bug Reports

If you encounter a bug, please [create a GitHub issue](https://github.com/TestFlowKit/testflowkit/issues).

## License

[MIT License](https://opensource.org/licenses/MIT)

## Author

[Marc-Henry Nanguy](https://github.com/marckent04)

## Acknowledgments

I would like to thank the following contributors and libraries:

- **Contributors:**
  - [Stéphane Salou](https://github.com/stephsalou)
- **Libraries:**
  - [godog](https://github.com/cucumber/godog)
  - [go-rod](https://github.com/go-rod/rod)
  - [go-yaml](https://github.com/goccy/go-yaml)
