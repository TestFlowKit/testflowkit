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

The `tkit init` command generates a `config.yml` file in your project root. You can edit it to match your needs.

## IDE agent (MCP)

TestFlowKit also ships an MCP server for Cursor and VS Code Copilot: [`@testflowkit/mcp`](https://www.npmjs.com/package/@testflowkit/mcp).

The MCP server is a companion to this CLI — it does not replace `tkit run`. It connects your IDE to the same step catalog and project config so AI assistants can draft valid Gherkin tests using only registered sentences. Under the hood it calls `tkit export-step-definitions` and `tkit export-config-schema` from your installed CLI.

Install the CLI first, then add the MCP server to `.cursor/mcp.json` or `.vscode/mcp.json`:

```json
{
  "mcpServers": {
    "testflowkit": {
      "command": "npx",
      "args": ["-y", "@testflowkit/mcp"]
    }
  }
}
```

See the [MCP package README](https://github.com/TestFlowKit/testflowkit/tree/main/packages/testflowkit-mcp) and the [IDE Agent guide](https://testflowkit.github.io/testflowkit/docs/guides/ide-agent) for setup, tools, and the optional `agent:` block in `testflowkit.yml`.

## 📚 Documentation

For comprehensive documentation, guides, examples, and best practices, please visit the **[TestFlowKit Web Documentation](https://testflowkit.github.io/testflowkit/)** on GitHub Pages.

The documentation includes:
- Getting Started Guide
- Configuration Options
- Step Definitions & Sentences
- Variable System
- API Testing
- Advanced Features
- FAQ & Troubleshooting

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
- [@testflowkit/mcp](https://www.npmjs.com/package/@testflowkit/mcp) — IDE agent (MCP server)
- [Issue Tracker](https://github.com/TestFlowKit/testflowkit/issues)
- [Releases](https://github.com/TestFlowKit/testflowkit/releases)
