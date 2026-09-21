---
title: Common Issues
description: Find your symptom, see the cause, apply the fix
navigation:
  title: Common Issues
---

Find the symptom, check the cause, apply the fix. Not listed? [Debug a failing scenario](/docs/how-to/debug-a-failing-scenario) shows how to get more information, and OS-specific problems are in [Platform Issues](/docs/troubleshooting/platform-issues).

## Before anything else

```bash
tkit validate
```

It catches most config, env reference and step typos in seconds, before a browser starts.

## Install and start

| Symptom | Cause | Fix |
|---|---|---|
| `tkit: command not found` | The binary or npm global bin is not on your `PATH` | Reinstall (`npm install -g @testflowkit/cli`) or add the folder to `PATH`, then check `tkit --version` |
| macOS "unidentified developer" | Gatekeeper blocks the binary | System Settings → Security → Allow Anyway, or `xattr -d com.apple.quarantine ./tkit` |
| Linux `permission denied` | The binary is not executable | `chmod +x ./tkit` |
| Windows SmartScreen block | Unsigned binary | **More info** → **Run anyway** |
| Browser will not start | Chrome or Edge missing (Rod), or Playwright browser not installed | Install Chrome or Edge. For Playwright, set `driver: "playwright"` and run `tkit install` |
| `tkit install` fails with `UNABLE_TO_GET_ISSUER_CERT_LOCALLY` | A corporate proxy or VPN inspects HTTPS and the bundled Node.js does not trust its certificate | See below |

### `UNABLE_TO_GET_ISSUER_CERT_LOCALLY`

`tkit install` downloads the Playwright browsers through a bundled Node.js driver, which uses Node's own certificate store instead of the OS one. Point it at your organization's root CA:

```bash
export NODE_EXTRA_CA_CERTS=/path/to/corporate-root-ca.pem
tkit install
```

Behind an HTTP(S) proxy, also set `HTTPS_PROXY` and `HTTP_PROXY`. To confirm the cause only (this disables certificate validation, never keep it):

```bash
NODE_TLS_REJECT_UNAUTHORIZED=0 tkit install
```

## Configuration

| Symptom | Cause | Fix |
|---|---|---|
| Config not found | `tkit` is run outside the project | Run from the project root, or `tkit run --config path/to/testflowkit.yml` |
| Invalid YAML | Tabs, bad indentation, unquoted special characters | Use spaces only and quote strings containing `:` or `#` |
| Element or page not found in config | The name in the step differs from the one under `frontend.elements` or `frontend.pages` (names are case-sensitive) | Align the two names |
| `{{ env.something }}` not resolved | The key is missing from the `env:` block or the env file in use | Define it, or pass the right file with `--env-file`. `tkit validate` reports it |

## UI steps

| Symptom | Cause | Fix |
|---|---|---|
| Element not found, or timeout | Wrong selector, wrong page, or slow page | Test the selector in DevTools with `document.querySelector('#selector')`. Add fallback selectors. Raise `frontend.default_timeout`. See [Wait for an element](/docs/how-to/wait-for-an-element) |
| Works locally, fails in CI | Headless timing or a missing env value | Rerun locally without `--headless=false` (same mode as CI), then [debug](/docs/how-to/debug-a-failing-scenario) |
| The browser window never appears | `tkit run` is headless by default | `tkit run --headless=false` |

Fallback selectors, most reliable first:

```yaml
frontend:
  elements:
    login:
      submit_button:
        - "[data-testid='submit']"
        - "#submit-btn"
        - "button[type='submit']"
```

To scroll to an element first: `the user scrolls to the "element" element`.

## API steps

| Symptom | Cause | Fix |
|---|---|---|
| Connection refused | Server not running, or wrong `base_url` | Start the server and check the URL |
| `401 Unauthorized` | No credentials sent | [Configure authentication](/docs/how-to/configure-authentication), or `I set the header "Authorization" to "Bearer {{auth_token}}"` |
| Timeout | Slow endpoint | Raise the endpoint, API or `apis.default_timeout` |
| API or endpoint unknown | `"api_name.endpoint_name"` does not match the config keys | Compare with `apis.definitions` |

## Variables

| Symptom | Cause | Fix |
|---|---|---|
| `{{name}}` shown literally | The variable was never stored, or is misspelled | Store it in an earlier step of the same scenario and check the spelling |
| Empty value | The storing step did not run yet, or stored an empty result | Print it with `I display the value of variable "name"` |
| Value missing in another scenario | Variables belong to one scenario | Use a [global variable](/docs/how-to/share-data-between-scenarios) or a [global hook](/docs/patterns/global-hooks) |

## Getting help

Open a [GitHub issue](https://github.com/TestFlowKit/testflowkit/issues) with the version (`tkit --version`), your OS, a minimal reproduction and the error output.

## See also

- [FAQ](/docs/troubleshooting/faq)
- [Debug a failing scenario](/docs/how-to/debug-a-failing-scenario)
- [CLI Reference](/docs/reference/cli)
