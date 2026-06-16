import assert from "node:assert/strict";
import { describe, it } from "node:test";
import type { StepDefinition } from "../config/types.js";
import { filterStepsByCategory, getStepCategories } from "./categories.js";

const steps: StepDefinition[] = [
  {
    sentence: "step one",
    description: "First step",
    categories: ["restapi", "assertions"],
    example: "Given step one",
  },
  {
    sentence: "step two",
    description: "Second step",
    categories: ["frontend"],
    example: "When step two",
  },
  {
    sentence: "step three",
    description: "Third step",
    categories: ["restapi"],
    example: "Then step three",
  },
];

describe("getStepCategories", () => {
  it("returns sorted unique categories", () => {
    assert.deepEqual(getStepCategories(steps), [
      "assertions",
      "frontend",
      "restapi",
    ]);
  });
});

describe("filterStepsByCategory", () => {
  it("returns steps that include the category", () => {
    const results = filterStepsByCategory(steps, "restapi");
    assert.equal(results.length, 2);
    assert.ok(results.every((step) => step.categories.includes("restapi")));
  });

  it("returns empty when category is unknown", () => {
    assert.equal(filterStepsByCategory(steps, "unknown").length, 0);
  });
});
