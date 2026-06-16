import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { toSemverBase } from "./version.js";

describe("toSemverBase", () => {
  it("returns version unchanged for stable", () => {
    assert.equal(toSemverBase("1.2.4"), "1.2.4");
  });

  it("strips canary suffix", () => {
    assert.equal(toSemverBase("3.6.1-canary.abc1234"), "3.6.1");
  });
});
