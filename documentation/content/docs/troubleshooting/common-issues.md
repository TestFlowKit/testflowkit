---
title: Common Issues
description: Solutions to frequently encountered problems
navigation:
  title: Common Issues
---

# Common Issues

## Installation

| Problem | Fix |
|---------|-----|
| macOS "unidentified developer" | System Preferences → Security → Allow Anyway, or `xattr -d com.apple.quarantine ./tkit` |
| Linux permission denied | `chmod +x ./tkit` |
| Windows SmartScreen block | Click **More info** → **Run anyway** |

## Configuration

| Problem | Fix |
|---------|-----|
| Config not found | Run from project root, or `tkit run --config path/to/testflowkit.yml` |
| Invalid YAML | Check indentation (spaces only), quote strings with special chars |
| Element not in config | Add to `frontend.elements` — names are case-sensitive |
| Undefined `{{ env.* }}` | Define in `env:` block or env file; run `tkit validate` |

## Element not found

1. Test selector in DevTools: `document.querySelector('#selector')`
2. Add fallback selectors (most reliable first)
3. Increase `frontend.default_timeout`
4. Scroll into view first: `the user scrolls to the "element" element`

```yaml
frontend:
  elements:
    login:
      submit_button:
        - "[data-testid='submit']"
        - "#submit-btn"
        - "button[type='submit']"
```

## Browser won't start

- Ensure Chrome/Edge is installed
- Try headless: `tkit run --headless`
- Playwright: run `tkit install` after setting `driver: "playwright"`

### `tkit install` fails with `UNABLE_TO_GET_ISSUER_CERT_LOCALLY`

`tkit install` downloads the Playwright browsers via a bundled Node.js driver, which uses Node's own certificate store instead of the OS trust store. This error means Node can't validate the TLS certificate chain — almost always caused by a corporate proxy/VPN that inspects HTTPS traffic (Zscaler, Netskope, Fortinet, etc.).

```bash
# point Node at your organization's root CA certificate
export NODE_EXTRA_CA_CERTS=/path/to/corporate-root-ca.pem
tkit install
```

If you're behind an HTTP(S) proxy, also set `HTTPS_PROXY`/`HTTP_PROXY`. To confirm the cause (do not use permanently — disables cert validation):

```bash
NODE_TLS_REJECT_UNAUTHORIZED=0 tkit install
```

## API issues

| Problem | Fix |
|---------|-----|
| Connection refused | Check server is running and `base_url` is correct |
| 401 Unauthorized | Set auth header or configure `security_schemes` |
| Timeout | Increase `apis.default_timeout` or endpoint timeout |
| Wrong endpoint | Verify `"api_name.endpoint_name"` matches config keys |

```gherkin
And I set the header "Authorization" to "Bearer {{auth_token}}"
```

## Variables

| Problem | Fix |
|---------|-----|
| Variable not substituted | Store before use; check spelling |
| Empty value | Verify the storing step ran successfully |
| Not available across scenarios | Use [global hooks](/docs/patterns/global-hooks) |

## Debugging

```bash
tkit run --debug                          # Headers and variable substitutions (verbosity 2)
tkit run --verbosity 3                    # Full bodies and all internals
tkit run --debug --debug-scope http       # HTTP layer only
tkit run --debug --debug-scenario "Login" # One scenario only
tkit run --debug --log-file debug.log     # Persist to file
tkit run --debug --log-format json | jq . # Structured output
tkit validate                             # Check config and env references
```

| Verbosity | Flag | Shows |
|-----------|------|-------|
| 1 | `--verbosity=1` | Scenario/step flow, timings |
| 2 | `--debug` | + HTTP headers, variable substitutions |
| 3 | `--verbosity=3` | + Full bodies, browser actions |

Use `--debug-scope` to reduce noise: `http`, `browser`, `variables`, `config`.

Enable screenshots on failure:

```yaml
frontend:
  screenshot_on_failure: true
```

## Getting help

Open a [GitHub issue](https://github.com/TestFlowKit/testflowkit/issues) with version (`tkit version`), OS, minimal reproduction, and error output.

## Next Steps

- [Platform Issues](/docs/troubleshooting/platform-issues) — OS-specific fixes
- [CLI Reference](/docs/reference/cli) — Command options
- [testflowkit.yml](/docs/config/overview) — Config reference
