import assert from "node:assert/strict";
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { describe, it } from "node:test";
import { resolveFrameworkDocsDir } from "./resolveDocsDir.js";

const packageRoot = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  "../..",
);

describe("resolveFrameworkDocsDir", () => {
  it("resolves the bundled docs directory from source modules", () => {
    const resolved = resolveFrameworkDocsDir();
    assert.equal(resolved, path.join(packageRoot, "docs/features"));
    assert.ok(fs.existsSync(resolved));
  });

  it("resolves the bundled docs directory from the dist entrypoint layout", () => {
    const distDir = path.join(packageRoot, "dist");
    const candidate = path.resolve(distDir, "../docs/features");
    assert.equal(candidate, path.join(packageRoot, "docs/features"));
    assert.ok(fs.existsSync(candidate));
  });
});
