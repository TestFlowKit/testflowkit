import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { AgentConfigSchema } from "./types.js";

describe("AgentConfigSchema", () => {
  it("applies defaults without release/cache configuration", () => {
    const config = AgentConfigSchema.parse({ version: 1 });

    assert.equal(config.step_catalog.url, undefined);
    assert.equal(config.step_catalog.file, undefined);
    assert.equal(config.step_catalog.cli_version, undefined);
    assert.equal((config.step_catalog as Record<string, unknown>).release, undefined);
    assert.equal((config.step_catalog as Record<string, unknown>).cache, undefined);
  });

  it("rejects user-defined step_catalog.cache", () => {
    const result = AgentConfigSchema.safeParse({
      version: 1,
      step_catalog: {
        cache: {
          path: ".testflowkit/cache/step-definitions.json",
        },
      },
    });

    assert.equal(result.success, false);
    if (!result.success) {
      const hasCacheError = result.error.issues.some((issue) =>
        issue.message.includes("cache"),
      );
      assert.equal(hasCacheError, true);
    }
  });

  it("rejects user-defined step_catalog.release", () => {
    const result = AgentConfigSchema.safeParse({
      version: 1,
      step_catalog: {
        release: {
          repository: "TestFlowKit/testflowkit",
          asset: "step-definitions.json",
        },
      },
    });

    assert.equal(result.success, false);
    if (!result.success) {
      const hasReleaseError = result.error.issues.some((issue) =>
        issue.message.includes("release"),
      );
      assert.equal(hasReleaseError, true);
    }
  });
});
