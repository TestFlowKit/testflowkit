import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { AgentConfigSchema } from "./types.js";

describe("AgentConfigSchema", () => {
  it("defaults agent.capabilities.macros to true", () => {
    const config = AgentConfigSchema.parse({ version: 1 });
    assert.equal(config.agent.capabilities.macros, true);
  });
});
