import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { resolveFeaturePath, writeFeatureFile } from "./writeFeature.js";

const GLOB = "features/**/*.feature";

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tfk-mcp-test-"));
}

describe("resolveFeaturePath", () => {
  it("accepts a valid feature path", () => {
    const dir = mkTmp();
    const result = resolveFeaturePath("features/a.feature", GLOB, dir);
    assert.equal(result, path.join(dir, "features", "a.feature"));
    fs.rmSync(dir, { recursive: true });
  });

  it("rejects path traversal above workspace root", () => {
    const dir = mkTmp();
    assert.throws(
      () => resolveFeaturePath("../../etc/passwd", GLOB, dir),
      { message: /Security/ }
    );
    fs.rmSync(dir, { recursive: true });
  });

  it("rejects path that does not match features_glob", () => {
    const dir = mkTmp();
    assert.throws(
      () => resolveFeaturePath("config/secret.yml", GLOB, dir),
      { message: /does not match features_glob/ }
    );
    fs.rmSync(dir, { recursive: true });
  });
});

describe("writeFeatureFile", () => {
  it("creates a feature file and its parents", () => {
    const dir = mkTmp();
    writeFeatureFile(
      "features/auth/login.feature",
      "Feature: login\n",
      GLOB,
      dir
    );
    const content = fs.readFileSync(
      path.join(dir, "features", "auth", "login.feature"),
      "utf-8"
    );
    assert.equal(content, "Feature: login\n");
    fs.rmSync(dir, { recursive: true });
  });
});
