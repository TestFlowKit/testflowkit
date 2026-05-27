import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { toSemverBase, isCanary, versionMismatchLevel } from "./version.js";

describe("toSemverBase", () => {
  it("returns version unchanged for stable", () => {
    assert.equal(toSemverBase("1.2.4"), "1.2.4");
  });

  it("strips canary suffix", () => {
    assert.equal(toSemverBase("3.6.1-canary.abc1234"), "3.6.1");
  });
});

describe("isCanary", () => {
  it("returns false for stable versions", () => {
    assert.equal(isCanary("1.2.4"), false);
  });

  it("returns true for canary versions", () => {
    assert.equal(isCanary("3.6.1-canary.abc1234"), true);
  });
});

describe("versionMismatchLevel", () => {
  it("ok when versions match", () => {
    assert.equal(versionMismatchLevel("1.2.4", "1.2.4"), "ok");
  });

  it("patch when only patch differs", () => {
    assert.equal(versionMismatchLevel("1.2.4", "1.2.3"), "patch");
  });

  it("minor when minor differs", () => {
    assert.equal(versionMismatchLevel("1.2.4", "1.1.4"), "minor");
  });

  it("major when major differs", () => {
    assert.equal(versionMismatchLevel("2.0.0", "1.0.0"), "major");
  });

  it("uses semver base of canary CLI", () => {
    assert.equal(versionMismatchLevel("1.2.4-canary.abc", "1.2.4"), "ok");
  });
});
