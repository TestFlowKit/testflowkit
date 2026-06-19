import assert from "node:assert/strict";
import { describe, it } from "node:test";
import {
  FRAMEWORK_DOCS_INDEX_URI,
  frameworkDocUri,
  registerFrameworkDocResources,
} from "./registerResources.js";
import type { FrameworkDocsIndex } from "./types.js";

const BUNDLED_DOC_SLUGS = [
  "api_testing",
  "frontend_testing",
  "global_hooks",
  "macros",
  "random_data",
  "variables",
] as const;

function createMockFrameworkDocsIndex(): FrameworkDocsIndex {
  const index: FrameworkDocsIndex = new Map();

  for (const slug of BUNDLED_DOC_SLUGS) {
    index.set(slug, {
      slug,
      title: slug,
      description: `Description for ${slug}`,
      filePath: `/fake/${slug}.md`,
      content: `---\ntitle: ${slug}\n---\n\n# ${slug}`,
    });
  }

  return index;
}

describe("registerFrameworkDocResources", () => {
  it("defines the index URI and one URI per bundled doc", () => {
    assert.equal(FRAMEWORK_DOCS_INDEX_URI, "docs://framework/features/index");
    assert.equal(frameworkDocUri("global_hooks"), "docs://framework/features/global_hooks");

    for (const slug of BUNDLED_DOC_SLUGS) {
      assert.equal(frameworkDocUri(slug), `docs://framework/features/${slug}`);
    }
  });

  it("registers static resources on the MCP server", () => {
    const index = createMockFrameworkDocsIndex();
    const registered: Array<{ name: string; uri: string }> = [];

    const server = {
      registerResource(
        name: string,
        uri: string,
        _config: unknown,
        _handler: unknown,
      ) {
        registered.push({ name, uri });
      },
    };

    registerFrameworkDocResources(server as never, index);

    assert.equal(registered.length, 7);
    assert.deepEqual(
      registered.map((entry) => entry.uri).sort(),
      [
        "docs://framework/features/api_testing",
        "docs://framework/features/frontend_testing",
        "docs://framework/features/global_hooks",
        "docs://framework/features/index",
        "docs://framework/features/macros",
        "docs://framework/features/random_data",
        "docs://framework/features/variables",
      ],
    );
  });
});
