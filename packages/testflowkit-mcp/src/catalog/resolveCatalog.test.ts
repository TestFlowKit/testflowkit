import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { resolveCatalog } from "./resolveCatalog.js";
import { AgentConfigSchema } from "../config/types.js";

function makeConfig() {
  return AgentConfigSchema.parse({ version: 1 });
}

function makeCatalogJson(label: string): string {
  return JSON.stringify([
    {
      sentence: `I use ${label}`,
      description: "Example step",
      categories: ["restapi"],
      example: "Given I use an example",
    },
  ]);
}

describe("resolveCatalog cache behavior", () => {
  it("writes release catalogs to a versioned cache file under .testflowkit/cache", async (t) => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      t.mock.method(globalThis, "fetch", async () => {
        return {
          ok: true,
          status: 200,
          text: async () => makeCatalogJson("release"),
        } as Response;
      });

      const result = await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.4",
      });

      const cachePath = path.join(
        tempDir,
        ".testflowkit/cache/step-definitions-1.2.4.json",
      );

      assert.equal(result.meta.source, "release");
      assert.equal(fs.existsSync(cachePath), true);
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("uses existing versioned cache before fetching", async (t) => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      const cachePath = path.join(
        tempDir,
        ".testflowkit/cache/step-definitions-1.2.4.json",
      );
      fs.mkdirSync(path.dirname(cachePath), { recursive: true });
      fs.writeFileSync(cachePath, makeCatalogJson("cached"), "utf-8");

      t.mock.method(globalThis, "fetch", async () => {
        throw new Error("fetch should not be called when cache exists");
      });

      const result = await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.4",
      });

      assert.equal(result.meta.source, "cache");
      assert.equal(result.steps[0]?.sentence, "I use cached");
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("separates cache files by resolved CLI base version", async (t) => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      t.mock.method(globalThis, "fetch", async (input) => {
        const url = String(input);
        const tag = url.split("/download/")[1]?.split("/")[0] ?? "unknown";
        return {
          ok: true,
          status: 200,
          text: async () => makeCatalogJson(tag),
        } as Response;
      });

      await resolveCatalog(makeConfig(), tempDir, { cliVersion: "1.2.3" });
      await resolveCatalog(makeConfig(), tempDir, { cliVersion: "1.3.0" });

      const cache123 = path.join(
        tempDir,
        ".testflowkit/cache/step-definitions-1.2.3.json",
      );
      const cache130 = path.join(
        tempDir,
        ".testflowkit/cache/step-definitions-1.3.0.json",
      );

      assert.equal(fs.existsSync(cache123), true);
      assert.equal(fs.existsSync(cache130), true);

      const cached123 = await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.3",
      });
      assert.equal(cached123.meta.source, "cache");
      assert.equal(cached123.steps[0]?.sentence, "I use 1.2.3");
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});

describe("resolveCatalog local CLI behavior", () => {
  it("returns source: local when a working CLI binary is provided", async () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      const catalogJson = makeCatalogJson("local-cli");
      const scriptPath = path.join(tempDir, "fake-tkit.sh");
      fs.writeFileSync(scriptPath, `#!/bin/sh\necho '${catalogJson}'`, {
        mode: 0o755,
      });

      const result = await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.4",
        cliBinary: scriptPath,
      });

      assert.equal(result.meta.source, "local");
      assert.equal(result.steps[0]?.sentence, "I use local-cli");

      // Verify the cache was written
      const cachePath = path.join(
        tempDir,
        ".testflowkit/cache/step-definitions-1.2.4.json",
      );
      assert.equal(fs.existsSync(cachePath), true);
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("falls through to GitHub release when the local CLI fails", async (t) => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      t.mock.method(globalThis, "fetch", async () => {
        return {
          ok: true,
          status: 200,
          text: async () => makeCatalogJson("release-fallback"),
        } as Response;
      });

      const result = await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.4",
        cliBinary: "/nonexistent/tkit",
      });

      assert.equal(result.meta.source, "release");
      assert.equal(result.steps[0]?.sentence, "I use release-fallback");
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("uses cache on second call even when CLI binary is available", async () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      const catalogJson = makeCatalogJson("local-cli-first");
      const scriptPath = path.join(tempDir, "fake-tkit.sh");
      fs.writeFileSync(scriptPath, `#!/bin/sh\necho '${catalogJson}'`, {
        mode: 0o755,
      });

      // First call — populates the cache via local CLI
      await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.4",
        cliBinary: scriptPath,
      });

      // Second call — must use cache, not call CLI
      fs.rmSync(scriptPath); // remove the script so CLI would fail if called
      const result = await resolveCatalog(makeConfig(), tempDir, {
        cliVersion: "1.2.4",
        cliBinary: scriptPath,
      });

      assert.equal(result.meta.source, "cache");
      assert.equal(result.steps[0]?.sentence, "I use local-cli-first");
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});
