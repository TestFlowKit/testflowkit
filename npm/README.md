# @testflowkit/cli

[![npm version](https://img.shields.io/npm/v/@testflowkit/cli.svg)](https://www.npmjs.com/package/@testflowkit/cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**TestFlowKit CLI** - A powerful end-to-end testing framework using Gherkin syntax.

## Installation

```bash
# Global installation (recommended)
npm install -g @testflowkit/cli

# Or use npx directly
npx @testflowkit/cli --version
```

## Usage

### Initialize a new test project

```bash
tkit init
```

### Run tests

```bash
tkit run
```

### Validate Gherkin files

```bash
tkit validate
```

### Check version

```bash
tkit --version
```

## Configuration

Create a `config.yml` file in your project root:

```yaml
mode: frontend
browser:
  type: chromium
  headless: true
features:
  - path: ./features
```

## Documentation

For full documentation, visit [TestFlowKit Documentation](https://testflowkit.github.io/testflowkit/).

## Supported Platforms

| OS      | Architecture |
|---------|--------------|
| Linux   | x64, arm64   |
| macOS   | x64, arm64   |
| Windows | x64, arm64   |

## Requirements

- Node.js >= 16.0.0

## License

MIT © [TestFlowKit](https://github.com/TestFlowKit)

## Links

- [GitHub Repository](https://github.com/TestFlowKit/testflowkit)
- [Issue Tracker](https://github.com/TestFlowKit/testflowkit/issues)
- [Releases](https://github.com/TestFlowKit/testflowkit/releases)
