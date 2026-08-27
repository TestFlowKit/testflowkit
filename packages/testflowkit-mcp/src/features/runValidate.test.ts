import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { runValidate } from "./runValidate.js";

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-runvalidate-"));
}

function writeFakeCliScript(dir: string, script: string): string {
  const scriptPath = path.join(dir, "fake-tkit.sh");
  fs.writeFileSync(scriptPath, `#!/bin/sh\n${script}`, { mode: 0o755 });
  return scriptPath;
}

describe("runValidate", () => {
  it("returns ok:true on a successful validation", () => {
    const dir = mkTmp();
    try {
      const scriptPath = writeFakeCliScript(
        dir,
        "echo 'All is good !'\nexit 0",
      );
      const result = runValidate(
        scriptPath,
        dir,
        path.join(dir, "testflowkit.yml"),
      );
      assert.equal(result.ok, true);
      assert.equal(result.exitCode, 0);
      assert.match(result.output, /All is good/);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("returns ok:false with captured output on validation failure, without throwing", () => {
    const dir = mkTmp();
    try {
      const scriptPath = writeFakeCliScript(
        dir,
        "echo 'List of undefined steps:' 1>&2\nexit 1",
      );
      const result = runValidate(
        scriptPath,
        dir,
        path.join(dir, "testflowkit.yml"),
      );
      assert.equal(result.ok, false);
      assert.equal(result.exitCode, 1);
      assert.match(result.output, /List of undefined steps/);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("passes tags through to the CLI command", () => {
    const dir = mkTmp();
    try {
      const scriptPath = writeFakeCliScript(dir, 'echo "args: $*"');
      const result = runValidate(
        scriptPath,
        dir,
        path.join(dir, "testflowkit.yml"),
        "@wip",
      );
      assert.match(result.output, /--tags @wip/);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("safely quotes tags containing shell metacharacters", () => {
    const dir = mkTmp();
    try {
      const scriptPath = writeFakeCliScript(dir, 'echo "args: $*"');
      const result = runValidate(
        scriptPath,
        dir,
        path.join(dir, "testflowkit.yml"),
        "@wip; echo injected",
      );
      assert.equal(result.ok, true);
      assert.doesNotMatch(result.output, /injected\n/);
      assert.match(result.output, /@wip; echo injected/);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });
});
