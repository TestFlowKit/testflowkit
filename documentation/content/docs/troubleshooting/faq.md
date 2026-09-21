---
title: FAQ
description: Short answers to the questions newcomers ask most
navigation:
  title: FAQ
---

## Do I need to know how to code?

No. You write scenarios in plain-English Gherkin and declare pages and APIs in YAML. See [Quick Start](/docs/getting-started/quick-start).

## Where is my report?

It depends on `settings.report_format`: `report.html` in the folder you ran `tkit run` from, or `report/report.json` or `report/report.xml`. See [Reporters](/docs/reference/reporters).

## I cannot see the browser

`tkit run` is headless by default. Use `tkit run --headless=false`. This flag wins over `frontend.headless` in the config.

## How do I run only some scenarios?

Tag them and filter: `tkit run --tags "@smoke"`. Expressions can combine tags: `"@login and not @slow"`. See [CLI Reference](/docs/reference/cli).

## How do I temporarily disable a scenario?

Add `@skip` above it: [Skip a test](/docs/how-to/skip-a-test).

## Which step should I use?

Search the [Step Catalog](/sentences) by keyword. Run `tkit validate` to find steps that do not exist before you run anything.

## Which browser driver should I pick?

Start with the default, `rod`: it is bundled and needs Chrome or Edge installed. Choose `playwright` only if you need it, and run `tkit install` once.

## Where do I put URLs and secrets that change per environment?

In an env file, selected with `--env-file`, and read as `{{ env.name }}`. See [Test multiple environments](/docs/how-to/test-multiple-environments).

## My value shows up as `{{ ... }}` in the test

The variable was never set, or its name is misspelled. See [Common Issues](/docs/troubleshooting/common-issues) and [Variables](/docs/patterns/variables).

## Can an AI assistant write tests for me?

Yes, through the MCP server: [IDE Agent](/docs/guides/ide-agent).

## Still stuck?

[Debug a failing scenario](/docs/how-to/debug-a-failing-scenario), then [Common Issues](/docs/troubleshooting/common-issues). If that does not help, open a [GitHub issue](https://github.com/TestFlowKit/testflowkit/issues).
