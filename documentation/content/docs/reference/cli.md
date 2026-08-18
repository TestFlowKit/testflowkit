---
title: CLI Reference
description: Complete command line interface reference for TestFlowKit
navigation:
  title: CLI Reference
---

# CLI Reference

```bash
tkit [command] [options]
```

Default config file: `testflowkit.yml` (falls back to legacy `config.yml`).

## Commands

| Command | Description |
|---------|-------------|
| `tkit run` | Execute test scenarios |
| `tkit validate` | Validate config and feature files |
| `tkit init` | Scaffold a new project |
| `tkit install` | Install browser driver (Playwright) |
| `tkit export-step-definitions` | Export step catalog as JSON |
| `tkit version` | Show version info |

## run

```bash
tkit run [options]
```

| Option | Short | Description | Default |
|--------|-------|-------------|---------|
| `--config` | `-c` | Config file path | `testflowkit.yml` |
| `--location` | `-l` | Feature files directory | from config |
| `--tags` | `-t` | Tag filter expression | all |
| `--env-file` | | Environment variables YAML | from config |
| `--headless` | | Headless browser mode | `true` |
| `--timeout` | | Request/element timeout (`10s`, `1m`) | from config |
| `--debug` | | Enable debug output (verbosity 2: headers and variable substitutions) | `false` |
| `--verbosity` | | Debug verbosity: `1`=summary, `2`=detailed, `3`=trace (overrides `--debug`) | `0` |
| `--debug-scope` | | Restrict debug to comma-separated scopes: `http`, `browser`, `variables`, `config` | all |
| `--debug-scenario` | | Restrict debug output to scenarios whose name or tag contains this value | all |
| `--log-file` | | Write debug output to this file path in addition to stdout | |
| `--log-format` | | Log output format: `text` or `json` | `text` |

```bash
tkit run
tkit run --tags @smoke
tkit run --tags "@login and not @slow"
tkit run --env-file .env.staging.yml
tkit run --config staging.yml --location e2e/features
tkit run --debug
tkit run --verbosity 3
tkit run --debug --debug-scope http
tkit run --debug --debug-scenario "Login"
tkit run --debug --log-format json | jq '.message'
tkit run --debug --log-file debug.log
```

## Debug verbosity levels

| Level | Flag | What it shows |
|-------|------|---------------|
| `0` | *(off)* | No debug output |
| `1` | `--verbosity=1` | Scenario names, step flow, timings |
| `2` | `--debug` or `--verbosity=2` | + HTTP headers, variable substitutions |
| `3` | `--verbosity=3` | + Full request/response bodies, browser actions |

## Debug scopes

Use `--debug-scope` to reduce noise when you only need to inspect one layer:

```bash
tkit run --debug --debug-scope http           # HTTP requests/responses only
tkit run --debug --debug-scope variables      # Variable resolution only
tkit run --debug --debug-scope http,browser   # Multiple scopes
```

| Scope | What it covers |
|-------|----------------|
| `http` | REST and GraphQL request/response logging |
| `browser` | Browser automation events |
| `variables` | Variable resolution and end-of-scenario dump |
| `config` | Configuration loading |

## JSON log format

The `--log-format json` flag emits one JSON object per line, useful for piping into `jq` or a log aggregator:

```bash
tkit run --debug --log-format json
# {"timestamp":"2026-08-18T10:30:00Z","level":"DEBUG","message":"→ GET /api/users"}

tkit run --debug --log-format json | jq 'select(.level == "ERROR")'
```

## validate

```bash
tkit validate [-c config] [-l location] [-t tags] [--env-file file]
```

## install

Installs Playwright when `frontend.driver: "playwright"` is set in config. Rod requires no install.

```bash
tkit install
tkit install --config staging.yml
```

## export-step-definitions

```bash
tkit export-step-definitions --format json > step-definitions.json
```

Used by the MCP server and for offline step catalog snapshots.

## Tag expressions

```bash
tkit run --tags @smoke
tkit run --tags "@smoke and @login"
tkit run --tags "@smoke or @regression"
tkit run --tags "not @slow"
tkit run --tags "@smoke and not @wip"
```

## Exit codes

| Code | Meaning |
|------|---------|
| `0` | All tests passed |
| `1` | Test failure |
| `2` | Configuration error |
| `3` | Invalid arguments |

## CI example

```yaml
# GitHub Actions
- run: |
    npm install -g @testflowkit/cli
    tkit run --headless --tags @smoke
```

## Next Steps

- [testflowkit.yml](/docs/config/overview) — Config file reference
- [Step Catalog](/sentences) — Searchable step catalog
