import assert from "node:assert/strict";
import path from "node:path";
import { describe, it } from "node:test";
import { buildFrameworkDocsIndex } from "./indexer.js";
import { resolveFrameworkDocsDir } from "./resolveDocsDir.js";

describe("buildFrameworkDocsIndex", () => {
  it("indexes all six bundled framework docs", () => {
    const index = buildFrameworkDocsIndex();
    assert.equal(index.size, 6);
    assert.ok(index.has("macros"));
    assert.ok(index.has("random_data"));
    assert.ok(index.has("global_hooks"));
    assert.ok(index.has("variables"));
    assert.ok(index.has("api_testing"));
    assert.ok(index.has("frontend_testing"));
  });

  it("parses title and description from frontmatter", () => {
    const index = buildFrameworkDocsIndex();
    const hooks = index.get("global_hooks");
    assert.ok(hooks);
    assert.equal(hooks.title, "Global Hooks");
    assert.match(hooks.description, /teardown/i);
    assert.equal(
      hooks.filePath,
      path.join(resolveFrameworkDocsDir(), "global_hooks.md"),
    );
    assert.match(hooks.content, /^---\n/);
  });
});
