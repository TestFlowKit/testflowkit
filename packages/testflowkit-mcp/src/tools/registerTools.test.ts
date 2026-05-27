import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { AgentConfigSchema } from "../config/types.js";
import { registerTools } from "./registerTools.js";
import { hasMacroTag } from "./write_macro.js";

type Handler = (args: any) => Promise<any>;

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tfk-mcp-tools-"));
}

function setupServerAndHandlers(
  config: unknown,
  configDir: string,
): Map<string, Handler> {
  const handlers = new Map<string, Handler>();
  const fakeServer = {
    registerTool(name: string, _meta: unknown, handler: Handler): void {
      handlers.set(name, handler);
    },
  };

  registerTools(
    fakeServer as unknown as McpServer,
    AgentConfigSchema.parse(config),
    configDir,
  );

  return handlers;
}

describe("hasMacroTag", () => {
  it("detects @macro tag", () => {
    assert.equal(hasMacroTag("@macro\nScenario: Reusable login"), true);
  });

  it("returns false for non macro content", () => {
    assert.equal(hasMacroTag("Feature: Auth\nScenario: Login"), false);
  });
});

describe("registerTools macro handlers", () => {
  it("get_guidelines returns macro concept rules", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
        agent: { capabilities: { macros: true } },
      },
      dir,
    );

    const getGuidelines = handlers.get("get_guidelines");
    assert.ok(getGuidelines, "get_guidelines handler should exist");

    const result = await getGuidelines!({ concept: "macro" });
    const text = result.content?.[0]?.text;
    assert.equal(typeof text, "string");

    const payload = JSON.parse(text) as {
      version: number;
      concept: string;
      guidelines: {
        capability: { key: string; enabled: boolean };
        authoring: { requiredTag: string; preferredTool: string };
        reading: { listTool: string; readTool: string };
      };
    };

    assert.equal(payload.version, 1);
    assert.equal(payload.concept, "macro");
    assert.equal(
      payload.guidelines.capability.key,
      "agent.capabilities.macros",
    );
    assert.equal(payload.guidelines.capability.enabled, true);
    assert.equal(payload.guidelines.authoring.requiredTag, "@macro");
    assert.equal(payload.guidelines.authoring.preferredTool, "write_macro");
    assert.equal(payload.guidelines.reading.listTool, "list_features");
    assert.equal(payload.guidelines.reading.readTool, "read_feature");

    fs.rmSync(dir, { recursive: true });
  });

  it("get_guidelines returns all guidelines for unknown concept", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
        agent: { capabilities: { macros: true } },
      },
      dir,
    );

    const getGuidelines = handlers.get("get_guidelines");
    assert.ok(getGuidelines, "get_guidelines handler should exist");

    const result = await getGuidelines!({ concept: "anything" });
    const payload = JSON.parse(result.content[0].text) as {
      version: number;
      concepts: { macro: { authoring: { preferredTool: string } } };
    };

    assert.equal(payload.version, 1);
    assert.equal(payload.concepts.macro.authoring.preferredTool, "write_macro");

    fs.rmSync(dir, { recursive: true });
  });

  it("get_guidelines reflects macros disabled state", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
        agent: { capabilities: { macros: false } },
      },
      dir,
    );

    const getGuidelines = handlers.get("get_guidelines");
    assert.ok(getGuidelines, "get_guidelines handler should exist");

    const result = await getGuidelines!({ concept: "macro" });
    const payload = JSON.parse(result.content[0].text) as {
      guidelines: { capability: { enabled: boolean } };
    };

    assert.equal(payload.guidelines.capability.enabled, false);

    fs.rmSync(dir, { recursive: true });
  });

  it("write_macro rejects when macro capability is disabled", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
        agent: { capabilities: { macros: false } },
      },
      dir,
    );

    const writeMacro = handlers.get("write_macro");
    assert.ok(writeMacro, "write_macro handler should exist");

    await assert.rejects(
      writeMacro!({
        path: "features/macros/auth.feature",
        content: "@macro\nScenario: Reusable login",
        createDirs: true,
      }),
      /Macro creation is disabled/,
    );

    fs.rmSync(dir, { recursive: true });
  });

  it("write_macro rejects content without @macro", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
      },
      dir,
    );

    const writeMacro = handlers.get("write_macro");
    assert.ok(writeMacro, "write_macro handler should exist");

    await assert.rejects(
      writeMacro!({
        path: "features/macros/auth.feature",
        content: "Feature: Auth\nScenario: Login",
        createDirs: true,
      }),
      /requires at least one scenario tagged with "@macro"/,
    );

    fs.rmSync(dir, { recursive: true });
  });

  it("write_feature blocks macro content when capability disabled", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
        agent: { capabilities: { macros: false } },
      },
      dir,
    );

    const writeFeature = handlers.get("write_feature");
    assert.ok(writeFeature, "write_feature handler should exist");

    await assert.rejects(
      writeFeature!({
        path: "features/macros/auth.feature",
        content: "@macro\nScenario: Reusable login",
        createDirs: true,
      }),
      /Macro creation is disabled/,
    );

    fs.rmSync(dir, { recursive: true });
  });

  it("write_macro writes valid macro content", async () => {
    const dir = mkTmp();
    const handlers = setupServerAndHandlers(
      {
        version: 1,
        project: { features_glob: "features/**/*.feature" },
        agent: { capabilities: { macros: true } },
      },
      dir,
    );

    const writeMacro = handlers.get("write_macro");
    assert.ok(writeMacro, "write_macro handler should exist");

    await writeMacro!({
      path: "features/macros/auth.feature",
      content:
        '@macro\nScenario: Reusable login\n  Given the user goes to the "login" page',
      createDirs: true,
    });

    const file = path.join(dir, "features", "macros", "auth.feature");
    assert.equal(fs.existsSync(file), true);

    fs.rmSync(dir, { recursive: true });
  });
});
