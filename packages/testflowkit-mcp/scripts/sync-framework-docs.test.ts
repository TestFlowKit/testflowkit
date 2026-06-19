import assert from "node:assert/strict";
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { describe, it } from "node:test";
import { parse as parseYaml } from "yaml";
import { syncFrameworkDocs } from "./sync-framework-docs.js";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const OUTPUT_DIR = path.resolve(__dirname, "../docs/features");

describe("syncFrameworkDocs", () => {
  it("writes six docs with expected frontmatter and link rewrites", () => {
    syncFrameworkDocs();

    const files = fs
      .readdirSync(OUTPUT_DIR)
      .filter((name) => name.endsWith(".md"))
      .sort();
    assert.deepEqual(files, [
      "api_testing.md",
      "frontend_testing.md",
      "global_hooks.md",
      "macros.md",
      "random_data.md",
      "variables.md",
    ]);

    const variables = fs.readFileSync(
      path.join(OUTPUT_DIR, "variables.md"),
      "utf8",
    );
    const frontmatter = parseYaml(
      variables.match(/^---\r?\n([\s\S]*?)\r?\n---/)![1],
    ) as Record<string, unknown>;

    assert.equal(frontmatter.title, "Variables");
    assert.ok(frontmatter.description);
    assert.equal(frontmatter.navigation, undefined);
    assert.match(variables, /\]\(\.\/random_data\.md\)/);

    const randomData = fs.readFileSync(
      path.join(OUTPUT_DIR, "random_data.md"),
      "utf8",
    );
    const randomFrontmatter = parseYaml(
      randomData.match(/^---\r?\n([\s\S]*?)\r?\n---/)![1],
    ) as Record<string, unknown>;
    assert.equal(randomFrontmatter.title, "Random Data");

    const apiTesting = fs.readFileSync(
      path.join(OUTPUT_DIR, "api_testing.md"),
      "utf8",
    );
    const apiFrontmatter = parseYaml(
      apiTesting.match(/^---\r?\n([\s\S]*?)\r?\n---/)![1],
    ) as Record<string, unknown>;
    assert.equal(apiFrontmatter.title, "API Testing");
    assert.ok(apiFrontmatter.description);
    assert.match(apiTesting, /\]\(\.\/variables\.md\)/);

    const frontendTesting = fs.readFileSync(
      path.join(OUTPUT_DIR, "frontend_testing.md"),
      "utf8",
    );
    const frontendFrontmatter = parseYaml(
      frontendTesting.match(/^---\r?\n([\s\S]*?)\r?\n---/)![1],
    ) as Record<string, unknown>;
    assert.equal(frontendFrontmatter.title, "Frontend Testing");
    assert.ok(frontendFrontmatter.description);
  });
});
