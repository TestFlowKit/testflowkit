---
title: Run Scenarios in Parallel
description: Shorten a long suite by running scenarios concurrently
navigation:
  title: Run in Parallel
---

## Goal

Cut the total run time by executing several scenarios at once.

## Set the concurrency

Parallelism is a config setting, not a command-line flag. Values from 1 (default) to 20:

```yaml
settings:
  concurrency: 4
```

Different values per machine? Keep a second config and pick it with `-c`:

```bash
tkit run -c testflowkit.ci.yml
```

## Make scenarios safe to run together

Scenarios must not depend on each other or on the order they run in.

- Create the data a scenario needs inside it, with unique values: `{{ rand:email }}` rather than a fixed `jane@example.com` that two scenarios would both register.
- Do not rely on a variable set by another scenario. Use [global hooks](/docs/patterns/global-hooks) for one-time setup, such as logging in once.
- More concurrent scenarios mean more browsers open at once, so watch CPU and memory. Start with 2 to 4.

## Verify

Run the suite twice, with `concurrency: 1` and with your target value, and compare:

- the total run time should drop
- the set of passed and failed scenarios should be identical

If a scenario fails only when parallel, it shares data or state with another one: rerun it alone with `tkit run --tags "@your_tag"`.

## See also

- [Run in CI](/docs/how-to/run-in-ci)
- [Share data between scenarios](/docs/how-to/share-data-between-scenarios)
- [Global Hooks](/docs/patterns/global-hooks)
