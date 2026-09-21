---
title: Changelog
description: User-facing changes and where to find release notes
navigation:
  title: Changelog
---

Release notes for each version are on [GitHub Releases](https://github.com/TestFlowKit/testflowkit/releases). Check your version with `tkit --version`.

## Recent changes

Listed from the project history, newest first. The release each one shipped in is on GitHub Releases.

- **Reports**: `report_format: json` now writes Cucumber JSON, and `junit` writes JUnit XML. **Breaking** for anything reading the previous JSON report. See [Reporters](/docs/reference/reporters) and the [Migration Guide](/docs/troubleshooting/migration-guide).
- **MCP server**: a `list_macros` tool lets an agent reuse existing `@macro` scenarios. See [Macros](/docs/patterns/macros).
- **Debug output**: verbosity levels (`--verbosity 1` to `3`) and scopes (`--debug-scope`), scenario filter, log file and JSON log format. See [CLI Reference](/docs/reference/cli).
- **Config schema**: `tkit export-config-schema` exports the config schema, also exposed by the MCP server.
- **Playwright driver**: the underlying library moved to `mxschmitt/playwright-go`. No config change is needed.
