import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { resolveCatalog } from "./resolveCatalog.js";

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

describe("resolveCatalog local CLI behavior", () => {
  it("returns source: local when a working CLI binary is provided", async () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-"));

    try {
      const catalogJson = makeCatalogJson("local-cli");
      const scriptPath = path.join(tempDir, "fake-tkit.sh");
      fs.writeFileSync(scriptPath, `#!/bin/sh\necho '${catalogJson}'`, {
        mode: 0o755,
      });

      const result = await resolveCatalog(tempDir, {
        cliBinary: scriptPath,
      });

      assert.equal(result.meta.source, "local");
      assert.equal(result.steps[0]?.sentence, "I use local-cli");
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});
