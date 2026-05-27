import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { registerResources } from "./registerResources.js";

type ResourceHandler = () => Promise<{
  contents: Array<{ uri: string; mimeType?: string; text?: string }>;
}>;

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tfk-mcp-res-"));
}

describe("registerResources", () => {
  it("registers documentation resources", () => {
    const registered: string[] = [];

    const fakeServer = {
      registerResource(
        name: string,
        _uri: string,
        _meta: unknown,
        _handler: ResourceHandler,
      ): void {
        registered.push(name);
      },
    };

    const dir = mkTmp();
    registerResources(fakeServer as unknown as McpServer, dir, {
      docsBaseDir: dir,
    });

    assert.deepEqual(registered.sort(), [
      "guidelines-copilot-instructions",
      "guidelines-ide-agent",
      "guidelines-macros",
    ]);
  });

  it("returns fallback content when documentation file is missing", async () => {
    const handlers = new Map<string, ResourceHandler>();
    const fakeServer = {
      registerResource(
        name: string,
        _uri: string,
        _meta: unknown,
        handler: ResourceHandler,
      ): void {
        handlers.set(name, handler);
      },
    };

    const dir = mkTmp();
    registerResources(fakeServer as unknown as McpServer, dir, {
      docsBaseDir: dir,
    });

    const macroHandler = handlers.get("guidelines-macros");
    assert.ok(macroHandler, "guidelines-macros handler should exist");

    const result = await macroHandler!();
    assert.match(result.contents[0].text ?? "", /Resource not available/);
    assert.match(result.contents[0].text ?? "", /get_guidelines/);

    fs.rmSync(dir, { recursive: true });
  });

  it("returns markdown file content when documentation file exists", async () => {
    const handlers = new Map<string, ResourceHandler>();
    const fakeServer = {
      registerResource(
        name: string,
        _uri: string,
        _meta: unknown,
        handler: ResourceHandler,
      ): void {
        handlers.set(name, handler);
      },
    };

    const dir = mkTmp();
    const macroPath = path.join(dir, "docs", "macros.md");
    fs.mkdirSync(path.dirname(macroPath), { recursive: true });
    fs.writeFileSync(macroPath, "# Macros\n\nLocal test doc\n", "utf-8");

    registerResources(fakeServer as unknown as McpServer, dir, {
      docsBaseDir: dir,
    });

    const macroHandler = handlers.get("guidelines-macros");
    assert.ok(macroHandler, "guidelines-macros handler should exist");

    const result = await macroHandler!();
    assert.match(result.contents[0].text ?? "", /Local test doc/);
    assert.equal(result.contents[0].mimeType, "text/markdown");

    fs.rmSync(dir, { recursive: true });
  });
});
