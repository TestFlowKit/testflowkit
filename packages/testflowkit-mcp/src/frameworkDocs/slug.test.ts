import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { titleToSlug } from "./slug.js";

describe("titleToSlug", () => {
  it("converts titles to snake_case slugs", () => {
    assert.equal(titleToSlug("Global Hooks"), "global_hooks");
    assert.equal(titleToSlug("Random Data"), "random_data");
    assert.equal(titleToSlug("Macros"), "macros");
    assert.equal(titleToSlug("Variables"), "variables");
  });
});
