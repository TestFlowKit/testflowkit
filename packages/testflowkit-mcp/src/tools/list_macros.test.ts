import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { ListMacrosTool } from "./list_macros.js";

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-listmacros-"));
}

describe("ListMacrosTool", () => {
  it("filters macros by nameContains", async () => {
    const dir = mkTmp();
    try {
      fs.writeFileSync(
        path.join(dir, "macros.feature"),
        [
          "Feature: macros",
          "",
          "  @macro",
          "  Scenario: Login as user",
          '    Given the user enters "${email}" into the "email" field',
          "",
          "  @macro",
          "  Scenario: Logout",
          "    Given the user clicks the logout button",
        ].join("\n"),
      );

      const tool = new ListMacrosTool();
      const result = await tool.handler({
        config: {
          configPath: path.join(dir, "testflowkit.yml"),
          featuresGlob: "*.feature",
        },
        configDir: dir,
        input: { nameContains: "login" },
      });

      const text = result.content[0]?.text;
      assert.ok(text);
      const payload = JSON.parse(text ?? "{}");
      assert.equal(payload.count, 1);
      assert.equal(payload.macros[0].name, "Login as user");
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("returns all macros when no filter is given", async () => {
    const dir = mkTmp();
    try {
      fs.writeFileSync(
        path.join(dir, "macros.feature"),
        [
          "Feature: macros",
          "",
          "  @macro",
          "  Scenario: Login as user",
          "    Given a step",
          "",
          "  @macro",
          "  Scenario: Logout",
          "    Given another step",
        ].join("\n"),
      );

      const tool = new ListMacrosTool();
      const result = await tool.handler({
        config: {
          configPath: path.join(dir, "testflowkit.yml"),
          featuresGlob: "*.feature",
        },
        configDir: dir,
        input: {},
      });

      const text = result.content[0]?.text;
      const payload = JSON.parse(text ?? "{}");
      assert.equal(payload.count, 2);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });
});
