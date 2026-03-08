"use strict";

const { describe, it } = require("node:test");
const assert = require("node:assert/strict");
const { getBinaryPath, SUPPORTED_PLATFORMS } = require("./getBinaryPath");

describe("SUPPORTED_PLATFORMS", () => {
  const expectedEntries = [
    ["linux-x64", "@testflowkit/cli-linux-x64", "tkit"],
    ["linux-arm64", "@testflowkit/cli-linux-arm64", "tkit"],
    ["darwin-x64", "@testflowkit/cli-darwin-x64", "tkit"],
    ["darwin-arm64", "@testflowkit/cli-darwin-arm64", "tkit"],
    ["win32-x64", "@testflowkit/cli-win32-x64", "tkit.exe"],
    ["win32-arm64", "@testflowkit/cli-win32-arm64", "tkit.exe"],
  ];

  it("covers exactly 6 platform/arch combinations", () => {
    assert.equal(Object.keys(SUPPORTED_PLATFORMS).length, 6);
  });

  for (const [key, expectedPkg, expectedBin] of expectedEntries) {
    it(`${key} maps to the correct package and binary name`, () => {
      assert.equal(SUPPORTED_PLATFORMS[key].pkg, expectedPkg);
      assert.equal(SUPPORTED_PLATFORMS[key].bin, expectedBin);
    });
  }

  it("uses tkit.exe only for win32 platforms", () => {
    for (const [key, entry] of Object.entries(SUPPORTED_PLATFORMS)) {
      if (key.startsWith("win32-")) {
        assert.equal(entry.bin, "tkit.exe");
      } else {
        assert.equal(entry.bin, "tkit");
      }
    }
  });
});

describe("getBinaryPath()", () => {
  it("throws a descriptive error for an unsupported platform", () => {
    assert.throws(
      () => getBinaryPath("freebsd", "x64"),
      (err) => {
        assert.ok(err instanceof Error);
        assert.ok(
          err.message.includes("Unsupported platform: freebsd-x64"),
          `Expected 'Unsupported platform' in: ${err.message}`
        );
        return true;
      }
    );
  });

  it("throws a descriptive error for an unsupported architecture", () => {
    assert.throws(
      () => getBinaryPath("linux", "mips"),
      (err) => {
        assert.ok(err instanceof Error);
        assert.ok(
          err.message.includes("Unsupported platform: linux-mips"),
          `Expected 'Unsupported platform' in: ${err.message}`
        );
        return true;
      }
    );
  });

  it("throws when native package is not installed (optional dep absent)", () => {
    // The native packages are not installed in the repo, so require.resolve
    // will throw — the wrapper should surface a clear install message.
    for (const [key] of Object.entries(SUPPORTED_PLATFORMS)) {
      const [plt, arc] = key.split("-");
      assert.throws(
        () => getBinaryPath(plt, arc),
        (err) => {
          assert.ok(err instanceof Error);
          assert.ok(
            err.message.includes("npm install") ||
              err.message.includes("Unsupported"),
            `Unexpected error message for ${key}: ${err.message}`
          );
          return true;
        }
      );
    }
  });
});
