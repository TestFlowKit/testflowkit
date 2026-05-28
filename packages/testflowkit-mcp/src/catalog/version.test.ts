import assert from "node:assert/strict";
import { describe, it } from "node:test";
import {
  toSemverBase,
  isCanary,
  hasGitDescribeSuffix,
  stripGitDescribeSuffix,
  versionMismatchLevel,
} from "./version.js";

describe("toSemverBase", () => {
  it("returns version unchanged for stable", () => {
    assert.equal(toSemverBase("1.2.4"), "1.2.4");
  });

  it("strips canary suffix", () => {
    assert.equal(toSemverBase("3.6.1-canary.abc1234"), "3.6.1");
  });

  it("strips git-describe and canary suffix", () => {
    assert.equal(toSemverBase("3.7.0-canary.fed5ed9-2-gba72690"), "3.7.0");
  });
});

describe("isCanary", () => {
  it("returns false for stable versions", () => {
    assert.equal(isCanary("1.2.4"), false);
  });

  it("returns true for canary versions", () => {
    assert.equal(isCanary("3.6.1-canary.abc1234"), true);
  });

  it("returns true for canary versions with git-describe suffix", () => {
    assert.equal(isCanary("3.7.0-canary.fed5ed9-2-gba72690"), true);
  });
});

describe("hasGitDescribeSuffix", () => {
  it("returns false for stable versions", () => {
    assert.equal(hasGitDescribeSuffix("1.2.4"), false);
  });

  it("returns false for plain canary versions", () => {
    assert.equal(hasGitDescribeSuffix("3.6.1-canary.abc1234"), false);
  });

  it("returns true for versions with git-describe suffix", () => {
    assert.equal(hasGitDescribeSuffix("3.7.0-canary.fed5ed9-2-gba72690"), true);
  });
});

describe("stripGitDescribeSuffix", () => {
  it("returns version unchanged when no git-describe suffix", () => {
    assert.equal(
      stripGitDescribeSuffix("3.6.1-canary.abc1234"),
      "3.6.1-canary.abc1234",
    );
  });

  it("strips the git-describe suffix, leaving the canary base", () => {
    assert.equal(
      stripGitDescribeSuffix("3.7.0-canary.fed5ed9-2-gba72690"),
      "3.7.0-canary.fed5ed9",
    );
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
