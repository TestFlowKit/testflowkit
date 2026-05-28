import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { searchSentences } from "./searchSentences.js";
import type { StepDefinition } from "../config/types.js";

const steps: StepDefinition[] = [
  {
    sentence: "I prepare a request to {string}",
    description: "Prepares a request to a named API endpoint",
    categories: ["restapi"],
    example: 'Given I prepare a request to "users_api.getUser"',
  },
  {
    sentence: "I send the request",
    description: "Sends the prepared request",
    categories: ["restapi"],
    example: "When I send the request",
  },
  {
    sentence: "the response status code should be {int}",
    description: "Asserts response status code",
    categories: ["restapi", "assertions"],
    example: "Then the response status code should be 200",
  },
  {
    sentence: "the user clicks the {string} button",
    description: "Clicks a button element",
    categories: ["frontend", "mouse"],
    example: 'When the user clicks the "login" button',
  },
];

describe("searchSentences", () => {
  it("returns all results when query is empty", () => {
    assert.equal(searchSentences(steps, "").length, 4);
  });

  it("filters by keyword in sentence", () => {
    const results = searchSentences(steps, "prepare request");
    assert.ok(results.length >= 1);
    assert.equal(results[0].sentence, "I prepare a request to {string}");
  });

  it("filters by keyword in description", () => {
    const results = searchSentences(steps, "clicks");
    assert.equal(results.length, 1);
    assert.equal(results[0].sentence, "the user clicks the {string} button");
  });

  it("filters by category", () => {
    const results = searchSentences(steps, "", "frontend");
    assert.equal(results.length, 1);
    assert.ok(results[0].categories.includes("frontend"));
  });

  it("combines keyword and category filter", () => {
    const results = searchSentences(steps, "response", "restapi");
    assert.equal(results.length, 1);
    assert.equal(
      results[0].sentence,
      "the response status code should be {int}"
    );
  });

  it("returns empty when nothing matches", () => {
    assert.equal(searchSentences(steps, "nonexistentXYZ").length, 0);
  });
});
