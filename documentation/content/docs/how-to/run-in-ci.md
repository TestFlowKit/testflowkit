---
title: Run in CI
description: Run TestFlowKit in a pipeline, fail the build on errors and publish reports
navigation:
  title: Run in CI
---

## Goal

Run your suite on every push, fail the build when a scenario fails, and keep the report.

## 1. Use a CI config

Reports are chosen in the config, so give CI its own file. `junit` is understood by most CI systems:

```yaml
# testflowkit.ci.yml: same as testflowkit.yml, except:
settings:
  gherkin_location: "features"
  report_format: "junit"     # writes report/report.xml
  concurrency: 2
```

Keep the URLs and secrets in an [environment file](/docs/how-to/test-multiple-environments) selected at run time.

## 2. Add the job

GitHub Actions:

```yaml
name: e2e
on: [push]
jobs:
  e2e:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm install -g @testflowkit/cli
      - run: tkit validate -c testflowkit.ci.yml
      - run: tkit run -c testflowkit.ci.yml --env-file .env.ci.yml --tags "not @wip"
      - uses: actions/upload-artifact@v4
        if: always()
        with:
          name: report
          path: report/
```

- `tkit run` is headless by default, so no display is needed.
- `if: always()` uploads the report even when tests fail.
- Using the Playwright driver? Add `- run: tkit install -c testflowkit.ci.yml` before the run.

## 3. Let the exit code gate the build

`tkit run` exits with `0` when all tests pass and non-zero otherwise (`1` test failure, `2` config error, `3` invalid arguments), so the job fails on its own. `tkit validate` catches config and step typos in seconds, before a browser starts.

## Verify

Push a branch with one failing assertion. The job must turn red and the `report` artifact must contain `report.xml`. Fix it and check the job goes green.

## See also

- [Reporters](/docs/reference/reporters): formats and paths
- [Run in parallel](/docs/how-to/run-in-parallel)
- [CLI Reference](/docs/reference/cli): tags, exit codes
