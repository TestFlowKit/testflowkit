import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { ValidateGherkinFilesTool } from "./validate_gherkin_files.js";

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-validate-"));
}

function writeFakeCliScript(dir: string, script: string): string {
  const scriptPath = path.join(dir, "fake-tkit.sh");
  fs.writeFileSync(scriptPath, `#!/bin/sh\n${script}`, { mode: 0o755 });
  return scriptPath;
}

describe("ValidateGherkinFilesTool", () => {
  it("reports ok:true when tkit validate succeeds", async () => {
    const dir = mkTmp();
    const previousCliPath = process.env.TESTFLOWKIT_CLI_PATH;
    try {
      const scriptPath = writeFakeCliScript(
        dir,
        "echo 'All is good !'\nexit 0",
      );
      process.env.TESTFLOWKIT_CLI_PATH = scriptPath;

      const tool = new ValidateGherkinFilesTool();
      const result = await tool.handler({
        config: {
          configPath: path.join(dir, "testflowkit.yml"),
          featuresGlob: "features/**/*.feature",
        },
        configDir: dir,
        input: {},
      });

      const payload = JSON.parse(result.content[0]?.text ?? "{}");
      assert.equal(payload.ok, true);
      assert.equal(payload.exitCode, 0);
    } finally {
      if (previousCliPath === undefined) {
        delete process.env.TESTFLOWKIT_CLI_PATH;
      } else {
        process.env.TESTFLOWKIT_CLI_PATH = previousCliPath;
      }
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("reports ok:false with output when tkit validate fails, without throwing", async () => {
    const dir = mkTmp();
    const previousCliPath = process.env.TESTFLOWKIT_CLI_PATH;
    try {
      const scriptPath = writeFakeCliScript(
        dir,
        "echo 'List of undefined steps: Given I call missing macro' 1>&2\nexit 1",
      );
      process.env.TESTFLOWKIT_CLI_PATH = scriptPath;

      const tool = new ValidateGherkinFilesTool();
      const result = await tool.handler({
        config: {
          configPath: path.join(dir, "testflowkit.yml"),
          featuresGlob: "features/**/*.feature",
        },
        configDir: dir,
        input: { tags: "@wip" },
      });

      const payload = JSON.parse(result.content[0]?.text ?? "{}");
      assert.equal(payload.ok, false);
      assert.equal(payload.exitCode, 1);
      assert.match(payload.output, /undefined steps/);
    } finally {
      if (previousCliPath === undefined) {
        delete process.env.TESTFLOWKIT_CLI_PATH;
      } else {
        process.env.TESTFLOWKIT_CLI_PATH = previousCliPath;
      }
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });
});
