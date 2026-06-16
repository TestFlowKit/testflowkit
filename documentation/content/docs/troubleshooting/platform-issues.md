---
title: Platform Issues
description: Platform-specific troubleshooting for Windows, macOS, and Linux
navigation:
  title: Platform Issues
---

# Platform Issues

## macOS

| Problem | Fix |
|---------|-----|
| Gatekeeper blocks binary | Security & Privacy → Allow Anyway, or `xattr -d com.apple.quarantine ./tkit` |
| Apple Silicon crash | Download `darwin-arm64` build, not Intel |
| Chrome not found | Install Chrome/Edge in `/Applications/` |

## Windows

| Problem | Fix |
|---------|-----|
| SmartScreen warning | More info → Run anyway, or Properties → Unblock |
| Antivirus quarantine | Add exclusion for `tkit.exe` and project folder |
| Path too long | Move project to shorter path (e.g. `C:\tests\`) |
| Line ending errors | Use LF in `.feature` files (`git config core.autocrlf input`) |

## Linux

| Problem | Fix |
|---------|-----|
| Permission denied | `chmod +x ./tkit` |
| Missing shared libraries | Install browser deps (Debian/Ubuntu): |

```bash
sudo apt-get install -y libnss3 libatk1.0-0 libatk-bridge2.0-0 \
  libdrm2 libxkbcommon0 libxcomposite1 libxdamage1 libxfixes3 \
  libxrandr2 libgbm1 libasound2
```

| Problem | Fix |
|---------|-----|
| No display (headless server) | `tkit run --headless` or `xvfb-run tkit run` |
| Docker browser fails | Use `--no-sandbox` flag or `--cap-add=SYS_ADMIN` |

```yaml
frontend:
  args:
    - "--no-sandbox"
    - "--disable-setuid-sandbox"
```

## CI

```yaml
# GitHub Actions (Ubuntu)
- run: |
    sudo apt-get update && sudo apt-get install -y libnss3 libatk1.0-0
    npm install -g @testflowkit/cli
    tkit run --headless
```

## Next Steps

- [Common Issues](/docs/troubleshooting/common-issues) — General troubleshooting
- [Installation](/docs/getting-started/installation) — Setup guide
