import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { ProjectConfigSchema, toFeaturesGlob } from "./types.js";

describe("ProjectConfigSchema", () => {
  it("defaults settings.gherkin_location to ./features", () => {
    const config = ProjectConfigSchema.parse({});
    assert.equal(config.settings.gherkin_location, "./features");
  });

  it("parses agent.default_tags_for_draft", () => {
    const config = ProjectConfigSchema.parse({
      agent: { default_tags_for_draft: "@wip @ai-generated" },
    });
    assert.equal(config.agent?.default_tags_for_draft, "@wip @ai-generated");
  });

  it("parses agent.step_catalog with file and url", () => {
    const config = ProjectConfigSchema.parse({
      agent: {
        step_catalog: {
          file: "./build/step-definitions.json",
          url: "https://example.com/steps.json",
        },
      },
    });
    assert.equal(
      config.agent?.step_catalog?.file,
      "./build/step-definitions.json",
    );
    assert.equal(
      config.agent?.step_catalog?.url,
      "https://example.com/steps.json",
    );
  });

  it("treats agent section as optional", () => {
    const config = ProjectConfigSchema.parse({});
    assert.equal(config.agent, undefined);
  });
});

describe("toFeaturesGlob", () => {
  it("appends a feature glob to a directory path", () => {
    assert.equal(toFeaturesGlob("./features"), "./features/**/*.feature");
  });

  it("keeps an existing glob unchanged", () => {
    assert.equal(
      toFeaturesGlob("features/**/*.feature"),
      "features/**/*.feature",
    );
  });
});
