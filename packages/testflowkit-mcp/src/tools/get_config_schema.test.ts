import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { GetConfigSchemaTool } from "./get_config_schema.js";

function makeTempDir(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-getconfigschema-"));
}

function writeFakeCliScript(dir: string, output: string): string {
  const scriptPath = path.join(dir, "fake-tkit.sh");
  fs.writeFileSync(scriptPath, `#!/bin/sh\necho '${output}'`, { mode: 0o755 });
  return scriptPath;
}

describe("GetConfigSchemaTool", () => {
  it("returns full config schema payload from CLI export", async () => {
    const tempDir = makeTempDir();
    const previousCliPath = process.env.TESTFLOWKIT_CLI_PATH;

    try {
      const payload = {
        root: "Config",
        version: "1.2.3",
        schema: {
          name: "Config",
          type: "object",
          properties: {
            settings: {
              yaml_key: "settings",
              type: "object",
            },
          },
        },
      };

      const scriptPath = writeFakeCliScript(tempDir, JSON.stringify(payload));
      process.env.TESTFLOWKIT_CLI_PATH = scriptPath;

      const tool = new GetConfigSchemaTool();
      const result = await tool.handler({
        config: {
          configPath: path.join(tempDir, "testflowkit.yml"),
          featuresGlob: "features/**/*.feature",
        },
        configDir: tempDir,
        input: {},
      });

      const text = result.content[0]?.text;
      assert.ok(text);
      assert.deepEqual(JSON.parse(text ?? "{}"), payload);
    } finally {
      if (previousCliPath === undefined) {
        delete process.env.TESTFLOWKIT_CLI_PATH;
      } else {
        process.env.TESTFLOWKIT_CLI_PATH = previousCliPath;
      }
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("throws when CLI output is not valid schema payload", async () => {
    const tempDir = makeTempDir();
    const previousCliPath = process.env.TESTFLOWKIT_CLI_PATH;

    try {
      const scriptPath = writeFakeCliScript(tempDir, '{"unexpected":true}');
      process.env.TESTFLOWKIT_CLI_PATH = scriptPath;

      const tool = new GetConfigSchemaTool();
      await assert.rejects(
        () =>
          tool.handler({
            config: {
              configPath: path.join(tempDir, "testflowkit.yml"),
              featuresGlob: "features/**/*.feature",
            },
            configDir: tempDir,
            input: {},
          }),
      );
    } finally {
      if (previousCliPath === undefined) {
        delete process.env.TESTFLOWKIT_CLI_PATH;
      } else {
        process.env.TESTFLOWKIT_CLI_PATH = previousCliPath;
      }
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});