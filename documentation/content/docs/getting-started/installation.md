---
title: Installation
description: Download and install TestFlowKit on your operating system
navigation:
  title: Installation
---

# Installation

## npm (recommended)

```bash
npm install -g @testflowkit/cli
tkit --version
```

Or run without installing:

```bash
npx @testflowkit/cli --version
```

Requires Node.js 16+.

## Direct download

Download the binary for your platform from [GitHub Releases](https://github.com/TestFlowKit/testflowkit/releases):

| Platform | File |
|----------|------|
| Windows x64 | `tkit-windows-amd64.zip` |
| macOS Intel | `tkit-darwin-amd64.tar.gz` |
| macOS Apple Silicon | `tkit-darwin-arm64.tar.gz` |
| Linux x64 | `tkit-linux-amd64.tar.gz` |

Extract, make executable (`chmod +x tkit` on macOS/Linux), then run `tkit --version`.

::alert{type="warning"}
**macOS:** If Gatekeeper blocks the binary, allow it in **System Preferences → Security & Privacy**, or run `xattr -d com.apple.quarantine ./tkit`.
::

## Browser drivers

| Driver | Setup |
|--------|-------|
| **rod** (default) | Bundled — no install step |
| **playwright** | Set `driver: "playwright"` in config, then run `tkit install` |

Playwright requires Go 1.19+ and downloads Chromium (~300MB) on first install.

## Next Steps

[Quick Start](/docs/getting-started/quick-start) — Create your first test project.
